package main

import (
	"errors"
	"os"

	"../compiler"
	"../prepSandbox"
)

func main() {
	argLength := len(os.Args)
	if argLength == 1 {
		panic(errors.New("No code data supplied"))
	} else if argLength == 2 {
		mountAndCompile(os.Args[1], "")
	} else if len(os.Args) > 2 {
		mountAndCompile(os.Args[1], os.Args[2])
	}
}

func mountAndCompile(codeData string, inputData string) {
	mountPath := prepSandbox.PrepSandbox(codeData, inputData)
	compiler.CompileAndRun(mountPath)
	os.RemoveAll(mountPath)
}
