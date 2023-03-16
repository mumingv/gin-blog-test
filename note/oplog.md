# 操作日志

## 项目初始化

生成go.mod

```azure
go mod init "github.com/mumingv/gin-blog"
```

创建目录和文件

```azure
mkdir conf controller dao logger models routers settings static templates util
touch conf/config.yaml
touch main.go
```

## 配置文件定义、读取、变更监控

配置文件&对应数据结构定义

```azure
vim config.yaml
touch settings/settings.go settings/settings_test.go
```

引入第三方库viper读取配置文件

```azure
go get github.com/spf13/viper
```

引入第三方库fsnotify监控配置文件变更

```azure
go get github.com/fsnotify/fsnotify
```

引入第三方库gin框架处理Http请求

```azure
go get github.com/gin-gonic/gin
```

验证配置文件读取&变更

```azure
go mod tidy
go run main.go
http://127.0.0.1:8080/hello
```

## 日志中间件

引入第三方库zap记录日志

```azure
go get github.com/natefinch/lumberjack
go get go.uber.org/zap
go get go.uber.org/zap/zapcore
```

验证日志记录情况

```azure
http://127.0.0.1:9002/hello
```

## MySQL

引入第三方库gorm访问MySQL

```azure
go get github.com/jinzhu/gorm
go mod tidy
```

创建数据库gin_blog

```azure
create database gin_blog
mysql -h127.0.0.1 -P3306 -uroot -p12345678
use gin_blog
source gin_blog.sql
```

## 后端管理平台

登录后端管理平台

```azure
http://127.0.0.1:9002/admin/login
用户名：admin
密码；admin
```

## 登录Session管理

```
go get github.com/gin-contrib/sessions
```
