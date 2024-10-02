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
// It returns an Option that modifies the serviceDir field of a ProtoStub.
//
// serviceDir: The directory path where the generated code will be saved.
func WithServiceDir(serviceDir string) Option {
	return func(ps *ProtoStub) {
		ps.ServiceDir = serviceDir
	}
}

// WithClientDir sets the destination directory for the generated client stub files.
// It returns an Option that modifies the clientDir field of a ProtoStub.
//
// clientDir: The directory path where the generated code will be saved.
func WithClientDir(clientDir string) Option {
	return func(ps *ProtoStub) {
		ps.ClientDir = clientDir
	}
}

// WithType sets the type of stub that will be generated (client or server).
// It returns an Option that modifies the type of the ProtoStub.
//
// typeName: The type of stub to generate, either ProtoStubServerType or ProtoStubClientType.
func WithType(typeName ProtoStubType) Option {
	return func(ps *ProtoStub) {
		ps.TypeName = typeName
	}
}
