/**
 * bigint 类型学习，内容要点：
 * 1. 定义语法:  let one: bigint = 1n 或 BigInt(1)
 *              在一个整数字面量后加 n 的方式定义一个 BigInt 或者 调用函数 BigInt()
 * 2. 关键字为:  bigint,  BigInt(), n
 * 3. JavaScript 中可以用 Number 表示的最大整数为 2^53 - 1，可以写为 Number.MAX_SAFE_INTEGER。
 * 4. bigint 数据类型是用来表示那些已经超出了 number 类型最大值的整数值，
 *     对于总是被诟病的整数溢出问题，使用了 bigint 后将完美解决
 * 5. BigInt 与 Number 的不同点：
 *      a. BigInt 不能用于 Math 对象中的方法。
 *      b. BigInt 不能和任何 Number 实例混合运算，两者必须转换成同一种类型。
 *      c. BigInt 变量在转换为 Number 变量时可能会丢失精度。 大转小 ，丢失精度
 *      d. BigInt 类型就是用来表示那些已经超出了 number 类型最大值的整数值，
 *         也就是这个容器还没满，在此基础上加上两个不同的值，其结果不相等。而number则会相等
 * 6. 使用 typeof 检测类型时，BigInt 对象返回 bigint       
 * 7. 应用场景：
 *      a. 不要在 number 和 bigint 两种类型之间进行相互转换
 *      b. 仅在值可能大于 2^53 - 1 时使用 BigInt
 *     
 */

// 1. bigint 类型定义
const theBiggestInt: bigint = 9007199254740991n
const alsoHuge: bigint = BigInt(9007199254740991)
const hugeString: bigint = BigInt("9007199254740991")

theBiggestInt === alsoHuge // true
theBiggestInt === hugeString // true

// 2.  number 最大值，最大精度就是这个容器已经完全满了，无论往上加多少都会溢出，所以这两个值是相等的。

const biggest: number = Number.MAX_SAFE_INTEGER

const biggest1: number = biggest + 1
const biggest2: number = biggest + 2

biggest1 === biggest2 // true 超过精度

// 3. bigint 达到 number 最大精度，再加也不会相等
const biggest11: bigint = BigInt(Number.MAX_SAFE_INTEGER)

const biggest3: bigint = biggest11 + 1n
const biggest4: bigint = biggest11 + 2n

biggest3 === biggest4 // false

// 4. 使用 typeof 检测类型时，BigInt 对象返回 bigint   
typeof 10n === 'bigint'         // true
typeof BigInt(10) === 'bigint'  // true