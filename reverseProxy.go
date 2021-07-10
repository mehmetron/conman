package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// ReverseProxy proxies requests
func (env *Env) ReverseProxy(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request Info: ", r.Host, r.Method, r.Proto)

	subdomain := strings.Split(r.Host, ".")[0]
	b, err := env.etcd.etcd.Get(context.TODO(), subdomain)
	if err != nil {
		fmt.Println(err)
	}

	var rawUrl string

	if len(b.Kvs) == 0 {
		fmt.Println("Subdomain key not in etcd", b.Kvs)
		rawUrl = env.createContainer()
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

func (env *Env) createContainer() string {
	fmt.Println("Creating container and storing in etcd then return url")

	return fmt.Sprintf("%s", "urlOfNewContainer")
}
