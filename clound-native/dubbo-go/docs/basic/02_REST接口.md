**provider步骤：**

- 1.填写配置：server.yml+log.yml
- 2.main.go 导包，一定要记得 `_ "github.com/apache/dubbo-go/config_center/nacos"`
- 3.启动nacos server
- 4.增加 dubbo.properties

**仍有问题：**
> panic: config center for nacos is not existing, make sure you have import the package.

**解决：**
- 一定要记得导包 `_ "github.com/apache/dubbo-go/config_center/nacos"`

参考：https.csdn.net/ra681t58cjxsgckj31/article/details/106110109

