# 组合DDD|CQRS|整洁架构

翻译自:[https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/](https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/)

In the previous articles, we introduced techniques like DDD Lite, CQRS, and Clean (Hexagonal) Architecture. Even if using them alone is beneficial, they work the best together. Like Power Rangers. Unfortunately, it is not easy to use them together in a real project. In this article, I will show you how to connect DDD Lite, CQRS, and Clean Architecture in the most pragmatic and efficient way.

在之前的文章中，我们介绍了DDD Lite、CQRS和Clean（六边形）架构等技术。即使单独使用它们是有益的，但它们在一起的效果是最好的。就像电力别动队一样。不幸的是，在实际项目中一起使用它们并不容易。在这篇文章中，我将向你展示如何以最实用、最有效的方式连接DDD Lite、CQRS和Clean Architecture。

## Why should I care? 我为什么要关心？

Working on a programming project is similar to planning and building a residential district. If you know that the district will be expanding in the near future, you need to keep space for future improvements. Even if at the beginning it may look like a waste of space. You should keep space for future facilities like residential blocks, hospitals, and temples. Without that, you will be forced to destroy buildings and streets to make space for new buildings. It’s much better to think about that earlier.

从事编程项目的工作类似于规划和建设一个住宅区。如果你知道该区在不久的将来会扩大，你需要为未来的改进保留空间。即使在开始的时候，这可能看起来是对空间的浪费。你应该为未来的设施如住宅区、医院和寺庙保留空间。如果没有这一点，你将被迫摧毁建筑物和街道，为新建筑腾出空间。早点考虑到这一点会好很多。

The situation is the same with the code. If you know that the project will be developed for longer than 1 month, you should keep the long term in mind from the beginning. You need to create your code in a way that will not block your future work. Even if at the beginning it may look like over-engineering and a lot of extra boilerplate, you need to keep in mind the long term.

代码的情况也是如此。如果你知道这个项目的开发时间将超过1个月，你就应该从一开始就把长期的事情放在心上。你需要以一种不会阻碍你未来工作的方式来创建你的代码。即使一开始看起来是过度工程和大量额外的模板，你也需要牢记长期的目标。

It doesn’t mean that you need to plan every feature that you will implement in the future – it’s actually the opposite one. This approach helps to adapt to new requirements or changing understanding of our domain. Big up front design is not needed here. It’s critical in current times, when the world is changing really fast and who can’t adapt to these changes can get simply out of business.

这并不意味着你需要计划你将来要实现的每一个功能--实际上恰恰相反。这种方法有助于适应新的需求或对我们领域不断变化的理解。这里不需要大的前期设计。这在目前的时代是很关键的，因为世界变化非常快，谁不能适应这些变化，谁就会被淘汰。

This is exactly what these patterns give you when they are combined – the ability to keep constant development speed. Without destroying and touching existing code too much.

这正是这些模式结合后给你带来的好处--保持恒定的开发速度的能力。而不需要破坏和过多地接触现有的代码。

Does it require more thinking and planning? Is it a more challenging way? Do you need to have extra knowledge to do that? Sure! But the long term result is worth that! Fortunately, you are in the right place to learn that. 😉

它是否需要更多的思考和计划？它是一种更有挑战性的方式吗？你需要有额外的知识才能做到这一点吗？当然！但长期的结果是值得的。幸运的是，你在正确的地方可以学到这些。

But let’s leave the theory behind us. Let’s go to the code. In this article, we will skip reasonings for our design choices. We described these already in the previous articles. If you did not read them yet, I recommend doing it – you will understand this article better.

但让我们把理论抛在脑后。让我们来看看代码。在这篇文章中，我们将跳过对我们的设计选择的推理。我们在以前的文章中已经描述了这些。如果你还没有读过这些文章，我建议你读一下--你会更好地理解这篇文章。

Like in previous articles, we will base our code on refactoring a real open-source project. This should make the examples more realistic and applicable to your projects.

像以前的文章一样，我们将把我们的代码建立在重构一个真正的开源项目上。这应该会使这些例子更加真实，并适用于你的项目。

Are you ready?

你准备好了吗？


