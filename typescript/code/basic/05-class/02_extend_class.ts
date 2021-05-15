/**
 * 继承，内容要点：
 * 1. 定义语法：class A  extend B{}
 * 2. 关键字：extend
 * 3. TypeScript 中的类类似于Java，Python中的类，有属性(成员变量)、构造函数、方法(成员方法)
 * 4. 通过 new 关键字可创建对象实例
 * 5. 在构造函数里访问 this 的属性之前，我们 一定要调用 super()。
 * 6. Go中并没有类的概念，而是结构体 struct
 *     
 */

// 1. 创建父类
class People {
    name: string
    constructor(name: string) {
        this.name = name
    }

    say(){
        console.log('People say...');
    }
}

// 2. Student 类继承 People
class Student extends People{
    constructor(name: string){
        // 在构造函数里访问 this 的属性之前，我们 一定要调用 super()。 
        // 这个是 TypeScript 强制执行的一条重要规则。
        super(name)
    }

    // 重写父类方法
    say(){
        console.log('Student say...');
    }
}

// 3. 创建实例
let people2 = new People('Tom')
people2.say()

// 多态应用：student 为People类型，但对象实例 student 为 Student，因此调用 say()时 执行的是 student 的 say()
let student: People = new Student('Bob') 
student.say()