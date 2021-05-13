/**
 * 元组 tuple 类型学习，内容要点：
 * 1. 定义语法: [数据类型1, 数据类型2 ...] eg: let x: [string, number]  = ["hello", 1]
 * 2. 关键字: tuple ，Go中没有此概念，但是和Python中的元组tuple基本一致
 * 3. 元组类型允许表示一个已知元素数量和类型的数组，各元素的类型不必相同
 * 4. 元组本质是一个数组，只不过支持多种数据类型的元素。
 * 5. 元组和数组的区别：
 *  a. 元组数据类型可以是多个；而数组的数据类型只能是单一的
 *  b. 元组的长度在定义时就已经确定，越界会报错(ts 3.1之后)；而数组并未确定，类似Go切片
 * 6. 注意：
 *  a.自从 TyeScript 3.1 版本之后，访问越界元素会报错，我们不应该再使用该特性。
 *  b.元组定时的数据类型顺序和实际赋值的顺序必须保持一致，数量也必须保持一致
 */

// 1. 定义
let tuple: [string,number] = ['hello', 1]

// tuple = [1, 'world'] // 报错: 数据类型的顺序必须一致
// tuple = ['hi',1,2] // 报错: 不能将类型“[string, number, number]”分配给类型“[string, number]”。源具有 3 个元素，但目标仅允许 2 个


// 2.数据操作
console.log(tuple[0]);
// console.log(tuple[2]); // 报错：访问越界


// 3.遍历元组
for(let i = 0; i < tuple.length; i++){
    console.log(`tuple[${i}]: `, tuple[i]);

}

tuple.forEach(element => {
    console.log('element: ', element);
});