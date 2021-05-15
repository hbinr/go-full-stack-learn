/**
 * interface 接口是如何工作的，内容要点：
 * 1. 类型检查器会查看 printLabel 的调用。
 * 2. printLabel 有一个参数，并要求这个对象参数有一个名为 label 类型为 string 的属性。 
 * 3. 需要注意的是，我们传入的对象参数实际上会包含很多属性，但是编译器只会检查那些必需的属性是否存在，以及其类型是否匹配。
 */

// labelledObj 对象有label属性
function printLabel(labelledObj: { label: string }) {
    console.log(labelledObj.label)
}


let myObj = { size: 10, label: 'Size 10 Object' }

printLabel(myObj) // 入参 myObj 对象中也包含了label属性，所以可以成功调用