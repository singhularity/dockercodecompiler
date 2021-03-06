package staging

import (
	"os"
	"path/filepath"

	"github.com/dockercodecompiler/compiler/configuration"
	"github.com/dockercodecompiler/compiler/utils/fileUtils"
	"github.com/dockercodecompiler/compiler/sandbox/staging/sanitizers"
)

func PrepSandbox(language string, codeData string, inputData string) string {
	appConfig := configuration.GetConfig()
	mountPoint, payloadSource := createSandboxPayloadMountPoint(appConfig)
	writePayloadFiles(appConfig, language, mountPoint, payloadSource, codeData, inputData)
	return mountPoint
}

func createSandboxPayloadMountPoint(appConfig configuration.AppConfiguration) (mountPoint string, payloadSource string) {
	sandBoxLocation := filepath.Join(fileUtils.GetMainDir(), appConfig.SandBoxLocation)
	mountPoint = filepath.Join(sandBoxLocation, appConfig.TempFileName, fileUtils.GetRandomFolderName(8))
	payloadSource = filepath.Join(sandBoxLocation, payloadSource)
	os.Mkdir(mountPoint, 0755)

	return
}

func writePayloadFiles(appConfig configuration.AppConfiguration,
	language string, mountPoint string, payloadSource string,
	codeData string, inputData string) (codeFileName string) {
	codeFileName = appConfig.CodeFileName + appConfig.LanguageExtensions[language].Extension
	codeFileWithPath := filepath.Join(mountPoint, codeFileName)
	inputFileWithPath := filepath.Join(mountPoint, appConfig.InputFileName)
	codeData = sanitizers.Sanitize(language, codeData)
	fileUtils.WriteDataToFile(codeFileWithPath, codeData, 0755)
	fileUtils.WriteDataToFile(inputFileWithPath, inputData, 0755)

	fileUtils.CopyFilesInDirectory(payloadSource, mountPoint, 0755)
	return
}
