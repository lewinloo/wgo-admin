# wgo-admin

## 项目介绍
本项目基于 [Golang](https://go.dev/) 语言 [Gin](https://gin-gonic.com/) 框架。

## 目录介绍

```text
.
├── README.md                       # 说明文件
├── cmd                             # 命令文件目录
│   └── app                         # app 模块
│       └── main.go                 # 运行服务入口
├── config                          # 配置文件
│   ├── config.go                   # 配置文件入口
│   ├── gorm_mysql.go               # mysql 配置
│   ├── jwt.go                      # jwt 配置
│   ├── redis.go                    # redis 配置
│   ├── system.go                   # 系统配置
│   └── zap.go                      # zap 配置
├── docs                            # swagger 自动生成文件
├── internal                        # 内部核心
│   └── app                         # app 核心
│       ├── api                     # api 层
│       │   ├── handler             # handler
│       │   ├── middleware          # 中间件
│       │   └── router.go           # 路由控制
│       ├── common                  # 公用模块
│       ├── core                    # 核心模块
│       ├── domain                  # domain 层
│       │   ├── model               # model 层
│       │   ├── repository          # repository 层
│       │   └── service             # service 层
│       ├── global                  # 全局变量
│       ├── utils                   # 工具类
│       ├── wire.go                 # 依赖注入代码
│       └── wire_gen.go             # 生成的依赖注入代码
├── pkg                             # 第三方包
│   ├── casbin.go                   # casbin
│   ├── cron.go                     # 定时任务
│   ├── gorm.go                     # gorm
│   ├── gorm_mysql.go               # gorm_mysql
│   ├── redis.go                    # redis
│   ├── viper.go                    # viper
│   └── zap.go                      # zap 日志
├── resource                        # 资源目录
│   ├── application.yaml            # 项目启动配置文件
│   └── rbac_model.conf             # rbac 匹配规则
├── scripts                         # 脚本
│   └── swagger.sh                  # 自动生成 swagger shell 脚本
├── test                            # 测试
├── .air.conf                       # 热更新配置
├── .gitignore                      # 忽略提交git
├── docker-compose.yml              
├── Dockerfile                      # Docker 打包镜像
├── go.mod                          
└── Makefile                        

```
