/**
 * 字面量类型学习，内容要点：
 * 1. 定义语法:   类型名 : 数字|字符串|true|false，
 *               声明字面量类型，注意这里是 : 不是 =。 = 等号是变量赋值，: 表示声明的类型。
 * 2. 关键字: :, | 
 * 3. 字面量（literal）是用于表达源代码中一个固定值的表示法（notation）。
 * 4. 通俗的讲，字面量也可以叫直接量，就是你看见什么，它就是什么。
 * 5. 字面量应用
 *      a. 单一值。这种情况并不怎么有用，就和 const没有什么区别了。let s: 'string' <=> const s = 'string'
 *      b. 联合其他类型。这才是字面量作用最大的地方，不仅可以联合字面量，还可以联合其他类型，比如对象类型接口
 * 6. boolean 底层就是 联合字面量 true | false 的别名
 */

// 1. 单一值字面量定义

let x: 'hello'

// x = 'world' // 报错：不能将类型“"world"”分配给类型“"hello"”。ts(2322)

// 2. 联合字面量定义

// 2.1 boolean 联合字面量
let isOK: true | false

isOK = true
function exists(exists: true | false) {
    if (exists) {
        console.log('isOK: ', isOK);
    }
}

// 2.3 数字联合字面量
function compare(a: string, b: string): -1 | 0 | 1 {
    return a === b ? 0 : a > b ? 1 : -1;
}

// 2.3 复杂联合字面量
interface Options {
    width: number
}

function configure(x: Options | 'auto') {
    console.log('x: ', x);
}

configure({ width: 100 });
configure("auto");
// configure("automatic"); // 报错：类型“"automatic"”的参数不能赋给类型“Options | "auto"”的参数。ts(2345)


// 3. boolean 字面量类型复杂应用
// 3.1 声明字面量
let success: true
let fail: false
let value: true | false // 等效于boolean

// 3.2 接口的返回值，会有正确返回和异常两种情况，这两种情况要有不同的数据返回格式：
type MyResponse = {
    success: true,
    code: number,
    data: object
} | {
    success: false
    code: number
    errMsg: string
}

let res: MyResponse = {
    success: false,
    code: 90001,
    errMsg: '该二维码已使用'
}

if (!res.success) {
    res.errMsg // OK
    // res.object // Error,类型“{ success: false; code: number; errMsg: string; }”上不存在属性“object”。ts(2339)
}