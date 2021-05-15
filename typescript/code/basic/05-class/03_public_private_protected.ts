/**
 * 公有public, 私有private, 受保护protected，内容要点：
 * 1. 定义语法：public/private/protected  属性/方法
 * 2. 关键字：public/private/protected 
 * 3. public 公有的，ts中默认修饰符，可省略不写。类内、类外都可以访问
 * 4. private 私有的，只能类自身内部访问，类外部 和 子类 都不允许访问
 * 5. protected 受保护的，只能在类自身内部及其子类中访问，外部不允许访问
 * 6. Go中并没有类的概念，而是结构体 struct
 *     
 */

// 1. 私有属性定义
class Animal2{
    private name: string

    constructor(name: string){
        this.name = name
    }
    private eat(){}
}

let animal = new Animal2('animal')
// 类外部不能访问 私有属性
// animal.name   // 报错：属性“name”为私有属性，只能在类“Animal2”中访问。
// animal.eat()  // 报错：属性“eat”为私有属性，只能在类“Animal2”中访问。

class Dog2 extends Animal2{
    color: string
    constructor(name:string, color: string){
        super(name)
        this.color = color
    }
    printParentName(){
        // 子类不能访问父类的私有属性
        // console.log('this.name: ', this.name);   // 报错：属性“name”为私有属性，只能在类“Animal2”中访问。
    }
}
let dog = new Dog2('阿黄','yellow')
dog.color = 'black'
// 类外部不能访问 私有属性，即使是子类也不行
// dog.name   // 报错：属性“name”为私有属性，只能在类“Animal2”中访问。
// dog.eat()  // 报错：属性“eat”为私有属性，只能在类“Animal2”中访问。



// 2. 受保护属性定义
class Cat{
    protected name: string

    constructor(name: string){
        this.name = name
    }
    protected eat(){}
}

let cat = new Cat('cat')
// 外部不允许访问
// cat.name // 报错：属性“name”受保护，只能在类“Cat”及其子类中访问。ts(2445)

class PersianCat extends Cat {
    constructor(name: string) {
        super(name)
    }
    printInfo(){
        // 子类内部是允许访问的
        this.name = '波斯猫' //此处的 name 就是 Cat 类中的 name
        this.eat()  //此处的 ear() 就是 Cat 类中的 eat()
    }
}

let persianCat = new PersianCat('波斯猫')
// 类外部不能访问 受保护属性，即使是子类也不行
// persianCat.name // 报错：属性“name”受保护，只能在类“Cat”及其子类中访问。ts(2445)
