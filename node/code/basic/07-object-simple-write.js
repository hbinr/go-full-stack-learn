/**
 * 对象的简化语法，内容要点
 * 1. 已声明的变量名和对象的字段名完全相同时，直接写字段名就行
 * 2. 省略 “赋值”过程
 * 3. 是ES6新提出的
 */

let name = "Bob"
let age = 20


let user = {
    name: name,
    age: age
}

// 对象简化语法
let student = { name, age }

console.log('student: ', student);