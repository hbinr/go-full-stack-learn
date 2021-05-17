/**
 * 字符串类型学习，内容要点：
 * 1. 定义语法 let str: string = "hello"
 * 2. 关键字为 string， 在Go中也为string
 * 3. 可以使用双引号 "" 或单引号 '' 表示字符串
 * 4. 模版字符串，可以定义多行文本和内嵌表达式， 这种字符串是被反引号包围（ `），并且以 ${ expr } 这种形式嵌入表达式
 * 
 */

// 1. 字符串定义
let hello: string = "hello"  // 双引号
let world: string = 'world'  // 单引号 
 

// 2. 模板字符串

let language: string = 'TypeScript'

// 支持内嵌表达式
let sayHi: string = `Hi ${language}` // 编译后：底层就是字符串的拼接

// 支持多行+内嵌表达式
let mutilLineStr: string = `
   Life is short,
   I use ${language}
`
 