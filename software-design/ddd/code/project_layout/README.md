以一个个人用户为例，
## 目录解释
```
├── api
├── cmd
│   ├── user-job
│   └── user-task
├── conf
├── internal
│   ├── application         应用层
│   │   ├── event             领域事件的发布和订阅类我建议放在应用层的 Event 目录结构下。
│   │   │   ├── publish       发布事件
│   │   │   └── subscribe     订阅事件
│   │   └── service
│   ├── domain              领域服务层
│   │   ├── aggregate01    
│   │   │   ├── entity
│   │   │   ├── event
│   │   │   ├── repository
│   │   │   └── service
│   │   └── aggregate02
│   │       ├── entity         实体类，一般情况下，领域模型的业务实体与微服务的数据库实体是一一对应的
│   │       ├── event          领域事件实体和处理类存放在领域层的 event 目录结构下，
│   │       │                  领域事件的发布和订阅类我建议放在应用层的 Event 目录结构下。
│   │       ├── repository     仓储接口及实现，仓储主要用来完成数据查询和持久化操作
│   │       └── service        领域服务类
│   ├── interfaces           用户接口层
│   │   ├── assembler
│   │   ├── dto
│   │   └── facade
│   └── pkg                  
│       ├── code
│       └── key
└── pkg
    ├── database
    ├── json
    ├── logger
    ├── mq
    └── time

```

比如执行一个创建用户的命令，依次执行:
- 1. 用户接口层： 
  - 1.1 Assembler->将CustomerDTO转换为CustomerEntity
  - 1.2 Dto->接收请求传入的数据CustomerDTO

  - 1.3 Facade->调用应用层创建用户方法

- 2. 应用层
  - 2.1 Event->发布用户创建事件给其它微服务
  - 2.2 Service:
     - 内部服务->创建用户
     - 外部服务->创建日志
- 3. 领域层
  - 3.1 Aggregate->进入用户聚合目录下(如：CustomerAggregate)
  - 3.2 Entity->用户聚合跟
  - 3.3 Event->创建用户事件
  - 3.4 Service->具体的创建用户逻辑，比如用户是否重复校验，分配初始密码等
  - 3.5 Repository->将用户信息保存到数据库
