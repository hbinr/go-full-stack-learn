/**
 * 类与接口的关系，内容要点：
 * 1. 定义语法 ： 同类 、接口的定义
 * 2. 关键字：class, interface
 * 3. 理清一下 类 和 接口的关系：
 *      a. 类实现接口的时候，必须声明接口中已声明的所有属性和方法。 类也可以定义自己的属性
 *      b. 接口只能约束 类 的 公有成员，因为接口中无法定义私有成员，只能是 public 的
 *      c. 接口不能约束 类 的 构造函数
 *      d. 接口继承 接口。接口 可以像 类 一样继承多个接口，子类必须实现所有接口中已声明的属性或方法
 *      e. 接口继承 类。它会继承类的成员但不包括其实现。就好像接口声明了所有类中存在的成员，但并没有提供具体实现一样。
 * 
 */

// 1. 类实现接口
interface Human {
    name: string
    sleep(): void

    // 接口不能约束 类 的构造函数
    // new(name: string): void // 报错：类型“Asian”提供的内容与签名“new (name: string): void”不匹配。ts(2420)

}

class Asian implements Human {
    constructor(name: string) {
        this.name = name
    }

    // 接口只能约束 类 的公有成员
    // private name: string // 报错：属性“name”在类型“Asian”中是私有属性，但在类型“Human”中不是

    name: string
    sleep() { }
}

// 2. 接口继承类

class Auto {
    state: number

    add(): void {
        console.log('类 Auto 实现add()');
    }
}

// 接口继承类， AutoInterface
interface AutoInterface extends Auto {

}

// 类C 实现接口 AutoInterface
class C implements AutoInterface {
    state = 2
    add() {
        console.log('类 C 实现add()');
    }
}

let testC = new C()

testC.add() // 类 C 实现add()

console.log(testC.state); // 2
