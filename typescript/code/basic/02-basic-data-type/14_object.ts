/**
 * object 非原始类型学习，内容要点：
 * 1. 定义语法: 
 * 2. 关键字: object
 * 3. 表示非原始类型，也就是除 number，string，boolean，symbol，null或undefined 之外的类型。
 * 4. 使用 object 类型，就可以更好的表示像 Object.create 这样的 API
 */

// 声明一个函数类型，入参为object或null，无返回值
declare function create(o: object): void;

create({ prop: 0 }) // OK  { prop: 0 }是对象类型

// create(42) // 报错: number 为原始类型
// create('string') // 报错: string 为原始类型
// create(false) // 报错: boolean 为原始类型
// create(null) //报错: null 为原始类型
// create(undefined) // 报错:  undefined为原始类型