/**
 * 交叉类型，内容要点：
 * 1. 定义语法: 类型一 & 类型二。
 * 2. 关键字: &
 * 3. 交叉类型是将多个类型合并为一个类型。
 *    这让我们可以把现有的多种类型叠加到一起成为一种类型，它包含了所需的所有类型的特性。
 * 4. 注意类型冲突：属性名称完全相同，但是属性类型不同，这会造成冲突，不可被赋值
 * 5. 应用场景：大多是在混入（Mixins）或其它不适合典型面向对象模型的地方看到交叉类型的使用
 */

// 1. 定义接口 
interface Admin {
    id: number,
    administrator: string,
    timestamp: string  // 属性名相同，类型不同
}

interface User {
    id: number,
    groups: number[],
    createLog: (id: number) => void,
    timestamp: number  // 属性名相同，类型不同
}
// 2. 合并(&)两个接口类型

let t: Admin & User
t!.administrator // 合法 Admin.administrator: string
t!.groups        // 合法 User.groups: number[]
t!.id            // 合法 id: number
t!.timestamp     // 合法 timestamp: never

// timestamp 属性名相同，类型不同，不可被赋值
// t.timestamp = 1  // 不能将类型“number”分配给类型“never”。ts(2322)


// 3. 合并两个传入对象的成员属性

// 3.1 定义合并方法
function extend<T, U>(first: T, second: U): T & U {
    for (const key in second) {
        (first as T & U)[key] = second[key] as any
    }
    return first as T & U
}

// 3.2 定义两个对象并合并
class Person {
    constructor(public name: string) { }
}
class ConsoleLogger {
    log() { }
}

let jim = extend(new Person('Jim'), new ConsoleLogger())
let n = jim.name
jim.log()