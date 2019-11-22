package sandBoxPrepper

import (
	"os"
	"path/filepath"

	"dockercodecompiler/utils/fileUtils"
)

func PrepSandbox(basePath string, language string, codeData string, inputData string) string {
	mountPoint, payloadSource := createSandboxPayloadMountPoint(basePath)
	writePayloadFiles(language, mountPoint, payloadSource, codeData, inputData)
	return mountPoint
}

func createSandboxPayloadMountPoint(basePath string) (mountPoint string, payloadSource string) {
	mountPoint = filepath.Join(basePath, "temp", fileUtils.GetRandomFolderName(8))
	payloadSource = filepath.Join(basePath, "payload")

	os.Mkdir(mountPoint, 0755)

	return
}

func writePayloadFiles(language string, mountPoint string, payloadSource string, codeData string, inputData string) {
	fileUtils.WriteDataToFile(filepath.Join(mountPoint, "file."+fileUtils.GetExtensionForLanguage(language)), codeData, 0755)
	fileUtils.WriteDataToFile(filepath.Join(mountPoint, "inputFile"), inputData, 0755)

	fileUtils.CopyFilesInDirectory(payloadSource, mountPoint, 0755)
}
