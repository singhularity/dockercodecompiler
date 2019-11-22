package main

import (
	"github.com/dockercodecompiler/compiler"
	"github.com/dockercodecompiler/server"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Must specify 'svc' or 'local' as an option!")
	}
	runType := os.Args[1]
	if runType == "svc" {
		server.Serve()
	} else if runType == "local" {
		fmt.Print(compiler.Compile(os.Args[1:]))
	} else {
		panic("Invalid run option!")
	}
}
