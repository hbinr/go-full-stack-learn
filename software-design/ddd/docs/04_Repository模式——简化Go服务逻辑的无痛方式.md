
# The Repository pattern: a painless way to simplify your Go service logic
Repository模式：简化Go服务逻辑的无痛方式

[原文](https://threedots.tech/post/repository-pattern-in-go/)

I’ve seen a lot of complicated code in my life. Pretty often, the reason of that complexity was application logic coupled with database logic. **Keeping logic of your application along with your database logic makes your application much more complex, hard to test, and maintain.**

我一生中见过很多复杂的代码。几乎很多时候，这种复杂性的原因是应用逻辑加上数据库逻辑。**将应用程序的逻辑与数据库逻辑一起保存，会让你的应用程序变得更加复杂，难以测试和维护。**

There is already a proven and simple pattern that solves these issues. The pattern that allows you to **separate your application logic from database logic**. It allows you to **make your code simpler and easier to add new functionalities**. As a bonus, you can **defer important decision** of choosing database solution and schema. Another good side effect of this approach is out of the box **immunity for database vendor lock-in**. The pattern that I have in mind is **repository**.

已经有一个成熟而简单的模式可以解决这些问题。该模式允许你**将你的应用逻辑和数据库逻辑分开**。它允许你**让你的代码更简单，更容易添加新的功能**。作为奖励，你可以**推迟选择**数据库解决方案和模式的重要决定。这种方法的另一个好的副作用是**对数据库供应商锁定的免疫力**。我想到的模式是`repository`。

When I’m going back in my memories to the applications I worked with, I remember that it was tough to understand how they worked. **I was always afraid to make any change there – you never know what unexpected, bad side effects it could have**. It’s hard to understand the application logic when it’s mixed with database implementation. It’s also the source of duplication.

当我在记忆中回溯到我工作过的应用程序时，我记得要理解它们的工作原理是很困难的。**我总是害怕在那里做任何改变--你永远不知道它会产生什么意想不到的、不好的副作用**。当应用逻辑与数据库实现混合在一起时，很难理解它。这也是重复的根源。

Some rescue here may be building end-to-end tests. But it’s hiding the problem instead of really solving it. Having a lot of E2E tests is slow, flaky, and hard to maintain. Sometimes they are even preventing us from creating new functionality, rather than helping.

这里的一些救援可能是构建端到端测试。但这是在隐藏问题，而不是真正解决问题。拥有大量的E2E测试是缓慢的、片面的、难以维护的。有时它们甚至阻碍了我们创建新功能，而不是帮助我们。

In today’s article, I will teach you how to apply this pattern in Go in a pragmatic, elegant, and straightforward way. I will also deeply cover a topic that is often skipped - **clean transactions handling**.

在今天的文章中，我将教你如何在Go中以一种务实、优雅、直接的方式应用这种模式(即repository)。我还将深入介绍一个经常被跳过的话题--**干净的事务处理**。

To prove that I prepared 3 implementations: Firestore, MySQL, and simple in-memory.

为了证明这一点，我准备了3个实现。Firestore, MySQL, 和简单的内存。

Without too long introduction, let’s jump to the practical examples!

不需要太长的介绍，让我们跳到实际的例子中去吧!

## Repository interface
The idea of using the repository pattern is:

使用`repository`模式的思路是:

**Let’s abstract our database implementation by defining interaction with it by the interface. You need to be able to use this interface for any database implementation – that means that it should be free of any implementation details of any database.**

让我们通过接口定义与数据库的交互来抽象我们的数据库实现。你需要能够将这个接口用于任何数据库的实现--这意味着它应该不包含任何数据库的实现细节。

Let’s start with the refactoring of trainer service. Currently, the service allows us to get information about hour availability via HTTP API and via gRPC. We can also change the availability of the hour via HTTP API and gRPC.

让我们从重构`trainer service`开始。目前，该服务允许我们通过`HTTP API`和通过`gRPC`获取`hour`的可用性信息。我们也可以通过`HTTP API`和`gRPC`来改变`hour`的可用性。

In the previous article, we refactored Hour to use DDD Lite approach. Thanks to that, we don’t need to think about keeping rules of when Hour can be updated. Our domain layer is ensuring that we can’t do anything “stupid”. We also don’t need to think about any validation. We can just use the type and execute necessary operations.

在[上一篇文章](https://threedots.tech/post/ddd-lite-in-go-introduction/)中，我们重构了`Hour`，使用`DDD Lite`方式。得益于此，我们不需要考虑保留`Hour`何时可以更新的规则。我们的领域层正在确保我们不能做任何 "愚蠢 "的事情。我们也不需要考虑任何验证。我们只需要使用类型并执行必要的操作。

We need to be able to get the current state of Hour from the database and save it. Also, in case when two people would like to schedule a training simultaneously, only one person should be able to schedule training for one hour.

我们需要能够从数据库中获取`Hour`的当前状态并保存。另外，在两个人同时想安排培训的情况下，只需要一个人能够安排一个小时的培训。

Let’s reflect our needs in the interface:

让我们在`interface`中体现我们的需求。

```go
package hour

type Repository interface {
   GetOrCreateHour(ctx context.Context, hourTime time.Time) (*Hour, error)
   UpdateHour(
      ctx context.Context,
      hourTime time.Time,
      updateFn func(h *Hour) (*Hour, error),
   ) error
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/domain/hour/repository.go#L8)


We will use GetOrCreateHour to get the data, and UpdateHour to save the data.

我们将使用`GetOrCreateHour`来获取数据，使用`UpdateHour`来保存数据。

We define the interface in the same package as the Hour type. Thanks to that, we can avoid duplication if using this interface in many modules (from my experience, it may often be the case). It’s also a similar pattern to io.Writer, where io package defines the interface, and all implementations are decupled in separate packages.

我们将这个接口定义在与`Hour`类型相同的包中（即hour结构体定义和repository接口定义在同一个包中）。多亏了这一点，如果在很多模块中使用这个接口，我们可以避免重复（根据我的经验，可能经常是这样）。这也是类似于`io.Writer`的模式，`io`包定义了接口，所有的实现都会在单独的包中进行分解。

How to implement that interface?

如何实现这个接口呢？
## Reading the data 读取数据

Most database drivers can use the ctx context.Context for cancellation, tracing, etc. It’s not specific to any database (it’s a common Go concept), so you should not be afraid that you spoil the domain. 😉

大多数数据库驱动都可以使用`ctx context.Context`来进行取消、追踪等操作。这不是任何数据库所特有的（这是一个通用的Go概念），所以你不要怕你把这个`domain`给弄坏了。 😉。

In most cases, we query data by using UUID or ID, rather than time.Time. In our case, it’s okay – the hour is unique by design. I can imagine a situation that we would like to support multiple trainers – in this case, this assumption will not be valid. Change to UUID/ID would still be simple. But for now, YAGNI!

在大多数情况下，我们查询数据是用UUID或ID，而不是用time.Time。在我们的案例中，没关系--`hour`是唯一的设计。我可以想象一种情况，我们希望支持多个培训师--在这种情况下，这个假设将不成立。改为UUID/ID仍然会很简单。但现在，[YAGNI](https://en.wikipedia.org/wiki/You_aren%27t_gonna_need_it)（XP一个编程原则：程序员在认为必要之前不应该增加功能）!

For clarity – this is how the interface based on UUID may look like:

为了清晰起见--这就是基于UUID的`interface`可能的样子。
```go
GetOrCreateHour(ctx context.Context, hourUUID string) (*Hour, error)
```

> You can find example of repository based on UUID in [Combining DDD, CQRS, and Clean Architecture article](https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/#repository-refactoring).

> 你可以在[Combining DDD, CQRS, and Clean Architecture article](https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/#repository-refactoring)文章中找到基于UUID的存储库的例子。


How is the interface used in the application?

`application`中如何使用`interface`？

```go
import (
    // ...
   "github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/domain/hour"
    // ...
)

type GrpcServer struct {
   hourRepository hour.Repository
}

// ...

func (g GrpcServer) IsHourAvailable(ctx context.Context, request *trainer.IsHourAvailableRequest) (*trainer.IsHourAvailableResponse, error) {
   trainingTime, err := protoTimestampToTime(request.Time)
   if err != nil {
      return nil, status.Error(codes.InvalidArgument, "unable to parse time")
   }

   h, err := g.hourRepository.GetOrCreateHour(ctx, trainingTime)
   if err != nil {
      return nil, status.Error(codes.Internal, err.Error())
   }

   return &trainer.IsHourAvailableResponse{IsAvailable: h.IsAvailable()}, nil
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/0249977c58a310d343ca2237c201b9ba016b148e/internal/trainer/grpc.go#L75)

No rocket science! We get hour.Hour and check if it’s available. Can you guess what database we use? No, and that is the point!

没有火箭科学！我们得到`hour.Hour`，然后检查它是否可用。你能猜到我们用的是什么数据库吗？不能，这就是重点!

As I mentioned, we can avoid vendor lock-in and be able to easily swap the database. If you can swap the database , **it’s a sign that you implemented the repository pattern correctly**. In practice, the situation when you are changing a database is rare. 😉 In case when you are using a solution that is not self-hosted (like Firestore), it’s more important to mitigate the risk and avoid vendor lock-in.

正如我所提到的，我们可以避免供应商锁定，并能够轻松地交换数据库。**如果你能交换数据库 ，就说明你正确实现了 repository 模式**。😉如果你使用的不是自托管的解决方案（比如Firestore），更重要的是降低风险，避免厂商锁定。

The helpful side effect of that is that we can defer the decision of which database implementation we would like to use. I name this approach Domain First. I described it in depth [in the previous article](https://threedots.tech/post/ddd-lite-in-go-introduction/#domain-first-approach). **Deferring the decision about the database for later can save some time at the beginning of the project. With more informations and context, we can also make a better decision**.

这样做的有益副作用是，我们可以推迟决定使用哪种数据库实现。我把这种方法命名为`Domain First`。我在[上一篇文章]((https://threedots.tech/post/ddd-lite-in-go-introduction/#domain-first-approach))中对它进行了深入的描述。将数据库的决定推迟到以后，可以在项目开始时节省一些时间。有了更多的信息和背景，我们也可以做出更好的决定。

When we use the Domain-First approach, the first and simplest repository implementation may be in-memory implementation.

当我们使用`Domain-First`方法时，第一个也是最简单的存储库实现可能是内存实现。

## Example In-memory implementation  内存实现示例


Our memory uses a simple map under the hood. getOrCreateHour has 5 lines (without a comment and one newline 😉)!

我们的内存使用的是一个简单的`hood`下的`map`，`getOrCreateHour`有5行（没有注释和一个新行😉）!

```go
type MemoryHourRepository struct {
   hours map[time.Time]hour.Hour
   lock  *sync.RWMutex

   hourFactory hour.Factory
}

func NewMemoryHourRepository(hourFactory hour.Factory) *MemoryHourRepository {
   if hourFactory.IsZero() {
      panic("missing hourFactory")
   }

   return &MemoryHourRepository{
      hours:       map[time.Time]hour.Hour{},
      lock:        &sync.RWMutex{},
      hourFactory: hourFactory,
   }
}

func (m MemoryHourRepository) GetOrCreateHour(_ context.Context, hourTime time.Time) (*hour.Hour, error) {
   m.lock.RLock()
   defer m.lock.RUnlock()

   return m.getOrCreateHour(hourTime)
}

func (m MemoryHourRepository) getOrCreateHour(hourTime time.Time) (*hour.Hour, error) {
   currentHour, ok := m.hours[hourTime]
   if !ok {
      return m.hourFactory.NewNotAvailableHour(hourTime)
   }

   // we don't store hours as pointers, but as values 
   // thanks to that, we are sure that nobody can modify Hour without using UpdateHour
   return &currentHour, nil
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_memory_repository.go#L11)

Unfortunately, memory implementation has some downsides. The biggest one is that it doesn’t keep the data of the service after a restart. 😉 It can be enough for the functional pre-alpha version. To make our application production-ready, we need to have something a bit more persistent.

不幸的是，内存实现有一些缺点。最大的一个缺点是它不能在重启后保留服务的数据。 😉对于功能上的`pre-alpha`版本来说，这已经足够了。为了使我们的应用程序可以投入生产，我们需要一些更持久的东西。

## Example MySQL implementation  MySQL的实现示例
We already know how our model looks like and how it behaves. Based on that, let’s define our SQL schema.

我们已经知道我们的模型是怎样的，它的行为是怎样的。在此基础上，让我们定义我们的`SQL schema`。
```sql
CREATE TABLE `hours`
(
    hour         TIMESTAMP                                                 NOT NULL,
    availability ENUM ('available', 'not_available', 'training_scheduled') NOT NULL,
    PRIMARY KEY (hour)
);
```
[sql定义](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/sql/schema.sql#L1)

When I work with SQL databases, my default choices are:
当我使用SQL数据库时，我的默认选择是:

- sqlx – for simpler data models, it provides useful functions that help with using structs to unmarshal query results. When the schema is more complex because of relations and multiple models, it’s time for…
- sqlx[](https://github.com/jmoiron/sqlx) -- 对于简单的数据模型，它提供了有用的功能，可以帮助使用结构来解开查询结果。当模式因为关系和多个模型而比较复杂时，就需要...


- SQLBoiler[](https://github.com/volatiletech/sqlboiler) - is excellent for more complex models with many fields and relations, it’s based on code generation. Thanks to that, it’s very fast, and you don’t need to be afraid that you passed invalid interface{} instead of another interface{}. 😉 Generated code is based on the SQL schema, so you can avoid a lot of duplication.
- SQLBoiler - 对于有许多字段和关系的更复杂的模型来说是非常好的，它是基于代码生成的。😉 生成的代码是基于SQL模式的，所以你可以避免大量的重复。

We currently have only one table. sqlx will be more than enough 😉. Let’s reflect our DB model, with “transport type”.

我们目前只有一张表，`sqlx`就足够了😉。让我们用 "transport type" 来体现我们的 `DB model`。
```go
type mysqlHour struct {
   ID           string    `db:"id"`
   Hour         time.Time `db:"hour"`
   Availability string    `db:"availability"`
}
```
> You may ask why not to add the db attribute to hour.Hour? From my experience, it’s better to entirely separate domain types from the database. It’s easier to test, we are not duplicating validation, and it doesn’t introduce a lot of boilerplate.

> 你可能会问为什么不给`hour.Hour`添加`db`属性？根据我的经验，最好将域类型和数据库完全分开。这样更容易测试，我们不会重复验证，也不会引入很多的模板。

> In case of any change in the schema, we can do it just in our repository implementation, not in the half of the project. Miłosz described a similar case in “Things to know about DRY” article.

> 万一模式有任何变化，我们可以只在我们的仓库实现中进行，而不是在项目的一半。**Miłosz**在 [Things to know about DRY](https://threedots.tech/post/things-to-know-about-dry/)一文中描述了类似的情况。

> I also described that rule deeper in the previous article about DDD Lite.

> 我在之前关于[DDD Lite](https://threedots.tech/post/ddd-lite-in-go-introduction/#the-third-rule---domain-needs-to-be-database-agnostic)的文章中也深入描述了这个规则。


How can we use this struct?
那我们如何使用这个 `struct`？ 下述代码展示了使用案例：

```go
// sqlContextGetter is an interface provided both by transaction and standard db connection
type sqlContextGetter interface {
   GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func (m MySQLHourRepository) GetOrCreateHour(ctx context.Context, time time.Time) (*hour.Hour, error) {
	return m.getOrCreateHour(ctx, m.db, time, false)
}

func (m MySQLHourRepository) getOrCreateHour(
	ctx context.Context,
	db sqlContextGetter,
	hourTime time.Time,
	forUpdate bool,
) (*hour.Hour, error) {
	dbHour := mysqlHour{}

	query := "SELECT * FROM `hours` WHERE `hour` = ?"
	if forUpdate {
		query += " FOR UPDATE"
	}

	err := db.GetContext(ctx, &dbHour, query, hourTime.UTC())
	if errors.Is(err, sql.ErrNoRows) {
		// in reality this date exists, even if it's not persisted
		return m.hourFactory.NewNotAvailableHour(hourTime)
	} else if err != nil {
		return nil, errors.Wrap(err, "unable to get hour from db")
	}

	availability, err := hour.NewAvailabilityFromString(dbHour.Availability)
	if err != nil {
		return nil, err
	}

	domainHour, err := m.hourFactory.UnmarshalHourFromDatabase(dbHour.Hour.Local(), availability)
	if err != nil {
		return nil, err
	}

	return domainHour, nil
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_mysql_repository.go#L40)

With the SQL implementation, it’s simple because we don’t need to keep backward compatibility. In previous articles, we used Firestore as our primary database. Let’s prepare the implementation based on that, with keeping backward compatibility.

用SQL实现，很简单，因为我们不需要保持向后兼容。在之前的文章中，我们使用`Firestore`作为主数据库。我们在保持向后兼容的前提下，在此基础上准备实现。

## Firestore implementation Firestore的实现
When you want to refactor a legacy application, abstracting the database may be a good starting point.
当你想重构一个遗留应用程序时，抽象化数据库可能是一个好的起点。

Sometimes, applications are built in a database-centric way. In our case, it’s an HTTP Response-centric approach 😉 – our database models are based on Swagger generated models. In other words – our data models are based on Swagger data models that are returned by API. Does it stop us from abstracting the database? Of course not! It will need just some extra code to handle unmarshaling.

有时，应用程序是以数据库为中心的方式构建的。在我们的案例中，这是一种以`HTTP Response`为中心的方法😉--我们的数据库模型是基于`Swagger`生成的模型。换句话说--我们的数据模型是基于API返回的`Swagger`数据模型。这是否会阻止我们对数据库进行抽象呢? 当然不会！我们的数据模型是基于`Swagger`生成的模型。它只需要一些额外的代码来处理`unmarshaling`。

**With Domain-First approach, our database model would be much better, like in the SQL implementation**. But we are where we are. Let’s cut this old legacy step by step. I also have the feeling that CQRS will help us with that. 😉

**如果采用Domain-First的方法，我们的数据库模型会好很多，就像在SQL的实现中一样**。但我们就是这样。让我们一步步的砍掉这个老旧的传统。我也感觉到CQRS会帮助我们解决这个问题。 😉。


> In practice, a migration of the data may be simple, as long as no other services are integrated directly via the database.

> 在实践中，只要没有其他服务直接通过数据库集成，数据的迁移可能很简单。

> Unfortunatly, it’s an optimistic assumption when we work with a legacy response/database-centric or CRUD service…

> 不幸的是，当我们使用传统的响应/以数据库为中心或CRUD服务时，这是一个乐观的假设......

```go
func (f FirestoreHourRepository) GetOrCreateHour(ctx context.Context, time time.Time) (*hour.Hour, error) {
   date, err := f.getDateDTO(
      // getDateDTO should be used both for transactional and non transactional query,
      // the best way for that is to use closure
      func() (doc *firestore.DocumentSnapshot, err error) {
         return f.documentRef(time).Get(ctx)
      },
      time,
   )
   if err != nil {
      return nil, err
   }

   hourFromDb, err := f.domainHourFromDateDTO(date, time)
   if err != nil {
      return nil, err
   }

   return hourFromDb, err
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_firestore_repository.go#L31)

```go
// for now we are keeping backward comparability, because of that it's a bit messy and overcomplicated
// todo - we will clean it up later with CQRS :-)
func (f FirestoreHourRepository) domainHourFromDateDTO(date Date, hourTime time.Time) (*hour.Hour, error) {
   firebaseHour, found := findHourInDateDTO(date, hourTime)
   if !found {
      // in reality this date exists, even if it's not persisted
      return f.hourFactory.NewNotAvailableHour(hourTime)
   }

   availability, err := mapAvailabilityFromDTO(firebaseHour)
   if err != nil {
      return nil, err
   }

   return f.hourFactory.UnmarshalHourFromDatabase(firebaseHour.Hour.Local(), availability)
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_firestore_repository.go#L120)

Unfortunately, the Firebase interfaces for the transactional and non-transactional queries are not fully compatible. To avoid duplication, I created `getDateDTO` that can handle this difference by passing `getDocumentFn`.

不幸的是，事务性查询和非事务性查询的`Firebase`接口并不完全兼容。为了避免重复，我创建了`getDateDTO`，可以通过传递`getDocumentFn`来处理这种差异。

```go
func (f FirestoreHourRepository) getDateDTO(
   getDocumentFn func() (doc *firestore.DocumentSnapshot, err error),
   dateTime time.Time,
) (Date, error) {
   doc, err := getDocumentFn()
   if status.Code(err) == codes.NotFound {
      // in reality this date exists, even if it's not persisted
      return NewEmptyDateDTO(dateTime), nil
   }
   if err != nil {
      return Date{}, err
   }

   date := Date{}
   if err := doc.DataTo(&date); err != nil {
      return Date{}, errors.Wrap(err, "unable to unmarshal Date from Firestore")
   }

   return date, nil
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_firestore_repository.go#L97)

Even if some extra code is needed, it’s not bad. And at least it can be tested easily.

即使需要一些额外的代码，也不坏。而且至少可以轻松测试。

## Updating the data   更新数据
As I mentioned before - it’s critical to be sure that **only one person can schedule a training in one hour**. To handle that scenario, we need to use optimistic locking and transactions. Even if `transactions` is a pretty common term, let’s ensure that we are on the same page with `Optimistic Locking`.

正如我之前提到的那样--**确保一个小时内只能安排一个人进行培训**是至关重要的。为了处理这种情况，我们需要使用乐观锁和事务。即使事务是一个很常见的术语，我们也要确保我们对乐观锁的理解是一致的。

> **Optimistic concurrency control** assumes that many transactions can frequently complete without interfering with each other. While running, transactions use data resources without acquiring locks on those resources. Before committing, each transaction verifies that no other transaction has modified the data it has read. If the check reveals conflicting modifications, the committing transaction rolls back and can be restarted.

> **乐观的并发控制**假设许多事务可以经常完成而不互相干扰。在运行时，事务使用数据资源而不获取这些资源的锁。在提交之前，每个事务都会验证是否没有其他事务修改过它所读取的数据。如果检查发现有冲突的修改，提交的事务就会回滚，可以重新开始。

Technically transactions handling is not complicated. The biggest challenge that I had was a bit different – how to manage transactions in a clean way that does not affect the rest of the application too much, is not dependent on the implementation, and is explicit and fast.

技术上事务处理并不复杂。我遇到的最大的挑战有点不同--如何以一种干净的方式管理事务，不对应用程序的其他部分造成太大影响，不依赖于实现，并且是显式的和快速的。

I experimented with many ideas, like passing transaction via `context.Context`, handing transaction on HTTP/gRPC/message middlewares level, etc. All approaches that I tried had many major issues – they were a bit magical, not explicit, and slow in some cases.

我尝试了很多想法，比如通过context.Context传递事务，在HTTP/gRPC/消息中间件级别上处理事务等等。我试过的所有方法都存在很多重大问题--它们有点魔幻，不显式，有些情况下还很慢。

Currently, my favorite one is an approach based on closure passed to the update function.

目前，我最喜欢的是一种基于传递给更新函数的闭包的方法。
```go
type Repository interface {
   // ...
   UpdateHour(
      ctx context.Context,
      hourTime time.Time,
      updateFn func(h *Hour) (*Hour, error),
   ) error
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/domain/hour/repository.go#L8)

The basic idea is that we when we run UpdateHour, we need to provide updateFn that can update the provided hour.

基本的思路是，我们在运行`UpdateHour`的时候，需要提供`updateFn`，可以更新提供的小时数。

So in practice in one transaction we:
所以在实践中，在一个事务中我们可以:

- get and provide all parameters for updateFn (h *Hour in our case) based on provided UUID or any other parameter (in our case hourTime time.Time)
- execute the closure (updateFn in our case)
- save return values (*Hour in our case, if needed we can return more)
- execute rollback in case of an error returned from the closure


- 根据所提供的`UUID`或任何其他参数（在我们的例子中是 `hourTime time.Time`）获取并提供 `updateFn (h *Hour in our case)`的所有参数（在我们的例子中是`h *Hour`）。
- 执行闭包(在我们的例子中是`updateFn`)
- 保存返回值（在我们的例子中是`*Hour`，如果需要，我们可以返回更多）。
- 在关闭过程中发生错误时执行回滚操作

How is it used in practice?

在实践中是如何使用的？

```go
func (g GrpcServer) MakeHourAvailable(ctx context.Context, request *trainer.UpdateHourRequest) (*trainer.EmptyResponse, error) {
   trainingTime, err := protoTimestampToTime(request.Time)
   if err != nil {
      return nil, status.Error(codes.InvalidArgument, "unable to parse time")
   }

   if err := g.hourRepository.UpdateHour(ctx, trainingTime, func(h *hour.Hour) (*hour.Hour, error) { // Closure functions
      if err := h.MakeAvailable(); err != nil {
         return nil, err
      }

      return h, nil
   }); err != nil {
      return nil, status.Error(codes.Internal, err.Error())
   }

   return &trainer.EmptyResponse{}, nil
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/0249977c58a310d343ca2237c201b9ba016b148e/internal/trainer/grpc.go#L20)

As you can see, we get Hour instance from some (unknown!) database. After that, we make this hour Available. If everything is fine, we save the hour by returning it. **As part of [previous article](https://threedots.tech/post/ddd-lite-in-go-introduction/), all validations were moved the domain level, so here we are sure that we aren’t doing anything “stupid”. It also simplified this code a lot.**

正如你所看到的，我们从某个（未知的！）数据库中获取`Hour`实例。之后，我们让这个`hour`成为`Available`。如果一切正常，我们通过返回来保存这个`hour`。**作为[前一篇文章](https://threedots.tech/post/ddd-lite-in-go-introduction/)的一部分，所有的验证都被移到了`domain`级别，所以在这里我们确信我们没有做任何 "愚蠢 "的事情。这也简化了很多这段代码。**

```go
+func (g GrpcServer) MakeHourAvailable(ctx context.Context, request *trainer.UpdateHourRequest) (*trainer.EmptyResponse, error) {
@ ...
-func (g GrpcServer) UpdateHour(ctx context.Context, req *trainer.UpdateHourRequest) (*trainer.EmptyResponse, error) {
-	trainingTime, err := grpcTimestampToTime(req.Time)
-	if err != nil {
-		return nil, status.Error(codes.InvalidArgument, "unable to parse time")
-	}
-
-	date, err := g.db.DateModel(ctx, trainingTime)
-	if err != nil {
-		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to get data model: %s", err))
-	}
-
-	hour, found := date.FindHourInDate(trainingTime)
-	if !found {
-		return nil, status.Error(codes.NotFound, fmt.Sprintf("%s hour not found in schedule", trainingTime))
-	}
-
-	if req.HasTrainingScheduled && !hour.Available {
-		return nil, status.Error(codes.FailedPrecondition, "hour is not available for training")
-	}
-
-	if req.Available && req.HasTrainingScheduled {
-		return nil, status.Error(codes.FailedPrecondition, "cannot set hour as available when it have training scheduled")
-	}
-	if !req.Available && !req.HasTrainingScheduled {
-		return nil, status.Error(codes.FailedPrecondition, "cannot set hour as unavailable when it have no training scheduled")
-	}
-	hour.Available = req.Available
-
-	if hour.HasTrainingScheduled && hour.HasTrainingScheduled == req.HasTrainingScheduled {
-		return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf("hour HasTrainingScheduled is already %t", hour.HasTrainingScheduled))
-	}
-
-	hour.HasTrainingScheduled = req.HasTrainingScheduled
-	if err := g.db.SaveModel(ctx, date); err != nil {
-		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to save date: %s", err))
-	}
-
-	return &trainer.EmptyResponse{}, nil
-}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/commit/0249977c58a310d343ca2237c201b9ba016b148e#diff-5e57cb39050b6e252711befcf6fb0a89L20)

In our case from `updateFn` we return only `(*Hour, error)` – **you can return more values if needed**. You can return event sourcing events, read models, etc.

在我们的例子中，从`updateFn`我们只返回`(*Hour, error)`--**如果需要的话，你可以返回更多的值**。你可以返回事件来源事件，读取模型等。


We can also, in theory, use the same `hour.*Hour`, that we provide to `updateFn`. I decided to not do that. Using the returned value gives us more flexibility (we can replace a different instance of `hour.*Hour` if we want).

理论上，我们也可以使用同样的`hour.*Hour`，我们提供给`updateFn`。我决定不这样做。使用返回的值给了我们更多的灵活性（如果我们愿意，我们可以替换不同的`hour.*Hour`实例）。

It’s also nothing terrible to have multiple `UpdateHour`-like functions created with extra data to save. Under the hood, the implementation should re-use the same code without a lot of duplication.

创建多个类似`UpdateHour`的函数，并保存额外的数据，也没什么可怕的。在底层，实现上应该重用相同的代码，而不会有很多重复的地方。

```go
type Repository interface {
   // ...
   UpdateHour(
      ctx context.Context,
      hourTime time.Time,
      updateFn func(h *Hour) (*Hour, error),
   ) error

    UpdateHourWithMagic(
      ctx context.Context,
      hourTime time.Time,
      updateFn func(h *Hour) (*Hour, *Magic, error),
   ) error
}
```

How to implement it now?
现在如何实现？



## In-memory transactions implementation  内存事务的实现
The memory implementation is again the simplest one. 😉 We need to get the current value, execute closure, and save the result.

内存的实现又是最简单的，😉我们需要获取当前值，执行闭包，然后保存结果。

What is essential, in the map, we store a copy instead of a pointer. Thanks to that, we are sure that without the “commit” (`m.hours[hourTime] = *updatedHour`) our values are not saved. We will double-check it in tests.

最重要的是，在`map`中，我们存储了一个副本而不是一个指针。多亏了这一点，我们可以确定，如果没有 "提交"(m.hours[hourTime] = *updatedHour)，我们的值就不会被保存。我们将在测试中仔细检查它。
```go
func (m *MemoryHourRepository) UpdateHour(
   _ context.Context,
   hourTime time.Time,
   updateFn func(h *hour.Hour) (*hour.Hour, error),
) error {
   m.lock.Lock()
   defer m.lock.Unlock()

   currentHour, err := m.getOrCreateHour(hourTime)
   if err != nil {
      return err
   }

   updatedHour, err := updateFn(currentHour)
   if err != nil {
      return err
   }

   m.hours[hourTime] = *updatedHour

   return nil
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_memory_repository.go#L48)

## Firestore transactions implementation Firestore事务的实现
Firestore implementation is a bit more complex, but again – it’s related to backward compatibility. Functions `getDateDTO`, `domainHourFromDateDTO`, `updateHourInDataDTO` could be probably skipped when our data model would be better. Another reason to not use Database-centric/Response-centric approach!

`Firestore`的实现比较复杂，但同样--这与后向兼容性有关。函数`getDateDTO`，`domainHourFromDateDTO`，`updateHourInDataDTO`可能可以跳过，当我们的数据模型会更好。这是另一个不使用以数据库为中心/以响应为中心的方法的原因!
```go
func (f FirestoreHourRepository) UpdateHour(
   ctx context.Context,
   hourTime time.Time,
   updateFn func(h *hour.Hour) (*hour.Hour, error),
) error {
   err := f.firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
      dateDocRef := f.documentRef(hourTime)

      firebaseDate, err := f.getDateDTO(
         // getDateDTO should be used both for transactional and non transactional query,
         // the best way for that is to use closure
         func() (doc *firestore.DocumentSnapshot, err error) {
            return transaction.Get(dateDocRef)
         },
         hourTime,
      )
      if err != nil {
         return err
      }

      hourFromDB, err := f.domainHourFromDateDTO(firebaseDate, hourTime)
      if err != nil {
         return err
      }

      updatedHour, err := updateFn(hourFromDB)
      if err != nil {
         return errors.Wrap(err, "unable to update hour")
      }
      updateHourInDataDTO(updatedHour, &firebaseDate)

      return transaction.Set(dateDocRef, firebaseDate)
   })

   return errors.Wrap(err, "firestore transaction failed")
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_firestore_repository.go#L52)

As you can see, we get `*hour.Hour`, call `updateFn`, and save results inside of `RunTransaction`.

如你所见，我们得到`*hour.Hour`，调用`updateFn`，并将结果保存在`RunTransaction`内部。

Despite some extra complexity, this implementation is still clear, easy to understand and test.

尽管有一些额外的复杂性，但这个实现还是很清晰的，易于理解和测试。

## MySQL transactions implementation  MySQL事务的实现
Let’s compare it with MySQL implementation, where we designed models in a better way. Even if the way of implementation is similar, the biggest difference is a way of handling transactions.

我们和MySQL的实现方式进行比较，我们设计模型的方式更好。即使实现方式相似，但最大的区别是处理事务的方式。

In the SQL driver, the transaction is represented by `*db.Tx`. We use this particular object to call all queries and do a rollback and commit. To ensure that we don’t forget about closing the transaction, we do rollback and commit in the `defer`.

在` SQL driver`中，事务由`*db.Tx`表示。我们使用这个特殊的对象来调用所有的查询，并进行回滚和提交。为了保证我们不忘记关闭事务，我们`在defer`中进行回滚和提交。

```go
func (m MySQLHourRepository) UpdateHour(
   ctx context.Context,
   hourTime time.Time,
   updateFn func(h *hour.Hour) (*hour.Hour, error),
) (err error) {
   tx, err := m.db.Beginx()
   if err != nil {
      return errors.Wrap(err, "unable to start transaction")
   }

   // Defer is executed on function just before exit.
   // With defer, we are always sure that we will close our transaction properly.
   defer func() {
      // In `UpdateHour` we are using named return - `(err error)`.
      // Thanks to that that can check if function exits with error.
      //
      // Even if function exits without error, commit still can return error.
      // In that case we can override nil to err `err = m.finish...`.
      err = m.finishTransaction(err, tx)
   }()

   existingHour, err := m.getOrCreateHour(ctx, tx, hourTime, true)
   if err != nil {
      return err
   }

   updatedHour, err := updateFn(existingHour)
   if err != nil {
      return err
   }

   if err := m.upsertHour(tx, updatedHour); err != nil {
      return err
   }

   return nil
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_mysql_repository.go#L82)

In that case, we also get the hour by passing `forUpdate == true` to `getOrCreateHour`. This flag is adding `FOR UPDATE` statement to our query. The `FOR UPDATE` statement is critical because without that, we will allow parallel transactions to change the hour.

在这种情况下，我们也可以通过传递`forUpdate == true`到`getOrCreateHour`来获取`hour`。这个标志是在我们的查询中添加`FOR UPDATE`语句。`FOR UPDATE`语句是至关重要的，因为如果没有这个语句，我们将允许并行事务来改变`hour`。

> SELECT ... FOR UPDATE

> For index records the search encounters, locks the rows and any associated index entries, the same as if you issued an UPDATE statement for those rows. Other transactions are blocked from updating those rows.

> 对于搜索遇到的索引记录，锁定行和任何相关的索引条目，就像对这些行发出UPDATE语句一样。其他事务被阻止更新这些行。

```go
func (m MySQLHourRepository) getOrCreateHour(
   ctx context.Context,
   db sqlContextGetter,
   hourTime time.Time,
   forUpdate bool,
) (*hour.Hour, error) {
   dbHour := mysqlHour{}

   query := "SELECT * FROM `hours` WHERE `hour` = ?"
   if forUpdate {
      query += " FOR UPDATE"
   }

    // ...
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_mysql_repository.go#L48)

I never sleep well when I don’t have an automatic test for code like that. Let’s address it later. 😉

当我没有自动测试这样的代码时，我总是睡不好觉。😉我们稍后再解决这个问题。


`finishTransaction` is executed, when `UpdateHour` exits. When commit or rollback failed, we can also override the returned error.

当`UpdateHour`退出时，`finishTransaction`被执行。当提交或回滚失败时，我们也可以覆盖返回的错误。

```go
// finishTransaction rollbacks transaction if error is provided.
// If err is nil transaction is committed.
//
// If the rollback fails, we are using multierr library to add error about rollback failure.
// If the commit fails, commit error is returned.
func (m MySQLHourRepository) finishTransaction(err error, tx *sqlx.Tx) error {
   if err != nil {
      if rollbackErr := tx.Rollback(); rollbackErr != nil {
         return multierr.Combine(err, rollbackErr)
      }

      return err
   } else {
      if commitErr := tx.Commit(); commitErr != nil {
         return errors.Wrap(err, "failed to commit tx")
      }

      return nil
   }
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_mysql_repository.go#L149)

```go
// upsertHour updates hour if hour already exists in the database.
// If your doesn't exists, it's inserted.
func (m MySQLHourRepository) upsertHour(tx *sqlx.Tx, hourToUpdate *hour.Hour) error {
   updatedDbHour := mysqlHour{
      Hour:         hourToUpdate.Time().UTC(),
      Availability: hourToUpdate.Availability().String(),
   }

   _, err := tx.NamedExec(
      `INSERT INTO 
         hours (hour, availability) 
      VALUES 
         (:hour, :availability)
      ON DUPLICATE KEY UPDATE 
         availability = :availability`,
      updatedDbHour,
   )
   if err != nil {
      return errors.Wrap(err, "unable to upsert hour")
   }

   return nil
}
```
[相关源码](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb/internal/trainer/hour_mysql_repository.go#L122)

## Summary 总结
Even if the repository approach is adding a bit more code, it’s totally worth making that investment. **In practice, you may spend 5 minutes more to do that, and the investment should pay you back shortly.**
即使仓库的方法是多加一点代码，也完全值得做这个投资。**在实践中，你可能会多花5分钟时间来做这件事，投资应该很快就会回报给你。**


In this article, we are missing one essential part – tests. Now adding tests should be much easier, but it still may not be obvious how to do it properly.
To not make a “monster” article, I will cover it in the next 1-2 weeks. 🙂 Anyway, the entire diff of this refactoring, including tests, is already available on GitHub.

在这篇文章中，我们缺少一个必不可少的部分--测试。现在，添加测试应该更容易了，但如何正确地进行测试，可能还是不明显。
为了不写出一篇 "怪兽 "文章，我将在接下来的1-2周内介绍它。 总之，包括测试在内的整个重构的差异已经在[`GitHub`](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/commit/34c74e9d2cbc80160b4ff26e59818a89d10aa1eb)上提供了。

And just to remind – you can also run the application with one command and find the entire source code on GitHub!

提醒一下--你也可以用一个命令](https://threedots.tech/post/serverless-cloud-run-firebase-modern-go-application/#running)来运行这个应用程序，并在[GitHub](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example)上找到整个源代码!

Another technique that works pretty well is Clean/Hexagonal architecture – Miłosz is already working on the article covering that. Stay tuned!

另一种效果不错的技术是Clean/Hexagonal架构--Miłosz已经在写关于这方面的文章了。敬请期待








