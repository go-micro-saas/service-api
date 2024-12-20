# saas service docs

### 端口划分原则:

1. 服务端口号用5位数字表示,避免与常用组件端口号冲突,最高不超65535
2. 服务编号用2位数表示,最后1位用于区分http和grpc,1表示http协议端口,2表示grpc协议端口
3. 错误码用9位数字表示,其中前3位与服务序号保持一致,第4为使用0作为分隔占位,后5位用于表示具体错误码的枚举值. (注意:
   错误码枚举值在error.proto有唯一约束,但在业务和代码中无意义,不应对外暴露)

| 服务              | 服务编号 | http惯用端口号 | grpc惯用端口号 | 错误码                 |
|:----------------|:-----|:----------|:----------|:--------------------|
| service-layout  | 99   | 9991      | 9992      | 99990001 : UNKNOWN  |
| dtm-manager     | 100  | 10001     | 10002     | 100000001 : UNKNOWN |
| ping-service    | 101  | 10101     | 10102     | 101000001 : UNKNOWN |
| nodeid-service  | 102  | 10201     | 10202     | 102000001 : UNKNOWN |
| account-service | 103  | 10301     | 10302     | 103000001 : UNKNOWN |
| admin-service   | 104  | 10401     | 10402     | 104000001 : UNKNOWN |

## testdata

* [testdata/configuration](./testdata/configuration/README.md)

## 安装依赖

请先安装必要的依赖：`make init`

### 查看帮助

查看帮助: `make help`

```text
Targets:
help                   show help
init                   init and install necessary software
echo                   echo test content
generate               generate : go:generate
protoc-api-protobuf    protoc :-->: generate api protobuf
protoc-specified-api   protoc :-->: example: make protoc-specified-api service=ping-service
protoc-ping-protobuf   protoc :-->: generate ping protobuf
```

### 编译`protobuf`文件

```shell

# 编译`protobuf`文件；示例如下：
make protoc-api-protobuf
make protoc-ping-protobuf
make protoc-ping-v1-protobuf

```

## 用户服务 account-service

初始化账户：

- 账户： user@user.user
- 密码： md5(User.123456) = b1c74a97bc4fbad404b16e193ffc3275

## 后台服务 admin-service

初始化账户：

- 账户： admin@admin.admin
- 密码： md5(Admin.123456) = b1c74a97bc4fbad404b16e193ffc3275
