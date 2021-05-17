/**
 * 空类型 void 学习，内容要点：
 * 1. 定义语法: let unusable: void = undefined 或者 function fn(): void{} 
 * 2. 关键字: void，类似Java中的void
 * 3. void 表示空类型，和any（任意类型）相反，void不表示任何类型
 * 4. void 函数返回值为空(没有返回值)时，可以使用void表示返回值
 * 
 */

// 函数返回值为空 
function noReturn():void {
    console.log("no result return");
}

// 声明一个 void 类型的变量没有什么大用，因为你只能为它赋予 undefined：
let unusable: void = undefined