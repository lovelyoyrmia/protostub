package protostub

// ServiceStub represents a template for generating gRPC service code.
// It holds metadata about a gRPC service, such as its name and the method details.
type ServiceStub struct {
	// ServiceName is the name of the gRPC service.
	ServiceName string

	// ProtoPackage is the name of protobuf package
	//  // Example:
	//  syntax = "proto3";
	//  package pb;
	ProtoPackage string

	// GoPackage is the name of protobuf go package
	//  // Example:
	//  option go_package = "github.com/lovelyoyrmia/protostub/examples/pb";
	GoPackage string

	// Package is the name of generated go package
	//  // Example:
	//  package examples
	Package string

	// Methods is a list of methods from the service
	Methods []Method
}

type Method struct {
	// ServiceName is the name of the gRPC service.
	ServiceName string

	// ProtoPackage is the name of protobuf package
	//  // Example:
	//  syntax = "proto3";
	//  package pb;
	ProtoPackage string

	// Method is the name of the method that the service provides.
	Method string

	// InputType is the type of the request message that the method accepts.
	InputType string

	// OutputType is the type of the response message that the method returns.
	OutputType string
}
