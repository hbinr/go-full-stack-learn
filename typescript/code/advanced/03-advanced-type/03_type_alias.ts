/**
 * 类型别名，内容要点：
 * 1. 定义语法: type 类型名 = 某类型或字面量， eg：type used = true | false
 * 2. 关键字: type
 * 3. 类型别名，通过关键字 type 给类型起个别名，类型别名较多应用于联合类型、交叉类型这种复合类型
 * 4. 类型别名会给类型起个新名字。类型别名有时和接口很像，但是可以作用于原始值，联合类型，元组以及其它任何你需要手写的类型
 * 5. 同接口一样，类型别名也可以是泛型
 * 6. 注意：类型别名不会新建一个类型，而是创建一个新名字来引用此类型。
 * 7. 类型别名和接口的区别：
 *      a. 接口可以实现 extends 和 implements，类型别名不行。
 *      b. 类型别名并不会创建新类型，是对原有类型的引用，而接口会定义一个新类型。
 *      c. 接口只能用于定义对象类型、函数类型、可索引类型，而类型别名的声明方式除了对象之外还可以定义交叉、联合、原始类型等
 * 8. 类型别名是最初 TypeScript 做类型约束的主要形式，后来引入接口之后，TypeScript 推荐我们尽可能的使用接口来规范我们的代码。
 * 9. Go中也有类型别名:
 *          type myInt32 int32    // 类型定义, myInt32定义为 int32
 *          type myInt64 = int32  // 类型别名, int64起个别名 myInt64, 区别只在是否有 '='
 */

// 1. 基础类型
type brand = string
type used = true | false

const str2: brand = 'imooc'
const state: used = true

// 2. 联合类型
type month = string | number

const currentMonth2: month = 'February'
const nextMonth: month = 3

// 3. 交叉类型
interface Admin {
    id: number,
    administrator: string,
    timestamp: string
}

interface User {
    id: number,
    groups: number[],
    createLog: (id: number) => void,
    timestamp: number
}

type T = Admin & User

// 4. 泛型，同接口一样，类型别名也可以是泛型
type Tree<T, U> = {
    left: T,
    right: U
}