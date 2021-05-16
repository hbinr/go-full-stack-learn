/**
 * 泛型，内容要点：
 * 1. 定义语法：function identity<T>(value: T): T{}
 * 2. 关键字：<>, T
 * 3. 泛型： 不预先确定的数据类型，具体的类型在使用的时候才能确定
 * 4. 泛型和any的区别：
 *      a. any 可以接收任何类型，传入的类型与返回的类型是任意的，不必保持一致。但是我们实践过程中是希望一致
 *      b. 泛型能够使 传入参数的类型 与返回值的类型 是相同的
 * 5. 泛型的使用：
 *      a. 指定具体类型，eg: let identity = identity<string>('myString')
 *      b. [更普遍]不指定具体类型，由 ts 自行类型推断，即编译器会根据传入的参数自动地帮助我们确定 T 的类型
 *         eg: let identity = identity('myString')
 *      c. 类型推论帮助我们保持代码精简和高可读性。
 *         如果编译器不能够自动地推断出类型的话，只能像上面那样明确的传入 T 的类型，在一些复杂的情况下，这是可能出现的。
 * 6. 泛型的好处：
 *      a. 函数和类可以轻松的支持多种数据类型，增强程序的扩展性
 *      b. 不必写多条函数重载，冗长的联合类型声明，增强代码可读性
 *      c. 灵活控制类型之间的约束
 */

// 1. 泛型函数定义
function identity<T>(value: T): T {
    return value
}


// 2. 泛型函数调用

// 2.1 明确指定数据类型
let output = identity<string>('myString')
console.log('output: ', output);

// 2.2 类型推断
let output2 = identity('myString')
console.log('output2: ', output2);