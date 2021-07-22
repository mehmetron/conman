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

func InitializeEtcd() (*clientv3.Client, clientv3.KV, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: dialTimeout,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize etcd: %v", err)
	}

	kv := clientv3.NewKV(cli)

	return cli, kv, nil
}

// DeleteAllKV Deletes all key:values
func (env *EtcdEnv) DeleteAllKV() (*clientv3.DeleteResponse, error) {
	delResp, err := env.etcd.Delete(context.TODO(), "", clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("failed to delete all kvs: %v", err)
	}
	return delResp, nil
}

// DeleteKV Deletes all key:values with _subdomain_ as prefix
func (env *EtcdEnv) DeleteKV(subdomain string) (*clientv3.DeleteResponse, error) {
	delResp, err := env.etcd.Delete(context.TODO(), subdomain, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("failed to delete kv: %v", err)
	}
	return delResp, nil
}

// CreateKV creates key:value pair
func (env *EtcdEnv) CreateKV(subdomain, port string) (*clientv3.PutResponse, error) {
	putResp, err := env.etcd.Put(context.TODO(), subdomain, fmt.Sprintf("http://localhost:%s", port))
	if err != nil {
		return nil, fmt.Errorf("failed to create kv: %v", err)
	}
	return putResp, nil
}

// GetKV gets value with passed in key
func (env *EtcdEnv) GetKV(subdomain string) (*clientv3.GetResponse, error) {
	getResp, err := env.etcd.Get(context.TODO(), subdomain)
	if err != nil {
		return nil, fmt.Errorf("failed to get kv: %v", err)
	}
	return getResp, nil
}

// GetKVPrefixed gets values with passed in prefix
func (env *EtcdEnv) GetKVPrefixed(prefix string) (*clientv3.GetResponse, error) {
	getResp, err := env.etcd.Get(context.TODO(), prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, fmt.Errorf("failed to get kv with prefix: %v", err)
	}
	return getResp, nil
}
