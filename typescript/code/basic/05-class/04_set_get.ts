/**
 * set, get 存取器，内容要点：
 * 1. 定义语法：set(),  get()
 * 2. 关键字：set, get
 * 3. TypeScript 支持通过 getters/setters 来截取对对象成员的访问。 它能帮助你有效的控制对对象成员的访问。
 * 4. 如果类中定义了set(), get()，也必须定义构造函数。 Java 定义了 get, set便不用定义构造函数了
 * 5. 使用注意事项：
 *      a. 首先，存取器要求你将编译器设置为输出 ECMAScript 5 或更高。 不支持降级到 ECMAScript 3
 *         编译时，增加编译参数： --target es5
 *      b. 其次，只带有 get 不带有 set 的存取器自动被推断为 readonly。
 *         这在从代码生成 .d.ts 文件时是有帮助的，因为利用这个属性的用户会看到不允许够改变它的值。
 *     
 */

// 1. 没有 使用存取器的案例
class Employee {
    fullName: string
    constructor(fullName: string){
        this.fullName = fullName
    }
}

let employee = new Employee('Tom')
employee.fullName = 'Bob Smith'

if (employee.fullName) {
    console.log(employee.fullName)
}

// 2. 使用存取器的案例，截取对对象成员的访问

let password = 'secret passcode'

class Employee2 {
    private _fullName: string

    constructor(_fullName: string){
        this._fullName = _fullName
    }
    
    get fullName(): string{
        return this._fullName
    }

    set fullName(newName: string){
        if (password && password === 'secret passcode') {
            this._fullName = newName
        }else{
            console.log('Error: Unauthorized update of employee!');
        }
    }
}


let employee2 = new Employee2('Tom')
employee2.fullName = 'Bob Smith'  // 如果修改了 password，那么便不能修改成功，还是Tom

if (employee2.fullName) {
  console.log(employee2.fullName)
}