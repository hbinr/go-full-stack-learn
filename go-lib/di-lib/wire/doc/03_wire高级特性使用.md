# wire高级特性

以开发一个web项目([配套代码](../code/advanced/))为基准，面向接口开发，最大限度的利用依赖注入`wire`工具。

先说下个人感受：

说实话，初次使用，学习成本和使用成本还是挺高的，尤其是结构体和接口之间的依赖多了以后，别人看你的代码时是无法知道项目如何初始化，如何处理依赖的。

虽然简化了`NewXXX()`初始化结构体等代码，但是可读性却大大降低了，对于大型项目可以考虑使用，小项目直接手撸`NewXXX`也无妨。

## ProviderSet
有时候可能多个类型有相同的依赖，我们每次都将相同的构造器传给`wire.Build()`不仅繁琐，而且不易维护，一个依赖修改了，所有传入`wire.Build()`的地方都要修改。为此，wire提供了一个`ProviderSet`（构造器集合），可以将多个构造器打包成一个集合，后续只需要使用这个集合即可。假设，我们有两个结构体需要构造，但是其字段都是一样的：
```go
// EndingA 结构体A
type EndingA struct {
  Player  Player
  Monster Monster
}

func NewEndingA(p Player, m Monster) EndingA {
  return EndingA{p, m}
}

func (p EndingA) Appear() {
  fmt.Println(p.Player.Name, p.Monster.Name)
}
// EndingB 结构体B
type EndingB struct {
  Player  Player
  Monster Monster
}

func NewEndingB(p Player, m Monster) EndingB {
  return EndingB{p, m}
}

func (p EndingB) Appear() {
  fmt.Println(p.Player.Name, p.Monster.Name)
}
```
编写两个注入器：

```go
func InitEndingA(name string) EndingA {
  wire.Build(NewMonster, NewPlayer, NewEndingA)
  return EndingA{}
}

func InitEndingB(name string) EndingB {
  wire.Build(NewMonster, NewPlayer, NewEndingB)
  return EndingB{}
}
```
我们观察到两次调用`wire.Build()`都需要传入`NewMonster`和`NewPlayer`。两个还好，如果很多的话写起来就麻烦了，而且修改也不容易。这种情况下，我们可以先定义一个`ProviderSet`：

```go
var monsterPlayerSet = wire.NewSet(NewMonster,NewPlayer)
```
 `wire.NewSet`返回的就是 `ProviderSet`
后续直接使用这个set：
```go
func InitEndingA(name string) EndingA {
  wire.Build(monsterPlayerSet, NewEndingA)
  return EndingA{}
}

func InitEndingB(name string) EndingB {
  wire.Build(monsterPlayerSet, NewEndingB)
  return EndingB{}
}
```
而后如果要添加或删除某个构造器，直接修改set的定义处即可。


## 结构构造器
因为我们的`EndingA`和`EndingB`的字段只有`Player`和`Monster`，我们就不需要显式为它们提供构造器，可以直接使用`wire`提供的**结构构造器**（Struct Provider）。结构构造器创建某个类型的结构，然后用参数或调用其它构造器填充它的字段。例如上面的例子，我们去掉`NewEndingA()`和`NewEndingB()`，然后为它们提供结构构造器：

```go
var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)

var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "Player", "Monster"))
var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "Player", "Monster"))

func InitEndingA(name string) EndingA {
  wire.Build(endingASet)
  return EndingA{}
}

func InitEndingB(name string) EndingB {
  wire.Build(endingBSet)
  return EndingB{}
}
```
结构构造器使用`wire.Struct`注入，第一个参数固定为new(结构名)，后面可接任意多个参数，表示需要为该结构的哪些字段注入值。上面我们需要注入`Player`和`Monster`两个字段。或者我们也可以使用通配符`*`表示注入所有字段：
```go
var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "*"))
var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "*"))
```

