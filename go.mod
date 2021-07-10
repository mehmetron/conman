module github.com/mehmetron/conman

go 1.16

// https://github.com/etcd-io/etcd/issues/11563
replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/containerd/containerd v1.5.1 // indirect
	github.com/docker/docker v20.10.6+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/ermos/chalk v0.0.0-20201126145954-216cfaa8732f // indirect
	github.com/ermos/cli v0.0.2 // indirect
	github.com/ermos/gomon v0.0.0-20201127000642-87ef922adbe5 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/gddo v0.0.0-20210115222349-20d68f94ee1f // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	google.golang.org/grpc v1.37.1 // indirect
)
