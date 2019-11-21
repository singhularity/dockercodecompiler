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
	language := os.Args[1]
	if language == "java" {
		handleJavaCode()
	} else if language == "python" {
		handlePythonCode()
	} else {
		panic("Unsupported Language")
	}
}

func handleJavaCode() {
	compileJavaCode()
	runJavaCode()
}

func handlePythonCode() {
	runPythonCode()
}

func runPythonCode() {
	inputContent := getInputFileContents()
	runCMD := exec.Command("python3", "file.py", inputContent)
	runCMD.Dir = "/usercode"
	runOut, runErr := runCMD.CombinedOutput()
	if runErr != nil {
		fmt.Printf("**********ERROR**********\n %v\n", string(runOut))
	} else {
		fmt.Print(string(runOut))
	}
}

func compileJavaCode() {
	compileCMD := exec.Command("javac", "file.java")
	compileCMD.Dir = "/usercode"
	out, err := compileCMD.CombinedOutput()
	if err != nil {
		fmt.Printf("**********ERROR COMPILING**********\n %v\n", string(out))
		panic(err)
	}
}

func runJavaCode() {
	inputContent := getInputFileContents()
	mainClassToRun := getMainJavaClass("/usercode")
	if mainClassToRun != "" {
		runCMD := exec.Command("java", mainClassToRun, inputContent)
		runCMD.Dir = "/usercode"
		runOut, runErr := runCMD.CombinedOutput()
		if runErr != nil {
			fmt.Printf("**********ERROR**********\n %v\n", string(runOut))
		} else {
			fmt.Print(string(runOut))
		}
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
			return extractMainClass(srcFolder, fileNameWithExtension)
		}
	}
	fmt.Printf("No runnable class files found in %s\n", srcFolder)
	return ""
}

func extractMainClass(srcFolder string, fileNameWithExtension string) string {
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

func getInputFileContents() string {
	inputFile := "/usercode/inputFile"
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
