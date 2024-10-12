# 存储配置文件到 consul

```shell

# ===== testing =====
# cluster service api
go run cmd/store-configuration/main.go \
  -conf ./testdata/configuration/consul \
  -source_dir ./testdata/configuration/testing/cluster_service_api \
  -store_dir go-micro-saas/general-configs/production
  
# general config
go run cmd/store-configuration/main.go \
  -conf ./testdata/configuration/consul \
  -source_dir ./testdata/configuration/testing/general-configs \
  -store_dir go-micro-saas/general-configs/production

# ping-service config
go run cmd/store-configuration/main.go \
  -conf ./testdata/configuration/consul \
  -source_dir ./testdata/configuration/testing/ping-service \
  -store_dir go-micro-saas/ping-service/production/latest

# nodeid-service config
go run cmd/store-configuration/main.go \
  -conf ./testdata/configuration/consul \
  -source_dir ./testdata/configuration/testing/nodeid-service \
  -store_dir go-micro-saas/nodeid-service/production/latest

```

## testing

**hosts**

```shell
vim /etc/hosts
```

```text
# develop
192.168.10.19 my-service-hostname
192.168.10.19 my-consul-hostname
192.168.10.19 my-etcd-hostname
192.168.10.19 my-jaeger-hostname
192.168.10.19 my-mysql-hostname
192.168.10.19 my-postgres-hostname
192.168.10.19 my-rabbitmq-hostname
192.168.10.19 my-redis-hostname
192.168.10.19 my-mongo-hostname
192.168.10.19 my-mongo1
192.168.10.19 my-mongo2
192.168.10.19 my-mongo3
```

**ping**

```shell

# export http_proxy="" 
# http_proxy="" curl http://my-service-hostname:20101/api/v1/ping/pong
curl http://my-service-hostname:20101/api/v1/ping/pong
curl http://my-service-hostname:20201/api/v1/nodeid/ping/pong
curl http://my-service-hostname:20201/api/v1/nodeid/get-service-info

```
