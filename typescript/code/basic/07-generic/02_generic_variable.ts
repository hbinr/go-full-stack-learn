/**
 * 泛型变量，内容要点：
 * 1. 定义语法：function identity<T>(value: T[]): T[]{}
 * 2. 关键字：<>, T, [] ...
 * 3. 泛型变量 T 当做类型的一部分使用，而不是整个类型
 *
 */

// 1. 使用错误的泛型
function loggingIdentity<T>(arg: T): T {
    // console.log(arg.length) // 报错：类型“T”上不存在属性“length”。ts(2339)
    return arg
}

// 2. 使用泛型变量
function loggingIdentity2<T>(arg: T[]): T[] { // 入参和出参都增加了 [] 
    console.log(arg.length)  // 正确。数组有 length 属性
    return arg
}

/**
 * 你可以这样理解 loggingIdentity 的类型：
 * 泛型函数 loggingIdentity，接收类型参数 T 和参数 arg，它是个元素类型是 T 的数组，并返回元素类型是T 的数组。
 * 如果我们传入数字数组，将返回一个数字数组，因为此时 T 的的类型为 number。
 * 这可以让我们把泛型变量 T 当做类型的一部分使用，而不是整个类型，增加了灵活性。
 *
 */