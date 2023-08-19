# hellosvr
a project for learning k8s


## build & run local

```bash
go build .
./hellosvr -h
./hellosvr server
```

## build docker

modify variables in build.sh, then run the shell

```bash
./build_push_docker.sh
```

## test 

```bash
curl localhost:8080/api/v1/echo
```

## how to use in k8s

```bash
kubectl apply -f prod-namespace.yaml
kubectl apply -f hellosvr-configmap.yaml
kubectl apply -f hellosvr-deployment.yaml
kubectl apply -f hellosvr-service.yaml
```

