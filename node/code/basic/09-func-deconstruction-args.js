/**
 * 函数形参解构机制及其默认值，内容要点
 * 1. 形参是对象时，也支持对象解构形式，省略 “赋值”过程
 * 2. 形参为对象时，可以设定默认值
 * 3. 注意入参为空和空对象的不同
 * 4. 是ES6新提出的
 */

let user = {
    name: "tom",
    age: 18
}
// 一. 函数形参解构机制
function getUser ({ name, age }) { // 等效于 {name, age} = user 解构格式
    // name和age值：就是user对象实例中属性的值
    console.log('name,age: ', name, age);
}


getUser(user) // 控制台会输出：name,age:  tom 18


// 注意1：如果函数入参为空，会报错
// getUser() // {name, age} = null  运行报错  TypeError: Cannot destructure property 'name' of 'undefined' as it is undefined.


// 注意2：如果函数入参确实没有具体值，但是还不想报错，可以传递空对象 {}
getUser({})  // name,age:  undefined undefined  不会报错，输出了undefined


// 二. 函数形参为对象时，可以设定默认值
function getUser2 ({ name, age } = {}) { // 默认值为空对象
    // 如果没有参数传递进来，那么使用默认值
    console.log('name,age: ', name, age);
}

getUser2() // 执行这行代码时不会像 第24行 代码那样报错了，效果等同于  getUser({})  