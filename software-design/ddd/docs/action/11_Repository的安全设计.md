# Repository secure by design: how to sleep better without fear of security vulnerabilities  
存储库的安全设计：如何在不担心安全漏洞的情况下睡得更好

Thanks to the tests and code review, you can make your project bug-free. Right? Well… actually, probably not. That would be too easy. 😉 These techniques lower the chance of bugs, but they can’t eliminate them entirely. But does it mean we need to live with the risk of bugs until the end of our lives?

由于有了测试和代码审查，你可以使你的项目没有错误。对吗？嗯......实际上，可能不是。那就太容易了。 😉这些技术降低了bug的机会，但它们不能完全消除它们。但这是否意味着我们需要带着bug的风险生活到我们生命的尽头？

Over one year ago, I found a pretty interesting PR in the harbor project. This was a fix for the issue that allowed to create admin user by a regular user. This was obviously a severe security issue. Of course, automated tests didn’t found this bug earlier.

一年多以前，我在Harbor项目中发现了一个相当有趣的PR。这是对允许普通用户创建管理用户的问题的一个修复。这显然是一个严重的安全问题。当然，自动化测试没有更早发现这个错误。

This is how the bugfix looks like:

修复该bug的代码大致如下:
```go
		ua.RenderError(http.StatusBadRequest, "register error:"+err.Error())
		return
	}
+
+	if !ua.IsAdmin && user.HasAdminRole {
+		msg := "Non-admin cannot create an admin user."
+		log.Errorf(msg)
+		ua.SendForbiddenError(errors.New(msg))
+		return
+	}
+
	userExist, err := dao.UserExists(user, "username")
	if err != nil {
```

One if statement fixed the bug. Adding new tests also should ensure that there will be no regression in the future. Is it enough? **Did it secure the application from a similar bug in the future? I’m pretty sure it didn’t.**

一个if语句修复了这个bug。添加新的测试也应该确保将来不会出现回归。这样做够吗？它是否保证了应用程序在未来不会出现类似的错误？我很确定它没有。


The problem becomes bigger in more complex systems with a big team working on them. What if someone is new to the project and forgets to put this if statement? Even if you don’t hire new people currently, they may be hired in the future. **You will probably be surprised how long the code you have written will live.** We should not trust people to use the code (that) we’ve created in the way it’s intended – they will not.

系统更复杂，团队也庞大，那么该问题变得更大。  如果有人是项目的新人，忘记了放这个if语句怎么办？即使你目前不雇佣新人，将来也可能会雇佣他们。**你可能会惊讶于你写的代码能活多久。**我们不应该相信人们会按照我们创建的代码的方式来使用它--他们不会。


**In some cases, the solution that will protect us from issues like that is good design. Good design should not allow using our code in an invalid way.** Good design should guarantee that you can touch the existing code without any fear. People new to the project will feel safer introducing changes.

**在某些情况下，能够保护我们不受这样的问题影响的解决方案是好的设计。好的设计不应该允许以无效的方式使用我们的代码。**好的设计应该保证你可以毫无顾忌地触摸现有的代码。新加入项目的人在引入修改时会感到更安全。

In this article, I’ll show how I ensured that only allowed people would be able to see and edit a training. In our case, a training can only be seen by the training owner (an attendee) and the trainer. I will implement it in a way that doesn’t allow to use our code in not intended way. By design.

在这篇文章中，我将展示我如何确保只有被允许的人才能看到和编辑培训。在我们的案例中，一个培训只能由培训所有者（与会者）和培训师看到。我将以一种不允许以非预期方式使用我们的代码的方式实现它。通过设计。

Our current application assumes that a repository is the only way how we can access the data. Because of that, I will add authorization on the repository level. **By that, we are sure that it is impossible to access this data by unauthorized users.**

我们目前的应用假设存储库是我们访问数据的唯一途径。正因为如此，我将在资源库层面上添加授权。**这样，我们就能确保未经授权的用户不可能访问这些数据。**

But wait, is the repository the right place to manage authorization? Well, I can imagine that some people may be skeptical about that approach. Of course, we can start some philosophical discussion on what can be in the repository and what shouldn’t. Also, the actual logic of who can see the training will be placed in the domain layer. I don’t see any significant downsides, and the advantages are apparent. In my opinion, pragmatism should win here.

但是等等，`repository`层是管理授权的正确位置吗？嗯，我可以想象，有些人可能会对这种方式持怀疑态度。当然，我们可以就什么可以放在`repository`层中，什么不应该放在`repository`层中展开一些哲学上的讨论。另外，谁能看到培训的实际逻辑将被放在领域层中。我没有看到任何明显的缺点，优点也很明显。在我看来，实用主义应该在这里获胜。

## Show me the code, please! 代码展示
To achieve our robust design, we need to implement three things:

为了实现我们强大的设计，我们需要实现三件事。

