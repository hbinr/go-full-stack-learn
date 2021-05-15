/**
 * interface 接口定义，内容要点：
 * 1. 定义语法：interface 接口名{}
 * 2. 关键字：interface
 * 3. TypeScript的核心原则是对值所具有的结构进行类型检查
 * 4. TS中的接口类型更像是 “鸭子类型” ducking type ，满足以下两点，就说该类型符合某接口，允许使用
 *    a. 接口中定义的 属性是否存在
 *    b. 接口中定义的 属性类型是否匹配
 * 注意：我们传入的对象参数实际上会包含很多属性，但是编译器只会检查那些必需的属性是否存在，以及其类型是否匹配
 * 5. 在TS里，接口的作用就是为这些类型命名和你的代码或第三方代码定义锲约。 简单的来说就是每一次定义或者输出的时候都要定义类型
 * 6. 和Go中的接口有相似之处，但也有很大不同：
 *    a. 相似之处：二者都是类似 “鸭子类型” 的
 *    b. Go中的接口定义一次，只要有类型实现了接口中的所有方法，就是该接口的实现类，不需要显示 implement interface
 *    c. Go语言中，接口是一系列方法的集合，是方法定义的契约。没有定义属性字段的概念
 */

interface LabelledValue{
    label: string
}

function printLabel2(labelValue:LabelledValue) {
    console.log('labelValue: ', labelValue.label);
}
let testObj2 = {
    label: "this is a label",
    labelName:"no"
}
// testObj2 包含 label 属性，且数据类型为 string 
printLabel2(testObj2)