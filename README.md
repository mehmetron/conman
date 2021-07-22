# conman

# Reverse-proxy + container manager

***
## Documentation
Conman gets request and either creates new pid1 image or 
proxy's the request to an existing container based on data in etcd.

Idea for stopping idle containers
https://github.com/sosedoff/docker-router/blob/master/monitor.go#L211

Idea for keeping pool of running containers
https://github.com/nathan-osman/i5/blob/master/dockmon/dockmon.go

### How replit does it
"We poll the container for published ports and the moment we see an open port we add a record to an etcd which stores the
routing state. We then send a command to the client that we published a port, which will react by opening an iframe.
Then the iframe or any request to the published url will hit our outer reverse proxy which will query etcd to find the
container and if the container is alive we will send the traffic to the relevant container manager which has another
reverse proxy which sends the traffic to the container."

```bash
# Install etcd
brew install etcd
# Run etcd
etcd
```

```
go get "github.com/coreos/etcd/clientv3"
```