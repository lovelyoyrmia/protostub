package protostub

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/lovelyoyrmia/protodoc"
)

// GenerateServices function to generate all protofiles and convert to `FileDescriptorProto`
// and mapping to service stub.
func (ps *ProtoStub) GenerateServices() ([]*ServiceStub, error) {
	if err := ps.generateProtoFiles(); err != nil {
		return nil, err
	}

	fileDesc, err := protodoc.GenerateDescriptor(protodoc.DefaultDescriptorFile)
	if err != nil {
		return nil, err
	}

	pbDoc := protodoc.New(
		protodoc.WithFileDescriptor(fileDesc),
		protodoc.WithType(protodoc.ProtodocTypeJson),
	)

	res, err := pbDoc.Generate()
	if err != nil {
		return nil, err
	}

	var apiDoc protodoc.APIDoc
	if err := json.Unmarshal(res, &apiDoc); err != nil {
		return nil, err
	}

	services := make([]*ServiceStub, 0)

	var packageName string

	switch ps.TypeName {
	case ProtostubServerType:
		packageName = filepath.Base(ps.ServiceDir)
	case ProtostubClientType:
		packageName = filepath.Base(ps.ClientDir)
	}

	if apiDoc.Services != nil {
		for _, service := range apiDoc.Services {
			methods := make([]Method, 0)

			for _, method := range service.Methods {
				methods = append(methods, Method{
					ServiceName:  service.Name,
					Method:       method.Name,
					InputType:    strings.Trim(method.InputType, "#"),
					OutputType:   strings.Trim(method.OutputType, "#"),
					ProtoPackage: apiDoc.Package,
				})
			}

			serviceStub := &ServiceStub{
				Package:      packageName,
				GoPackage:    apiDoc.GoPackage,
				ProtoPackage: apiDoc.Package,
				ServiceName:  service.Name,
				Methods:      methods,
			}

			services = append(services, serviceStub)
		}
	}

	return services, nil
}

func (ps *ProtoStub) getAllProtoFiles() ([]string, error) {
	var protoFiles []string

	files, err := os.ReadDir(ps.ProtoDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".proto" {
			protoFiles = append(protoFiles, filepath.Join(ps.ProtoDir, file.Name()))
		}
	}

	return protoFiles, nil
}

func (ps *ProtoStub) generateProtoFiles() error {
	// Gather all .proto files
	protoFiles, err := ps.getAllProtoFiles()
	if err != nil {
		return err
	}

	// Prepare the protoc command with all proto files
	cmdArgs := append([]string{
		"--proto_path=" + ps.ProtoDir,
		"--descriptor_set_out=" + protodoc.DefaultDescriptorFile,
		"--go_out=" + ps.DestDir,
		"--go_opt=paths=source_relative",
		"--go-grpc_opt=paths=source_relative",
		"--go-grpc_out=" + ps.DestDir,
	}, protoFiles...)

	// Exec command protoc to generate descriptor file
	cmd := exec.Command("protoc", cmdArgs...)

	// Capture output and error
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// Function to extract existing method names from the file (ignoring the body).
func (ps *ProtoStub) extractMethodSignatures(filePath string) (map[string]bool, error) {
	existingMethods := make(map[string]bool)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Define a regex to match method signatures, e.g., `func (s *ServiceImpl) MethodName`
	methodRegex := regexp.MustCompile(`func \(s \*([a-zA-Z]+)\) ([a-zA-Z]+)\(`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := methodRegex.FindStringSubmatch(line)
		if len(match) > 2 {
			existingMethods[match[2]] = true
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return existingMethods, nil
}
