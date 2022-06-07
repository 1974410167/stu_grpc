# stu_grpc

1. 学习Grpc
2. 学习Nacos(代码主要在main.go)

### 实现两个服务的互相调用

一个服务A内部调用另一个服务B,然后客户端C来请求A。

#### 此demo的实现

1. user_server.go提供了GetUserMessage和AddUserMessage两个方法
2. manage_server.go也提供了GetUserMessage方法
3. manage_server.go的GetUserMessage内部调用user_server.go的GetUserMessage。(具体代码在https://github.com/1974410167/stu_grpc/blob/main/service/manage.go)
4. 编写manager_client.go来调用manage_server.go的GetUserMessage

### Nacos安装与基本操作
1. 仿照教程https://nacos.io/zh-cn/docs/quick-start.html  下载解压
2. bash startup.sh -m standalone 运行nacos
3. http://127.0.0.1:8848/nacos 可登陆可视化界面，账号密码均为nacos


- 服务注册

POST http://127.0.0.1:8848/nacos/v1/ns/instance?serviceName=hello_server&ip=1.1.1.1&port=8081
- 服务发现

GET http://127.0.0.1:8848/nacos/v1/ns/instance/list?serviceName=hello_server

- 发布配置

POST http://127.0.0.1:8848/nacos/v1/cs/configs?dataId=nacos.cfg.dataId&group=test&content=HelloWorld

- 获取配置

GET http://127.0.0.1:8848/nacos/v1/cs/configs?dataId=nacos.cfg.dataId&group=test


- 切换数据库为mysql

1. 创建nacos_config数据库
2. 执行{nacos_path}/conf/nacos-mysql.sql文件，即在数据库中执行
> source {nacos_path}/conf/nacos-mysql.sql 命令
3. 改{nacos_path}/conf/application.properties配置文件，如下

        ### If use MySQL as datasource:
        spring.datasource.platform=mysql
        
        ### Count of DB:
        db.num=1
        
        ### Connect URL of DB:
        db.url.0=jdbc:mysql://gz-cynosdbmysql-grp-0ku3m9v9.sql.tencentcdb.com:29978/nacos_config?characterEncoding=utf8&connectTimeout=1000&socketTimeout=3000&autoReconnect=true&useUnicode=true&useSSL=false&serverTimezone=UTC
        db.user.0=root
        db.password.0=wsghy.5637

