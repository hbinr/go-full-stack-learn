/**
 * 对象展开，内容要点：
 * 1. 将对象中的属性一一展开，从左至右，这就意味着出现在展开对象后面的属性会覆盖前面的属性
 * 2. 
 */

let spicyObj = {
    food: "spicy",
    price: 10
}

// 1. 对象展开
let search = {...spicyObj}
console.log('search: ', search); // search:  { food: 'spicy', price: 10 }


// 2. 前置展开
let search2 = {...spicyObj, food: 'rich'}
console.log('search2: ', search2); // search2:  { food: 'rich', price: 10 }

// 3. 后置展开，如果前面存在重复的属性名，那么后置展开后会覆盖前面的属性值
let search3 = {food: 'rich', ...spicyObj }  // 编辑器提示报错，但是也能编译成功
console.log('search3: ', search3); // search3:  { food: 'spicy', price: 10 } ， spicyObj中food字段值覆盖了rich
