/**
 * 静态属性，内容要点：
 * 1. 定义语法：static 属性名:属性类型, eg: [public/private/protected] static name: string
 * 2. 关键字：static
 * 3. 静态属性： 存在于类本身上面而不是类的实例上，通过类名去调用，而不是 this.
 * 4. 类的成员(普通属性)： 类的实例成员，那些仅当类被实例化的时候才会被初始化的属性
 * 5. typeof 可以获取类的类型，即构造函数的类型。 这个类型包含了类的 所有静态成员 和 构造函数 
 * 6. 通过typeof 获取到类的所有静态成员和构造函数后，便可以操作 静态成员了
 * 
 */

 class Grid {
    //  静态属性，不需要实例化对象进行初始化，本身就是类的属性
    static origin = {x: 0, y: 0} // 定义时便初始化了
    static num: number // 定义时未初始化，需要通过 typeof 获取类的静态属性，然后再初始化
  
    scale: number
  
    constructor (scale: number) {
        this.scale = scale
    }
  
    calculateDistanceFromOrigin(point: {x: number; y: number}) {
        console.log(Grid.num);
        
        // 通过 '类名.' 去调用静态属性 -> Grid.origin
        let xDist = point.x - Grid.origin.x
        let yDist = point.y - Grid.origin.y
        return Math.sqrt(xDist * xDist + yDist * yDist) * this.scale
    }
  }
  
let grid1 = new Grid(1.0)  // 1x scale
let grid2 = new Grid(5.0)  // 5x scale

console.log(grid1.calculateDistanceFromOrigin({x: 3, y: 4}))
console.log(grid2.calculateDistanceFromOrigin({x: 3, y: 4}))


// 通过typeof Grid 获取 Grid 类的类型，而不是实例的类型。
// 或者更确切的说，"告诉我 Grid 标识符的类型"，也就是构造函数的类型。 这个类型包含了类的所有静态成员和构造函数
let GridMaker: typeof Grid = Grid
GridMaker.num = 100

console.log(Grid.num); // 100