# Go 微服务项目中应用DDD
【自己学习留存用，因为每次访问都非常耗时，且容易超时】

当我开始从事Go工作时，社区对DDD（领域驱动设计）和清洁架构等技术并不看好。我听到过很多次。"不要在Golang里做Java！"，"我在Java里见过这个，请不要！"。

这些时候，我已经有了近10年的PHP和Python的经验。我已经在那里看到了太多糟糕的事情。我记得所有这些 “Eight-thousanders”（代码超过8千行的方法😉）和没有人愿意维护的应用程序。我在检查这些丑陋怪物的旧git历史，它们在开始时都看起来无害。但随着时间的推移，小的、无辜的问题开始变得越来越大，越来越严重。**我也看到了DDD和清洁架构是如何解决这些问题的**。

也许Golang是不同的？也许用Golang编写微服务可以解决这个问题？
## It was supposed to be so beautiful  本应该是美好的
现在，在与多人交流经验并有能力看到很多代码库后，我的观点比3年前要干净一些。不幸的是，我现在远不认为仅仅使用Golang和微服务就能把我们从之前遇到的这些问题中拯救出来。我开始真正地从以前的、糟糕的时代中flashbacks(闪回、病理性重现)。

因为代码库相对年轻，所以不那么明显。因为Golang的设计，它不太明显。但我相信，随着时间的推移，我们会有越来越多的遗留Golang应用程序，没有人愿意去维护。

幸运的是，3年前，尽管受到了冷遇，但我并没有放弃。我决定尝试使用DDD和相关技术，这些技术之前在Go中对我很有效。与Milosz一起，我们领导的团队在3年内成功地使用了DDD、Clean Architecture以及所有相关的、在Golang中并不流行的技术。**他们让我们有能力以恒定的速度开发我们的应用程序和产品，无论代码的年龄如何**。

从一开始就很明显，从其他技术中**1:1转移模式是行不通的**。最重要的是，我们没有放弃惯用的Go代码和微服务架构--它们完美地结合在了一起

今天，我想与你分享第一个最直接的技术--DDD lite。

## State of DDD in Golang Go在DDD中状况
在开始写这篇文章之前，我在谷歌上查了几篇关于Go中的DDD的文章。我将在此粗暴(brutal)地指出：它们都忽略了使DDD发挥作用的最关键点。**如果我想象自己在没有任何DDD知识的情况下阅读这些文章，我不会被鼓励在我的团队中使用它们。这种肤浅的方法也可能是DDD在Go界仍未广泛应用的原因。**

在这个系列中，我们试图展示所有的基本技术，并以最实用的方式进行。在描述任何模式之前，我们以一个问题开始：**它能给我们带来什么？这是一个挑战我们当前思维的绝佳方式**。

我相信，我们可以改变 Go 社区对这些技术(DDD)的接受程度。我们相信，它们是实现复杂商业项目的最佳方式。我相信，我们将帮助确立Go的地位，使其成为一种不仅适用于构建基础设施，而且适用于商业软件的伟大语言。

## You need to go slow, to go fast  你需要走得慢，才能走得快
以最简单的方式实施你所从事的项目可能是很诱人的。当你感受到来自 "高层 "的压力时，这就更加诱人了。不过，我们是否使用微服务？如果需要的话，我们会不会直接重写服务？我听过多次这样的故事，而且很少有圆满的结局。 😉**走捷径的确可以节省一些时间。但那只是在短期内**。

让我们考虑一下任何种类的测试的例子。你可以在项目开始的时候跳过写测试。你显然会节省一些时间，管理层也会很高兴。**计算起来似乎很简单--因为这样做项目交付得更快。**

但从长远来看，这种捷径是不值得的。当项目增长时，你的团队会开始害怕做任何改变。最后，你所花费的时间总和将比从一开始就实施测试要多。**从长远来看，你会因为在一开始就为快速提高性能而牺牲质量而放慢速度**。另一方面，如果一个项目不是很关键，需要快速创建，你可以跳过测试。这应该是一个务实的决定，而不仅仅是 "我们知道得更多，而且我们不会产生bug"。

对于DDD来说，情况也是如此。当你想使用DDD时，你在开始时需要更多的时间，但长期的节省是巨大的。然而，并非每个项目都复杂到可以使用DDD这样的高级技术。

