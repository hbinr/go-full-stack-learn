/**
 * 函数类型 接口，内容要点：
 * 1. 定义语法：interface SearchFunc {
 *                  (source: string, subString: string): boolean  // 函数签名
 *              }
 * 2. 关键字：interface, ()
 * 3. 函数类型，表示该接口是一个函数，只要实现了其中的函数定义，就能调用
 * 4. 函数类型就像是一个只有参数列表和返回值类型的函数定义。参数列表里的每个参数都需要名字和类型。
 * 5. 注意：
 *      a. 函数类型的类型检查来说，函数的参数名不需要与接口里定义的名字相匹配。 本质就是形参名，可以自定义
 *      b. 尽管参数名不会检查，但是参数类型+参数个数，这两个要始终和接口中的函数定义一致
 *      c. 使用接口中的函数定义时，可以不写返回值，因为ts会自动类型推断
 * 6. 和Go语言很大不同：
 *      a. ts 中，函数类型的接口，可以当作函数来直接调用
 *      b. Go 中，接口并没有函数类型的概念，是一系列方法的契约
 */

// 1. 定义函数类型的接口
interface SearchFunc{
    // 检查字符串是否存在
    (source: string, subString: string):boolean
}

// 2. 实现接口
let mySearch: SearchFunc

// 2.1 最简实现
mySearch = function (source: string, subString: string): boolean { 
    return source.search(subString) > -1
}

// 2.2 函数的参数名不需要与接口里定义的名字相匹配
mySearch = function (src: string, sub: string): boolean { 
    return src.search(sub) > -1
}


// 2.3 函数返回值可以不写
mySearch = function (src: string, sub: string){ 
    return src.search(sub) > -1
}

console.log(mySearch('tom','om'));
