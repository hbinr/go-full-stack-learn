/**
 * 数组解构，内容要点：(类似Python中 列表 推导式，列表就是动态数组)
 * 1. 解构顺序是一一对应的关系，从索引为0的位置开始，依次解构
 * 2. 支持嵌套数组，始终记住：数组解构时是一一对应的
 * 3. 可以是逗号 ',' 作为占位符，表示不获取对应位置元素
 */

// 一. 数组解构
let arr1: number[] = [10, 20, 30]

// 完全解构
let [a, b, c] = arr1
console.log('a, b, c: ', a, b, c);  // a, b, c:  10 20 30

// 部分解构， 使用 , 作为占位符
let [, d] = arr1
console.log('d: ', d);  // d:  20


// 在数组里使用 ... 语法创建剩余变量
let [first, ...last] = [1,2,3,4,5]

console.log('first: ', first); // first: 1
console.log('last: ', last);   // last:  [ 2, 3, 4, 5 ]

// 嵌套数组
let arr2 = [10, 20, 30, arr1]
let [i, j, k, res] = arr2
console.log('i, j, k: ', i, j, k);  // i, j, k:  10 20 30
console.log('res: ', res);  // res:  [ 10, 20, 30 ]



