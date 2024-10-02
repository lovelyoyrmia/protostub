package main

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/lovelyoyrmia/protostub"
)

const helpMessage = `
This library provides a simple way to generate Service Stub from proto files.

EXAMPLE: Basic Command
protostub --proto_dir=protos/

See https://github.com/lovelyoyrmia/protostub for more details.
`

// Version returns the currently running version of protodoc
func Version() string {
	return protostub.VERSION
}

// Flags contains details about the CLI invocation of protodoc
type Flags struct {
	appName     string
	flagSet     *flag.FlagSet
	err         error
	showHelp    bool
	showVersion bool
	protoDir    string
	destDir     string
	serviceDir  string
	writer      io.Writer
}

// Code returns the status code to exit with after handling the supplied flags
func (f *Flags) Code() int {
	if f.err != nil {
		return 1
	}

	return 0
}

// ShowHelp determines whether or not to show the help message
func (f *Flags) ShowHelp() bool {
	return f.err != nil || f.showHelp
}

// ShowVersion determines whether or not to show the version message
func (f *Flags) ShowVersion() bool {
	return f.showVersion
}

// CheckRequiredArgs function to check args are required
func (f *Flags) CheckRequiredArgs(fields map[string]string) bool {
	var missingFields []string
	for fieldName, fieldValue := range fields {
		if fieldValue == "" {
			missingFields = append(missingFields, fieldName)
		}
	}
	if len(missingFields) > 0 {
		f.err = fmt.Errorf("error: The following fields must not be empty: %s", strings.Join(missingFields, ", "))
		return false
	}
	return true
}

// PrintHelp prints the usage string including all flags to the `io.Writer` that was supplied to the `Flags` object.
func (f *Flags) PrintHelp() {
	fmt.Fprintf(f.writer, "Usage of %s:\n", f.appName)
	fmt.Fprintf(f.writer, "%s\n", helpMessage)
	fmt.Fprintf(f.writer, "FLAGS\n")
	f.flagSet.PrintDefaults()
}

// PrintVersion prints the version string to the `io.Writer` that was supplied to the `Flags` object.
func (f *Flags) PrintVersion() {
	fmt.Fprintf(f.writer, "%s version %s\n", f.appName, Version())
}

// PrintError prints the error string
func (f *Flags) PrintError() {
	if f.err == nil {
		return
	}
	fmt.Println(f.err)
	fmt.Println("Use --help or -h for usage information.")
}

// ParseFlags parses the supplied options are returns a `Flags` object to the caller.
func ParseFlags(w io.Writer, args []string) *Flags {
	f := Flags{appName: args[0], writer: w}

	f.flagSet = flag.NewFlagSet(args[0], flag.ContinueOnError)
	f.flagSet.StringVar(&f.protoDir, "proto_dir", "", "proto_dir is the directory of the all protobuf files.")
	f.flagSet.StringVar(&f.destDir, "dest_dir", ".", "dest_dir is the output directory of the go and grpc services.")
	f.flagSet.StringVar(&f.serviceDir, "service_dir", ".", "service_dir is the output directory of the services stub.")

	f.flagSet.BoolVar(&f.showHelp, "help", false, "Show this help message")
	f.flagSet.BoolVar(&f.showVersion, "version", false, fmt.Sprintf("Print the current version (%v)", Version()))
	f.flagSet.SetOutput(w)

	// prevent showing help on parse error
	f.flagSet.Usage = func() {}

	f.err = f.flagSet.Parse(args[1:])
	return &f
}
