/**
 * 泛型约束，内容要点：
 * 1. 定义语法: loggingIdentity<T extends Lengthwise>(arg: T): T,  Lengthwise为接口
 * 2. 关键字: interface, <>, T, extends
 * 3. 泛型约束， 对于 T 进行约束。让它具有某种行为或属性
 * 4. 比如定义一个接口来描述约束条件，创建一个包含 .length 属性的接口，使用这个接口和 extends 关键字来实现约束
 * 5. 在泛型约束中使用类型参数
 * 
 */
// 1. 未使用泛型约束
function loggingIdentity3<T>(arg: T): T {
    // console.log(arg.length); // 报错：类型“T”上不存在属性“length”。ts(2339)

    return arg
}

// 2. 定义约束接口
interface Lengthwise {
    length: number
}

// 3. 定义泛型函数，增加约束
function loggingIdentity4<T extends Lengthwise>(arg: T): T {
    console.log(arg.length);
    return arg
}


// loggingIdentity4(3) // 报错

loggingIdentity4([1, 2, 3]) // ok
loggingIdentity4('string')  // ok

// 4. 在泛型约束中使用类型参数
// prop 函数的作用，该函数用于获取某个对象中指定属性的属性值。
// 因此我们期望用户输入的属性是对象上已存在的属性
function prop<T extends object, K extends keyof T>(obj: T, key: K) {
    // 上述函数定义等效于：function prop<T, K extends keyof T>(obj: T, key: K) {
    return obj[key];
}

/**
 * 首先定义了 T 类型并使用 extends 关键字约束该类型必须是 object 类型的子类型，
 * 然后使用 keyof 操作符获取 T 类型的所有键，其返回类型是联合类型，
 * 最后利用 extends 关键字约束 K 类型必须为 keyof T 联合类型的子类型。
 */

type Todo = {
    id: number;
    text: string;
    done: boolean;
}

const todo: Todo = {
    id: 1,
    text: "Learn TypeScript keyof",
    done: false
}

const id = prop(todo, "id");
console.log('id: ', id); // id:  1

const text = prop(todo, "text");
console.log('text: ', text); // text:  Learn TypeScript keyof


const done = prop(todo, "done");
console.log('done: ', done); // done:  false
