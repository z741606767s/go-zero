### go-zero
    简介
    参考：https://github.com/Mikaelemmmm/go-zero-looklook
### tree
    项目根目录
    |-- README.md
    |-- app
    |-- common
    |-- data
    |-- deploy
    |-- docker-compose-demo.yml
    |-- docker-compose.yml
    |-- go.mod

docker-compose.yml 创建docker容器的配置文件，启动服务：
```shell
$ docker-compose up -d
```
deploy docker容器服务的配置文件目录。

api模版生成命令：
```shell
 $ goctl api go -api *.api -dir ../  -style=goZero
```
rpc模版生成命令:
```shell
$  goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../
$  sed -i "" 's/,omitempty//g' *.pb.go
```



