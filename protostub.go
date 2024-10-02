package protostub

import (
	"os"
	"path/filepath"
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

	// ClientDir is the directory where the generated client stub will be placed.
	ClientDir string

	// TypeName is the type of stub that will be generated (client or server)
	TypeName ProtoStubType
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
func (s *ProtoStub) Generate() error {
	services, err := s.GenerateServices()
	if err != nil {
		return err
	}

	for _, service := range services {
		data, err := RenderTemplate(s.TypeName, service)
		if err != nil {
			return err
		}

		fileName := toSnakeCase(service.ServiceName) + s.TypeName.extractExtension()

		var serverFile string
		switch s.TypeName {
		case ProtostubClientType:
			serverFile = filepath.Join(s.ClientDir, fileName)
		case ProtostubServerType:
			serverFile = filepath.Join(s.ServiceDir, fileName)
		}

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
