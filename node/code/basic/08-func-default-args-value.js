/**
 * 函数参数默认值，内容要点  基本和Python写法一致
 * 1. 通过 ‘=’ 给某个字段指定默认值
 * 2. 调用函数时，如果形参未传递任何值，那么则使用函数定义时的默认值
 * 3. 调用函数时，如果形参已经传递了真正的值(实参)，那么则使用实参
 */

function getUser (name = "bob", age) {
    console.log('name,age: ', name, age);
}

// 形参已经传递了真正的值(实参)，那么则使用实参
getUser("tom", 20) // name,age:  tom 20



// 如果形参未传递任何值，那么则使用函数定义时的默认值
getUser() // name,age:  bob undefined

// 如果形参只传递了部分呢？ 按参数定义的顺序解析，实参值依次赋值
getUser(30)  // name,age:  30 undefined