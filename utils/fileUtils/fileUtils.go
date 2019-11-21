package fileUtils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func CopyFilesInDirectory(srcFolder string, destFolder string, mode os.FileMode) {

	err := filepath.Walk(srcFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			copyFile(path, filepath.Join(destFolder, info.Name()), mode)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func copyFile(srcFilePath string, destFilePath string, mode os.FileMode) (status bool) {
	content, err := ioutil.ReadFile(srcFilePath)
	if err != nil {
		fmt.Print(err)
		return false
	}

	return WriteDataToFile(destFilePath, string(content), mode)
}

func WriteDataToFile(path string, data string, mode os.FileMode) bool {
	os.Create(path)
	err := ioutil.WriteFile(path, []byte(data), mode)
	if err != nil {
		fmt.Print(err)
		return false
	}
	return true
}

func GetRandomFolderName(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
