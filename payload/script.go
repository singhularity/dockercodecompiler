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
	var inputContent string

	inputFile := "/usercode/input"
	if fileExists(inputFile) {
		content, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Print(err)
		}
		inputContent = string(content)
	}

	compileCMD := exec.Command("javac", "file.java")
	compileCMD.Dir = "/usercode"
	out, err := compileCMD.CombinedOutput()
	if err == nil {
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
		} else {
			fmt.Printf("**********ERROR COMPILING**********\n %v\n", string(out))
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getMainJavaClass(srcFolder string) string {

	files, err := ioutil.ReadDir(srcFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range files {
		if filepath.Ext(info.Name()) == ".class" {
			classNameWithExtension := filepath.Base(info.Name())
			classWithPath := filepath.Join(srcFolder, classNameWithExtension)
			javapCommandString := "javap -public " + classWithPath + " | fgrep -q 'public static void main(java.lang.String[])'"
			javapCmd := exec.Command("bash", "-c", javapCommandString)
			_, err := javapCmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Error when trying to find runnable classfile: %s\n", err)
				return ""
			}
			return strings.TrimSuffix(classNameWithExtension, filepath.Ext(classNameWithExtension))
		}
	}
	fmt.Printf("No runnable class files found in %s\n", srcFolder)
	return ""
}
