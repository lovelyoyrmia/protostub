package main

import (
	"fmt"

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
}
