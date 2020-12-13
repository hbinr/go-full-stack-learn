# rest api中的model， vo， dto之间的关系

**model:**用于接收数据库中的数据，通过mybatis 的ORM对象关系映射来获取数据

service主要是来处理业务逻辑，返回数据（数据的返回不会做如何处理，只是将model中的数据进行整合然后保证返回的数据完整性，比如在service中返回一个user model 和一个 role model, 会提供一个UserROleBO来接收user 和 role的数据）

**dto:** dto主要是用来作为传输数据，在我们的项目中是用来接收远程调用接收响应的对象，如使用`grpc`时，需要将传输过来的数据进行转换，根据业务需求，要么转为model落库，要么转为vo返回给UI界面

**vo：**主要是我们本地调用的对象， 用来向用户显示（BO中的user和role只需要显示userName和roleName,会通过VO来进行显示）

在实际的开发中，VO对应页面上需要显示的数据，DO对应于数据库中储存的数据（表列，也就是model）,DTO对应于除二者之外需要传递的数据。

 DTO(data transfer object):数据传输对象，以前被称为值对象(VO,value object)，作用仅在于在应用程序的各个子系统间传输数据，在表现层展示。与POJO对应一个数据库实体不同，DTO并不对应一个实体，可能仅存储实体的部分属性或加入符合传输需求的其他的属性。