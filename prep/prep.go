package prep

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"../utils/fileUtils"
)

func Prepcode(codeData string, inputData string) string {
	// create a TestDir directory on current working directory
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	mountPath := filepath.Join(path, "..", "temp", getRandomFolderName(8))
	payloadSource := filepath.Join(path, "..", "payload")

	os.Mkdir(mountPath, 0755)

	fileUtils.WriteDataToFile(filepath.Join(mountPath, "file.java"), codeData, 0755)
	fileUtils.WriteDataToFile(filepath.Join(mountPath, "input"), inputData, 0755)

	fileUtils.CopyFilesInDirectory(payloadSource, mountPath, 0755)

	return mountPath
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func getRandomFolderName(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
