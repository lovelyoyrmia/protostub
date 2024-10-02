package main

import (
	"fmt"
	"os"

	"github.com/lovelyoyrmia/protodoc"
	"github.com/lovelyoyrmia/protostub"
)

// This is the main examples for the protostub generator
func main() {
	protoDir := "."
	destDir := "./pb"
	serviceDir := "./service"

	ps := protostub.New(
		protostub.WithProtoDir(protoDir),
		protostub.WithDestDir(destDir),
		protostub.WithServiceDir(serviceDir),
	)

	if err := ps.Generate(); err != nil {
		fmt.Printf("failed to generate service stub, err=%v\n", err)
		return
	}

	if err := os.Remove(protodoc.DefaultDescriptorFile); err != nil {
		fmt.Printf("failed to execute, err=%v\n", err)
		return
	}
}
