# account
这条街最靓的仔，账户管理项目

## 项目目录说明（自定义框架目录）

> account-serbice

>>config 放一些配置文件，如db、 grpc服务的信息、redis、服务发现连接等
>>> tls tls/ssl加密

>>proto 里面定义是用protobuf定义服务的微服务
>>>account 

>>service  protobuf服务的具体实现

>>model orm模型定义

>>lib自定义的库,或者系统

>> client 测试用的客户端存放处，里面存放的是golang的实现，别的语言的自己想办法

>> gateway 为grpc网关服务器的反向代理存放处
>>>config grpc网关的服务器配置

>>client 测试

>> rpc grpc服务存放处



## config 文件配置说明

config文件夹中必须拥有一下配置
- db.json
- redis.json
- service.json


### db.json配置说明

```
{
  "start": true          #开启mysql
  "ip":"126.123.76.73",  #连接ip
  "port":"3306",         #端口
  "user": "myaccount",   #连接账号
  "passwd": "123456",    #连接密码
  "db": "mydb"           #使用的数据
}

```

### redis.json配置说明

```
{
  "start": true,         #是否开启redis
  "ip": "126.123.76.73", #redis连接地址
  "port": "6379",        #端口
  "db": 0,                #选择的数据库
  "passwd":"123456"     #密码
}
```


### grpc_service.json grpc服务器信息
```
{
  "ip": "localhost",  #服务所在ip
  "port": "50051",   #服务所在端口
  "describe": "",   #服务描述
  "name": "zldz.account",  #服务发现的名字
  "checkPort":"8080",   #服务发现心跳检测的端口
  "tags":"zldz account ly 最靓的仔 账户管理" #服务发现的tag，可用来过滤查询,中文会报错，不知道怎么解决
}
```


### consul 服务发现连接信息配置

```
{
  "start": true,     #服务发现是否开启
  "ip": "localhost", #连接服务发现的ip
  "port": "8500"     #连接服务发现的端口
}

```

### grpc_gateway_service.json grpc网关服务器信息
```
{
  "start": true,           #是否开启了grpc网关
  "gateway_port": "50051"  #grpc网关服务器端口
}
```