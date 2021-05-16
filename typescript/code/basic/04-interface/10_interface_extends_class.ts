/**
 * 接口继承类 接口，内容要点：
 * 1. 定义语法：interface A extends class B  {
 *             }
 * 2. 关键字：interface , extends
 * 3. 当接口继承了一个类类型时，它会继承类的成员但不包括其实现。 
 * 4. 就好像接口声明了所有类中存在的成员，但并没有提供具体实现一样。 
 * 5. 接口同样会继承到类的 private 和 protected 成员。 
 * 6. 这意味着当你创建了一个接口继承了一个拥有 私有或受保护 的成员的类时，这个接口类型只能被这个类或其子类所实现（implement）。
 * 
 */

// 1. 定义类
class Control {
    private state: any
}

// 2. 定义接口，并继承 Control 类
interface SelectableControl extends Control {
    select(): void
}

// 3. 定义 Control 的子类
class TextBox extends Control {
    select() { }
}

// 4. 定义 Control 的子类，并实现 SelectableControl
class Button extends Control implements SelectableControl {
    select() { }
}

// Error：“ImageC”类型缺少“state”属性。
// ImageC 不是 Control 的子类，无法实现 SelectableControl接口
class ImageC implements SelectableControl {
    select() { }
}