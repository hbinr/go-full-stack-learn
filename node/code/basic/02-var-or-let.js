/**
 * 内容要点：
 *  1. ES5中，var的弊端介绍和案例
 *  2. ES6中，新关键字let学习和案列，为了解决上述问题
 *  3. let 用于定义变量
 */

// var 弊端 1. var 声明变量的时候会有预解析(变量提升)，造成逻辑混乱，可以先使用，后声明
// 在Go中则严格相反，必须先声明，后使用，而且如果你声明了，但是未使用也会报错

console.log(a); // undefined， 正常反应应该是报错，而不是可以解释并执行成功
var a = 10;

//  var 弊端 2. var 可以重复定义同一个名字的变量，也会造成逻辑混乱，逻辑错误。
// 正常情况下，第二次使用已定义的变量应该是更新操作，而不是重新定义 

var b = 10
var b = 20   // 定义还是修改？
console.log('b: ', b); // b: 20

//  var 弊端 3. var 变量作用域污染问题。比如在for循环中，定义一个局部变量
for (var i = 0; i < 5; i++) {
    console.log('i: ', i);
}
console.log('------------------');
// 下面这行代码会输出i=5，正常情况应该是编译报错 undefined: i，这才是正常逻辑，如Go，Java都是这样处理的
console.log('i: ', i); // i:  5 ，es5 并不会报错，而且还有值



//  var 弊端 4. var 在声明的时候没有块级作用域的概念，这会造成变量作用域污染的问题。ES5中，本身就只有全局、局部作用域。

var c = 10 //全局作用域

function fn () {
    var d = 20 // 局部作用域
}

{
    // 正常情况应该是编译报错 undefined: e，这才是正常逻辑
    var e = 30 // 块级作用域，go语言中会编译报错： e declared but not used
}
console.log('e: ', e); // e: 30，，es5、es6 并不会报错，而且还有值，go语言中会编译报错：  undefined: e


// let ES6 新关键字，为了解决上述问题


// var 弊端 1 解决：
console.log('f: ', f);  // 运行报错：ReferenceError: Cannot access 'f' before initialization
let f = 1

// var 弊端 2解决：
let g = 10
let g = 20   // 运行报错：SyntaxError: Identifier 'g' has already been declared


// var 弊端 3 解决：
for (let j = 0; j < 5; j++) { // j 使用let定义后，就是一个临时变量，用完就回收，在for外面使用就会报错
    console.log('j: ', j);
}
console.log('------------------');
console.log('j: ', j); // 运行报错：ReferenceError: j is not defined

// var 弊端 4 解决：

{
    let h = 30 // VS Code也会高亮提示，表示该字段未声明，只能在该块级作用域中使用
}
console.log('h: ', h); //  运行报错：ReferenceError: h is not defined