/**
 * 可选属性，内容要点：
 * 1. 定义语法：interface{ name?:string } , name属性就是可选的，可以不传递或者不使用
 * 2. 关键字：无, 通过 '?' 来定义
 * 3. 接口里的属性不全都是必需的。 有些是只在某些条件下存在，或者根本不存在。例如给函数传入的参数对象中只有部分属性赋值了。
 * 4. 好处之一是可以对可能存在的属性进行预定义
 * 5. 好处之二是可以捕获引用了不存在的属性时的错误。
 */

interface Square{
    color: string,
    area: number
}

interface SquareConfig{
    color?: string, // color 可选属性
    width?: number  // width 可选属性
}

function createSquare(config:SquareConfig): Square{
    let newSquare = {color: "white", area: 100}
    // 如果color属性存在，则重新赋值
    if (config.color) {
        newSquare.color = config.color
    }
    // 如果width属性存在，则重新计算square的面积 area
    if (config.width) {
        newSquare.area = config.width * config.width
    }

    return newSquare
}


let mySquare = createSquare({color: 'black'})
console.log('mySquare: ', mySquare); //  mySquare:  { color: 'black', area: 100 }
 