## Let’s refactor 开始重构
Let’s start our refactoring with the Domain-First approach. We will start with introduction of a domain layer. Thanks to that, we will be sure that implementation details do not affect our domain code. We can also put all our efforts into understanding the business problem. Not on writing boring database queries and API endpoints.

让我们用领域优先的方法开始我们的重构。我们将从引入领域层开始。有了它，我们将确保实现细节不会影响我们的领域代码。我们也可以把所有的精力放在理解业务问题上。而不是去写无聊的数据库查询和API端点。

Domain-First approach works good for both rescue (refactoring 😉) and greenfield projects.

领域优先的方法对救援（重构😉）和全新项目都很有效。

To start building my domain layer, I needed to identify what the application is actually doing. This article will focus on refactoring of trainings Wild Workouts microservice. I started with identifying use cases handled by the application. After previous refactoring to Clean Architecture, we can find it in the TrainingService. When I work with a messy application, I look at RPC and HTTP endpoints to find supported use cases.

为了开始构建我的领域层，我需要确定应用程序实际上在做什么。本文将重点介绍培训Wild Workouts微服务的重构。我从识别应用所处理的用例开始。在之前重构为Clean Architecture之后，我们可以在TrainingService中找到它。当我在处理一个混乱的应用程序时，我会查看RPC和HTTP端点以找到支持的用例。

One of functionalities that I identified is the approval of training reschedule. In Wild Workouts, a training reschedule approval is required if it was requested less than 24h before its date. If a reschedule is requested by the attendee, the approval needs to be done by the trainer. When it’s requested by the trainer, it needs to be accepted by the attendee.

我发现的一个功能是批准培训的重新安排。在Wild Workouts中，如果在培训日期前24小时内申请，则需要批准培训改期。如果参加培训的人要求改期，则需要由培训师进行批准。当培训师提出要求时，需要由学员接受。
```go
- func (c TrainingService) ApproveTrainingReschedule(ctx context.Context, user auth.User, trainingUUID string) error {
-  return c.repo.ApproveTrainingReschedule(ctx, trainingUUID, func(training Training) (Training, error) {
-     if training.ProposedTime == nil {
-        return Training{}, errors.New("training has no proposed time")
-     }
-     if training.MoveProposedBy == nil {
-        return Training{}, errors.New("training has no MoveProposedBy")
-     }
-     if *training.MoveProposedBy == "trainer" && training.UserUUID != user.UUID {
-        return Training{}, errors.Errorf("user '%s' cannot approve reschedule of user '%s'", user.UUID, training.UserUUID)
-     }
-     if *training.MoveProposedBy == user.Role {
-        return Training{}, errors.New("reschedule cannot be accepted by requesting person")
-     }
-
-     training.Time = *training.ProposedTime
-     training.ProposedTime = nil
-
-     return training, nil
-  })
- }
```
## Start with the domain 从领域层开始
Even if it doesn’t look like the worst code you’ve seen in your life, functions like ApproveTrainingReschedule tend to get more complex over time. More complex functions mean more potential bugs during future development.

即使它看起来不像是你一生中见过的最糟糕的代码，像ApproveTrainingReschedule这样的函数往往会随着时间的推移变得越来越复杂。更复杂的函数意味着在未来的开发过程中可能出现更多的bug。

It’s even more likely if we are new to the project, and we don’t have the “shaman knowledge” about it. You should always consider all the people who will work on the project after you, and make it resistant to be broken accidentally by them. That will help your project to not become the legacy that everybody is afraid to touch. You probably hate that feeling when you are new to the project, and you are afraid to touch anything to not break the system.

如果我们是项目的新手，没有相关的 "萨满知识"，这种可能性就更大了。你应该总是考虑到在你之后将从事该项目工作的所有人员，并使其能够抵抗被他们意外地破坏。这将有助于你的项目不至于成为所有人都不敢碰的遗产。你可能讨厌这种感觉，当你是项目的新人时，你不敢碰任何东西，以免破坏系统。

It’s not uncommon for people to change their job more often than every 2 years. That makes it even more critical for long-term project development.

人们更换工作的频率超过每两年一次，这并不罕见。这使得项目的长期发展更加关键
。
If you don’t believe that this code may become complex, I recommend checking the Git history of the worst place in the project you work on. In most cases, that worst code started with “just a couple simple ifs”. 😉 The more complex the code will be, the more difficult it will be to simplify it later. We should be sensitive to emerging complexity and try to simplify it as soon as we can.

