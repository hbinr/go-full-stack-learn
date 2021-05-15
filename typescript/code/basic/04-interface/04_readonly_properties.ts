/**
 * 只读属性，内容要点：
 * 1. 定义语法：interface{ readonly x: number} 属性只读，
 * 2. 关键字：readonly
 * 3. 只读属性只能在对象刚刚创建的时候修改其值，也就是初识化后就不能再修改了
 * 4. readonly vs const 什么时候用哪个？最简单判断该用 readonly 还是 const 的方法是
 * 5. 看要把它做为变量使用还是做为一个属性。
 *      a. 若做为变量使用的话用 const，
 *      b. 若做为属性则使用 readonly。 属性一般是在 interface 或对象 中
 */

interface Point{
    readonly x: number
    readonly y: number
}

let p1: Point = {
    x:10,
    y:15
}

// p1.x = 20  // 报错: 无法分配到 "x" ，因为它是只读属性。ts(2540)
// p1.y = 25  // 报错: 无法分配到 "y" ，因为它是只读属性。ts(2540)
