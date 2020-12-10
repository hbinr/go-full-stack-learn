# KubeVela安装

[官网]()
## Minikube 下安装

### 一、启动k8s集群
#### 1.启动
```sh
minikube start --driver=docker --base-image="anjone/kicbase" --image-mirror-country='cn' \
--image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers' \
--registry-mirror=https://guuu392p.mirror.aliyuncs.com \
--image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers' \
--logtostderr
```
#### 2.安装入口
```sh
$ minikube addons enable ingress
```
### 二、下载KubeVela
macOS/Linux

```sh
$ curl -fsSl https://kubevela.io/install.sh | bash
```
如果嫌慢的话，可以尝试：
- 1.手动下载[安装包](https://github.com/oam-dev/kubevela/releases)
- 2.解压，进入解压后的目录，如`cd linux-amd64` 
- 3.将解压后 `vela`二进制文件移动到`bin`目录下， `sudo mv vela /usr/local/bin/vela`
### 三、安装KubeVela
```sh
$ vela install
```
这将安装KubeVela服务器组件及其依赖组件，后台会一直在安装，耐心等待

#### 检查是否已安装`Vela Helm Chart`
```sh
$ helm list -n vela-system

NAME      NAMESPACE   REVISION  ...
kubevela  vela-system 1         ...
```
#### 检查是否已安装所有依赖项组件

需要等一会，5-10分钟显示
```sh
$ helm list --all-namespaces

NAME                  NAMESPACE   REVISION  UPDATED                               STATUS    CHART                       APP VERSION
flagger               vela-system 1         2020-11-10 18:47:14.0829416 +0000 UTC deployed  flagger-1.1.0               1.1.0
keda                  keda        1         2020-11-10 18:45:15.6981827 +0000 UTC deployed  keda-2.0.0-rc3              2.0.0-rc2
kube-prometheus-stack monitoring  1         2020-11-10 18:45:37.9608079 +0000 UTC deployed  kube-prometheus-stack-9.4.4 0.38.1
kubevela              vela-system 1         2020-11-10 10:44:20.663582 -0800 PST  deployed
```

未来引入`vela system health`命令来检查依赖关系，便不需要使用上述命令了

### 四、卸载
**执行：**
```sh
$ helm uninstall -n vela-system kubevela
$ rm -r ~/.vela
```
将卸载KubeVela服务器组件+其依赖组件+本地CLI缓存。

然后清理CRD（默认情况下，不会通过via helm删除CRD）：
```sh
$ kubectl delete crd \

applicationconfigurations.core.oam.dev \
applicationdeployments.core.oam.dev \
autoscalers.standard.oam.dev \
certificaterequests.cert-manager.io \
certificates.cert-manager.io \
challenges.acme.cert-manager.io \
clusterissuers.cert-manager.io \
components.core.oam.dev \
containerizedworkloads.core.oam.dev \
healthscopes.core.oam.dev \
issuers.cert-manager.io \
manualscalertraits.core.oam.dev \
metricstraits.standard.oam.dev \
orders.acme.cert-manager.io \
podspecworkloads.standard.oam.dev \
routes.standard.oam.dev \
scopedefinitions.core.oam.dev \
servicemonitors.monitoring.coreos.com \
traitdefinitions.core.oam.dev \
workloaddefinitions.core.oam.dev
```