package main

import (
	"errors"
	"os"

	"../compiler"
	"../prepSandbox"
)

func main() {
	argLength := len(os.Args)
	if argLength < 3 {
		panic(errors.New("Inssuficient arguments to run compiler"))
	} else if argLength == 3 {
		mountAndCompile(os.Args[1], os.Args[2], "")
	} else {
		mountAndCompile(os.Args[1], os.Args[2], os.Args[3])
	}
}

func mountAndCompile(language string, codeData string, inputData string) {
	mountPath := prepSandbox.PrepSandbox(language, codeData, inputData)
	compiler.CompileAndRun(language, mountPath)
	os.RemoveAll(mountPath)
}
