namespace Shape {
    // 我们可以定义任意个变量，但是只在 Shape这个命名空间下可见
    const PI = Math.PI

    // 如果想让命名空间的成员在全局内可用的话，就通过export导出
    export function circle(r: number) {
        return PI * r ** 2
    }
}

Shape.circle(1)
Shape.square(1)

// 为了方便函数调用，我们可以给命名空间的成员起别名
// 注意不要和模块中的import混淆，这里只是起别名
import circle = Shape.circle
circle(1)  // 等效于：Shape.circle(1)