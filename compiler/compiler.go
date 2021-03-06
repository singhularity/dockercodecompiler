package compiler

import (
	"os"

	"github.com/dockercodecompiler/compiler/dockerCompiler"
	"github.com/dockercodecompiler/compiler/sandbox/staging"
)

func Compile(params []string) string {
	argLength := len(params)
	if argLength < 3 {
		return "Insufficient arguments to run compiler"
	}

	if argLength == 3 {
		return mountAndCompile(params[1], params[2], "")
	} else {
		return mountAndCompile(params[1], params[2], params[3])
	}
}

func mountAndCompile(language string, codeData string, inputData string) string {
	mountPath := staging.PrepSandbox(language, codeData, inputData)
	runOutput := dockerCompiler.CompileAndRun(language, mountPath)
	os.RemoveAll(mountPath)
	return runOutput
}
