/**
 * 继承接口 ，内容要点：
 * 1. 定义语法：interface A extend B {
 *             }
 * 2. 关键字：interface , extend
 * 3. 继承接口，和类一样，接口也可以相互继承。这让我们能够从一个接口里复制成员到另一个接口里,更灵活地将接口分割到可重用的模块里
 * 4. 继承接口时，属性和方法都会继承，一个接口可以继承一个或多个接口 
 * 5. 在Go中，接口之间是组合关系，没有继承概念 
 */

// 1. 单继承
interface Shape{
    color: string
}

// 定义长方形 Rectangle 接口继承形状 Shape 接口
interface Rectangle extends Shape{
    sideLength: number
}

let rectangle = {} as Rectangle

rectangle.color = 'black'
rectangle.sideLength = 10


// 2. 多继承，一个接口可以继承多个接口，创建出多个接口的合成接口。

interface PenStroke{
    penWidth: number
}

interface Rectangle2 extends Shape, PenStroke{
    sideLength: number
}

let rectangle2 = {} as Rectangle2

rectangle2.color = 'red'
rectangle2.sideLength = 10
rectangle2.penWidth = 8