如果你不相信这段代码可能会变得复杂，我建议查看你工作的项目中最糟糕的地方的Git历史。在大多数情况下，那段最糟糕的代码是从 "只有几个简单的ifs "开始的。 😉代码越复杂，以后就越难简化它。我们应该对新出现的复杂性保持敏感，并试图尽快简化它。

### `Training` domain entity
During analyzing the current use cases handled by the trainings microservice, I found that they are all related to a training. It is pretty natural to create a Training type to handle these operations.
在分析当前由`trainings`微服务处理的用例时，我发现它们都与`training`有关。创建一个`Training`类型来处理这些操作是相当自然的。

![](https://cdn.nlark.com/yuque/0/2021/png/2774323/1629114836529-24c1481c-266a-4e52-b5b8-0ea0042d860d.png?x-oss-process=image%2Fresize%2Cw_833)

> noun == entity 名词==实体

Is it a valid approach to discover entities? Well, not really.

这是一个发现实体的有效方法吗？嗯，不尽然。

DDD provides tools that help us to model complex domains without guessing (Strategic DDD Patterns, Aggregates). We don’t want to guess how our aggregates look like – we want to have tools to discover them. Event Storming technique is extremely useful here… but it’s a topic for an entire separate article.

DDD提供的工具可以帮助我们对复杂的领域进行建模，而无需猜测（战略DDD模式，聚合）。我们不希望猜测我们的聚合体是什么样子的--我们希望有工具来发现它们。事件风暴技术在这里非常有用......但这是另一篇文章的主题。

The topic is complex enough to have a couple articles about that. And this is what we will do shortly. 😉

这个话题很复杂，足以有几篇关于这个的文章。这就是我们即将要做的事情。

Does it mean that you should not use these techniques without Strategic DDD Patterns? Of course not! The current approach can be good enough for simpler projects. Unfortunately (or fortunately 😉), not all projects are simple.

这是否意味着你不应该在没有战略DDD模式的情况下使用这些技术？当然不是! 对于比较简单的项目来说，目前的方法已经足够好了。不幸的是（或幸运的是😉），不是所有的项目都很简单。

```go
package training

// ...

type Training struct {
   uuid string

   userUUID string
   userName string

   time  time.Time
   notes string

   proposedNewTime time.Time
   moveProposedBy  UserType

   canceled bool
}
```

All fields are private to provide encapsulation. This is critical to meet “always keep a valid state in the memory” rule from the article about DDD Lite.

所有字段都是私有的，以提供封装。这对于满足DDD Lite文章中 "在内存中始终保持有效状态 "的规则至关重要。

Thanks to the validation in the constructor and encapsulated fields, we are sure that Training is always valid. Now, a person that doesn’t have any knowledge about the project is not able to use it in a wrong way.

由于构造函数和封装字段的验证，我们可以确保培训始终有效。现在，一个对该项目没有任何了解的人无法以错误的方式使用它。

The same rule applies to any methods provided by Training.

同样的规则也适用于 `Training` 所提供的任何方法。

```go
package training

func NewTraining(uuid string, userUUID string, userName string, trainingTime time.Time) (*Training, error) {
   if uuid == "" {
      return nil, errors.New("empty training uuid")
   }
   if userUUID == "" {
      return nil, errors.New("empty userUUID")
   }
   if userName == "" {
      return nil, errors.New("empty userName")
   }
   if trainingTime.IsZero() {
      return nil, errors.New("zero training time")
   }

   return &Training{
      uuid:     uuid,
      userUUID: userUUID,
      userName: userName,
      time:     trainingTime,
   }, nil
}
```
## Approve reschedule in the domain layer 批准重新安排——在领域层处理
As described in DDD Lite introduction, we build our domain with methods oriented on behaviours. Not on data. Let’s model ApproveReschedule on our domain entity.

正如`DDD Lite`介绍中所述，我们用面向行为的方法构建我们的领域。而不是在数据上。让我们在我们的领域实体上建立`ApproveReschedule`模型。
```go
package training

// ...s

func (t *Training) IsRescheduleProposed() bool {
   return !t.moveProposedBy.IsZero() && !t.proposedNewTime.IsZero()
}

var ErrNoRescheduleRequested = errors.New("no training reschedule was requested yet")

func (t *Training) ApproveReschedule(userType UserType) error {
   if !t.IsRescheduleProposed() {
      return errors.WithStack(ErrNoRescheduleRequested)
   }

   if t.moveProposedBy == userType {
      return errors.Errorf(
         "trying to approve reschedule by the same user type which proposed reschedule (%s)",
         userType.String(),
      )
   }

   t.time = t.proposedNewTime

   t.proposedNewTime = time.Time{}
   t.moveProposedBy = UserType{}

   return nil
}
```

If you had no chance to read:
如果你没有机会阅读：

- DDD Lite introduction, 
- Introducing Clean Architecture,
- Introducing basic CQRS,
- Repository pattern.

I highly recommend checking them. It will help you understand this article better. They explain the decisions and techniques that we combine in this article.

我强烈建议你查看它们。它将帮助你更好地理解这篇文章。它们解释了我们在这篇文章中结合的决定和技术。

### Orchestrate with command  用 command 进行协调
Now the application layer can be responsible only for the orchestration of the flow. There is no domain logic there. We hide the entire business complexity in the domain layer. This was exactly our goal.

现在，应用层可以只负责流程的协调。那里没有领域逻辑。我们将整个业务的复杂性隐藏在领域层中。这正是我们的目标。

For getting and saving a training, we use the Repository pattern.

为了获得和保存训练，我们使用了`Repository`模式。

```go
package command

// ...

func (h ApproveTrainingRescheduleHandler) Handle(ctx context.Context, cmd ApproveTrainingReschedule) (err error) {
   defer func() {
      logs.LogCommandExecution("ApproveTrainingReschedule", cmd, err)
   }()

   return h.repo.UpdateTraining(
      ctx,
      cmd.TrainingUUID,
      cmd.User,
      func(ctx context.Context, tr *training.Training) (*training.Training, error) {
         originalTrainingTime := tr.Time()

         if err := tr.ApproveReschedule(cmd.User.Type()); err != nil {
            return nil, err
         }

         err := h.trainerService.MoveTraining(ctx, tr.Time(), originalTrainingTime)
         if err != nil {
            return nil, err
         }

         return tr, nil
      },
   )
}
```

## Refactoring of training cancelation 重构培训取消服务
Let’s now take a look at CancelTraining from TrainingService.

现在让我们来看看`TrainingService`的`CancelTraining`

The domain logic is simple there: you can cancel a training up to 24h before its date. If it’s less than 24h before the training, and you want to cancel it anyway:

这里的领域逻辑很简单：你可以在培训日期前24小时内取消培训。如果在培训前不到24小时，而你还是想取消培训。

- if you are the trainer, the attendee will have his training “back” plus one extra session (nobody likes to change plans on the same day!)
- if you are the attendee, you will lose this training
  
- 如果你是培训师，会员将得到他的培训 "回报"，外加一次额外的课程（没有人喜欢在同一天改变计划！）。
- 如果你是会员，你将失去这次培训。

This is how the current implementation looks like:

这就是目前的实施情况：

```go
func (c TrainingService) CancelTraining(ctx context.Context, user auth.User, trainingUUID string) error {
-  return c.repo.CancelTraining(ctx, trainingUUID, func(training Training) error {
-     if user.Role != "trainer" && training.UserUUID != user.UUID {
-        return errors.Errorf("user '%s' is trying to cancel training of user '%s'", user.UUID, training.UserUUID)
-     }
-
-     var trainingBalanceDelta int
-     if training.CanBeCancelled() {
-        // just give training back
-        trainingBalanceDelta = 1
-     } else {
-        if user.Role == "trainer" {
-           // 1 for cancelled training +1 fine for cancelling by trainer less than 24h before training
-           trainingBalanceDelta = 2
-        } else {
-           // fine for cancelling less than 24h before training
-           trainingBalanceDelta = 0
-        }
-     }
-
-     if trainingBalanceDelta != 0 {
-        err := c.userService.UpdateTrainingBalance(ctx, training.UserUUID, trainingBalanceDelta)
-        if err != nil {
-           return errors.Wrap(err, "unable to change trainings balance")
-        }
-     }
-
-     err := c.trainerService.CancelTraining(ctx, training.Time)
-     if err != nil {
-        return errors.Wrap(err, "unable to cancel training")
-     }
-
-     return nil
-  })
- }
```
You can see some kind of “algorithm” for calculating training balance delta during cancelation. That’s not a good sign in the application layer.

你可以看到某种 "算法"，用于计算取消期间的训练平衡三角洲。这在应用层不是一个好兆头。

Logic like this one should live in our domain layer. If you start to see some `if`'s related to logic in your application layer, you should think about how to move it to the domain layer. It will be easier to test and re-use in other places.

像这样的逻辑应该在我们的领域层中。如果你开始看到一些与应用层中的逻辑有关的`if`，你应该考虑如何将它移到领域层。这将更容易测试和在其他地方重新使用。

It may depend on the project, but often domain logic is pretty stable after the initial development and can live unchanged for a long time. It can survive moving between services, framework changes, library changes, and API changes. Thanks to that separation, we can do all these changes in a much safer and faster way.

这可能取决于项目，但通常领域逻辑在最初的开发之后是相当稳定的，可以长期保持不变。它可以在服务、框架变化、库变化和API变化之间生存。由于这种分离，我们可以以更安全、更快速的方式进行所有这些改变。

Let’s decompose the `CancelTraining` method to multiple, separated pieces. That will allow us to test and change them independently.

让我们把`CancelTraining`方法分解成多个独立的部分。这将使我们能够独立地测试和改变它们。

First of all, we need to handle cancelation logic and marking Training as canceled.

首先，我们需要处理取消的逻辑，并将培训标记为已取消。
```go
package training

func (t Training) CanBeCanceledForFree() bool {
   return t.time.Sub(time.Now()) >= time.Hour*24
}

var ErrTrainingAlreadyCanceled = errors.New("training is already canceled")

func (t *Training) Cancel() error {
   if t.IsCanceled() {
      return ErrTrainingAlreadyCanceled
   }

   t.canceled = true
   return nil
}
```
Nothing really complicated here. That’s good!

这里没有真正复杂的东西。这很好!

The second part that requires moving is the “algorithm” of calculating trainings balance after cancelation. In theory, we could put it to the Cancel() method, but IMO it would break the Single Responsibility Principle and CQS. And I like small functions.

第二个需要移动的部分是取消后计算训练余额的 "算法"。理论上，我们可以把它放在Cancel()方法中，但我认为这将破坏单一职责原则和CQRS。而且我喜欢小函数。

But where to put it? Some object? A domain service? In some languages, like the one that starts with J and ends with ava, it would make sense. But in Go, it’s good enough to just create a simple function.


但要把它放在哪里呢？某个对象？一个领域服务？在某些语言中，比如以J开头，以ava结尾的语言，这将是有意义的。但在Go中，只要创建一个简单的函数就够了。

```go
package training

// CancelBalanceDelta return trainings balance delta that should be adjusted after training cancelation.
func CancelBalanceDelta(tr Training, cancelingUserType UserType) int {
   if tr.CanBeCanceledForFree() {
      // just give training back
      return 1
   }

   switch cancelingUserType {
   case Trainer:
      // 1 for cancelled training +1 "fine" for cancelling by trainer less than 24h before training
      return 2
   case Attendee:
      // "fine" for cancelling less than 24h before training
      return 0
   default:
      panic(fmt.Sprintf("not supported user type %s", cancelingUserType))
   }
}
```

The code is now straightforward. I can imagine that I could sit with any non-technical person and go through this code to explain how it works.

现在的代码是直截了当的。我可以想象，我可以和任何非技术人员坐在一起，通过这段代码来解释它是如何工作的。

What about tests? It may be a bit controversial, but IMO tests are redundant there. Test code would replicate the implementation of the function. Any change in the calculation algorithm will require copying the logic to the tests. I would not write a test there, but if you will sleep better at night – why not!

那么测试呢？这可能有点争议，但是IMO的测试在这里是多余的。测试代码将复制该函数的实现。计算算法的任何改变都需要将逻辑复制到测试中。我不会在那里写一个测试，但如果你晚上能睡得更好--为什么不呢！？

### Moving CancelTraining to command 将 CancelTraining 服务移动的 command 中
Our domain is ready, so let’s now use it. We will do it in the same way as previously:

我们的域名已经准备好了，所以现在让我们来使用它。我们将以与之前相同的方式进行。

1. getting the entity from the repository, 从资源库中获取实体。
2. orchestration of domain stuff,  协调域的事情 (可以理解为: 拼装domain中的业务数据)
3. calling external trainer service to cancel the training (this service is the point of truth of “trainer’s calendar”),  调用外部`trainer`服务来取消培训（这个服务是 "培训师日历 "的真实点）。
4. returning entity to be saved in the database.  返回实体以保存在数据库中。

```go
package command

// ...

func (h CancelTrainingHandler) Handle(ctx context.Context, cmd CancelTraining) (err error) {
   defer func() {
      logs.LogCommandExecution("CancelTrainingHandler", cmd, err)
   }()

   return h.repo.UpdateTraining(
      ctx,
      cmd.TrainingUUID,
      cmd.User,
      func(ctx context.Context, tr *training.Training) (*training.Training, error) {
         if err := tr.Cancel(); err != nil {
            return nil, err
         }

         if balanceDelta := training.CancelBalanceDelta(*tr, cmd.User.Type()); balanceDelta != 0 {
            err := h.userService.UpdateTrainingBalance(ctx, tr.UserUUID(), balanceDelta)
            if err != nil {
               return nil, errors.Wrap(err, "unable to change trainings balance")
            }
         }

         if err := h.trainerService.CancelTraining(ctx, tr.Time()); err != nil {
            return nil, errors.Wrap(err, "unable to cancel training")
         }

         return tr, nil
      },
   )
}
```
## Repository refactoring 将repository重构
The initial implementation of the repository was pretty tricky because of the custom method for every use case.

存储库的最初实现是相当棘手的，因为每个用例都有自定义的方法。
```go
- type trainingRepository interface {
-  FindTrainingsForUser(ctx context.Context, user auth.User) ([]Training, error)
-  AllTrainings(ctx context.Context) ([]Training, error)
-  CreateTraining(ctx context.Context, training Training, createFn func() error) error
-  CancelTraining(ctx context.Context, trainingUUID string, deleteFn func(Training) error) error
-  RescheduleTraining(ctx context.Context, trainingUUID string, newTime time.Time, updateFn func(Training) (Training, error)) error
-  ApproveTrainingReschedule(ctx context.Context, trainingUUID string, updateFn func(Training) (Training, error)) error
-  RejectTrainingReschedule(ctx context.Context, trainingUUID string, updateFn func(Training) (Training, error)) error
- }
```
Thanks to introducing the `training.Training` entity, we can have a much simpler version, with one method for adding a new training and one for the update.

由于引入了`training.Training`实体，我们可以有一个更简单的版本，用一个方法来添加一个新的训练，一个方法来更新。

```go
package training

// ...

type Repository interface {
   AddTraining(ctx context.Context, tr *Training) error

   GetTraining(ctx context.Context, trainingUUID string, user User) (*Training, error)

   UpdateTraining(
      ctx context.Context,
      trainingUUID string,
      user User,
      updateFn func(ctx context.Context, tr *Training) (*Training, error),
   ) error
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainings/domain/training/repository.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/8d9274811559399461aa9f6bf3829316b8ddfb63/internal/trainings/domain/training/repository.go#L16)

As in the previous article, we implemented our repository using Firestore. We will also use Firestore in the current implementation. Please keep in mind that this is an implementation detail – you can use any database you want. In the previous article, we have shown example implementations using different databases.

在上一篇文章中，我们使用 `Firestore` 实现了我们的 `repository` 。我们在当前的实现中也将使用`Firestore`。请记住，这是一个实现细节--你可以使用任何你想要的数据库。在之前的文章中，我们已经展示了使用不同数据库的实例实现。

```go
package adapters

// ...

func (r TrainingsFirestoreRepository) UpdateTraining(
   ctx context.Context,
   trainingUUID string,
   user training.User,
   updateFn func(ctx context.Context, tr *training.Training) (*training.Training, error),
) error {
   trainingsCollection := r.trainingsCollection()

   return r.firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
      documentRef := trainingsCollection.Doc(trainingUUID)

      firestoreTraining, err := tx.Get(documentRef)
      if err != nil {
         return errors.Wrap(err, "unable to get actual docs")
      }

      tr, err := r.unmarshalTraining(firestoreTraining)
      if err != nil {
         return err
      }

      if err := training.CanUserSeeTraining(user, *tr); err != nil {
         return err
      }

      updatedTraining, err := updateFn(ctx, tr)
      if err != nil {
         return err
      }

      return tx.Set(documentRef, r.marshalTraining(updatedTraining))
   })
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainings/adapters/trainings_firestore_repository.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/8d9274811559399461aa9f6bf3829316b8ddfb63/internal/trainings/adapters/trainings_firestore_repository.go#L83)

