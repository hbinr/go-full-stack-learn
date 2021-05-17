/**
 * 任意类型 any 类型学习，内容要点：
 * 1. 定义语法: let x:any = 1
 * 2. 关键字: any，在Go中可以通过空接口 interface{} 来指定任意类型
 * 3. any 类型更多是用在编程阶段还不清楚类型的变量，通过 any 指定该变量可以是任意类型。
 * 4. ts编译器在编译时不会检查 any 类型的变量和其相关操作。
 * 5. 使用场景：一般any用在动态的内容，比如来自用户输入或第三方代码库，并不确定真正的类型，
 *             我们不希望类型检查器对这些值进行检查，而是直接让它们通过编译阶段的检查
 * 6. 并不建议大规模使用 any，否则就相当于直接使用 JavaScript 了，没有强类型的优势了
 * 
 */

let notSure: any = 1 // notSure 默认推导为number类型
notSure = 'string'   // notSure 可以为string类型
notSure = true       // notSure 也可以为boolean类型



let arr1: number[] = [1, 2, 3]
let book = {
    name: 'tom'
}
// 包含任意类型的数组
let arrAny: any[] = [1, 'string', true, arr1, book]