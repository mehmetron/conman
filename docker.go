package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

type Request struct {
	LangID int `json:"lang_id"`
}

// type ContainerOutput struct {
// 	ContainerID string `json:"output"`
// 	DemoHost    string `json:"demohost"`
// }

type Language struct {
	LangID int
	Name   string
	Image  string
}

// func Test(docker *client.Client) {
// 	ctx := context.Background()
// 	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	reader, err := docker.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	io.Copy(os.Stdout, reader)

// 	resp, err := docker.ContainerCreate(ctx, &container.Config{
// 		Image: "alpine",
// 		Cmd:   []string{"sleep", "5"},
// 		Tty:   false,
// 	}, nil, nil, nil, "")
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := docker.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
// 		panic(err)
// 	}

// 	statusCh, errCh := docker.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
// 	select {
// 	case err := <-errCh:
// 		if err != nil {
// 			panic(err)
// 		}
// 	case <-statusCh:
// 	}

// 	out, err := docker.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
// 	if err != nil {
// 		panic(err)
// 	}

// 	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
// }

// Create container
func Create(env *Env, demoPort string, apiPort string, langID int) string {

	n := make(map[int]Language)
	n[1] = Language{LangID: 1, Name: "python", Image: "1c4e16e817b4"}
	n[2] = Language{LangID: 2, Name: "go", Image: ""}
	n[3] = Language{LangID: 3, Name: "javascript", Image: ""}
	n[4] = Language{LangID: 4, Name: "java", Image: ""}

	image := n[langID].Image

	// ctx := context.Background()
	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	// check(err)

	// memoryLimit := int64(30 * 1024 * 1024)
	// resourceConfig := container.Resources{Memory: memoryLimit}

	var ApiContainer nat.Port = "8090/tcp"
	// ApiHost := "8000"

	var DemoContainer nat.Port = "8000/tcp"
	// DemoHost := "8090"

	config := &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			DemoContainer: struct{}{},
			ApiContainer:  struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			DemoContainer: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: demoPort,
				},
			},
			ApiContainer: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: apiPort,
				},
			},
		}, //Resources: resourceConfig,
	}

	resp, err := env.docker.ContainerCreate(context.TODO(), config, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	err = env.docker.ContainerStart(context.TODO(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	// resp, err := cli.ContainerCreate(ctx, &container.Config{
	// 	Image: image,
	// 	Tty:   true,
	// }, &container.HostConfig{Resources: resourceConfig}, nil, nil, "")
	// check(err)

	// if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
	// 	panic(err)
	// }

	// demohosturl := "subdomain123"
	// apihosturl just appends "galatatower" to subdomain of demohosturl so no need to send it

	// return ContainerOutput{
	// 	ContainerID: resp.ID,
	// 	DemoHost:    demohosturl,
	// }
	return resp.ID
}

// Remove container
func DeleteContainer(env *Env, containerID string) (string, error) {

	err := env.docker.ContainerStop(context.TODO(), containerID, nil)
	if err != nil {
		fmt.Println(err)
		return "Failed to remove container", err
	}

	err = env.docker.ContainerRemove(context.TODO(), containerID, types.ContainerRemoveOptions{})
	check(err)

	return "Success removing container!", nil
}
