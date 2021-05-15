/**
 * class 类定义，内容要点：
 * 1. 定义语法：class 类名{属性|构造函数|方法}
 * 2. 关键字：class , new 
 * 3. TypeScript 中的类类似于Java，Python中的类，有属性(成员变量)、构造函数、方法(成员方法)
 * 4. 通过 new 关键字可创建对象实例
 * 5. Go中并没有类的概念，而是结构体 struct
 *     
 */

class Greeter{
    // 属性
    greeting: string

    // 构造函数(构造器)
    constructor(msg: string){
        this.greeting = msg
    }

    // 方法
    greet(){
        console.log('hello ' + this.greeting );
    }
}

let greeter2 = new Greeter('typescript')
greeter2.greet()
