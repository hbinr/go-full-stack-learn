/**
 * 函数形参解构机制及其默认值，内容要点
 * 1. 形参是对象或元组时，也支持解构，省略 “赋值”过程
 * 2. 形参为对象时，可以设定默认值
 *      a.普通默认值: 参数定义时，便初识化了默认值
 *      b.解构默认值: 在解构属性上给予一个默认或可选的属性用来替换主初始化列表
 * 3. 注意入参为空和空对象的不同
 * 4. 解构表达式要尽量保持小而简单
 */

// 1. 函数参数为元组时，解构示例
function f([a, b]:[number, number]){
    console.log('a, b: ', a, b);
}

let input:[number, number] = [1, 2]
f(input)


// 2.  形参为对象，设定默认值
// 现在，即使 b 为 undefined ， keepWholeObject 函数的变量 wholeObject 的属性 a 和 b 都会有值。
function keepWholeObject(wholeObject: {a: string, b?:number}){
    let {a,b = 1001} = wholeObject
    console.log('a,b : ', a,b );
}
let testObj = {
    a: "str" // 只有字段a
}

keepWholeObject(testObj) // a,b :  str 1001

// 3. 解构也能用于函数声明
type C = { a: string, b?: number }

function cTest({a, b}: C): void{
    console.log('a, b: ', a, b);
}
cTest(testObj)  // a, b:  str undefined


// 4. 对象解构，普通默认值和解构默认值
// 普通默认值：参数定义时，便初识化了默认值
function deconObjectDefaultVal({ a = 'default val', b = 0} = {}) : void{
    console.log('a, b: ', a, b);
}
deconObjectDefaultVal()  // a, b:  default val 0，入参为空，a,b 都是使用了普通默认值


// 解构默认值：在解构属性上给予一个默认或可选的属性用来替换主初始化列表
function deconObjectDefaultVal2({ a, b = 0} = {a: "decon default value"}) : void{
    console.log('a, b: ', a, b);
}
deconObjectDefaultVal2()  // a, b:  decon default value 0，入参为空，a的值是解构默认值， b则是普通默认值

deconObjectDefaultVal2({a: 'yes'})  // a, b:  yes 0  ， a的值为实参，b的值为普通默认值

// deconObjectDefaultVal2({}) // 编译报错：一旦传入参数则 a 是必须的 


