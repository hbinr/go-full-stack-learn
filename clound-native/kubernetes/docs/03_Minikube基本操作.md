# Minikube基本操作

### 检测集群状态  `kubectl cluster-info`
```sh
$ kubectl cluster-info
 
Kubernetes master is running at https://192.168.99.100:8443
KubeDNS is running at https://192.168.99.100:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

Ubuntu下，Minikube的配置文件在如下路径

~/.minikube/machines/minikube/config.json

### 查看配置文件内容 `kubectl config view`
```sh
$ kubectl config view

apiVersion: v1
clusters:
- cluster:
    certificate-authority: /home/james/.minikube/ca.crt
    server: https://192.168.99.100:8443
  name: minikube
contexts:
- context:
    cluster: minikube
    user: minikube
  name: minikube
current-context: minikube
kind: Config
preferences: {}
users:
- name: minikube
  user:
    client-certificate: /home/james/.minikube/client.crt
    client-key: /home/james/.minikube/client.key
```
 



## 检验Node状态 `kubectl get nodes`

```sh
$ kubectl get nodes
NAME       STATUS   ROLES    AGE   VERSION
minikube   Ready    master   11m   v1.15.0
```

## 使用ssh进入Minikube虚机 `sudo minikube ssh`

```sh
$ sudo minikube ssh
                         _             _            
            _         _ ( )           ( )           
  ___ ___  (_)  ___  (_)| |/')  _   _ | |_      __  
/' _ ` _ `\| |/' _ `\| || , <  ( ) ( )| '_`\  /'__`\
| ( ) ( ) || || ( ) || || |\`\ | (_) || |_) )(  ___/
(_) (_) (_)(_)(_) (_)(_)(_) (_)`\___/'(_,__/'`\____)

```

## 停止运行中的kubernetes集群 `minikube stop`

```
$ minikube stop
```

## 删除本地的kubernetes集群  `minikube delete`

```sh
$ minikube delete
🔥  正在删除 docker 中的“minikube”…
🔥  正在移除 /home/hblock/.minikube/machines/minikube…
🔥  尝试删除无效的配置文件 minikube
```


## 删除本地的kubernetes集群+所有配置 `minikube delete --all`
```sh
$ sudo minikube delete --all
🔄  正在使用 kubeadm 卸载 Kubernetes v1.19.4…
🔥  正在删除 none 中的“minikube”…
💀  Removed all traces of the "minikube" cluster.
🔥  成功删除所有配置文件
```
## 打开Kubernetes控制台，直接在默认浏览器上打开 `minikube dashboard`

## 获取仪表板的URL `minikube dashboard --url`

```sh
$ minikube dashboard --url
http://192.168.39.117:30000
```

一旦Minikube虚拟机启动，用户就可以使用熟悉的Kubectl CLI在Kubernetes集群上执行操作。

通过打开您最喜欢的浏览器上的URL访问Kubernetes Dashboard。进一步阅读，请查看:

- 你好Minikube系列: https://kubernetes.io/docs/tutorials/stateless-application/hello-minikube/
- minkube新手指南: https://kubernetes.io/docs/getting-started-guides/minikube/