/**
 * 抽象类，内容要点：
 * 1. 定义语法：abstract class Animal {
 *                 abstract makeSound(): void
 *                    
 *                 move(): void {
 *                     console.log('roaming the earth...')
 *                 }
 *             }
 * 2. 关键字：abstract
 * 3. 抽象类作为其它派生类的基类使用。 它们只能被继承，不能被实例化
 * 4. 抽象类中定义的抽象方法，只需要声明方法签名，不包含具体实现，并且必须在 派生类 中实现
 * 5. 抽象方法和接口方法区别：
 *      a. 两者都是定义方法签名但不包含方法体
 *      b. 抽象方法必须包含 abstract 关键字并且可以包含访问修饰符。 接口方法不需包含 abstract 和 访问修饰符
 * 6. 抽象类好处：
 *      a. 抽离出事务的共性，有利于代码的复用和扩展。类似Go中接口的作用
 *      b. 抽象类可以实现多态：父类定义一个抽象方法声明，其子类就可以去实现这个方法，不同的子类可以有不同的实现，
 *         这样在程序运行过程中，根据不同的对象有不同的实现，形成了多态
 */

// 1. 定义抽象类 Department
abstract class Department {
    name: string

    constructor(name: string) {
        this.name = name
    }

    printName(): void {
        console.log('Department name: ' + this.name)
    }
    // 定义抽象方法，必须在派生类中实现
    abstract printMeeting(): void 
}

// 2. 子类 AccountingDepartment 继承抽象类 Department
class AccountingDepartment extends Department {
    constructor() {
        super('Accounting and Auditing') // 在派生类的构造函数中必须调用 super()
    }
    // 实现抽象类中的方法
    printMeeting(): void {
        console.log('The Accounting Department meets each Monday at 10am.')
    }

    generateReports(): void {
        console.log('Generating accounting reports...')
    }
}

let department: Department // 允许创建一个对抽象类型的引用
//   department = new Department() // 错误: 不能创建一个抽象类的实例

department = new AccountingDepartment() // 允许对一个抽象子类进行实例化和赋值
department.printName()
department.printMeeting()
//   department.generateReports() // 错误: 方法在声明的抽象类中不存在

// 3. 多态： 实现父类抽象方法

abstract class Animal3 {
    abstract sleep():void
}

class Dog3  extends Animal3 {
    sleep(){
        console.log('Dog sleep');
    }
}

class Cat2  extends Animal3 {
    sleep(){
        console.log('Cat sleep');
    }
}
let dog3= new Dog3()
let cat2= new Cat2()


let animals: Animal3[] = [dog3, cat2]

animals.forEach( animal =>{
    animal.sleep()  // 根据不同的对象有不同的实现，形成了多态
})


// 4. 多态： this多态，链式调用

class WorkFlow{
    step1() {
        return this
    }

    step2() {
        return this
    }
}

// 普通链式调用
new WorkFlow().step1().step2()

class MyFlow extends WorkFlow{
    myStep1() {
        return this
    }

    myStep2() {
        return this
    }
}

// this 多态
new MyFlow().myStep1().step1().step2().myStep2()