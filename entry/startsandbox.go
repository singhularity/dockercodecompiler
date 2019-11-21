package main

import (
	"errors"
	"fmt"
	"os"

	"../compiler"
	"../prep"
)

func main() {
	argLength := len(os.Args)
	if argLength == 1 {
		fmt.Print(errors.New("No code data to process"))
	} else if argLength == 2 {
		mountAndCompile(os.Args[1], "")
	} else if len(os.Args) > 2 {
		mountAndCompile(os.Args[1], os.Args[2])
	}
}

func mountAndCompile(codeData string, inputData string) {
	mountPath := prep.Prepcode(codeData, inputData)
	compiler.Compile(mountPath)
	os.RemoveAll(mountPath)
}
