/**
 * never学习，内容要点：
 * 1. 定义语法: 
 * 2. 关键字: never
 * 3. never 类型表示的是那些永不存在的值的类型，或者说表示哪些永远没有返回值的类型
 * 4. 使用场景：
 *  a.函数返回异常: function err(msg: string): never{ throw new Error(msg)}
 *  b.死循环: function endless(){ while (true){} } 
 * 5. 值的注意的是：
 *  a.never 类型是任何类型的子类型，也可以赋值给任何类型
 *  b.然而，没有类型是 never 的子类型或可以赋值给never 类型（除了 never 本身之外）。 
 *  c.即使 any 也不可以赋值给 never。
 * 6. TypeScript中的never类型具体有什么用？ https://www.zhihu.com/question/354601204
 */


// 1. 没有任何返回值，只会抛出异常
function error(msg:string): never {
    throw new Error(msg)
}

// 2. 死循环
function endLess(): never {
    while (true){}
}


// 3.

interface Foo {
    type: 'foo'
}
  
interface Bar {
    type: 'bar'
}

type All = Foo | Bar
//  type All = Foo | Bar | Baz // 如果有一天多了 Baz

function handleValue(val: All) {
    switch (val.type) {
        case 'foo':
        // 这里 val 被收窄为 Foo
        break
        case 'bar':
        // val 在这里是 Bar
        break
        default:
        // val 在这里是 never
        const exhaustiveCheck: never = val
        break
    }
}

/**
 * 然而他忘记了在 handleValue 里面加上针对 Baz 的处理逻辑，
 * 这个时候在 default branch 里面 val 会被收窄为 Baz，导致无法赋值给 never，产生一个编译错误。
 * 所以通过这个办法，你可以   
 * 重点：确保 handleValue 总是穷尽 (exhaust) 了所有 All 的可能类型。
 */