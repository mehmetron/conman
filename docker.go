package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	//"github.com/docker/docker/client"
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

type DockerEnv struct {
	docker *client.Client
}

func InitializeDocker() (*client.Client, error) {
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return docker, err
}

// CreateContainer creates a docker container
func (env *DockerEnv) CreateContainer(demoPort string, apiPort string) string {

	//n := make(map[int]Language)
	//n[1] = Language{LangID: 1, Name: "python", Image: "nginx"}
	//n[2] = Language{LangID: 2, Name: "go", Image: ""}
	//n[3] = Language{LangID: 3, Name: "javascript", Image: ""}
	//n[4] = Language{LangID: 4, Name: "java", Image: ""}
	//
	//image := n[langID].Image
	image := "bob"

	// memoryLimit := int64(30 * 1024 * 1024)
	// resourceConfig := container.Resources{Memory: memoryLimit}

	var ApiContainer nat.Port = "8080/tcp"
	// ApiHost := "8000"

	var DemoContainer nat.Port = "8090/tcp"
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
		fmt.Printf("89 error %s", err)
		panic(err)
	}

	err = env.docker.ContainerStart(context.TODO(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	// return ContainerOutput{
	// 	ContainerID: resp.ID,
	// 	DemoHost:    demohosturl,
	// }
	return resp.ID
}

// DeleteContainer deletes a docker container
func (env *DockerEnv) DeleteContainer(containerID string) (string, error) {

	err := env.docker.ContainerStop(context.TODO(), containerID, nil)
	if err != nil {
		fmt.Println(err)
		return "Failed to remove container", err
	}

	err = env.docker.ContainerRemove(context.TODO(), containerID, types.ContainerRemoveOptions{})
	if err != nil {
		fmt.Println(err)
	}

	return "Success removing container!", nil
}

// DeleteAllContainers deletes all docker containers
func (env *DockerEnv) DeleteAllContainers() {

	containers, err := env.docker.ContainerList(context.TODO(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Print("Stopping container ", container.ID[:10], "... ")
		env.DeleteContainer(container.ID)
		fmt.Println("Success")
	}

}

// PauseContainer pauses a docker container
func (env *DockerEnv) PauseContainer(containerID string) {

	err := env.docker.ContainerPause(context.TODO(), containerID)
	if err != nil {
		fmt.Println(err)
	}
}

// UnPauseContainer unpauses a docker container
func (env *DockerEnv) UnPauseContainer(containerID string) {

	err := env.docker.ContainerUnpause(context.TODO(), containerID)
	if err != nil {
		fmt.Println(err)
	}
}
