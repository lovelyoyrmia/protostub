package main

import (
	"fmt"
	"os"

	"github.com/lovelyoyrmia/protodoc"
	"github.com/lovelyoyrmia/protostub"
)

func main() {
	flags := ParseFlags(os.Stdout, os.Args)

	if flags.handleFlags() {
		os.Exit(flags.Code())
	}

	ps := protostub.New(
		protostub.WithProtoDir(flags.protoDir),
		protostub.WithDestDir(flags.destDir),
		protostub.WithServiceDir(flags.serviceDir),
	)

	if err := ps.Generate(); err != nil {
		fmt.Printf("failed to generate service stub, err=%v\n", err)
		return
	}

	// Clean Up
	if err := os.Remove(protodoc.DefaultDescriptorFile); err != nil {
		fmt.Printf("failed to remove desc file: err=%v\n", err)
		return
	}

	fmt.Println("✅ Generate Code Success !")
}

// handleFlags checks if there's a match and returns true if it was "handled"
func (f *Flags) handleFlags() bool {
	if f.ShowHelp() {
		f.PrintHelp()
		return true
	}

	if f.ShowVersion() {
		f.PrintVersion()
		return true
	}

	// Check all required fields
	if !f.CheckRequiredArgs(map[string]string{
		"proto_dir":   f.protoDir,
		"service_dir": f.serviceDir,
	}) {
		f.PrintError()
		return true
	}

	return false
}
