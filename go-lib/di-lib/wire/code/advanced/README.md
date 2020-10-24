# wire依赖注入库开发示例

面向接口开发web项目，使用wire依赖注入库来管理各层之间的依赖。


```
├── config.yaml              
├── internal                          # 具体业务 
│   └── app
│       └── user                      # 按职责划分了该业务模块，方便后续拓展微服务，其内部还是按CLD模式分层
│           ├── cmd
│           │   ├── cmd
│           │   ├── config-test.yaml
│           │   ├── main.go     
│           │   ├── wire_gen.go       # wire工具自动生成的代码
│           │   └── wire.go           # 创建依赖关系，实现依赖注入函数
│           ├── conf
│           │   └── conf.go
│           ├── controller
│           │   └── user.go
│           ├── dao
│           │   ├── config-test.yaml  # 简单测试用
│           │   ├── user.go
│           │   └── user_test.go
│           ├── model
│           │   ├── params.go
│           │   └── user.go
│           ├── service
│           │   └── user.go
│           └── set.go                # user模块的依赖provider
├── pkg                               # 公共的，与业务无关的的模块
│   ├── database
│   │   └── mysql.go
│   ├── ginx
│   │   └── response.go
│   └── middleware
│       └── jwt.go
├── README.md
└── setting                           # 项目初始设置，主要是初始化配置
    └── conf.go
```