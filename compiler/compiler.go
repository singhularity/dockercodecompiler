package compiler

import (
	"context"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func CompileAndRun(mountPath string) {
	imageName := "compiler_machine"

	dockerClient := createDockerClient()
	backgoundContext := context.Background()

	containerConfig, containerHostConfig := buildConfigs(imageName, mountPath)

	createdContainer := createContainer(imageName, dockerClient, containerConfig, containerHostConfig, backgoundContext)

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

func buildContainerConfig(imageName string) *container.Config {
	return &container.Config{
		Image: imageName,
		Entrypoint: []string{"go",
			"run", "/usercode/codehandler.go"},
	}
}

func buildContainerHostConfig(mountPath string) *container.HostConfig {
	return &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: mountPath,
				Target: "/usercode",
			},
		},
	}
}

func buildConfigs(imageName string, mountPath string) (containerConfig *container.Config, containerHostConfig *container.HostConfig) {
	containerConfig = buildContainerConfig(imageName)
	containerHostConfig = buildContainerHostConfig(mountPath)
	return
}

func createContainer(imageName string, dockerClient client.APIClient, containerConfig *container.Config,
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
	out, err := dockerClient.ContainerLogs(backgoundContext, createdContainer.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}
