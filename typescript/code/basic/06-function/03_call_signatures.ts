/**
 * 调用签名，内容要点：
 * 1. 定义语法：type DescribableFunction = {
 *                 description: string;         // 函数作用描述；是函数自身属性
 *                 (someArg: number): boolean;  // 注意没有 => ，更类似接口中函数类型定义
 *             };
 * 2. 关键字：type, (), {}
 * 3. 在JavaScript中， 函数除了可以被调用外，还可以有自己的属性。
 * 4. 与函数表达式不同的是，增加了 属性定义 + 没有箭头，使用 ':'
 *      a. 函数表达式：type addFunc4 = (x: number, y: number) => void; 
 *      b. 调用签名：  type addFunc5 = { addDescrip: string; (x: number, y: number): void} // 
 * 
 */

type DescribableFunction = {
    description: string         // 函数作用描述
    (someArg: number): boolean  // 函数调用签名定义
}

// 函数作为参数
function doSomething(fn:DescribableFunction) {
    // 使用函数签名 + 调用函数
    console.log(fn.description + 'returned ' + fn(1));
    
}