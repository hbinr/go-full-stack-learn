# PB类型与go语言对应类型汇总


| .proto Type |                                   Notes                                    | Go Type |
| :---------: | :------------------------------------------------------------------------: | :-----: |
|   double    |                                                                            | float64 |
|    float    |                                                                            | float32 |
|    int32    | 使用变长编码，对于负值的效率很低，如果你的域有可能有负值，请使用sint64替代 |  int32  |
|   uint32    |                                使用变长编码                                | uint32  |
|   uint64    |                                使用变长编码                                | uint64  |
|   sint32    |               使用变长编码，这些编码在负值时比int32高效的多                |  int32  |
|   sint64    |          使用变长编码，有符号的整型值。编码时比通常的int64高效。           |  int64  |
|   fixed32   |     总是4个字节，如果数值总是比总是比228大的话，这个类型会比uint32高效     | uint32  |
|   fixed64   |     总是8个字节，如果数值总是比总是比256大的话，这个类型会比uint64高效     | uint64  |
|  sfixed32   |                                总是4个字节                                 |  int32  |
|  sfixed64   |                                总是8个字节                                 |  int64  |
|    bool     |                                                                            |  bool   |
|   string    |             一个字符串必须是UTF-8编码或者7-bit ASCII编码的文本             | string  |
|    bytes    |                         可能包含任意顺序的字节数据                         | []byte  |

如果想看其他语言的类型，[详见官网](https://developers.google.com/protocol-buffers/docs/proto3?hl=zh-cn#scalar)

# 默认值
解析消息时，如果编码的消息不包含特定的单数元素，则已解析对象中的相应字段将设置为该字段的默认值。 这些默认值是特定于类型的：

- 对于strings，默认值是一个空string
- 对于bytes，默认值是一个空bytes
- 对于bools，默认值是false
- 对于数值类型，默认值是0
- 对于枚举，默认值为第一个定义的枚举值，必须为0
- 对于消息字段（message），未设置该字段。 它的确切值取决于语言。 有关详细信息，[请参见generated code guide](https://developers.google.com/protocol-buffers/docs/reference/go-generated?hl=zh-cn)


**注意：**
- 重复字段的默认值为空（通常为相应语言的空列表），就是加了repeated修饰的字段，其默认值为对应修饰类型的空值。
- 对于标量消息字段，一旦解析了一条消息，就无法告诉该字段是显式设置为默认值（例如，布尔值是否设置为false）还是根本没有设置：您应该在 定义消息类型时要注意。 例如，如果您不希望默认情况下也发生这种情况，则当布尔值设置为false时，没有布尔值会打开某些行为。
- 如果将标量消息字段设置为其默认值，则该值将不会在线路上被序列化。


