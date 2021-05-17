/**
 * null 和 undefined 学习，内容要点：
 * 1. 定义语法: let u: undefined = undefined ，let n: null = null
 * 2. 关键字: undefined   null
 * 3. null 表示为空，在Go中用nil表示。在实际开发中，常判断返回值或参数是否为null
 * 4. undefined 表示未定义
 * 5. 默认情况下 null 和 undefined 是所有类型的子类型。 就是说你可以把 null 和 undefined 赋值给 number 类型的变量。
 * 6. 编译ts代码时，添加 --strictNullChecks 参数，表示严格检查null，遇到null就会编译报错，这能规避很多问题
 */

let un: undefined = undefined
let nu: null = null

function testNull(name: string) {

}

//  testNull(null) // VSCode编辑器会报错: 类型“null”的参数不能赋给类型“string”的参数
                // 但是使用tsc xx.ts 却能通过。
                // 添加 --strictNullChecks 参数，即 tsc xx.ts --strictNullChecks 时就能检查null了，编译就报错了

