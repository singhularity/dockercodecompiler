package prepSandbox

import (
	"log"
	"os"
	"path/filepath"

	"../utils/fileUtils"
)

func PrepSandbox(codeData string, inputData string) string {
	mountPoint, payloadSource := createSandboxPayloadMountPoint()
	writePayloadFiles(mountPoint, payloadSource, codeData, inputData)
	return mountPoint
}

func createSandboxPayloadMountPoint() (mountPoint string, payloadSource string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	mountPoint = filepath.Join(path, "..", "temp", fileUtils.GetRandomFolderName(8))
	payloadSource = filepath.Join(path, "..", "payload")

	os.Mkdir(mountPoint, 0755)

	return
}

func writePayloadFiles(mountPoint string, payloadSource string, codeData string, inputData string) {
	fileUtils.WriteDataToFile(filepath.Join(mountPoint, "codeFile.java"), codeData, 0755)
	fileUtils.WriteDataToFile(filepath.Join(mountPoint, "inputFile"), inputData, 0755)

	fileUtils.CopyFilesInDirectory(payloadSource, mountPoint, 0755)
}
