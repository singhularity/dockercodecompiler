package configuration

import (
	"dockercodecompiler/compiler/utils/fileUtils"
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type AppConfiguration struct {
	ImageName          string
	MountPoint         string
	Languages          []string
	SandBoxRunParams   []string
	SandBoxLocation    string
	PayloadLocation    string
	TempFileName       string
	CodeFileName       string
	InputFileName      string
	LanguageExtensions map[string]LanguageSpecs
}

type LanguageSpecs struct {
	Extension   string
	ExecCommand string
}

var config AppConfiguration

func init() {
	configLocation := fileUtils.GetConfigLocation()
	if _, err := toml.DecodeFile(filepath.Join(configLocation, "config.toml"), &config); err != nil {
		fmt.Println(err)
		return
	}
}

func GetConfig() AppConfiguration {
	return config
}
