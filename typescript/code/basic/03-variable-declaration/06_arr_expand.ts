/**
 * 数组展开，内容要点：
 * 1. 将数组中的元素一一展开，从左至右
 * 2. 展开操作底层是浅拷贝，原值不会被展开操作所改变
 */

let firstArr: number[] = [1, 2, 3]
let secondtArr: number[] = [4, 5, 6]

let bothPlusArr = [0,...firstArr,...secondtArr,7]
console.log('bothPlusArr: ', bothPlusArr); // bothPlusArr:  [ 0, 1, 2, 3, 4, 5, 6, 7]

bothPlusArr[2] = 20 
console.log('bothPlusArr: ', bothPlusArr); // bothPlusArr:  [ 0, 1, 20, 3, 4, 5, 6, 7]
console.log('firstArr: ', firstArr); // firstArr:  [ 1, 2, 3 ]  firstArr中的 firstArr[1]还是为2，并未改变
