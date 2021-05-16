/**
 * this 使用和注意事项，内容要点：
 * 1. 定义语法：(参数1:数据类型行, ....) => 返回值数据类型
 * 2. 关键字：this
 * 3. this 在学习Class的时候经常用到， 用来获取对象实例的 属性和方法。
 * 4. 在函数中，也有this的用法，JavaScript里，this 的值在函数被调用的时候才会指定.
 *    要理清函数调用的上下文是什么，这样才能明白 this 的作用
 * 5. this作用域：
 *      a. this作用域和 this 所在的当前函数有关，比如普通函数，回调函数，this 只在当前函数作用域内
 *      b. 可以使用箭头函数解决上述问题 ()=>{}，这样就有更大的作用域了
 * 
 */


// 1. this 作用域问题， 使用箭头函数： ()=>{} 代替function关键字
let deck = {
    suits: ['hearts', 'spades', 'clubs', 'diamonds'],
    cards: Array(52),
    createCardPicker: function() {
        // return function() {  // 原报错代码
        return () => {  // 改成箭头函数就不会报错了
        let pickedCard = Math.floor(Math.random() * 52)
        let pickedSuit = Math.floor(pickedCard / 13)
        //  报错： 因为 createCardPicker 返回的函数里的 this 被设置成了 global 而不是 deck 对象
        return {suit: this.suits[pickedSuit], card: pickedCard % 13}
        }
    }
}

let cardPicker = deck.createCardPicker()
let pickedCard = cardPicker()

console.log('card: ' + pickedCard.card + ' of ' + pickedCard.suit)

// 2.在上述的例子中 this.suits[pickedSuit] 的类型为 any，这是因为 this 来自对象字面量里的函数表达式。 
//   修改的方法是，提供一个显式的 this 参数。 this 参数是个假的参数，它出现在参数列表的最前面

interface Card { // 新增接口 Card
    suit: string
    card: number
}

interface Deck { // 新增接口 Deck
    suits: string[]
    cards: number[]

    createCardPicker (this: Deck): () => Card
}

let deck2: Deck = { // 新增类型注解
    suits: ['hearts', 'spades', 'clubs', 'diamonds'],
    cards: Array(52),
    createCardPicker: function(this: Deck) {  // 新增参数 this: Deck
        return () => {  
        let pickedCard = Math.floor(Math.random() * 52)
        let pickedSuit = Math.floor(pickedCard / 13)
        //  报错： 因为 createCardPicker 返回的函数里的 this 被设置成了 global 而不是 deck 对象
        return {suit: this.suits[pickedSuit], card: pickedCard % 13}
        }
    }
}

let cardPicker2 = deck2.createCardPicker()
let pickedCard2 = cardPicker()

console.log('card: ' + pickedCard.card + ' of ' + pickedCard.suit)

// 3. 

class Handler {
    type: string

    constructor(type: string){
        this.type = type
    }
    
    // 错误 
    onClickBad = function(e: Event){
        // 读取不到 this
        // this.type = e.type 
      }

    // 正确
    onClickGood = (e: Event) => {
      this.type = e.type 
    }
}