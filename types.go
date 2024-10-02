package protostub

import (
	"errors"
)

var ErrNoType = errors.New("no type defined")

// ProtoStubType represents the type of stub that will be generated (client or server).
type ProtoStubType uint8

const (
	ProtostubServerType ProtoStubType = iota
	ProtostubClientType
)

func (p ProtoStubType) String() string {
	switch p {
	case ProtostubServerType:
		return "server"
	case ProtostubClientType:
		return "client"
	}

	return "client"
}

func (p ProtoStubType) extractExtension() string {
	switch p {
	case ProtostubClientType:
		return "_client.go"
	case ProtostubServerType:
		return "_impl.go"
	}

	return ""
}

func (p ProtoStubType) renderer() (string, error) {
	switch p {
	case ProtostubClientType:
		return string(clientTmpl), nil
	case ProtostubServerType:
		return string(serviceTmpl), nil
	}

	return "", ErrNoType
}

func RenderType(typeName string) (ProtoStubType, error) {
	switch typeName {
	case "client":
		return ProtostubClientType, nil
	case "server":
		return ProtostubServerType, nil
	}

	return 0, ErrNoType
}
