/**
 * 类型推断，内容要点：
 * 1. 类型推断：TypeScript 里，在有些没有明确指出类型的地方，类型推断会根据某些规则自动地为其推断出一个类型
 * 2. 类型推断之基础类型推断：变量定义，函数定义时，未指定数据类型， 从右向左
 * 3. 类型推断之最佳通用类型： 从多个类型中选出一个最合适的类型，兼容性最好的类型
 * 4. 类型推断之上下文类型： 根据程序的上下文来推断类型。 从左向右
 *    上下文类型的出现和表达式的类型以及所处的位置相关，常见的上下文：
 *      a. 参数、赋值的右侧
 *      b. 类型断言
 *      c. 对象和数组字面的成员
 *      d. 返回语句
 * 5. 上下文类型也作为最佳通用类型中的一个候选类型
 * 6. 如果我们明确理解上下文，能够确定数据类型，可以通过类型断言来显示指定数据类型
 */

// 1. 基础类型推断————从右向左
let a = 1
let arr = [1, 2, 3]
let str = 'this is a string'

let fn = (x = 1) => x + 1   // 入参和出参都被推断为 number

// 2. 通用类型推断
class Animal {
    numLegs: number
}

class Bee extends Animal {
}

class Lion extends Animal {
}

let zoo = [new Bee(), new Lion()]  // zoo会被推断为  (Bee | Lion)[] 联合类型

// 我们想让 zoo 被推断为 Animal[] 类型，但是这个数组里没有对象是 Animal 类型的，因此不能推断出这个结果。
// 为了更正，我们可以明确的声明我们期望的类型
let zoo2: Animal[] = [new Bee(), new Lion()]

// 3. 上下文类型推断————从左向右。  以下官方案例有问题， tsc -v 4.2.4  并没有报错

// TypeScript 类型检查器使用 window.onmousedown 函数的类型来推断右边函数表达式的类型。
window.onmousedown = (event) => {  // event: any 推断出 event 为 any 类型
    console.log(event.clickTime)
}


window.onmousedown = function (mouseEvent) {
    console.log(mouseEvent.button); //<- OK
    console.log(mouseEvent.kangaroo); //<- Error!   tsc -v 4.2.4  并没有报错
};


// 4. 使用类型断言指定数据类型

interface Foo {
    bar: number
}

let foo = {} as Foo // 通过类型断言，把  foo 指定为接口类型 Foo
foo.bar = 1


// 类型断言指定数据类型是不能滥用的，如果 foo.bar 并没有赋值，也不会报错，但是这就不符合接口的契约性了
// 没有按照接口的严格约定，给 foo 赋值 bar 属性的值
let foo2 = {} as Foo

// 为了解决上述 不遵循接口约定 的问题，我们可以在定义 foo 的时候就指定其数据类型

let foo3: Foo = {
    bar: 1
}