1. Logic who can see the training (domain layer) 可以看到培训的逻辑（领域层）
2. Functions used to get the training (GetTraining in the repository), 用来获取训练的函数（repository 层）。
3. Functions used to update the training (UpdateTraining in the repository. 用于更新训练的函数（repository 层).
### Domain layer Domain 层
The first part is the logic responsible for deciding if someone can see the training. Because it is part of the domain logic (you can talk about who can see the training with your business or product team), it should go to the domain layer. It’s implemented with CanUserSeeTraining function.

第一部分是负责决定某人是否能看到培训的逻辑。因为它是领域逻辑的一部分（你可以和你的业务或产品团队讨论谁可以看到培训），所以它应该归入领域层。它是通过`CanUserSeeTraining`函数实现的。

It is also acceptable to keep it on the repository level, but it’s harder to re-use. I don’t see any advantage of this approach – especially if putting it to the domain doesn’t cost anything. 😉

把它放在`repository`层面也是可以接受的，但它更难被重复使用。我看不出这种方法有什么好处--特别是如果把它放到`domain`层中不需要花费什么。 😉

```go
package training

// ...

type User struct {
	userUUID string
	userType UserType
}

// ...

type ForbiddenToSeeTrainingError struct {
	RequestingUserUUID string
	TrainingOwnerUUID  string
}

func (f ForbiddenToSeeTrainingError) Error() string {
	return fmt.Sprintf(
		"user '%s' can't see user '%s' training",
		f.RequestingUserUUID, f.TrainingOwnerUUID,
	)
}

func CanUserSeeTraining(user User, training Training) error {
	if user.Type() == Trainer {
		return nil
	}
	if user.UUID() == training.UserUUID() {
		return nil
	}

	return ForbiddenToSeeTrainingError{user.UUID(), training.UserUUID()}
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainings/domain/training/user.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/22c0a25b67c4669d612a2fa4a434ffae8e35e65a/internal/trainings/domain/training/user.go#L92)

### Repository
Now when we have the `CanUserSeeTraining` function, we need to use this function. Easy like that.
现在当我们有`CanUserSeeTraining`函数时，我们需要使用这个函数，如下: 
```go
func (r TrainingsFirestoreRepository) GetTraining(
	ctx context.Context,
	trainingUUID string,
+	user training.User,
) (*training.Training, error) {
	firestoreTraining, err := r.trainingsCollection().Doc(trainingUUID).Get(ctx)

	if status.Code(err) == codes.NotFound {
		return nil, training.NotFoundError{trainingUUID}
	}
	if err != nil {
		return nil, errors.Wrap(err, "unable to get actual docs")
	}

	tr, err := r.unmarshalTraining(firestoreTraining)
	if err != nil {
		return nil, err
	}
+
+	if err := training.CanUserSeeTraining(user, *tr); err != nil {
+		return nil, err
+	}
+
	return tr, nil
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainings/adapters/trainings_firestore_repository.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/186a2c4a912e485ac7bb4d18c2892df7617e9ec9/internal/trainings/adapters/trainings_firestore_repository.go#L57)

Isn’t it too simple? Our goal is to create a simple, not complex, design and code. This is an excellent sign that it is deadly simple.

这不是太简单了吗？我们的目标是创造一个简单而不是复杂的设计和代码。这是一个极好的迹象，说明它是极其简单。

We are changing `UpdateTraining` in the same way.

我们正在以同样的方式改变`UpdateTraining`业务。

```go
func (r TrainingsFirestoreRepository) UpdateTraining(
	ctx context.Context,
	trainingUUID string,
+	user training.User,
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
+
+		if err := training.CanUserSeeTraining(user, *tr); err != nil {
+			return err
+		}
+
		updatedTraining, err := updateFn(ctx, tr)
		if err != nil {
			return err
		}

		return tx.Set(documentRef, r.marshalTraining(updatedTraining))
	})
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainings/adapters/trainings_firestore_repository.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/186a2c4a912e485ac7bb4d18c2892df7617e9ec9/internal/trainings/adapters/trainings_firestore_repository.go#L83)

And… that’s all! Is there any way that someone can use this in a wrong way? As long as the User is valid – no.

OK，调整就是这样! 只要用户是有效的，那就没有什么办法可以让人以错误的方式使用它。


This approach is similar to the method presented in the DDD Lite introduction article. It’s all about creating code that we can’t use in a wrong way.

这种方法类似于[《DDD Lite介绍》](https://threedots.tech/post/ddd-lite-in-go-introduction/)文章中介绍的方法。这都是为了创建我们不能用错的代码。

This is how usage of `UpdateTraining` now looks like:

这是 `UpdateTraining`的用法，如下:
```go
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
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainings/app/command/approve_training_reschedule.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/22c0a25b67c4669d612a2fa4a434ffae8e35e65a/internal/trainings/app/command/approve_training_reschedule.go#L39)

Of course, there are still some rules if Training can be rescheduled, but this is handled by the Training domain type. It’s covered in details in the DDD Lite introduction article. 😉

当然，如果培训可以重新安排，仍有一些规则，但这是由培训领域类型处理的。这在DDD Lite介绍文章中有详细介绍。

## Handling collections
Even if this approach works perfectly for operating on a single training, you need to be sure that access to a collection of trainings is properly secured. There is no magic here:

即使这种方法对单一培训的操作非常有效，你也需要确保对一系列培训的访问得到适当的保障。代码如下，简单的查询，没有其他魔法操作:
```go
func (r TrainingsFirestoreRepository) FindTrainingsForUser(ctx context.Context, userUUID string) ([]query.Training, error) {
	query := r.trainingsCollection().Query.
		Where("Time", ">=", time.Now().Add(-time.Hour*24)).
		Where("UserUuid", "==", userUUID).
		Where("Canceled", "==", false)

	iter := query.Documents(ctx)

	return r.trainingModelsToQuery(iter)
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainings/adapters/trainings_firestore_repository.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/186a2c4a912e485ac7bb4d18c2892df7617e9ec9/internal/trainings/adapters/trainings_firestore_repository.go#L182)

Doing it on the application layer with the `CanUserSeeTraining` function will be very expensive and slow. It’s better to create a bit of logic duplication.

在应用层调用 `CanUserSeeTraining` 函数做这件事会非常昂贵和缓慢。最好是建立一点逻辑的重复。

If this logic is more complex in your application, you can try to abstract it in the domain layer to the format that you can convert to query parameters in your database driver. I did it once, and it worked pretty nicely.

如果这个逻辑在你的应用程序中比较复杂，你可以尝试在`domain`层中把它抽象成你可以在数据库驱动中转换为查询参数的格式。我曾经这样做过，而且效果很好。

But in `Wild Workouts`, it will add unnecessary complexity - let’s Keep It Simple, Stupid.

但是在 `Wild Workouts` 项目中，它会增加不必要的复杂性--让我们保持简单，愚蠢。

## Handling internal updates  处理内部update
We often want to have endpoints that allow a developer or your company operations department to do some “backdoor” changes. The worst thing that you can do in this case is creating any kind of “fake user” and hacks.

我们经常希望有一些端点，允许开发人员或你的公司运营部门做一些 "后门 "的改变。在这种情况下，你可以做的最糟糕的事情是创建任何形式的 "假用户 "和黑客。

It ends with a lot of `if` statements added to the code from my experience. It also obfuscates the audit log (if you have any). Instead of a “fake user”, it’s better to create a special role and explicitly define the role’s permissions.

根据我的经验，它以大量的if语句添加到代码中结束。它还混淆了审计日志（如果你有的话）。与其说是 "假用户"，不如说是创建一个特殊的角色并明确定义该角色的权限。

If you need repository methods that don’t require any user (for Pub/Sub message handlers or migrations), it’s better to create separate repository methods. In that case, naming is essential – we need to be sure that the person who uses that method knows the security implications.

如果你需要不要求任何用户角色的`repository`方法（比如于Pub/Sub消息处理程序或迁移），最好创建独立的`repository`方法。在这种情况下，命名是必不可少的--我们需要确保使用该方法的人知道其安全含义。

From my experience, if updates are becoming much different for different actors, it’s worth to even introduce a separate CQRS Commands per actor. In our case it may be UpdateTrainingByOperations.

根据我的经验，如果更新对于不同的角色来说变得非常不同，甚至值得为每个角色引入一个单独的`CQRS`命令。在我们的例子中，它可能是 `UpdateTrainingByOperations` 。

## Passing authentication via `context.Context` 使用`context.Context`传递身份验证 
As far as I know, some people are passing authentication details via context.Context.

据我所知，有些人是通过`context.Context`传递认证细节。

I highly recommend not passing anything required by your application to work correctly via context.Context. The reason is simple – when passing values via context.Context, we lose one of the most significant Go advantages – static typing. It also hides what exactly the input for your functions is.

我强烈建议不要通过`context.Context`来传递你的应用程序所需的任何东西，以使其正常工作。原因很简单--当通过`context.Context`传递数值时，我们就失去了Go最显著的优势之一--静态类型化。它还隐藏了你的函数的具体输入内容。


If you need to pass values via context for some reason, it may be a symptom of a bad design somewhere in your service. Maybe the function is doing too much, and it’s hard to pass all arguments there? Perhaps it’s the time to decompose that?

如果你因为某些原因需要通过上下文传递数值，这可能是你的服务中某个地方设计不良的症状。也许函数做得太多了，很难在那里传递所有的参数？也许现在是分解这个的时候了？


## And that’s all for today!
As you see, the presented approach is straightforward to implement quickly.

I hope that it will help you with your project and give you more confidence in future development.

Do you see that it can help in your project? Do you think that it may help your colleagues? Don’t forget to share it with them!

正如你所看到的，所提出的方法是直截了当的，可以快速实施。

我希望它能对你的项目有所帮助，让你在未来的发展中更有信心。

你认为它对你的项目有帮助吗？你认为它可以帮助你的同事吗？不要忘记与他们分享!