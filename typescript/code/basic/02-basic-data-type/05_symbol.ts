/**
 * symbol 类型学习，内容要点：
 * 1. 定义语法:  Symbol([description])
 *          Symbol() 函数会返回 symbol 类型的值。
 * 2. 关键字为:  symbol, Symbol()
 * 3. 唯一性：每个从 Symbol() 返回的 symbol 值都是唯一的。
 * 4. symbol 需要在 es2015甚至更高的版本才能使用，编译时加: --target es2015
 * 5. 使用场景：
 *      a. 当一个对象有较多属性时（往往分布在不同文件中由模块组合而成），
 *      很容易将某个属性名覆盖掉，使用 Symbol 值可以避免这一现象
 *      b. 判断是否可以用 for...of 迭代
 *      c. symbol描述。Symbol([description]) 中可选的字符串即为这个 Symbol 的描述
 *         TIPS： description 属性是 ES2019 的新标准，Node.js 最低支持版本 11.0.0
 */

// 1. symbol 定义

let sym1: symbol = Symbol()
let sym2: symbol = Symbol()
let sym3: symbol = Symbol()

// 每个 Symbol() 方法返回的值都是唯一的，所以，sym2 和 sym3 不相等。
console.log(sym2 === sym3) // false


// 2. 判断是否可以用 for...of 迭代
let iterable: number[] = [1, 2, 3]
if (Symbol.iterator in iterable) {
    for (let n of iterable) {
        console.log(n) // 1, 2, 3
    }
}

// 3. symbol描述。 Symbol.prototype.description
// 编译时，需增加参数 --target es2019

const sym: symbol = Symbol('imooc')

console.log(sym);               // Symbol(imooc)
console.log(sym.toString());    // Symbol(imooc)
console.log(sym.description);   // imooc