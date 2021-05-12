// 类定义
var User = /** @class */ (function () {
    function User(firstName, lastName) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.fullName = firstName + ' ' + lastName;
    }
    return User;
}());
// 方法
function greeter(person) {
    return 'Hello' + person.fullName;
}
// 调用方法
var user = new User('haoBin', 'duan');
greeter(user);
console.log('greeter(user): ', greeter(user));
