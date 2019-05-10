# account
这条街最靓的仔，账户管理项目

项目目录说明

> account-serbice
>>config 放一些配置文件，如db、 grpc服务的信息、redis、服务发现连接等

>>proto 里面定义是用protobuf定义服务的微服务
>>>account 

>>service  具体实现protobuf的struct

>>model orm模型定义

>>lib自定义的库 





## config 里面的配置文件是json格式的

demo
```
项目不大，不需要弄读写分离
{
  "start": true          #开启mysql
  "ip":"126.123.76.73",  #连接ip
  "port":"3306",         #端口
  "user": "myaccount",   #连接账号
  "passwd": "123456",    #连接密码
  "db": "mydb"           #使用的数据
}

```