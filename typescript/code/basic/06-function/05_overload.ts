/**
 * 函数重载，内容要点：
 * 1. 定义语法：和函数定义一致
 * 2. 关键字： function
 * 3. 函数重载：函数名一模一样，但是参数类型、参数数量、返回值类型不一样
 * 
 */


function add(...rest: number[]): number

function add(...rest: string[]): string

function add(...rest: any[]): any{
    let first = rest[0]

    if (typeof first === 'string') {
        return rest.join('')
    }

    if (typeof first === 'number') {
        return rest.reduce((pre, cur) => pre + cur)
    }
}

