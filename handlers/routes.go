package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	docker "github.com/mehmetron/conman/docker"
	etcd "github.com/mehmetron/conman/etcd"
	"log"
	"net/http"
	"time"
)

type Env struct {
	docker docker.DockerEnv
	etcd   etcd.EtcdEnv
}

func Routes() {
	dockerEnv := docker.DockerEnv{}
	docker, err := docker.InitializeDocker()
	if err != nil {
		fmt.Println(err)
	}
	dockerEnv.Docker = docker

	etcdEnv := etcd.EtcdEnv{}
	cli, kv, err := etcd.InitializeEtcd()
	if err != nil {
		fmt.Println(err)
	}
	etcdEnv.Etcd = kv
	defer cli.Close()

	env := Env{dockerEnv, etcdEnv}

	go func() {

		router2 := http.NewServeMux()

		router2.HandleFunc("/", env.ReverseProxy)
		srv := &http.Server{
			Handler:      router2,
			Addr:         ":8001",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		log.Fatal(srv.ListenAndServe())
	}()

	router := mux.NewRouter()

	router.HandleFunc("/allinstances", env.AllInstances).Methods("GET")
	router.HandleFunc("/create", env.Create).Methods("POST")
	router.HandleFunc("/deleteall", env.DeleteAll).Methods("DELETE")
	router.HandleFunc("/delete", env.Delete).Methods("DELETE")
	router.HandleFunc("/whereto", env.WhereTo).Methods("GET")

	router.HandleFunc("/pause", env.Pause).Methods("PUT")
	router.HandleFunc("/unpause", env.UnPause).Methods("PUT")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
