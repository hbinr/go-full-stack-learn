/**
 * const 声明常量，，内容要点：
 * 1. const 用于定义常量
 * 2. const 使用规范
 * 3. const 引入某个模块的某个对象时，用const定义
 * 
 * 注意：在node中(规范还是遵守ES6及之后)，const定义的数据并不是不能修改的，如下
 * 4. const 定义 对象 类型的常量，其属性是可以进行修改的，但整体不能修改
 * 5. const 定义 数组 类型的常量，其元素是可以进行修改的，但整体不能修改
 */

// 1. const 用于定义常量。修改就报错
const PI = 3.141592653
console.log('PI: ', PI);

// 2. const 定义常量，一般都会用全大写来书写，如果涉及多个单词，则用下划线 '_' 分隔

// 3. const 引入某个模块的某个对象时，用const定义。这种情况下，一般小写命名即可

// 4. const 对于对象型的常量，其属性(或者说字段)是可以进行修改的，但整体无法修改。Go中结构体是变量，不是常量
const people = {
    name: "tom",
    age: 18, // ES5、6 甚至更高，对象最后一个字段有没有 ','都可以，但是在Go中却必须用 ','来结尾 
}

people.age = 20  // 属性可以更改，因为people.age 指向的内存地址并没有改变，还是指向people整体所在的地址，所以可以修改属性
console.log('people.age: ', people.age);

// 注意：但是people本身是不允许修改的，people指向的内存地址会发生变化
// people = {
//     name: "bob", //  运行报错： TypeError: Assignment to constant variable.
//     age: 18,
// }


// 5. 数组型的常量中的每一项数据是可以修改的，但整体无法修改。Go中数组是变量，不是常量
const ARR = [10, 20, 30]
console.log('ARR[1]: ', ARR[1]);  // ARR[1]:  20
ARR[1] = 21
console.log('ARR[1]: ', ARR[1]);  // ARR[1]:  21

// ARR = [11, 20, 30]  // 运行报错：TypeError: Assignment to constant variable.