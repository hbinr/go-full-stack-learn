/**
 * 枚举 enum 类型学习，内容要点：
 * 1. 定义语法: enum 枚举名{枚举值1,枚举值2.....}
 * 2. 关键字: enum，在Go中没有枚举的概念，更多是通过struct来实现
 * 3. enum 类型是对 JavaScript 标准数据类型的一个补充，使用枚举类型可以为一组数值赋予友好的名字。
 * 4. 默认情况下，从 0 开始为元素编号，且递增步长 1。当然，你可以手动指定成员的数值
 * 5. enum 数据操作
 */


enum VIP{
    GreenVIP, // 从0开始编号，每次递增1
    BlackVIP,
    RedVIP
}

// 1. 获取枚举值
let greenVIP:VIP = VIP.GreenVIP
console.log('green: ', greenVIP); // green: 0

// 2. 根据枚举值得到枚举名称
let greenVIPName: string = VIP[0] // 类似数组的操作
console.log('greenVIPName: ', greenVIPName);


// 3. 可以手动指定枚举值的编号
enum Color{
    Green = 101 ,
    Black = 103,
    Red = 105
}

console.log('Color: ', Color);
// 输出结果：
// Color:  {
//     '101': 'Green',
//     '103': 'Black',
//     '105': 'Red',
//     Green: 101,
//     Black: 103,
//     Red: 105
//   }
  

