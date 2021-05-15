# conman

# Reverse-proxy + container manager


https://github.com/nathan-osman/i5/blob/master/dockmon/dockmon.go
- maintain a map of containers


https://github.com/sosedoff/docker-router/blob/master/monitor.go#L211
- Interesting solution to stopping idle containers


TODO: Editor that handles all cases (...open iframe on port, call api to recreate sandbox if response idle or non existant)

TODO: A user opens their repl which opens the IDE and attempts to connect to the backend evaluation server via a WebSocket.

TODO: A healthy, living conman gets the request. Conman notices that the request is for a container that is living on a different conman and proxies the request there.



```
We poll the container for published ports and the moment we see an open port we add a record to an etcd which stores the 
routing state. We then send a command to the client that we published a port, which will react by opening an iframe. 
Then the iframe or any request to the published url will hit our outer reverse proxy which will query etcd to find the 
container and if the container is alive we will send the traffic to the relevant container manager which has another 
reverse proxy which sends the traffic to the container.

If the container however is dead (from idling or because of an error) we revive via picking a container out of one 
the pools and going through the initialization phase described above.

```



***

1. Install then run etcd in separate terminal

```
brew install etcd
etcd
```

```
go get "github.com/coreos/etcd/clientv3"
```

