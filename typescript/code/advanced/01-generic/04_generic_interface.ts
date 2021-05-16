/**
 * 泛型接口，内容要点：
 * 1. 定义语法：<T>(value: T): T
 * 2. 关键字：interface, <>, T,
 * 3. 类型别名: type 
 * 4. 泛型接口
 *      a. 约束接口的函数类型
 *      b. 约束接口的所有成员。 在实现接口的时候，必须指定类型
 */

function log<T>(val: T): T {
    return val
}

// 1. 类型别名
type Log = <T>(value: T) => T

let myLog: Log = log

// 2. 泛型接口

// 2.1 约束接口的函数类型
interface Log2 {
    <T>(val: T): T
}

// 2.2 约束接口的所有成员
interface Log3<T> {
    (val: T): T
}

// Log3<string>, 必须指定类型
let myLog3: Log3<string> = log
myLog3('myLog3...')


// 2.3 约束接口的所有成员， 并且指定默认类型
interface Log4<T = number> { // 泛型默认为 number
    (val: T): T
}

let myLog4: Log4 = log // 无需指定类型，默认为number类型
myLog4(4)


