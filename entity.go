package protostub

// ServiceStub represents a template for generating gRPC service code.
// It holds metadata about a gRPC service, such as its name and the method details.
type ServiceStub struct {
	// ServiceName is the name of the gRPC service.
	ServiceName string

	// Package is the name of protobuf package
	Package string

	// GoPackage is the name of protobuf go package
	GoPackage string

	// Method is the name of the method that the service provides.
	Method string

	// InputType is the type of the request message that the method accepts.
	InputType string

	// OutputType is the type of the response message that the method returns.
	OutputType string
}