`wire`为我们生成正确的代码，非常棒：
```go
func InitEndingA(name string) EndingA {
  player := NewPlayer(name)
  monster := NewMonster()
  endingA := EndingA{
    Player:  player,
    Monster: monster,
  }
  return endingA
}
```
## 绑定值
有时候，我们需要为某个类型绑定一个值，而不想依赖构造器每次都创建一个新的值。有些类型天生就是单例，例如配置，数据库对象（sql.DB）。这时我们可以使用`wire.Value`绑定值，使用`wire.InterfaceValue`绑定接口。例如，我们的怪兽一直是一个Kitty，我们就不用每次都去创建它了，直接绑定这个值就 ok 了：
```go
var kitty = Monster{Name: "kitty"}

func InitEndingA(name string) EndingA {
  wire.Build(NewPlayer, wire.Value(kitty), NewEndingA)
  return EndingA{}
}

func InitEndingB(name string) EndingB {
  wire.Build(NewPlayer, wire.Value(kitty), NewEndingB)
  return EndingB{}
}
```
注意一点，这个值每次使用时都会拷贝，需要确保拷贝无副作用：
```go
// wire_gen.go
func InitEndingA(name string) EndingA {
  player := NewPlayer(name)
  monster := _wireMonsterValue 
  endingA := NewEndingA(player, monster)
  return endingA
}

var (
  _wireMonsterValue = kitty
)
```
结构字段作为构造器
有时候我们编写一个构造器，只是简单的返回某个结构的一个字段，这时可以使用`wire.FieldsOf`简化操作。现在我们直接创建了`Mission`结构，如果想获得`Monster`和`Player`类型的对象，就可以对`Mission`使用`wire.FieldsOf`：
```go
func NewMission() Mission {
  p := Player{Name: "dj"}
  m := Monster{Name: "kitty"}

  return Mission{p, m}
}

// wire.go
func InitPlayer() Player {
  wire.Build(NewMission, wire.FieldsOf(new(Mission), "Player"))
}

func InitMonster() Monster {
  wire.Build(NewMission, wire.FieldsOf(new(Mission), "Monster"))
}

// main.go
func main() {
  p := InitPlayer()
  fmt.Println(p.Name)
}
```
同样的，第一个参数为`new(结构名)`，后面跟多个参数表示将哪些字段作为构造器，`*`表示全部。

## 清理函数
构造器可以提供一个清理函数，如果后续的构造器返回失败，前面构造器返回的清理函数都会调用：
```go
func NewPlayer(name string) (Player, func(), error) {
  cleanup := func() {
    fmt.Println("cleanup!")
  }
  if time.Now().Unix()%2 == 0 {
    return Player{}, cleanup, errors.New("player dead")
  }
  return Player{Name: name}, cleanup, nil
}

func main() {
  mission, cleanup, err := InitMission("dj")
  if err != nil {
    log.Fatal(err)
  }

  mission.Start()
  cleanup()
}

// wire.go
func InitMission(name string) (Mission, func(), error) {
  wire.Build(NewMonster, NewPlayer, NewMission)
  return Mission{}, nil, nil
}
```

跟返回错误类似，将`provider`的第二个返回参数设置成`func()`用于返回`cleanup function`，上述例子中在第三个参数中返回了error，但这是可选的：

`wire`对`provider`的返回值个数和顺序有所规定：
- 第一个参数是需要生成的依赖对象
- 如果返回2个返回值，第二个参数必须是`func()`或者error
- 如果返回3个返回值，第二个参数必须是`func()`，第三个参数则必须是`error`
## 区分类型
由于`injector`的函数中，不允许出现重复的参数类型，否则`wire`将无法区分这些相同的参数类型，比如：
```go
// 两个字段都为string
type FooBar struct {
	foo string
	bar string
}

func NewFooBar(foo string, bar string) FooBar {
	return FooBar{
	    foo: foo,  
	    bar: bar,
	}
}
```
injector函数签名定义:
```go
// wire无法得知入参a,b跟 FooBar.foo,FooBar.bar 的对应关系
func InitializeFooBar(a string, b string) FooBar {
	wire.Build(NewFooBar)
	return FooBar{}
}
```
如果使用上面的`provider`来生成`injector`,`wire`会报如下错误：
> provider has multiple parameters of type string
因为入参均是字符串类型，`wire`无法得知入参a,b跟`FooBar.foo`,`FooBar.bar`的对应关系。 所以我们使用不同的类型来避免冲突：
```go
type Foo string
type Bar string
type FooBar struct {
	foo Foo
	bar Bar
}

func NewFooBar(foo Foo, bar Bar) FooBar {
	return FooBar{
	    foo: foo,
	    bar: bar,
	}
}
```
injector函数签名定义：
```go
func InitializeFooBar(a Foo, b Bar) FooBar {
	wire.Build(NewFooBar)
	return FooBar{}
}
```
其中基础类型和通用接口类型是最容易发生冲突的类型，如果它们在`provider`函数中出现，最好统一新建一个别名来代替它(尽管还未发生冲突)，例如：
```go
type MySQLConnectionString string
type FileReader io.Reader
```
# Options 结构体
如果一个`provider`方法包含了许多依赖，可以将这些依赖放在一个`options`结构体中，从而避免构造函数的参数太多：
```go
type Message string

// Options
type Options struct {
	Messages []Message
	Writer   io.Writer
	Reader   io.Reader
}
type Greeter struct {
}

// NewGreeter Greeter的provider方法使用Options以避免构造函数过长
func NewGreeter(ctx context.Context, opts *Options) (*Greeter, error) {
	return nil, nil
}
// GreeterSet 使用wire.Struct设置Options为provider
var GreeterSet = wire.NewSet(wire.Struct(new(Options), "*"), NewGreeter)
```
injector函数签名：
```go
func InitializeGreeter(ctx context.Context, msg []Message, w io.Writer, r io.Reader) (*Greeter, error) {
	wire.Build(GreeterSet)
	return nil, nil
}
```
## 接口绑定
根据依赖倒置原则（Dependence Inversion Principle），对象应当依赖于接口，而不是直接依赖于具体实现。

