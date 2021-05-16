/**
 * 联合类型，内容要点：
 * 1. 定义语法: 类型1 | 类型2
 * 2. 关键字: |
 * 3. 联合类型是使用管道符 | 把多个类型连起来，表示它可能是这些类型中的其中一个。我们把 | 理解成 or，这样便于轻松记忆
 * 4. 联合类型和交叉类型区别：
 *      a. 联合类型表示取值为多种中的一种类型
 *      b. 交叉类型每次都是多个类型的合并类型
 * 5. 基础类型，字面量都可以应用联合类型
 * 6. 访问联合类型成员：如果一个值是联合类型，那么只能访问联合类型的共有属性或方法
 * 7. 可辨识联合（Discriminated Union），根据其不同的字符串字面量类型引导到不同的 case 分支 
 *
 */

//  1. 基础类型

// currentMonth 的值可以是 string 类型或者 number 类型中的一种
let currentMonth: string | number

currentMonth = 'February'
currentMonth = 2

// 2. 字面量类型  

// 类型别名 Scanned 可以是 true 或者 false 两种布尔字面量中的一种
type Scanned = true | false

// 类型别名 Result 可以是 { status: 200, data: object } 或者 { status: 500, request: string} 两个对象字面量中的一种。
type Result = { status: 200, data: object } | { status: 500, request: string }

// 3. 访问联合类型成员 (共有的)
interface Dog {
    name: string,
    eat: () => void,
    destroy: () => void
}

interface Cat {
    name: string,
    eat: () => void,
    climb: () => void
}

let pet: Dog | Cat
pet!.name    // OK
pet!.eat()   // OK
// pet!.climb() // Error  climb()是接口 Cat 独有的，无法访问

// 4. 可辨识联合案例，根据其不同的字符串字面量类型引导到不同的 case 分支 
interface Rectangle {
    type: 'rectangle',
    width: number,
    height: number
}

interface Circle {
    type: 'circle',
    radius: number
}

interface Parallelogram {
    type: 'parallelogram',
    bottom: number,
    height: number
}

function area(shape: Rectangle | Circle | Parallelogram) {
    switch (shape.type) {
        case 'rectangle':
            return shape.width * shape.height
        case 'circle':
            return Math.PI * Math.pow(shape.radius, 2)
        case 'parallelogram':
            return shape.bottom * shape.height
    }
}

let shape: Circle = {
    type: 'circle',
    radius: 10
}

console.log(area(shape))