package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Env struct {
	docker DockerEnv
	etcd   EtcdEnv
}

func main() {
	fmt.Println("working...")

	dockerEnv := DockerEnv{}
	docker, err := InitializeDocker()
	if err != nil {
		fmt.Println(err)
	}
	dockerEnv.docker = docker

	etcdEnv := EtcdEnv{}
	cli, kv, err := InitializeEtcd()
	if err != nil {
		fmt.Println(err)
	}
	etcdEnv.etcd = kv
	defer cli.Close()

	env := Env{dockerEnv, etcdEnv}

	go func() {

		router2 := http.NewServeMux()

		router2.HandleFunc("/", env.ReverseProxy)
		srv := &http.Server{
			Handler:      router2,
			Addr:         ":8081",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		log.Fatal(srv.ListenAndServe())
		//log.Fatal(http.ListenAndServe(":8081", router2))
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
		Handler: router,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
