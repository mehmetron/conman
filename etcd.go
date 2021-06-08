package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type EtcdEnv struct {
	etcd clientv3.KV
}

var (
	dialTimeout = 5 * time.Second
)

func InitializeEtcd() (*clientv3.Client, clientv3.KV) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: dialTimeout,
	})
	if err != nil {
		fmt.Println(err)
	}

	kv := clientv3.NewKV(cli)

	return cli, kv
}

// Delete all key:values with _subdomain_ as prefix
func (env *EtcdEnv) DeleteKV(subdomain string) (*clientv3.DeleteResponse, error) {
	delResp, err := env.etcd.Delete(context.TODO(), subdomain, clientv3.WithPrefix())
	return delResp, err
}

func (env *EtcdEnv) CreateKVs(subdomain, port string) (*clientv3.PutResponse, error) {
	putResp, err := env.etcd.Put(context.TODO(), subdomain, fmt.Sprintf("http://localhost:%s", port))
	return putResp, err
}

func (env *EtcdEnv) GetKV(subdomain string) (*clientv3.GetResponse, error) {
	getResp, err := env.etcd.Get(context.TODO(), subdomain)
	return getResp, err
}

func (env *EtcdEnv) GetKVPrefixed(prefix string) (*clientv3.GetResponse, error) {
	getResp, err := env.etcd.Get(context.TODO(), prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	return getResp, err
}
