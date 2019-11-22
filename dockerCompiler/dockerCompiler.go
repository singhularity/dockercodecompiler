package dockerCompiler

import (
	"context"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"

	"dockercodecompiler/utils/configuration"
)

func CompileAndRun(language string, mountPath string) {
	dockerClient := createDockerClient()
	backgoundContext := context.Background()

	containerConfig, containerHostConfig := buildConfigs(language, mountPath)

	createdContainer := createContainer(dockerClient, containerConfig, containerHostConfig, backgoundContext)

	startContainer(dockerClient, backgoundContext, createdContainer)

	waitForContainerToStopWithTimeout(dockerClient, backgoundContext, createdContainer)

	printContainerLogs(dockerClient, backgoundContext, createdContainer)
}

func createDockerClient() client.APIClient {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return dockerClient
}

func buildConfigs(language string, mountPath string) (containerConfig *container.Config,
	containerHostConfig *container.HostConfig) {
	appConfig := configuration.GetConfig()
	containerConfig = buildContainerConfig(language, appConfig)
	containerHostConfig = buildContainerHostConfig(mountPath, appConfig)
	return
}

func buildContainerConfig(language string, appConfig configuration.AppConfiguration) *container.Config {
	languageSpecs := appConfig.LanguageExtensions[language]
	return &container.Config{
		Image: appConfig.ImageName,
		Entrypoint: append(appConfig.SandBoxRunParams, languageSpecs.ExecCommand,
			appConfig.CodeFileName+languageSpecs.Extension, appConfig.MountPoint, appConfig.InputFileName),
	}
}

func buildContainerHostConfig(mountPath string, appConfig configuration.AppConfiguration) *container.HostConfig {
	return &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: mountPath,
				Target: appConfig.MountPoint,
			},
		},
	}
}

func createContainer(dockerClient client.APIClient, containerConfig *container.Config,
	containerHostConfig *container.HostConfig, backgoundContext context.Context) container.ContainerCreateCreatedBody {
	createdContainer, err := dockerClient.ContainerCreate(backgoundContext, containerConfig, containerHostConfig, nil, "")
	if err != nil {
		panic(err)
	}

	return createdContainer
}

func startContainer(dockerClient client.APIClient, backgoundContext context.Context, createdContainer container.ContainerCreateCreatedBody) {
	if err := dockerClient.ContainerStart(backgoundContext, createdContainer.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}

func waitForContainerToStopWithTimeout(dockerClient client.APIClient, backgoundContext context.Context, createdContainer container.ContainerCreateCreatedBody) {
	statusCh, errCh := dockerClient.ContainerWait(backgoundContext, createdContainer.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}
}

func printContainerLogs(dockerClient client.APIClient, backgoundContext context.Context, createdContainer container.ContainerCreateCreatedBody) {
	out, err := dockerClient.ContainerLogs(backgoundContext, createdContainer.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}
