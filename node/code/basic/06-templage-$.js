/**
 * 模板语法，内容要点
 * 1. 模板语法： `${变量}XXX`
 * 2. 多个变量时取last位置。`${变量1,变量2...变量last}XXX`  这种情况下，只取${}中最后一个变量的值
 * 3. 和jQuery没有关系，是ES6新提出的
 */

let user = {
    name: "Tom",
    age: 18,
}

console.log(`${user.name}的年龄为${user.age}岁`);
console.log(`${user.name, user.age}的年龄为${user.age}岁`);