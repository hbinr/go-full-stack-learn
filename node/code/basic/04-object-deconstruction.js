/**
 * 对象解构，内容要点 (类似Python种dict推导式)
 * 1. 字段值获取： 对象实例.字段名 或者 对象示例[字段名]，类似键值对写法
 * 2. 解构，解析对象的字段，并自动赋值给变量， let {varibale1, varibale2, ... } = obj
 * 3. 解构注意事项： 
 *    a.字段名和变量名要保持一致，否则会undefined
 *    b.变量定义的顺序和字段名的顺序可以不一致
 *    c.支持部分解构.可以获取部分字段的值
 *    d.起别名。为了解决变量名冲突问题，使用了多次let定义同名变量
 */

let user = {
    name: "bob",
    age: 18,
}
// 第一种方式： '.' 运算符
let name = user.name
let age = user.age
console.log('name: ', name);
console.log('age: ', age);

// 第二种方式： 键值对方式
name = user["name"]
age = user["age"]
console.log('name: ', name);
console.log('age: ', age);


// 第三种方式： 解构 等效于第一种方式
let order = {
    price: 50.0,
    num: 18,
}

// 变量名和字段名保持一致
let { price, num } = order

// 变量顺序可以任意
// let { num, price } = order 
// let { num }= order // 部分解构
console.log('price, num: ', price, num);


// 起别名：解决变量名冲突问题
let { price: bookPrice, num: bookNum } = order
console.log('bookPrice, bookNum: ', bookPrice, bookNum);