# 导入proto实战


[.proto文件详情请见](../code/demo_product/pbfile/)

有3个proto文件 `base_model.proto` `prod_model.proto`、`product.proto`文件内容如下：

**base_model.proto：**
```go
syntax = "proto3";

package model;
option go_package = "go-full-stack-learn/clound-native/grpc/code/demo_product/pbfile";

message BaseModel {
    int64 UUID = 1;
}
```

**prod_model.proto：**
```go
syntax = "proto3";

package model;

option go_package = "go-full-stack-learn/clound-native/grpc/code/demo_product/pbfile";
import "base_model.proto";

message Product {
    BaseModel BaseModel = 1;
    int32 ProdID = 2;
    string ProdName = 3;
    float ProdPrice = 4;
}
```

**product.proto：**
限于篇幅，只列了一部分

```go
syntax = "proto3";

package product_service;

option go_package = "go-full-stack-learn/clound-native/grpc/code/demo_product/pbfile";

message ProdRequest {
  int32 prodID = 1;
}

service ProdService{
  // 接口定义......
}
```

`product.proto`中需要增加一个接口——返回商品信息，返回字段类型是`Product`。

需要在 `product.proto` 中 `import "prod_model.proto"`，注意导入是字符串格式， 如下所示： 
```go
syntax = "proto3";

package product_service;

// 改动一：引入 prod_model.proto 定义的内容
import public "prod_model.proto";

option go_package = "go-full-stack-learn/clound-native/grpc/code/demo_product/pbfile";

message ProdRequest {
  int32 prodID = 1;
}

service ProdService{
  // 改动：新增获取商品信息接口  
  rpc GetProdInfo(ProdRequest) returns(model.Product); // 需要加上包名 model
}
```

如果还想使用 `base_model.proto` 中的 `BaseModel` ， 你不能因为 `product.proto` 引用了 `prod_model.proto`， `prod_model.proto`又引用了 `base_model.proto`，就可以直接在 `product.proto` 中使用 `base_model.proto` 的对象，这里不具有传递性， `base_model.proto` 在 `product.proto` 中依旧不可见。

有两种方法可以解决：
  
- 在 `product.proto` 中 `import "base_model.proto"`;
- 在 `product.proto` 中 `import public "prod_model.proto"`;  推荐这种方式
  
**需要提醒一下：**

> 在使用Goland编写 proto 文件时，尽管已经引入了插件 **Protocol Buffer Editor**，但是 import 其他proto的时候会报红，但是不影响正常使用。