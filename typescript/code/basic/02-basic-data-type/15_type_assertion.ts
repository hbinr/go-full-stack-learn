/**
 * 类型断言 学习，内容要点：
 * 1. 定义语法一: <string> someValue  
 * 2. 定义语法二(推荐): someValue as string， 推荐是因为当你在 TypeScript 里使用 JSX 时，只有 as 语法断言是被允许的。
 * 3. 关键字: as, !
 * 4. 使用类型断言的前提是：你不知道某变量的具体类型(any 定义的)，想判断该变量是否符合自己的预期类型
 * 5. 类型断言好比其它语言里的类型转换，但是不进行特殊的数据检查和解构。更类似Go中的类型断言
 * 6. 类型断言没有运行时的影响，只是在编译阶段起作用。
 * 7. 好处：
 *      a. 类型断言可以增加代码的灵活性
 *      b. 在改造旧代码的时候非常有效
 *      c. 但是类型断言避免滥用，要对上下文有充足的预判，没有任何根据的类型断言会给代码带来安全隐患。
 * 8. Go中的类型断言，有 ok (bool)返回，如果断言失败，开发者还能搂底处理，但是ts中并没有这种机制，增加了心智负担
 */


let someValue: any = 'this is a string'

// 1. <> 断言  
let strVal: string = <string>someValue


// 2. as 断言 推荐
strVal = someValue as string

// 3. ! 非空断言
// 如果编译器不能够去除 null 或 undefined，可以使用非空断言 ! 手动去除。
function fixed(name: string | null): string {
    /**
     * postfix() 是一个嵌套函数，
     * 因为编译器无法去除嵌套函数的 null (除非是立即调用的函数表达式)
     * 所以 TypeScript 推断 name!.charAt(0) 的 name 可能为空。
     * @param epithet string
     * @returns string
     */
    function postfix(epithet: string) {
        return name!.charAt(0) + '.  the ' + epithet; // name 被断言为非空
    }
    // 而 name = name || "Bob" 这行代码已经明确了 name 不为空，所以可以直接给 name 断言为非空（ name!.charAt(0)）
    name = name || "Bob"
    return postfix("great")
}

console.log(fixed('hello')); // h.  the great

console.log(fixed(null));    // B.  the great ，并没有报 null 错误
