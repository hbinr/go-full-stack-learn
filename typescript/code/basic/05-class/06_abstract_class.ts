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
 * 3. 抽象类作为其它派生类的基类使用。 它们一般不会直接被实例化
 * 4. 抽象类中定义的抽象方法，只需要声明方法签名，不包含具体实现，并且必须在 派生类 中实现
 * 5. 抽象方法和接口方法区别：
 *      a. 两者都是定义方法签名但不包含方法体
 *      b. 抽象方法必须包含 abstract 关键字并且可以包含访问修饰符。 接口方法不需包含 abstract 和 访问修饰符
 * 
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