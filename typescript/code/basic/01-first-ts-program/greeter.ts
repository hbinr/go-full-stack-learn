// 接口定义
interface Person{
    firstName: string  // : string 表示类型注解，string类型
    lastName: string
    fullName: string
}


// 类定义
class User {
    firstName: string 
    lastName: string
    fullName: string
    constructor(firstName:string,lastName:string) {
        this.firstName = firstName
        this.lastName = lastName
        this.fullName = firstName + ' ' + lastName
    }
}

// 方法
function greeter(person:Person) {
    return 'Hello' + person.fullName
}
// 调用方法

let user = new User('haoBin','duan')

greeter(user)
console.log('greeter(user): ', greeter(user));