原文：https://blog.drewolson.org/dependency-injection-in-go/
翻译：devabel

我最近用Go开发了一个小项目。过去几年我一直使用Java进行开发，立即被Go生态系统中依赖注入（DI）背后的动力所打动。我决定尝试使用Uber的dig(https://github.com/uber-go/dig)库来构建我的项目，给我留下了深刻的印象。

我发现依赖注入DI帮助解决了我在以前的Go应用程序中遇到的很多问题 - 过度使用init函数，滥用全局变量和复杂的应用程序设置。

在本文中，我将介绍DI，然后通过一个示例，来看看使用DI框架（通过Uber的dig库）前后的不同。

## DI的简要概述

依赖注入是你的组件（比如go语言中的structs）在创建时应该接收它的依赖关系的思想。这与在初始化期间构建其自己的依赖关系的组件的相关反模式相反。我们来看一个例子。

假设你有一个Server结构体需要一个Config结构体来实现它的行为。实现的一种方法就是在初始化期间Server自行构建Config。

```go
type Server struct {
  config *Config
}

func New() *Server {
  return &Server{
    config: buildMyConfigSomehow(),
  }
}
```
这似乎很方便。组件调用者不必知道我们Server，甚至不需要操作Config。这对我们的组件调用者用户来说都是隐藏的。

但是，这用做有一些缺点。首先，如果我们想改变我们Config建造的方式，我们将不得不改变所有调用这个组件的代码。例如，假设我们的buildMyConfigSomehow功能现在需要一个参数。每个调用者都需要访问该参数，并需要将其传递到构建函数中。

另外，mock我们的行为变得非常困难Config。我们必须以某种方式进入我们的New功能内部，以便创建Config。

下面看看使用依赖注入的方式来实现：

```go
type Server struct {
  config *Config
}

func New(config *Config) *Server {
  return &Server{
    config: config,
  }
}
```
现在我们Server的创造与创造的创造脱离了关系Config。我们可以使用任何我们想要创建的逻辑Config，然后将结果数据传递给我们的New函数。

此外，如果Config是一个接口，这给我们一个容易的嘲弄路线。New只要它实现我们的界面，我们就可以传递我们想要的任何东西。这使得我们Server用Config简单的模拟实现进行测试。

主要的缺点是，Config在我们可以创建之前必须手动创建它是一种痛苦Server。我们在这里创建了一个依赖图 - 我们必须创建我们的Config第一个，因为Server它依赖于它。在实际应用中，这些依赖关系图可能会变得非常大，并且会导致构建应用程序需要完成其工作的所有组件的复杂逻辑。

这是DI框架可以提供帮助的地方。DI框架通常提供两种功能：

一种“提供”新组件的机制。简而言之，这将告诉DI框架您需要构建自己的其他组件（您的依赖关系）以及在拥有这些组件后如何构建自己。
一种“检索”构建组件的机制。
DI框架通常基于您所讲述的“提供者”构建一个图并确定如何构建您的对象。这在摘要中很难理解，所以让我们通过一个中等大小的示例。

## 示例应用程序

我们将要审查一个HTTP服务器的代码，当客户端发出GET请求时，它会提供JSON响应/people。我们将逐个检查代码。为了简单起见，它全都存在于相同的包（main）中。请不要在真正的Go应用程序中执行此操作。这个例子的完整代码可以在[这里](https://gitlab.com/drewolson/go_di_example/-/blob/master/example.go)找到。

首先，让我们看看我们的Person结构。除了一些JSON标签外，它没有任何行为:

```go
type Person struct {                    
  Id   int    `json:"id"`
  Name string `json:"name"`                              
  Age  int    `json:"age"`                   
}  
```
A Person有Id，Name和Age。而已。
接下来让我们看看我们的Config。类似于Person，它没有依赖关系。不像Person，我们会提供一个构造函数。

```go
type Config struct {               
  Enabled      bool         
  DatabasePath string        
  Port         string                       
}                                         
                     
func NewConfig() *Config {                            
  return &Config{            
    Enabled:      true,                           
    DatabasePath: "./example.db",       
    Port:         "8000",
  }                 
}   
```
Enabled告诉我们我们的应用程序是否应该返回实际数 DatabasePath告诉我们数据库在哪里（我们正在使用sqlite）。Port告诉我们将运行我们的服务器的端口。

这里是我们用来打开数据库连接的函数。它依赖于我们Config并返回一个*sql.DB。

```go
func ConnectDatabase(config *Config) (*sql.DB, error) {
  return sql.Open("sqlite3", config.DatabasePath)
} 
```
接下来我们会看看我们的PersonRepository。该结构将负责从我们的数据库中提取人员并将这些数据库结果反序列化为合适的Person结构。

```go
type PersonRepository struct {                           
  database *sql.DB                                              
}                                        
                                      
func (repository *PersonRepository) FindAll() []*Person {            
  rows, _ := repository.database.Query(
    `SELECT id, name, age FROM people;`
  )   
  defer rows.Close()                                           
                                                                                  
  people := []*Person{}                          
                                                       
  for rows.Next() {    
    var (               
      id   int      
      name string
      age  int              
    )                   
                                          
    rows.Scan(&id, &name, &age)         
                                           
    people = append(people, &Person{
      Id:   id,
      Name: name,
      Age:  age,
    })                 
  }                                         
                                          
  return people
}                                                     
                                   
func NewPersonRepository(database *sql.DB) *PersonRepository {
  return &PersonRepository{database: database}
}   
```
PersonRepository需要建立数据库连接。它公开了一个单独的函数FindAll，它使用我们的数据库连接返回一个Person表示数据库中数据的结构列表。

为了在我们的HTTP服务器和PersonRepository我们之间提供一个图层，我们将创建一个PersonService。

```go
type PersonService struct {                                                          
  config     *Config                                                  
  repository *PersonRepository    
}                                                                                                        
                                                                  
func (service *PersonService) FindAll() []*Person {
  if service.config.Enabled {     
    return service.repository.FindAll()
  }                                                
                                                                      
  return []*Person{}               
}                                     
                                                                     
func NewPersonService(config *Config, repository *PersonRepository) *PersonService {
  return &PersonService{config: config, repository: repository}
}
```
我们PersonService依赖于Config和PersonRepository。它公开了一个被称为“ FindAll有条件地调用PersonRepository应用程序是否被启用” 的函数。

最后，我们有我们的Server。这是负责运行一个HTTP服务器并委托给我们的合适的请求PersonService。

```go
type Server struct {                                   
  config        *Config
  personService *PersonService
}                                     
                                   
func (s *Server) Handler() http.Handler {
  mux := http.NewServeMux()   
                                          
  mux.HandleFunc("/people", s.people)
                                           
  return mux                              
}                                
                                   
func (s *Server) Run() {     
  httpServer := &http.Server{
    Addr:    ":" + s.config.Port,
    Handler: s.Handler(),
  }            
                                                      
  httpServer.ListenAndServe()             
}                                    
                          
func (s *Server) people(w http.ResponseWriter, r *http.Request) {
  people := s.personService.FindAll()
  bytes, _ := json.Marshal(people)
                               
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)          
  w.Write(bytes)               
}                                                  
                              
func NewServer(config *Config, service *PersonService) *Server {
  return &Server{   
    config:        config,
    personService: service,            
  }                                                                           
} 
```
这Server是依赖于PersonService和Config。

好的，我们知道我们系统的所有组件。现在我们究竟如何初始化它们并启动我们的系统呢？

## 可怕的main（）

首先，让我们main()以旧式的方式编写我们的功能。
```go
func main() {
  config := NewConfig()

  db, err := ConnectDatabase(config)

  if err != nil {
    panic(err)
  }

  personRepository := NewPersonRepository(db)

  personService := NewPersonService(config, personRepository)

  server := NewServer(config, personService)

  server.Run()
}
```
首先，我们创造我们的Config。然后，使用Config，我们创建我们的数据库连接。从那里我们可以创造我们的PersonRepository这使我们能够创造我们的PersonService。最后，我们可以用它来创建Server并运行它。

唷，那很复杂。更糟的是，随着我们的应用程序变得越来越复杂，我们的main意志会越来越复杂。每当我们为我们的任何组件添加一个新的依赖关系时，我们都必须通过main函数中的顺序和逻辑来反映该依赖关系，以构建该组件。

正如您可能已经猜到的那样，依赖注入框架可以帮助我们解决这个问题。让我们来看看如何。

## 建立一个容器

术语“容器”通常用于DI框架中，用于描述添加“提供程序”的内容，您可以从中获取完整构建对象。该dig库为我们Provide提供了添加提供程序的Invoke功能以及从容器中检索完全构建的对象的功能。

首先，我们建立一个新的容器。
```go
container := dig.New()
```

现在我们可以添加新的提供者。为此，我们调用Provide容器上的函数。它只需要一个参数：一个函数。该函数可以有任意数量的参数（表示要创建的组件的依赖关系）以及一个或两个返回值（表示该函数提供的组件以及可选的错误）。

```go
container.Provide(func() *Config {
  return NewConfig()
})
```
上面的代码说：“我为Config容器提供了一个类型，为了构建它，我不需要其他东西。” 现在我们已经展示了容器如何构建一个Config类型，我们可以使用它来构建其他类型。

```go
container.Provide(func(config *Config) (*sql.DB, error) {
  return ConnectDatabase(config)
})
```
这段代码说：“我为*sql.DB容器提供了一个类型，为了构建它，我需要一个Config。我也可以选择返回一个错误。”

在这两种情况下，我们都比所需的更冗长。因为我们已经有了NewConfig和ConnectDatabase定义的功能，我们可以直接使用他们作为供应商的容器。

```go
container.Provide(NewConfig)
container.Provide(ConnectDatabase)
```
现在，我们可以要求容器给我们提供我们提供的任何类型的完全构建的组件。我们这样做使用该Invoke功能。该Invoke函数接受一个参数 - 一个包含任意数量参数的函数。函数的参数是我们希望容器为我们构建的类型。

```go
container.Invoke(func(database *sql.DB) {
  // sql.DB is ready to use here
})
```
容器做了一些非常聪明的东西。以下是发生的情况：

- 容器认识到我们要求一个 *sql.DB
- 它确定我们的功能ConnectDatabase提供了这种类型
- 接下来确定我们的ConnectDatabase函数具有依赖性Config
- 它找到了Config这个NewConfig函数的提供者
- NewConfig 没有任何依赖关系，所以它被调用
- 结果NewConfig是一个Config传递给ConnectDatabase
- 结果ConnectionDatabase是*sql.DB被传递回调用者Invoke

这是容器为我们做的很多工作。事实上，它做得更多。该容器足够聪明，可以构建每个类型的一个实例，并且只有一个实例。这意味着如果我们在多个地方使用它（如多个存储库），我们绝不会意外创建第二个数据库连接。

## 更好的main（）

现在我们知道dig容器的工作原理了，让我们用它来构建一个更好的主体。

```go
func BuildContainer() *dig.Container {
  container := dig.New()

  container.Provide(NewConfig)
  container.Provide(ConnectDatabase)
  container.Provide(NewPersonRepository)
  container.Provide(NewPersonService)
  container.Provide(NewServer)

  return container
}

func main() {
  container := BuildContainer()

  err := container.Invoke(func(server *Server) {
    server.Run()
  })

  if err != nil {
    panic(err)
  }
}
```
我们以前没有见过的唯一的东西就是error来自的返回值Invoke。如果任何使用的提供者Invoke返回错误，我们的调用Invoke将停止并返回错误。

尽管这个例子很小，但应该很容易看出这种方法在我们的“标准”主体上的一些好处。随着我们的应用程序越来越大，这些好处变得更加明显。

最重要的好处之一就是创建我们的组件与创建它们的依赖关系的解耦。比如说，我们PersonRepository现在需要访问Config。我们所要做的就是改变我们的NewPersonRepository构造函数以包含Config作为参数。我们的代码中没有其他的变化。

其他的巨大好处是缺乏全局状态，缺少调用init（需要时懒惰地创建依赖关系，只创建一次，无需易错的init设置），并且易于对各个组件进行测试。想象一下，在测试中创建容器并要求完全构建的对象进行测试。或者，使用所有依赖关系的模拟实现来创建一个对象。所有这些在DI方法中都更容易。

## 一个值得传播的理念

我相信依赖注入有助于构建更强大和可测试的应用程序。随着这些应用程序规模的扩大，情况尤其如此。Go非常适合构建大型应用程序，并拥有一个很棒的DI工具dig。我相信Go社区应该拥抱DI并将其用于更多的应用程序中。