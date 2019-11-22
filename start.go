package main

import (
	"errors"
	"os"

	"dockercodecompiler/dockerCompiler"
	"dockercodecompiler/sandbox/sandBoxPrepper"
)

func main() {
	argLength := len(os.Args)
	if argLength < 3 {
		panic(errors.New("Inssuficient arguments to run compiler"))
	}

	if argLength == 3 {
		mountAndCompile(os.Args[1], os.Args[2], "")
	} else {
		mountAndCompile(os.Args[1], os.Args[2], os.Args[3])
	}
}

func mountAndCompile(language string, codeData string, inputData string) {
	mountPath := sandBoxPrepper.PrepSandbox(language, codeData, inputData)
	dockerCompiler.CompileAndRun(language, mountPath)
	os.RemoveAll(mountPath)
}
