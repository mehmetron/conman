package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ReverseProxy proxies requests
func (env *Env) ReverseProxy(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request Info: ", r.Host, r.Method, r.Proto)

	subdomain := strings.Split(r.Host, ".")[0]
	fmt.Println("21 subdomain ", subdomain)
	//subdomain = strings.Replace(subdomain, "pid1", "", 1)
	//fmt.Println("23 subdomain ", subdomain)
	b, err := env.etcd.etcd.Get(context.TODO(), subdomain)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("26 ", b)

	var rawUrl string

	if len(b.Kvs) == 0 {
		fmt.Println("Subdomain key not in etcd", b.Kvs)
		rawUrl = env.CreateContainer()
	} else {
		rawUrl = string(b.Kvs[0].Value)
	}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedUrl)
	proxy.ServeHTTP(w, r)

}

func (env *Env) CreateContainer() string {
	fmt.Println("Creating container and storing in etcd then return url")

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
	putDemoResp, err := env.etcd.CreateKV(demoSubdomain, demoPort)
	//putDemoResp, err := env.etcd.CreateKV(demoSubdomain, "http://localhost:9003")
	fmt.Println("CreateKV putDemoResp: ", putDemoResp)
	if err != nil {
		fmt.Println(err)
	}

	putApiResp, err := env.etcd.CreateKV(apiSubdomain, apiPort)
	fmt.Println("CreateKV putApiResp: ", putApiResp)
	if err != nil {
		fmt.Println(err)
	}

	// Create docker container
	res, err := env.docker.CreateContainer(demoPort, apiPort)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("res", res)

	//_, err = fmt.Fprintf(w, "demoSubdomain: %s \n apiSubdomain: %s \n res: %s", demoSubdomain, apiSubdomain, res)
	//if err != nil {
	//	fmt.Println(err)
	//}

	return demoSubdomain
}
