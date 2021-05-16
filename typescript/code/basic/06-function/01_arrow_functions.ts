/**
 * 箭头函数，内容要点：
 * 1. 定义语法：() =>{}
 * 2. 关键字：=>
 * 3. 箭头函数是 ES6引入的概念，可以省略 function 关键字，简化代码，在特殊情况下，甚至还可以省略 () 和 {}
 * 4. 书写箭头函数，形参依旧放在小括号 () 中，多个形参及默认值的书写格式和之前相同
 * 5. 省略情况：
 *      a. 如果只有 1 个形参，并且没有带默认值，可以省略小括号 () 不写。 但是在typescript中不行，因为要指定参数类型
 *      b. 如果函数体 {} 内里面只有一条语句，可以省略大括号 {} 不写
 *      c. 在 b 的基础上，会返回一个值，给到函数的调用结果。return 当前 => 后面的值
 * 6. 作用域问题： typescript中不存在这个问题？
 *      a. 箭头函数没有自己的作用域，即箭头函数中的 this 指向其外层作用域         
 */

// 1. 普通函数改造成箭头函数
function func(){
    console.log('hello');
}   

// 改成箭头函数形式， 无参数无返回值
let func2 = () => {
    console.log('hello'); 
}

// 2.1 带参数的箭头函数
let func3 = (x: number,y: number) =>{
    console.log('x,y: ', x,y);
}
func3(1, 2)

// 2.2 只带一个参数的箭头函数，小括号 () 不能省略，因为要指定参数类型
let func4 = (x: number) =>{
    console.log('x: ', x);
}
func4(1)

// 2.3 如果函数体 {} 内里面只有一条语句，可以省略大括号 {} 不写，有参数无返回值
let func5 = (x: number) =>  console.log('x: ', x);
console.log('func5(): ',func5(1));  // func5():  undefined


// 2.4 省略 大括号{} 的返回值问题， 有参数有返回值
let func6 = (x:number  ) => x +1
console.log('func6(): ',func6(99));  // func6():  100


// 2.5 无参数有返回值
let func7 = () => 9 +1
console.log('func7(): ',func7());  // func9():  100
