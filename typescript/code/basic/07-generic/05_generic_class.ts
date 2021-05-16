/**
 * 泛型类，内容要点：
 * 1. 定义语法: class GenericNumber<T>{ x:T}
 * 2. 关键字: class, <>, T,
 * 3. 与接口一样，直接把泛型类型放在类后面，可以帮助我们确认类的所有属性都在使用相同的类型。
 * 4. 类有两部分：静态部分和实例部分。 泛型类指的是实例部分的类型，所以类的静态属性不能使用这个泛型类型。
 *
 */

// 1. 定义泛型类
class GenericNum<T>{
    zeroVal: T
    add: (x: T, y: T) => T
}

let myGenericNum = new GenericNum<number>()
myGenericNum.zeroVal = 0
myGenericNum.add = (x, y) => {
    return x + y
}

myGenericNum.add(1, 2)
console.log('myGenericNum.add(1, 2): ', myGenericNum.add(1, 2));

// 2. 类的静态属性不能使用泛型

class GenericNum2<T>{
    // static zeroVal: T // 报错： 静态成员不能引用类类型参数。ts(2302)

    add: (x: T, y: T) => T
}
