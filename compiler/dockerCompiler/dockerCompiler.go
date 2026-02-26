package dockerCompiler

import (
	"context"
	"io"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"

	"github.com/dockercodecompiler/compiler/configuration"
)

func CompileAndRun(language string, mountPath string) string {
	dockerClient := createDockerClient()
	backgoundContext := context.Background()

	containerConfig, containerHostConfig := buildConfigs(language, mountPath)

	createdContainer := createContainer(dockerClient, containerConfig, containerHostConfig, backgoundContext)

	startContainer(dockerClient, backgoundContext, createdContainer)

	waitForContainerToStopWithTimeout(dockerClient, backgoundContext, createdContainer)

	return getContainerLogs(dockerClient, backgoundContext, createdContainer)
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
	containerHostConfig *container.HostConfig, backgoundContext context.Context) container.CreateResponse {
	createdContainer, err := dockerClient.ContainerCreate(backgoundContext, containerConfig, containerHostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	return createdContainer
}

func startContainer(dockerClient client.APIClient, backgoundContext context.Context, createdContainer container.CreateResponse) {
	if err := dockerClient.ContainerStart(backgoundContext, createdContainer.ID, container.StartOptions{}); err != nil {
		panic(err)
	}
}

func waitForContainerToStopWithTimeout(dockerClient client.APIClient, backgoundContext context.Context, createdContainer container.CreateResponse) {
	statusCh, errCh := dockerClient.ContainerWait(backgoundContext, createdContainer.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}
}

func getContainerLogs(dockerClient client.APIClient, backgoundContext context.Context, createdContainer container.CreateResponse) string {
	containerLogs, err := dockerClient.ContainerLogs(backgoundContext, createdContainer.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return err.Error()
	}

	return parseDockerLogsToString(containerLogs)
}

func parseDockerLogsToString(containerLogs io.ReadCloser) string {
	defer containerLogs.Close()

	//read the first 8 bytes to ignore the HEADER part from docker container logs
	p := make([]byte, 8)
	containerLogs.Read(p)
	content, _ := io.ReadAll(containerLogs)

	return string(content)
}