在quickstart的例子中的依赖均是具体实现，现在我们来看看在`wire`中如何处理接口依赖：
```go
// UserService 
type UserService struct {
	userRepo UserRepository // <-- UserService依赖UserRepository接口
}

// UserRepository 存放User对象的数据仓库接口,比如可以是mysql,restful api ....
type UserRepository interface {
	// GetUserByID 根据ID获取User, 如果找不到User返回对应错误信息
	GetUserByID(id int) (*User, error)
}
// NewUserService *UserService构造函数
func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo:userRepo,
	}
}

// mockUserRepo 模拟一个UserRepository实现
type mockUserRepo struct {
	foo string
	bar int
}
// GetUserByID UserRepository接口实现
func (u *mockUserRepo) GetUserByID(id int) (*User,error){
	return &User{}, nil
}
// NewMockUserRepo *mockUserRepo构造函数
func NewMockUserRepo(foo string,bar int) *mockUserRepo {
	return &mockUserRepo{
		foo:foo,
		bar:bar,
	}
}
// MockUserRepoSet 将 *mockUserRepo与UserRepository绑定
var MockUserRepoSet = wire.NewSet(NewMockUserRepo,wire.Bind(new(UserRepository), new(*mockUserRepo)))
```

在这个例子中，`UserService`依赖`UserRepository`接口，其中`mockUserRepo`是`UserRepository`的一个实现，由于在Go的最佳实践中，更推荐返回具体实现而不是接口。所以`mockUserRepo`的`provider`函数返回的是*`mockUserRepo`这一具体类型。`wire`无法自动将具体实现与接口进行关联，我们需要显示声明它们之间的关联关系。通过`wire.NewSet`和`wire.Bind`将*`mockUserRepo`与`UserRepository`进行绑定：
```go
// MockUserRepoSet 将 *mockUserRepo与UserRepository绑定
var MockUserRepoSet = wire.NewSet(NewMockUserRepo,wire.Bind(new(UserRepository), new(*mockUserRepo)))

```

## 一些细节和记忆技巧
- 首先，我们调用wire生成wire_gen.go之后，如果wire.go文件有修改，只需要执行go generate即可；
- `wire.Struct((new(*User)),"*")` 相当于初始化结构体或者说创建实例(对象)，平时我们得一个一个的对字段进行赋值，可以使用`*`来指定全部字段，如果指向初始化某个字段，
只需将 `"*"`改为，`"字段名"`，如:`wire.Struct((new(*User)),"Name","Age")`
- `wire.Bind(interface,impl)`，顾名思义，将接口和实现绑定在一起。因为go中接口实现是隐式的，只要一个接口体实现了某个接口的全部方法，就说该
结构体实现了这个接口。`wire.Bind(interface,impl)`最简单的理解就是相当于Java中的类`implment`接口，需要显示绑定。但要注意，user和*User实现接口是不同的，需要特别指明
- `wire.NewSet()`，返回`ProviderSet`，这个可以理解为初始化复杂结构体，该结构体的字段类型是结构体或者其他，也就是依赖于其他结构体先初始化，然后再将其赋值给
对应的结构类型的字段。
- `wire.Build()`，构建依赖并生成代码，需要我们将之前的`NewSet()`的值一个一个注入进去，但是要注意两点：
    - 1.注入对象时，要注意顺序，从上层到下层依次注入，否则会报 no provider + 对应的对象
    - 2.没有用到对象不用注入

# 总结

wire是随着go-cloud的示例 [guestbook](https://github.com/google/go-cloud/tree/master/samples/guestbook)一起发布的，可以阅读guestbook看看它是怎么使用wire的。与dig不同，wire只是生成代码，不使用reflect库，性能方面是不用担心的。因为它生成的代码与你自己写的基本是一样的。如果生成的代码有性能问题，自己写大概率也会有 。


参考：

https://zhuanlan.zhihu.com/p/110453784

https://blog.csdn.net/uisoul/article/details/108776073