/**
 * 函数类型，内容要点：
 * 1. 定义语法：(参数1:数据类型行, ....) => 返回值数据类型
 * 2. 关键字：(), =>
 * 3. 函数类型，顾名思义，首先其是一个类型，是函数形式的，之前我们学过number, string。道理类似
 * 4. 函数类型包含两部分：参数类型和返回值类型。
 * 5. 函数也可以作为类型，本质上也是 values，可以作为入参，出参，函数式编程；Go中也有函数类型，一切皆类型，支持函数式编程
 * 
 */

// 1. 定义函数类型， 类似定义函数签名，没有函数体
let addFunc : (x: number, y: number) => number

// 2. 函数体具体实现
addFunc = function(x: number, y: number): number{
    return x + y
}

addFunc(1,2)

// 3. 上述二者可以合并写：
let addFunc2: (baseVale: number, increment: number) => number =  // 多了个等号，相当于直接实现了
function (x: number, y: number): number {
    return x + y
}

// 4. 类型推断，更加简化代码； 函数具体实现中，入参和出参不指定类型，ts也能自动推断出
let addFunc3: (baseVale: number, increment: number) => number = function(x, y ) { return x + y}

// 5. 函数式编程

// 5.1 定义函数类型
type addFunc4 = (x: number, y: number) => void;

// 5.2 函数作为参数
function greeter(add: addFunc4) {
    add(1, 2)
}
  

