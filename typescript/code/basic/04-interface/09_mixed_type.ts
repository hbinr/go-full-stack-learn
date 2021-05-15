/**
 *  混合类型 接口，内容要点：
 * 1. 定义语法：interface A extend B {
 *             }
 * 2. 关键字：interface , [], ()
 * 3. 混合类型接口，表示同一个接口中定义多种类型， 因为 JavaScript 其动态灵活的特点，一个对象可以同时具有上面提到的多种类型。
 * 
 */

// 1. 一个对象可以同时做为函数和对象使用，并带有额外的属性。
interface Counter{
    // 函数类型
    (start: number): string

    // 属性类型，带有额外的属性
    internal: string
    

    // 类类型，带有方法
    reset():void
}

// 2. 创建 Counter
function createCounter(): Counter {
    // 1. 初始化函数类型
    let counter = (function (start: number): string {
        console.log("counter start",start);
        return ''
    }) as Counter

    // 2. 初始化私有属性
    counter.internal = 'hello'
    // 3. 初始化 接口方法
    counter.reset = function():void{
        console.log("counter reset: ", counter.internal);
    }

    return counter
}


// 3. 使用Counter

let counter: Counter = createCounter()
counter(1)
counter.internal += ' world'
counter.reset()


