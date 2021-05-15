/**
 * 可索引类型 接口，内容要点：
 * 1. 定义语法：interface StringArray {
 *                [index: number]: string
 *             }
 * 2. 关键字：interface , []
 * 3. 可索引类型，类似map，根据key获取value，
 * 4. 可索引类型具有一个 索引签名，它描述了对象索引的类型，还有相应的索引返回值类型。 
 * 5. TypeScript 支持两种索引签名：
 *      a. 字符串， index 为字符串类型，   [index: string]: string，赋值时是 dict/map 形式
 *      b. 数字， index 为数字类型，  [index: number]: string， 赋值时是数组形式
 *      c. 可同时使用这两种类型，但是数字索引的返回值必须是字符串索引返回值类型的子类型。 
 *         这是因为当使用 number 来索引时，JavaScript 会将它转换成string 然后再去索引对象。
 * 6. readonly 修饰，将索引签名设置为只读，这样就防止了给索引赋值
 */

// 1. 可索引类型定义--数字索引
interface StringArray {
    // 索引key为  number 类型，索引value为 string类型
    [idx: number]: string
}

let myStrArr: StringArray

myStrArr = ['Tom', 'Bob']  // 赋值时是数组形式

console.log('myStrArr[0]', myStrArr[0]); // 索引为整型 0

// 2. 可索引类型定义--字符串索引
interface StringArray2{
    // 索引key为  string 类型，索引value为 string类型
    [idx: string]: string
}
let myStrArr2: StringArray2

myStrArr2 = {key1: 'Tom', key2: 'Bob'}  // 赋值时是dict形式

console.log("myStrArr2['key1']: ", myStrArr2['key1']); // myStrArr2['key1']:  Tom  ，注意：索引字符串类型


// 3. 数组索引和字符串索引共存，数字索引的返回值必须是字符串索引返回值类型的子类型
class Animal{
    name: string

    //  需要写构造器初识化属性
    constructor(name:string){
        this.name = name
    }
}

//  Dog 是 Animal的子类
class Dog extends Animal{
    breed: string
    
    //  需要写构造器初识化属性
    constructor(name:string,breed:string){
        super(name) 
        this.breed = breed
    }
}

// 错误：使用数值型的字符串索引，有时会得到完全不同的Animal!
// interface NotOkay{
//     [index: number]: Animal
//     [index: string]: Dog
// }

// 正确 
interface Okay{
    [index: number]: Dog
    [index: string]: Animal
}

// 4. 字符串索引，描述 dict 模式 ，确保所有属性与其返回值类型相匹配。
interface NumberDict{
    [index: string]: number
    length: number     //可以，length是number类型
    // name: string       // 错误，`name`的类型与索引类型返回值的类型不匹配
}

// 5. 只读索引
interface ReadOnlyStringArray{
    readonly [index: number]: string 
}


let myReadOnlyStringArr: ReadOnlyStringArray

myReadOnlyStringArr = ['tom','bob']

// myReadOnlyStringArr[0] = 'ronger' // 无法修改

