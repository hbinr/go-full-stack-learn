/**
 * 类(class)类型 接口，内容要点：
 * 1. 定义语法：interface StringArray {
 *                [index: number]: string
 *             }
 * 2. 关键字：interface , constructor, new 
 * 3. 类类型接口，定义了属性\方法，能够用它来明确的强制一个类去符合某种契约。
 * 4. 注意：接口描述了类的公共部分，而不是公共和私有两部分。 它不会帮你检查类是否具有某些私有成员。
 * 5. 类静态部分与实例部分的区别
 *      a. 当你操作类和接口的时候，你要知道类是具有两个类型的：静态部分的类型和实例的类型
 *      b. 静态部分的类型： constructor / new() 所以不在检查的范围内。
 *      c. 实例的类型：字面量定义的属性和方法
 */

// 1. 类类型，只定义属性

interface ClockInterface{
    currentTime: Date
}

// 类Clock实现  ClockInterface 接口
class Clock implements ClockInterface{
    currentTime: Date

    constructor(currentTime: Date){
        this.currentTime = currentTime
    }
}

// 2. 类类型，定义属性 + 方法

interface ClockInterface2{
    currentTime: Date
    setTime(d: Date):void
}

// 类Clock实现  ClockInterface 接口
class Clock2 implements ClockInterface2{
    // 定义 ClockInterface2 属性
    currentTime: Date

    // 实现 ClockInterface2 中的方法
    setTime(d:Date):void{
        console.log('setTime success: ', d);
    }

    constructor(currentTime: Date){
        this.currentTime = currentTime
    }
}

// 3. 类静态部分与实例部分的区别
// 类静态部分：构造器 constructor 型接口，new()函数定义
interface ClockConstructor{
    new (h: number, m: number): ClockInterface3  // 类静态部分，构造器
}
// 实例部分：实例型 接口，普通方法定义
interface ClockInterface3{
    tick():void
}

// 实例创建方法: 为了方便我们定义一个构造函数 createClock，它用传入的类型创建实例。
function createClock(ctor: ClockConstructor, h: number, m: number):ClockInterface3 {
    return new ctor(h, m)
}


// 创建类去实现 实例型接口 ClockInterface3
class DigitalClock implements ClockInterface3 {
    constructor(h: number, m: number) {    }

    tick(): void{
        console.log('DigitalClock...');
    }
}


class AnalogClock implements ClockInterface3 {
    constructor(h: number, m: number) {    }

    tick(): void{
        console.log('AnalogClock...');
    }
}

// 类 DigitalClock 实例化及调用
let digitalClock: DigitalClock = createClock(DigitalClock, 12, 17)
digitalClock.tick()

// 类 AnalogClock 实例化及调用
let analogClock: AnalogClock = createClock(AnalogClock, 7, 35)
analogClock.tick()

// 因为 createClock 的第一个参数是 ClockConstructor 类型，
// 在 createClock(AnalogClock, 7, 32) 里，会检查 AnalogClock 是否符合构造函数签名。