## Connecting everything 
 
How to use our code now? What about our ports layer? Thanks to the refactoring that Miłosz did in refactoring to Clean Architecture article, our ports layer is decoupled from other layers. That’s why, after this refactoring, it doesn’t require almost any significant changes. We just call the application command instead of the application service.

现在如何使用我们的代码？我们的`ports`层怎么样了？由于 Miłosz 在 [refactoring to Clean Architecture](https://threedots.tech/post/introducing-clean-architecture/)一文中所做的重构，我们的`ports`层已经与其他层解耦。这就是为什么，在这次重构之后，它几乎不需要任何重大改变。我们只是调用`application command`而不是`application service`。

![](https://cdn.nlark.com/yuque/0/2021/png/2774323/1629115267919-4f3714b0-e620-40ef-9df7-9dc4e8f023a9.png?x-oss-process=image%2Fresize%2Cw_1500)

```go
package ports

// ...

type HttpServer struct {
   app app.Application
}

// ...

func (h HttpServer) CancelTraining(w http.ResponseWriter, r *http.Request) {
   trainingUUID := r.Context().Value("trainingUUID").(string)

   user, err := newDomainUserFromAuthUser(r.Context())
   if err != nil {
      httperr.RespondWithSlugError(err, w, r)
      return
   }

   err = h.app.Commands.CancelTraining.Handle(r.Context(), command.CancelTraining{
      TrainingUUID: trainingUUID,
      User:         user,
   })
   if err != nil {
      httperr.RespondWithSlugError(err, w, r)
      return
   }
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainings/ports/http.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/8d9274811559399461aa9f6bf3829316b8ddfb63/internal/trainings/ports/http.go#L87)
## How to approach such refactoring in a real project?  如何在一个真实的项目中进行这样的重构？
It may not be obvious how to do such refactoring in a real project. It’s hard to do a code review and agree on the team level on the refactoring direction.

在一个真实的项目中，如何进行这样的重构可能并不明显。要做代码审查并在团队层面上就重构方向达成一致是很难的。

From my experience, the best approach is Pair or Mob programming. Even if, at the beginning, you may feel that it is a waste of time, the knowledge sharing and instant review will save a lot of time in the future. Thanks to great knowledge sharing, you can work much faster after the initial project or refactoring phase.

根据我的经验，最好的方法是Pair或Mob编程。即使在开始的时候，你可能觉得这是浪费时间，但知识共享和即时审查将在未来节省大量的时间。由于伟大的知识共享，在最初的项目或重构阶段之后，你可以更快地工作。
```text
Mob programming: 整个团队在同一时间,同一空间和同一台计算机上工作
Pair programming: 两个人在同一时间,同一空间和同一台计算机上工作
```

In this case, you should not consider the time lost for Mob/Pair programming. You should consider the time that you may lose because of not doing it. It will also help you finish the refactoring much faster because you will not need to wait for the decisions. You can agree on them immediately.

在这种情况下，你不应该考虑 Mob或Pair 编程的时间损失。你应该考虑的是你可能因为不做而损失的时间。这也会帮助你更快地完成重构，因为你不需要等待决定。你可以立即就它们达成一致。

Mob and pair programming also work perfectly while implementing complex, greenfield projects. Knowledge sharing is especially important investment in that case. I’ve seen multiple times how this approach allowed to go very fast in the project in the long term.

在实施复杂的绿地项目时，Mob、Pair编程也能完美地发挥作用。在这种情况下，知识共享是特别重要的投资。我曾多次看到这种方法是如何让项目在长期内快速发展的。

When you are doing refactoring, it’s also critical to agree on reasonable timeboxes. And keep them. You can quickly lose your stakeholders’ trust when you spend an entire month on refactoring, and the improvement is not visible. It is also critical to integrate and deploy your refactoring as fast as you can. Perfectly, on a daily basis (if you can do it for non-refactoring work, I’m sure that you can do it for refactoring as well!). If your changes stay unmerged and undeployed for a longer time, it will increase the chance of breaking functionalities. It will also block any work in the refactored service or make changes harder to merge (it is not always possible to stop all other development around).

当你在做重构的时候，就合理的时间框架达成一致也很关键。并且要遵守。如果你花了一整个月的时间在重构上，而改进却不明显的话，你会很快失去利益相关者的信任。尽可能快地集成和部署你的重构也是至关重要的。完美的，每天都要做（如果你能对非重构工作做到这一点，我相信你对重构也能做到！）。如果你的修改停留在未合并和未部署的时间较长，它将增加破坏功能的机会。它也会阻碍重构服务中的任何工作，或者使变化更难合并（并不总是可能停止周围所有其他的开发）。

But when to know if the project is complex enough to use mob programming? Unfortunately, there is no magic formula for that. But there are questions that you should ask yourself:

但是，什么时候才能知道项目是否复杂到可以使用mob编程？不幸的是，这并没有什么神奇的公式。但有一些问题你应该问自己。

- do we understand the domain?  我们是否了解这个领域？
- do we know how to implement that?  我们知道如何实现它吗？
- will it end up with a monstrous pull request that nobody will be able to review?  最终会不会产生一个没有人能够审查的巨大的拉动请求？
- can we risk worse knowledge sharing while not doing mob/pair programming? 我们能不能在不做mob/pair编程的同时冒着更大的知识共享风险？

## Summary
And we come to an end. 😄
我们就这样结束了。😄

The entire diff for the refactoring is available on our Wild Workouts GitHub (watch out, it’s huge!).

整个重构的diff可以在我们的 [Wild Workouts GitHub](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/commit/8d9274811559399461aa9f6bf3829316b8ddfb63) 上找到（注意，它很庞大！）。

If you had no chance to read previous articles yet, you know what to do! Even if some of the used approaches are simplified, you should already be able to use them in your project and see value from them.

如果你还没有机会阅读以前的文章，你就知道该怎么做了! 即使一些使用的方法被简化了，你也应该已经能够在你的项目中使用它们并从中看到价值。

I hope that after this article, you also see how all introduced patterns are working nicely together. If not yet, don’t worry. It took me 3 years to connect all the dots. But it was worth the time spent. After I understood how everything is connected, I started to look at new projects in a totally different way. It allowed me and my teams to work more efficiently in the long-term.

我希望在这篇文章之后，你也能看到所有引入的模式是如何很好地结合在一起的。如果还没有，请不要担心。我花了3年时间来连接所有的点。但这是值得花时间的。在我理解了所有事情的联系之后，我开始以一种完全不同的方式来看待新项目。这使我和我的团队能够更有效地长期工作。

It is also important to mention, that as all techniques, this combination is not a silver bullet. If you are creating project that is not complex and will be not touched any time soon after 1 month of development, probably it’s enough to put everything to one main package. 😉 Just keep in mind, when this 1 month of development will become one year!

同样重要的是要提到，和所有的技术一样，这种组合并不是银弹。如果你创建的项目并不复杂，并且在开发1个月后不会很快被触及，可能把所有东西都放在一个主要包里就足够了。

We will also continue these topics in the next articles. We will be shortly drifting to Strategic DDD Patterns, which should also help you gain a more high-level perspective on your projects.

我们还将在接下来的文章中继续这些话题。我们很快就会转向战略DDD模式，这也应该有助于你对你的项目获得更高层次的看法。

Did this article help you to understand how to connect DDD, Clean Architecture, and CQRS? Is something still not clear? Please let us know in the comments! We are happy to discuss all your doubts!

这篇文章是否帮助你理解了如何连接DDD、清洁架构和CQRS？有什么地方还不清楚吗？请在评论中告诉我们! 我们很乐意讨论你所有的疑惑!