package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/Ow1Dev/QustionOfToday/cmd"
)

var (
	Version     = "development" // program version for this build
	Tags        = ""            // the Golang build tags
	MakeVersion = ""            // "make" program version if built with make
)

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(stdout io.Writer, args []string) error {
	app := cmd.NewMainApp(cmd.AppVersion{Version: Version, Extra: formatBuiltWith()})
	return cmd.RunMainApp(app, os.Args...)
}

func formatBuiltWith() string {
	version := runtime.Version()
	if len(MakeVersion) > 0 {
		version = MakeVersion + ", " + runtime.Version()
	}
	if len(Tags) == 0 {
		return " built with " + version
	}

	return " built with " + version + " : " + strings.ReplaceAll(Tags, " ", ", ")
}