**不存在质量与速度的权衡。如果你想长期快速发展，你需要保持高质量。**

![](https://cdn.nlark.com/yuque/0/2021/png/2774323/1626660246716-5ed6a14f-8921-4e3d-b85b-c4c0589bda46.png)

> 图来自: 'Is High Quality Software Worth the Cost?' from [martinfowler.com](https://martinfowler.com/articles/is-quality-worth-cost.html)

## That’s great, but do you have any evidence it works?  这很好，但你有证据证明它很有效吗？
如果你两年前问我这个问题，我会说。"嗯，我觉得它的效果更好！"。但仅仅相信我的话可能是不够的。 😉**有许多教程展示了一些愚蠢的想法，并在没有任何证据的情况下声称它们有效--我们不要盲目地相信它们**。

幸运的是，2年前  [《Accelerate》: The Science of Lean Software and DevOps: Building and Scaling High Performing Technology Organizations ](https://www.amazon.com/Accelerate-Software-Performing-Technology-Organizations/dp/1942788339)。一书发布。简而言之，这本书描述了哪些因素影响开发团队的绩效。但这本书成名的原因是，它不仅仅是一套未经验证的想法--**它是基于科学研究的**。

我主要是对**展示什么能让团队成为顶级绩效团队**的部分感兴趣。这本书展示了一些显而易见的事实，比如介绍了DevOps、CI/CD和松散耦合的架构，这些都是高绩效团队的一个基本因素。

《Accelerate》一书介绍了高效率团队是什么样的：
> 我们发现，只要系统以及构建和维护它们的团队是松散耦合的，各种系统都可以实现高性能。

> 这一关键的架构属性使团队能够轻松地测试和部署单个组件或服务，即使组织和其运营的系统数量在增长--也就是说，它允许组织在扩大规模时提高其生产力。

所以，让我们使用微服务，我们就完成了？如果这就够了，我就不会写这篇文章了。 😉

书中还提到：
- > 对他们的系统设计进行大规模的修改，而不依赖其他团队对他们的系统进行修改或为其他团队创造大量的工作
- > 在不与团队以外的人沟通和协调的情况下完成他们的工作
- > 按需部署和发布他们的产品或服务，而不考虑它所依赖的其他服务
- > 按需进行大部分测试，不需要集成的测试环境 在正常工作时间内进行部署，停机时间可以忽略不计

> 不幸的是，在现实生活中，许多所谓的面向服务的架构并不允许测试和部署相互独立的服务，因此不会使团队获得更高的性能。

> [省略其他内容......]如果你忽略了这些特性，即使是采用部署在容器上的最先进的微服务架构并不能保证获得更高的性能。
> [省略其他内容......]为了实现这些特性，设计系统是松散耦合的--也就是说，这戏系统之间可以相互独立地改变和验证。

**仅仅使用微服务架构和将服务分割成小块是不够的。如果方式不对，就会增加额外的复杂性，并拖慢团队的速度**。在此，DDD可以帮助我们。

我多次提到DDD一词。DDD究竟是什么？

## What is DDD (Domain-Driven Design)  什么是DDD
让我们先看下Wiki上的定义：
> Domain-driven design (DDD) is the concept that the structure and language of your code (class names, class methods, class variables) should match the business domain. For example, if your software processes loan applications, it might have classes such as LoanApplication and Customer, and methods such as AcceptOffer and Withdraw.

是不是有以下类似的感觉：
[](https://cdn.nlark.com/yuque/0/2021/png/2774323/1626674028766-ecc13eff-c818-4a07-b63c-30ba3c5d0e97.png?x-oss-process=image%2Fresize%2Cw_800)

嗯~，这不是一个完美的描述。😅 它仍然缺少一些最重要的点。

还值得一提的是，**DDD是在2003年推出的**。那是相当长的时间了。一些提炼可能有助于将DDD放在2020年和Go的背景中。

我对DDD的简单定义是：
- 确保你以最佳方式解决**有效问题**。之后，以你的**企业能够理解的方式实施解决方案，而不需要任何额外的技术语言翻译**。

那如何实现这一目标？

## Coding is a war, to win you need a strategy! 编码是一场战争，要想取胜，你需要一个策略!

我喜欢说，"5天的编码可以节省15分钟的计划"。 (“5 days of coding can save 15 minutes of planning”.)

在开始写任何代码之前，你应该确保你正在解决一个有效的问题。这听起来很明显，但从我的经验来看，实际上并不像听起来那么容易。通常的情况是，工程师所创造的解决方案实际上并没有解决业务所要求的问题。在这个领域，有一套模式可以帮助我们，它被命名为战略DDD模式(Strategic DDD patterns.)。

根据我的经验，DDD战略模式经常被跳过。原因很简单：我们都是开发人员，我们喜欢写代码，而不是与 "业务人员 "交谈。 

😉不幸的是，当我们封闭在地下室而不与任何业务人员交谈时，这种方法有很多弊端: 
- 缺乏来自企业的信任
- 缺乏对系统工作原理的了解（从企业和工程方面）
- 解决错误的问题
这些只是一些最常见的问题。

好消息是，在大多数情况下，它是由缺乏适当的技术如事件风暴(Event Storming)造成的。它们可以给双方带来优势。同样令人惊讶的是，与企业交谈可能是工作中最令人愉快的部分之一

除此以外，我们将从适用于代码的模式开始。它们可以给我们带来DDD的一些优势。它们也会更快地对你有用。

**如果没有战略模式，我想说的是，你将只拥有DDD所能提供的30%的优势。我们将在接下来的文章中再谈战略模式。**

## DDD Lite in Go
在经历了相当长的介绍之后，现在终于到了接触一些代码的时候了 在这篇文章中，我们将介绍**Go中战术领域驱动设计模式**的一些基础知识。请记住，这只是一个开始。还需要几篇文章来涵盖整个主题。

战术DDD(Tactical DDD)最关键的部分之一是试图在代码中直接反映领域逻辑。

但这仍然是一些不具体的定义--而且在这一点上不需要。我也不想从描述什么是值对象(Value Objects)、实体(Entities)、聚合(Aggregates)开始。让我们最好从实际的例子开始。

## Wild workouts 一个示例项目
我还没有提到，特别是为了这些文章，我们创建了一个名为 `Wild Workouts` 的整个应用程序。

有趣的是，我们在这个应用程序中引入了一些微妙的问题，以便有东西可以重构。如果 `Wild Workouts` 看起来像你正在开发的一个应用程序--最好和我们一起呆一会儿😉。

### Refactoring of trainer service  重构 trainer service
我们要重构的第一个（微）服务是trainer service。我们现在将不碰其他服务--我们将在以后回到它们上面去。

这个服务具体业务: 
- 负责保持trainer的时间表，确保我们在一个小时内只能有一个培训计划。
- 它还保留了关于可用时间的信息（trainer的时间表）。

最初的实现并不是最好的。即使不是很多的逻辑，代码的某些部分也开始变得很乱。根据经验，我也有一些感觉，随着时间的推移，情况会越来越糟。

```go
func (g GrpcServer) UpdateHour(ctx context.Context, req *trainer.UpdateHourRequest) (*trainer.EmptyResponse, error) {
   trainingTime, err := grpcTimestampToTime(req.Time)
   if err != nil {
      return nil, status.Error(codes.InvalidArgument, "unable to parse time")
   }

   date, err := g.db.DateModel(ctx, trainingTime)
   if err != nil {
      return nil, status.Error(codes.Internal, fmt.Sprintf("unable to get data model: %s", err))
   }

   hour, found := date.FindHourInDate(trainingTime)
   if !found {
      return nil, status.Error(codes.NotFound, fmt.Sprintf("%s hour not found in schedule", trainingTime))
   }

   if req.HasTrainingScheduled && !hour.Available {
      return nil, status.Error(codes.FailedPrecondition, "hour is not available for training")
   }

   if req.Available && req.HasTrainingScheduled {
      return nil, status.Error(codes.FailedPrecondition, "cannot set hour as available when it have training scheduled")
   }
   if !req.Available && !req.HasTrainingScheduled {
      return nil, status.Error(codes.FailedPrecondition, "cannot set hour as unavailable when it have no training scheduled")
   }
   hour.Available = req.Available

   if hour.HasTrainingScheduled && hour.HasTrainingScheduled == req.HasTrainingScheduled {
      return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf("hour HasTrainingScheduled is already %t", hour.HasTrainingScheduled))
   }

   hour.HasTrainingScheduled = req.HasTrainingScheduled
   // do something
}
```
> Full source: github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/grpc.go

即使这不是有史以来最糟糕的代码，它也让我想起了我在检查我工作的代码的 git 历史时看到的情况。我可以想象，一段时间后，一些新的功能会出现，情况会更糟糕。

这里也很难模拟依赖关系，所以也没有单元测试。

### The First Rule - reflect your business logic literally  第一条规则--从字面上反映你的业务逻辑
在实现你的领域时，你应该停止把结构想成是假的数据结构或 "像ORM一样 "的实体，有一串`getter`和`setter`。相反，你应该把它们看成是有**行为的类型(types with behavior)**。

当你和你的商业利益相关者交谈时，他们会说 "I’m scheduling training on 13:00"，而不是 "I’m setting the attribute state to ‘training scheduled’ for hour 13:00."。

他们也不会说。"you can’t set attribute status to ‘training_scheduled’"。相反，它是 "You can’t schedule training if the hour is not available"。

但是如何把它直接写在代码中？

```go
func (h *Hour) ScheduleTraining() error {
   if !h.IsAvailable() {
      return ErrHourNotAvailable
   }

   h.availability = TrainingScheduled
   return nil
}
```
> Full source: github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/domain/hour/availability.go

一个可以帮助我们实施的问题是："如果没有任何额外的技术术语的翻译, business 会理解我的代码吗？"。你可以在这个以上代码中看到(语义化很丰富)，**当你安排培训时，即使不是技术人员也能理解**。

这种方法的成本(构建思路和写代码时间)并不高，而且有助于解决复杂性，使规则更容易理解。即使变化不大，我们也摆脱了这堵将来会变得更加复杂的ifs墙。

我们现在也能够轻松地添加单元测试，如下：
```go
func TestHour_ScheduleTraining(t *testing.T) {
   h, err := hour.NewAvailableHour(validTrainingHour())
   require.NoError(t, err)

   require.NoError(t, h.ScheduleTraining())

   assert.True(t, h.HasTrainingScheduled())
   assert.False(t, h.IsAvailable())
}

func TestHour_ScheduleTraining_with_not_available(t *testing.T) {
   h := newNotAvailableHour(t)
   assert.Equal(t, hour.ErrHourNotAvailable, h.ScheduleTraining())
}
```
> Full source: github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/domain/hour/availability_test.go
我们不需要在这里模拟任何东西。测试也是一个文档，帮助我们理解 `Hour` 的行为方式。

现在，如果有人问 "我什么时候可以安排培训 "的问题，你可以很快回答。

在一个更大、更复杂的系统中，这种问题的答案(指根据代码命名便能清晰理解业务)甚至不那么明显。你很可能需要花几个小时试图找到所有的地方，其中一些对象被以意外的方式使用，并且可能需要经历多次上述情况

下一条规则将更多地帮助我们解决这个问题。
### Testing Helpers  测试帮手
在测试中拥有一些创建我们的领域实体的辅助工具是很有用的。例如：`newExampleTrainingWithTime`，`newCanceledTraining`等。这也使得我们的领域测试更具有可读性。


自定义断言函数，比如 `assertTrainingsEquals`，这可以节省大量的重复工作。另外，[`github.com/google/go-cmp`](https://threedots.tech/post/ddd-lite-in-go-introduction/github.com/google/go-cmp)库对于比较复杂的结构极为有用。

它允许我们用私有字段比较我们的域类型，[跳过一些字段验证](https://godoc.org/github.com/google/go-cmp/cmp/cmpopts#IgnoreFields) 或 [实现自定义验证函数](https://pkg.go.dev/github.com/google/go-cmp/cmp?utm_source=godoc#Comparer)。

```go
func assertTrainingsEquals(t *testing.T, tr1, tr2 *training.Training) {
	cmpOpts := []cmp.Option{
		cmpRoundTimeOpt,
		cmp.AllowUnexported(
			training.UserType{},
			time.Time{},
			training.Training{},
		),
	}

	assert.True(
		t,
		cmp.Equal(tr1, tr2, cmpOpts...),
		cmp.Diff(tr1, tr2, cmpOpts...),
	)
}
```
为经常使用的构造函数提供 "Must"版本也是一个好主意，例如 "MustNewUser"。与普通的构造函数相比，如果参数无效，它们就会`panic`（对于测试来说，这并不是一个问题）。

```go
func NewUser(userUUID string, userType UserType) (User, error) {
	if userUUID == "" {
		return User{}, errors.New("missing user UUID")
	}
	if userType.IsZero() {
		return User{}, errors.New("missing user type")
	}

	return User{userUUID: userUUID, userType: userType}, nil
}

func MustNewUser(userUUID string, userType UserType) User {
	u, err := NewUser(userUUID, userType)
	if err != nil {
		panic(err)
	}

	return u
}
```

### The Second Rule: always keep a valid state in the memory  第二条规则：在内存中始终保持有效状态

> I recognize that my code will be used in ways I cannot anticipate, in ways it was not designed, and for longer than it was ever intended.
>  -- [The Rugged Manifesto](https://ruggedsoftware.org/)
> 
> 我认识到，我的代码将以我无法预料的方式被使用，以它没有被设计的方式被使用，而且使用的时间比它曾经被设计的时间更长。  -- [The Rugged Manifesto](https://ruggedsoftware.org/)

如果每个人都能考虑到这句话，世界就会变得更好。

根据我的观察，当你确信你所使用的对象总是有效的时候，它有助于避免很多的`if`和`bug`。你也会感到更有信心，因为你知道你不可能用当前的代码做任何愚蠢的事情。

我有很多回想，我不敢做一些改变，因为我不确定它的副作用。**如果不相信你能正确地使用代码，开发新功能的速度就会慢很多**

我们的目标是只在一个地方进行验证（良好的DRY），并确保没有人可以改变`Hour`的内部状态。该对象的唯一公共API应该是描述行为的方法。没有愚蠢的`getters`和`setters`！。我们还需要把我们的类型放到独立的包中，并使所有的属性成为私有。

```go
type Hour struct {
   hour time.Time 

   availability Availability
}

// ...

func NewAvailableHour(hour time.Time) (*Hour, error) {
   if err := validateTime(hour); err != nil {
      return nil, err
   }

   return &Hour{
      hour:         hour,
      availability: Available,
   }, nil
}
```
> Full source: github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/domain/hour/hour.go

我们还应该确保我们的类型里面没有违反任何规则。

**Bad example:**
```go
h := hour.NewAvailableHour("13:00") 

if h.HasTrainingScheduled() {
    h.SetState(hour.Available)
} else {
    return errors.New("unable to cancel training")
}
```
**Good example:**
```go
func (h *Hour) CancelTraining() error {
   if !h.HasTrainingScheduled() {
      return ErrNoTrainingScheduled
   }

   h.availability = Available
   return nil
}

// ...

h := hour.NewAvailableHour("13:00") 
if err := h.CancelTraining(); err != nil {
    return err
}
```
### The Third Rule - domain needs to be database agnostic  第三条规则--领域需要与数据库无关
这里有多种流派--有些人告诉我们，数据库客户端对`domain`定义的影响是可以的。根据我们的经验，严格保持`domain`不受任何数据库的影响效果最好。

最重要的原因是：
- 领域类型不是由所使用的数据库解决方案决定的--它们应该只由业务规则决定
- 我们可以以一种更理想的方式在数据库中存储数据
- 由于Go的设计和缺乏像注解那样的 "魔法"，ORM或任何数据库解决方案都会产生更重要的影响。

> **Domain-First approach**
> 
> 如果项目足够复杂，我们甚至可以花2-4周的时间在领域层工作，只用内存数据库实现。在这种情况下，我们可以更深入地探索这个想法，并推迟决定是否选择数据库。我们所有的实现都只是基于单元测试。

> 我们试过几次这种方法，总是很顺利。在这里有一些时间框架也是一个好主意，不要花太多的时间。

> 请记住，这种方法需要一个良好的关系和来自企业的大量信任！如果你与企业的关系远非如此。**如果你与业务的关系还不够好，战略DDD模式将改善这一状况。Been there, done that!**

为了不使本文冗长，我们只介绍一下`Repository`接口，并假设它能工作。 😉 

我将在下一篇文章中更深入地介绍这个话题。

```go
type Repository interface {
   GetOrCreateHour(ctx context.Context, time time.Time) (*Hour, error)
   UpdateHour(
      ctx context.Context,
      hourTime time.Time,
      updateFn func(h *Hour) (*Hour, error),
   ) error
}
```
> Full source: github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/domain/hour/repository.go

你可能会问为什么`UpdateHour`有`updateFn func(h *Hour) (*Hour, error)`--我们会用它来很好地处理事务。你可以在[the article about repositories](https://threedots.tech/post/repository-pattern-in-go/)中了解更多。

### Using domain objects 使用领域对象

我对我们的gRPC端点做了一个小的重构，以提供一个更 "behavior-oriented"的API，而不是CRUD。

它更好地反映了领域的新特性。根据我的经验，维护多个小方法比维护一个允许我们更新一切的 "God"方法要容易得多。以下提交记录，展示了我去掉了"大而全"的方法，使用了更有原子性的"小"方法：
```go
--- a/api/protobuf/trainer.proto
+++ b/api/protobuf/trainer.proto
@@ -6,7 +6,9 @@ import "google/protobuf/timestamp.proto";
 
 service TrainerService {
   rpc IsHourAvailable(IsHourAvailableRequest) returns (IsHourAvailableResponse) {}
-  rpc UpdateHour(UpdateHourRequest) returns (EmptyResponse) {}
+  rpc ScheduleTraining(UpdateHourRequest) returns (EmptyResponse) {}
+  rpc CancelTraining(UpdateHourRequest) returns (EmptyResponse) {}
+  rpc MakeHourAvailable(UpdateHourRequest) returns (EmptyResponse) {}
 }
 
 message IsHourAvailableRequest {
@@ -19,9 +21,6 @@ message IsHourAvailableResponse {
 
 message UpdateHourRequest {
   google.protobuf.Timestamp time = 1;
-
-  bool has_training_scheduled = 2;
-  bool available = 3;
 }
 
 message EmptyResponse {}
```
> Full source: github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/commit/0249977c58a310d343ca2237c201b9ba016b148e#diff-15fd9ad3f3992b0210090109b82c5594

现在的实现要简单得多，也更容易理解。我们在这里也没有逻辑--只是一些协调工作。我们的gRPC处理程序现在有18行并且还没领域逻辑
```go
func (g GrpcServer) MakeHourAvailable(ctx context.Context, request *trainer.UpdateHourRequest) (*trainer.EmptyResponse, error) {
   trainingTime, err := protoTimestampToTime(request.Time)
   if err != nil {
      return nil, status.Error(codes.InvalidArgument, "unable to parse time")
   }

   if err := g.hourRepository.UpdateHour(ctx, trainingTime, func(h *hour.Hour) (*hour.Hour, error) {
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
> Full source: github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/grpc.go

> **No more Eight-thousanders**
> 在我的记忆中，许多 Eight-thousanders 实际上是在controllers，在HTTP controllers 中含有大量的domain逻辑。

> 通过将复杂性隐藏在我们的domain types里面，并保持我所描述的规则，我们可以防止这个地方不受控制地增长。

## That’s all for today
我不想把这篇文章写得太长--让我们一步一步来吧!

如果你等不及了，整个重构的工作差异可以在GitHub上找到。在下一篇文章中，我将介绍这里没有解释的部分：repository。

即使这仍然是个开始，我们的代码中的一些简化是可见的。

目前的领域实现也不完美--这很好! 你永远不会从一开始就实施完美的model。**最好是准备好轻松地改变这个model，而不是浪费时间来使它完美**。

在我添加了model的测试，并将其与应用程序的其他部分分开后，我可以毫无顾虑地改变它。

## Can I already put that I know DDD to my CV? 我是否已经可以在我的简历中写上我知道DDD？
还没有。

在我听到DDD之后，我需要3年的时间才能把所有的点连接起来（那是在我听到Go之前）。 😉之后，我看到了为什么我们将在接下来的文章中描述的所有技术是如此重要。

但是在连接这些点之前，需要一些耐心和信任，相信会有效果。这是很值得的! 

你不会像我一样需要3年时间，但我们目前计划了大约10篇关于战略和战术模式的文章。 😉 这是在Wild Workouts中剩下的很多新功能和要重构的部分

我知道，现在很多人都承诺，只要看了一篇文章或一段视频10分钟，你就能成为某个领域的专家。如果能做到这一点，世界将是美好的，但在现实中，这并不是那么简单。

幸运的是，我们所分享的知识有很大一部分是通用的，可以应用于多种技术，而不仅仅是Go。

你可以把这些知识当作对你的职业生涯和心理健康的长期投资。 😉 没有什么比解决正确的问题，而不与不可维护的代码作斗争更好的了！。