package protostub

// Option is a function type that modifies a ProtoStub.
// It allows setting optional parameters for ProtoStub configuration.
type Option func(*ProtoStub)

// WithProtoDir sets the directory where the .proto files are located.
// It returns an Option that modifies the ProtoDir field of a ProtoStub.
//
// protoDir: The directory path to the .proto files.
func WithProtoDir(protoDir string) Option {
	return func(ps *ProtoStub) {
		ps.ProtoDir = protoDir
	}
}

// WithDestDir sets the destination directory for the generated Go and gRPC files.
// It returns an Option that modifies the DestDir field of a ProtoStub.
//
// destDir: The directory path where the generated code will be saved.
func WithDestDir(destDir string) Option {
	return func(ps *ProtoStub) {
		ps.DestDir = destDir
	}
}

// WithServiceDir sets the destination directory for the generated service stub files.
// It returns an Option that modifies the SserviceDir field of a ProtoStub.
//
// serviceDir: The directory path where the generated code will be saved.
func WithServiceDir(serviceDir string) Option {
	return func(ps *ProtoStub) {
		ps.ServiceDir = serviceDir
	}
}
