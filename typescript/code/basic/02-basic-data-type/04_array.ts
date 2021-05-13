/**
 * 数组类型学习，内容要点：
 * 1. 定义语法一: 数据类型后面加上[]  eg: let numArr number[] = [1,2,3]
 * 2. 定义语法二: 使用数组泛型，Array<数据类型>  eg: let numArr Array[number] = [1,2,3]
 * 3. 关键字: Array，在Go中定义数组，除了要指定数据类型之外，还必须指定数组长度，eg: numArr := [3]int{1,2,3}
 * 4. 注意：新增元素，不会报数据越界；这点需要注意，ts的数组和Go中切片slice最类似
 * 
 */

// 1. 定义方式一：数据类型[]
let numArr: number[] = [1,2,3] 
 

// 2. 定义方式一：Array<数据类型>

let numArr2: Array<string> = ['one','two','three']

// 3. 元素操作
let one: number = numArr[0] // 获取元素
console.log('one: ', one); // one: 1

numArr[2] = 33 // 修改元素
console.log('numArr: ', numArr); // numArr:  [ 1, 2, 33 ]

numArr[3] = 10 // 新增元素，不会报数据越界，这点需要注意
console.log('numArr: ', numArr); // numArr:  [ 1, 2, 33, 10 ]

// 4. 遍历数组

for(let i = 0; i < numArr.length;i++){  // 普通for遍历
   console.log(`numArr[${i}]: `, numArr[i]);
}


numArr.forEach(num => { // forEach遍历
   console.log('num: ', num);
});
