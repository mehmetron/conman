package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type DeleteSandbox struct {
	ContainerID string `json:"container_id"`
	FromHost    string `json:"from_host"`
}

type CreateSandbox struct {
	LangID int `json:"lang_id"`
}

type redirectHostStruct struct {
	FromHost string `json:"fromHost"`
	ToHost   string `json:"toHost"`
	LangID   int    `json:"lang_id"`
}

type ListOfetcd struct {
	Prefix string               `json:"prefix"`
	Data   []redirectHostStruct `json:"data"`
}

// WhereTo tells where the subdomain goes to? || What is the value of this key?
func (env *Env) WhereTo(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query()["q"][0]

	b, err := env.etcd.GetKV(key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("length of ", len(b.Kvs))
	if len(b.Kvs) == 0 {
		_, err = fmt.Fprintf(w, "%s doesn't exist", key)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		found := string(b.Kvs[0].Value)
		fmt.Println("keys", found)

		_, err = fmt.Fprintf(w, "Found World! %s", found)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Create creates key:value pair
func (env *Env) Create(w http.ResponseWriter, r *http.Request) {

	//// Serialize request body
	//var toCreate CreateSandbox
	//err := json.NewDecoder(r.Body).Decode(&toCreate)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// Generate subdomains
	rand.Seed(time.Now().UnixNano())
	demoSubdomain := RandStringRunes(10)
	apiSubdomain := fmt.Sprintf("%spid1", demoSubdomain)
	fmt.Println("gen subdomains", demoSubdomain, apiSubdomain)

	// Find ports
	port1, port2 := GeneratePorts()
	demoPort := strconv.Itoa(port1)
	apiPort := strconv.Itoa(port2)

	//fmt.Println("demoSubdomain, apiSubdomain, demoPort, apiPort, toCreate.LangID", demoSubdomain, apiSubdomain, demoPort, apiPort, toCreate.LangID)
	fmt.Println("demoSubdomain, apiSubdomain, demoPort, apiPort", demoSubdomain, apiSubdomain, demoPort, apiPort)

	// Put demoPort and apiPort into etcd
	putDemoResp, err := env.etcd.CreateKVs(demoSubdomain, demoPort)
	//putDemoResp, err := env.etcd.CreateKVs(demoSubdomain, "http://localhost:9003")
	fmt.Println("CreateKVs putDemoResp: ", putDemoResp)
	if err != nil {
		fmt.Println(err)
	}

	putApiResp, err := env.etcd.CreateKVs(apiSubdomain, apiPort)
	fmt.Println("CreateKVs putApiResp: ", putApiResp)
	if err != nil {
		fmt.Println(err)
	}

	// Create docker container
	res := env.docker.CreateContainer(demoPort, apiPort)
	fmt.Println("res", res)

	_, err = fmt.Fprintf(w, "demoSubdomain: %s \n apiSubdomain: %s \n res: %s", demoSubdomain, apiSubdomain, res)
	if err != nil {
		fmt.Println(err)
	}

}

func (env *Env) DestroyAll(w http.ResponseWriter, r *http.Request) {

	delKVResp, err := env.etcd.DeleteAllKV()
	fmt.Println("DeleteKV delKVResp: ", delKVResp)
	if err != nil {
		fmt.Println(err)
	}

	_, err = fmt.Fprintf(w, "Destroy World!")
	if err != nil {
		fmt.Println(err)
	}
}

// Destroy deletes a key:value pair
func (env *Env) Destroy(w http.ResponseWriter, r *http.Request) {

	var p DeleteSandbox
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Println(err)
	}

	delKVResp, err := env.etcd.DeleteKV(p.FromHost)
	fmt.Println("DeleteKV delKVResp: ", delKVResp)
	if err != nil {
		fmt.Println(err)
	}

	delContainerResp, err := env.docker.DeleteContainer(p.ContainerID)
	print("DeleteContainer delContainerResp: ", delContainerResp)
	if err != nil {
		fmt.Println(err)
	}

	_, err = fmt.Fprintf(w, "Destroy World!")
	if err != nil {
		fmt.Println(err)
	}
}

// Puase pauses a docker container
func (env *Env) Pause(w http.ResponseWriter, r *http.Request) {

	containerID := "fakeid"

	env.docker.PauseContainer(containerID)
}

// UnPuase unpauses a docker container
func (env *Env) UnPause(w http.ResponseWriter, r *http.Request) {

	containerID := "fakeid"

	env.docker.UnPauseContainer(containerID)
}

// AllInstances gets all instances
func (env *Env) AllInstances(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)
	prefix := ""

	list := ListOfetcd{
		Prefix: prefix,
		Data:   []redirectHostStruct{},
	}

	getResp, err := env.etcd.GetKVPrefixed(prefix)
	if err != nil {
		fmt.Println(err)
	}

	for _, ev := range getResp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)

		penis := redirectHostStruct{
			FromHost: string(ev.Key),
			ToHost:   string(ev.Value),
		}
		list.Data = append(list.Data, penis)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Fprintf(w, "List World!", list)
}
