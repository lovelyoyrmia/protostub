package protostub

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lovelyoyrmia/protodoc"
)

// ProtoStub represents the necessary information for handling
// Protocol Buffers (.proto) files, including their source directory
// and the destination directory for the generated files.
type ProtoStub struct {
	// ProtoDir is the directory where the .proto files are located.
	ProtoDir string

	// DestDir is the destination directory where the generated
	// Go and gRPC code will be placed after running protoc.
	DestDir string

	// ServiceDir is the directory where the generated service stub will be placed.
	ServiceDir string
}

func New(opts ...Option) *ProtoStub {
	ps := &ProtoStub{
		ProtoDir: defaultProtoDir,
		DestDir:  defaultDestDir,
	}

	for _, opt := range opts {
		opt(ps)
	}

	return ps
}

// Generate generates additional server scaffolding code.
func (ps *ProtoStub) Generate() error {
	services, err := ps.GenerateServices()
	if err != nil {
		return err
	}

	for _, service := range services {
		data, err := RenderTemplate(service)
		if err != nil {
			return err
		}

		fileName := toSnakeCase(service.ServiceName)
		serverFile := filepath.Join(ps.ServiceDir, fileName)

		f, err := os.Create(serverFile)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.Write(data); err != nil {
			return err
		}
	}

	return nil
}

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

	if apiDoc.Services != nil {
		for _, service := range apiDoc.Services {
			serviceStub := &ServiceStub{
				GoPackage:   apiDoc.GoPackage,
				Package:     apiDoc.Package,
				ServiceName: service.Name,
				Method:      service.Methods[0].Name,
				InputType:   strings.TrimPrefix(service.Methods[0].InputType, "#"),
				OutputType:  strings.TrimPrefix(service.Methods[0].OutputType, "#"),
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
