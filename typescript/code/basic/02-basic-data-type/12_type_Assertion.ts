/**
 * 类型断言 学习，内容要点：
 * 1. 定义语法一: <string> someValue  
 * 2. 定义语法二(推荐): someValue as string， 推荐是因为当你在 TypeScript 里使用 JSX 时，只有 as 语法断言是被允许的。
 * 3. 关键字: as ，Go中也有类型断言，判断空接口类型的变量是否符合自己的预期
 * 4. 使用类型断言的前提是：你不知道某变量的具体类型(any 定义的)，想判断该变量是否符合自己的预期类型
 * 5. 类型断言好比其它语言里的类型转换，但是不进行特殊的数据检查和解构。更类似Go中的类型断言
 * 6. 类型断言没有运行时的影响，只是在编译阶段起作用。
 */

let someValue: any = 'this is a string'

let strVal: string = <string> someValue

strVal = someValue as string