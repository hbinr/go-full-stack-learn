/**
 * 数组/字符串解构，内容要点 (类似Python种 列表 推导式，列表就是动态数组)
 * 1. 元素获取是一一对应的关系，从索引为0的位置开始，依次解构
 * 2. 支持嵌套数组，始终记住：数组解构时是一一对应的
 * 3. 可以是逗号 ',' 作为占位符，表示不获取对应位置元素
 */

// 一. 数组解构
let arr = [10, 20, 30]

// 完全解构
let [a, b, c] = arr
console.log('a, b, c: ', a, b, c);

// 部分解构， 使用 , 作为占位符
let [, d] = arr
console.log('d: ', d);

// 嵌套数组
let arr2 = [10, 20, 30, [1, 2]]
let [i, j, k, [x, y]] = arr2
console.log('i, j, k: ', i, j, k);
console.log('x, y: ', x, y);


// 一. 字符串解构  字符串底层也是数组
let str = "hello"

let [s1, s2, s3] = str  // 部分解构
console.log('s1, s2, s3: ', s1, s2, s3);

str[1] = "E" // 无法更改，但是不报错。Go在编译阶段就报错了
console.log('str[1]: ', str[1]); // e

