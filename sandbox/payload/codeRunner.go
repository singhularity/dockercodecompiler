package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 5 {
		panic("Not enough arguments to run code!")
	}
	execCommand := os.Args[1]
	codeFileName := os.Args[2]
	mountDir := os.Args[3]
	inputFileName := os.Args[4]
	if execCommand == "java" {
		handleJavaCode(mountDir, codeFileName, inputFileName)
	} else {
		runCode(mountDir, execCommand, codeFileName, inputFileName)
	}
}

func handleJavaCode(mountDir string, codeFileName string, inputFileName string) {
	compileJavaCode(mountDir, codeFileName)
	mainClassToRun := getMainJavaClass(mountDir)
	runCode(mountDir, "java", mainClassToRun, inputFileName)
}

func runCode(mountDir string, execCommand string, codeFile string, inputFileName string) {
	inputContent := getInputFileContents(mountDir, inputFileName)
	runCMD := exec.Command(execCommand, codeFile, inputContent)
	runCMD.Dir = mountDir
	runOut, runErr := runCMD.CombinedOutput()
	if runErr != nil {
		fmt.Printf("**********ERROR**********\n %v\n%v\n", runErr, string(runOut))
	} else {
		fmt.Print(string(runOut))
	}
}

func compileJavaCode(mountDir string, codeFileName string) {
	compileCMD := exec.Command("javac", codeFileName)
	compileCMD.Dir = mountDir
	out, err := compileCMD.CombinedOutput()
	if err != nil {
		fmt.Printf("**********ERROR COMPILING**********\n %v\n", string(out))
		panic(err)
	}
}

func getMainJavaClass(srcFolder string) string {
	files, err := ioutil.ReadDir(srcFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range files {
		if filepath.Ext(info.Name()) == ".class" {
			fileNameWithExtension := filepath.Base(info.Name())
			return extractMainClassFromClassFile(srcFolder, fileNameWithExtension)
		}
	}
	fmt.Printf("No runnable class files found in %s\n", srcFolder)
	return ""
}

func extractMainClassFromClassFile(srcFolder string, fileNameWithExtension string) string {
	classWithPath := filepath.Join(srcFolder, fileNameWithExtension)
	javapCommandString := "javap -public " + classWithPath + " | fgrep -q 'public static void main(java.lang.String[])'"
	javapCmd := exec.Command("bash", "-c", javapCommandString)
	_, err := javapCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error when trying to find runnable classfile: %s\n", err)
		return ""
	}
	return strings.TrimSuffix(fileNameWithExtension, filepath.Ext(fileNameWithExtension))
}

func getInputFileContents(mountDir string, inputFileName string) string {
	inputFile := filepath.Join(mountDir, inputFileName)
	if fileExists(inputFile) {
		content, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Print(err)
		}
		return string(content)
	}

	return ""
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
