## 四组基本概念
### Pod/Pod控制器
- pod是k8s能够运行的最小单元（原子单元）
- 一个pod运行多个容器，他们共享UTS+NET+IPC名称空间（sidecar 边车模式）
- pod控制器是pod启动的一种模板，用来保证在k8s里启动的pod都按照预期运行（副本数，生命周期，健康状态检查）
- k8s提供了很多的pod控制器，常用的是Deployment DaemonSet 。。。
### Name/Namespace
- name 对应“资源”，每个“资源”都应该有它对应的名称。资源通常包括以下信息，api版本，类别，元数据，定义清单(spec)，状态,等配置信息，name通常定义在资源的元数据中。
- k8s中隔离各种资源的方法是，名称空间，可以理解为虚拟集群组
- 不同名称空间的资源名可以相同，同一名称空间的同种资源名字不能重复
- k8默认的名称空间有 default、kube-system、kube-public 等，并且查询相应的资源要带上相应的名称空间
### Lable/Lable选择器
- 一个资源对应多个标签，一个标签也可能对应多个资源，是多对多的关系。类似的还有注解
- 标签组成 key=value。
- 给资源打上标签后，可以通过标签选择器过滤标签
- 标签选择器目前有两个：1.基于等值关系，等于或者不等于，2.基于集合关系，属于，不属于，存在。
- 许多资源支持内嵌标签选择器：matchlable， matchexpressions
### Service/Ingress
- 一个service可以看做是一组提供相同服务的pod的对外接口
- service作用于那些pod是根据标签选择器确定的
- service只能用于IPV4的调度：ip+port
- ingress是应用对外暴露的接口，作用于不同的业务域，不同url访问路径的业务流量
## k8s核心组件
- 配置存储中心——etcd服务
- 主控节点，即（master）节点
      - kube-apiserver服务
      - kube-controller-manage服务
      - kube-scheduler服务
- 运算节点，即(node)节点
      - kube-kubelet服务
      - kube-proxy服务
- CLI客户端
      - kubectl
## k8s核心附件
- CNI网路插件——flannel/calico
- 网路发现插件——coredns
- 服务暴露插件——traefik
- GUI管理插件——Dashboard
### Pod
k8s最小的运行单元，其内可以运行一个或多个容器，每个pod都有自己的IP

### Service
Service里可以包含多个pod，有自己的IP

### Label
Label 标签(k-v结构)，用来区分是哪个Service，每个pod都会打上标签，这样pod就会被划分到对应的service下
### Node

部署的节点，包含了service+pod

### Deployment


### Worker

### Master 
核心中核心，管理多个worker