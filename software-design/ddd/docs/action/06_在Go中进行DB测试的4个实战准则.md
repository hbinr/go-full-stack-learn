# 4 practical principles of high-quality database integration tests in Go
原文地址: https://threedots.tech/post/database-integration-testing/

Did you ever hear about a project where changes were tested on customers that you don’t like or countries that are not profitable? Or even worse – did you work on such project?
你是否听说过这样一个项目——在你不喜欢的客户或不盈利的国家进行更改测试？ 或者比这更糟——你参与过这样的项目吗？ 

It’s not enough to say that it’s just not fair and not professional. It’s also hard to develop anything new because you are afraid to make any change in your codebase.
仅仅说它不公平和不专业是不够的。开发任何新东西也很困难，因为您害怕对代码库进行任何更改。 


In 2019 HackerRank Developer Skills Report *Professional growth & learning* was marked as the most critical factor during looking for a new job. Do you think you can learn anything and grow when you test your application in that way?
在 2019 年 HackerRank 开发人员技能报告中，*专业成长和学习*被标记为寻找新工作的最关键因素。当您以这种方式测试您的应用程序时，您认为您可以学到任何东西并成长吗？ 

**It’s all leading to frustration and burnout.**  这一切都会导致沮丧和倦怠 


To develop your application easily and with confidence, you need to have a set of tests on multiple levels. **In this article, I will cover in practical examples how to implement high-quality database integration tests. I will also cover basic Go testing techniques, like test tables, assert functions, parallel execution, and black-box testing.**
为了轻松自信地开发应用程序，您需要在多个级别上进行一组测试。 **在本文中，我将通过实际示例介绍如何实现高质量的数据库集成测试。我还将介绍基本的 Go 测试技术，如测试表、断言函数、并行执行和黑盒测试。** 

What it actually means that test quality is high?
测试质量高究竟意味着什么？ 

## 4 principles of high-quality tests  高质量测试的4个原则 

I prepared 4 rules that we need to pass to say that our integration tests quality is high.
我准备了 4 条规则，我们需要通过这些规则来说明我们的集成测试质量很高。 

### 1. 快 Fast  
Good tests **need to be fast. There is no compromise here.**
好的测试**需要快速。这里没有妥协。** 

Everybody hates long-running tests. Let’s think about your teammates’ time and mental health when they are waiting for test results. Both in CI and locally. It’s terrifying.
每个人都讨厌长时间运行的测试。让我们想想你的队友在等待测试结果时的时间和心理健康。无论是 CI 和本地测试。这太可怕了。 

When you wait for a long time, you will likely start to do anything else in the meantime. After the CI passes (hopefully), you will need to switch back to this task. Context switching is one of the biggest productivity killers. It’s very exhausting for our brains. We are not robots.
当你正漫长等等测试结果，在此期间很可能会开始做其他事情。当 CI 通过后（希望如此），你需要切换回之前的任务。上下文切换或者说注意力转移是最大的生产力杀手之一。这对我们的大脑来说非常累人。我们不是机器人。

I know that there are still some companies where tests can be executing for 24 hours. We don’t want to follow this approach. 😉 You should be able to run your tests locally in less than 1 minute, ideally in less than 10s. I know that sometimes it may require some time investment. It’s an investment with an excellent ROI (*Return Of Investment*)! In that case, you can really quickly check your changes. Also, deployment times, etc. are much shorter.
我知道仍有一些公司可以执行 24 小时测试。我们不想遵循这种方法。 😉 您应该能够在不到 1 分钟的时间内在本地运行您的测试，最好在 10 秒内。我知道有时可能需要一些时间投入。这是一项具有出色 ROI（*投资回报*）的投资！在这种情况下，您可以非常快速地检查您的更改。此外，部署时间等要短得多。 

It’s always worth trying to find quick wins that can reduce tests execution the most from my experience. Pareto principle ([80/20 rule](https://en.wikipedia.org/wiki/Pareto_principle)) works here perfectly!
根据我的经验，尝试找到可以最大程度减少测试执行的快速胜利总是值得的。二八原则（[80/20 规则](https://en.wikipedia.org/wiki/Pareto_principle)）在这里完美适用！ 


## 2. 在各个层面测试足够多的场景  Testing enough scenarios on all levels  

I hope that you already know that 100% test coverage is not the best idea  (as long as it is not a simple/critical library).
我希望你已经知道 100% 的测试覆盖率不是最好的主意（只要它不是一个简单/关键的库）。

It’s always a good idea to ask yourself the question “how easily can it break?". It’s even more worth to ask this question if you feel that the test that you are implementing starts to look exactly as a function that you test. At the end we are not writing tests because tests are nice, but they should save our ass!
问自己这个问题“它有多容易破坏？”总是一个好主意。如果你觉得你正在实现的测试开始看起来和你测试的功能完全一样，那么问这个问题就更有价值了。最后，我们不是在编写测试，因为测试很好，但它们应该拯救我们的屁股！

From my experience, **coverage like 70-80% is a pretty good result in Go.**
根据我的经验，**像 70-80% 这样的覆盖率在 Go 中是一个不错的结果。**

It’s also not the best idea to cover everything with *component* or *end-to-end tests*. 
- First – you will not be able to do that because of the inability to simulate some error scenarios, like rollbacks on the repository.
- Second – it will break the first rule. These tests will be slow.
用*组件*或*端到端测试*覆盖所有内容也不是最好的主意：
- 首先 —— 你将无法这样做，因为无法模拟某些错误场景，例如存储库上的回滚。
- 第二 —— 它会打破第一条规则。 这些测试会很慢。 

![](https://cdn.nlark.com/yuque/0/2021/png/2774323/1627002980538-5fbaa7d8-75b1-41a9-9209-96105041d1ae.png)

![](https://cdn.nlark.com/yuque/0/2021/png/2774323/1627003003077-121c4b40-8afe-4aec-87b4-9d1059abf689.png)

Tests on several layers should also overlap, so we will know that integration is done correctly.
多个层上的测试也应该重叠，这样我们就可以知道集成是正确完成的。

You may think that solution for that is simple: the test pyramid! And that’s true…sometimes. Especially in applications that handle a lot of operations based on writes.
你可能认为这个问题的解决方案很简单：测试金字塔！没错，的确可以…  比如在基于写入处理大量操作的应用程序中。

![](https://cdn.nlark.com/yuque/0/2021/png/2774323/1627003051862-0bcc78f3-4026-4a77-bd39-0dfa72a173e7.png)

But what, for example, about applications that aggregate data from multiple other services and expose the data via API? It has no complex logic of saving the data. Probably most of the code is related to the database operations. 
但是，例如，从多个其他服务聚合数据并通过 API 公开数据的应用程序呢？ 它没有保存数据的复杂逻辑。可能大部分代码都与数据库操作相关。

In this case, we should use reversed test pyramid (it actually looks more like a christmas tree). When big part of our application is connected to some infrastructure (for example: database) it’s just hard to cover a lot of functionality with unit tests.
在这种情况下，我们应该使用反向测试金字塔(它实际上看起来更像一棵圣诞树)。当我们的应用程序的很大一部分连接到一些基础设施(例如：数据库)时，很难用单元测试覆盖很多功能。

![](https://cdn.nlark.com/yuque/0/2021/png/2774323/1627003094600-d9febbdf-17f2-4635-8323-0016a2b43f84.png)

### 3. Tests need to be robust and deterministic  测试需要是稳健的和确定的
Do you know that feeling when you are doing some urgent fix, tests are passing locally, you push changes to the repository and … after 20 minutes they fail in the CI? It’s incredibly frustrating. It also discourages us from adding new tests. It’s also decreasing our trust in them.
你知道那种感觉吗？当你在做一些紧急修复时，测试在本地通过了，你把修改推送到版本库，但......20分钟后，它们在CI中失败了。这真是令人难以置信的沮丧。这也使我们不愿意增加新的测试。这也降低了我们对他们的信任。

You should fix that issue as fast as you can. In that case, [Broken windows theory](https://en.wikipedia.org/wiki/Broken_windows_theory) is really valid.
你应该尽可能快地解决这个问题。在这种情况下。[Broken windows theory](https://en.wikipedia.org/wiki/Broken_windows_theory) 确实有效
### 4. You should be able to execute most of the tests locally 你应该能够在本地执行大部分的测试
Tests that you run locally should give you enough confidence that the feature that you developed or refactored is still working. E2E tests should just double-check if everything is integrated correctly.
本地运行的测试应该给你足够的信心，你开发或重构的功能仍然在工作。E2E测试应该只是重复检查是否所有的东西都被正确地整合了。

You will have also much more confidence when contracts between services are [robust because of using gRPC](https://threedots.tech/post/robust-grpc-google-cloud-run/), protobuf, or OpenAPI.
在使用[gRPC](https://threedots.tech/post/robust-grpc-google-cloud-run/)或OpenAPI后，服务之间的通信是健壮的，这时你也会有更大的信心

This is a good reason to cover as much as we can at lower levels (starting with the lowest): unit, integration, and component tests. Only then E2E.
这是一个很好的理由，让我们在较低的层次（从最低的开始）尽可能多地覆盖：单元、集成和组件测试。然后才是E2E。

## Implementation 如何实现？
We have some common theoretical ground. But nobody pays us for being the master of theory of programming. Let’s go to some practical examples that you can implement in your project.
我们有一些共同的理论基础。但是没有人因为我们是编程理论的大师而付钱给我们。让我们来看看一些实际的例子，你可以在你的项目中实施。

Let’s start with the repository pattern that I described in the previous article. You don’t need to read the rest of the articles from the series, but it’s a good idea to check at least the [previous one](https://threedots.tech/post/repository-pattern-in-go/). It will be much more clear for you how our repository implementation is working.
让我们从我在上一篇文章中描述的`repository pattern`开始。你不需要阅读这个系列的其他文章，但至少回顾一下[前面的文章](https://threedots.tech/post/repository-pattern-in-go/)。这将使你更清楚我们的仓库实现是如何工作的。

The way how we can interact with our database is defined by the `hour.Repository` interface. It assumes that our repository implementation is stupid. All complex logic is handled by domain part of our application. 
我们如何与我们的数据库互动的方式是由`hour.Repository`接口定义的。但是我们的存储库实现是愚蠢的，因为所有复杂的逻辑都由我们应用程序的领域部分处理。

**It should just save the data without any validations, etc. One of the significant advantages of that approach is the simplification of the repository and tests implementation.**
`Repository`模式下的实现应该只是保存数据，而不需要任何验证，等等。这种方方式的一个显著优点是简化了资源库和测试的实现。

In the previous article I prepared three different database implementations: MySQL, Firebase, and in-memory. We will test all of them. All of them are fully compatible, so we can have just one test suite.
在上一篇文章中，我准备了三种不同的数据库实现方式。MySQL，Firebase，和内存中。我们将对它们全部进行测试。所有这些都是完全兼容的，所以我们可以只有一个测试套件。
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
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/domain/hour/repository.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/521fdb5d6aa4f1f7ff18ec33f50ce6710906d73b/internal/trainer/domain/hour/repository.go#L8)

Because of multiple repository implementations, in our tests we iterate through a list of them.
由于有多个`repository`的实现，在我们的测试中，我们会在其中的一个列表中迭代。


All tests that we will write will be black-box tests. In other words – we will only cover public functions with tests. To ensure that, all our test packages have the `_test` suffix. That forces us to use only the public interface of the package. **It will pay back in the future with much more stable tests, that are not affected by any internal changes.** If you cannot write good black-box tests, you should consider if your public APIs are well designed.
我们将编写的所有测试都是黑盒测试。换句话说 - 我们将只用测试来覆盖公共函数。为了确保这一点，我们所有的测试包都有`_test`后缀。这迫使我们只使用包的公共接口。如果你不能写出好的黑盒测试，你应该考虑你的公共API是否设计得很好。**这将在未来得到回报，因为它更稳定，不受任何内部变化的影响。**

All our repository tests are executed in parallel. Thanks to that, they take less than 200ms. After adding multiple test cases, this time should not increase significantly.
我们所有的存储库测试都是并行执行的。多亏了这一点，它们的时间不到200ms。在添加多个测试案例后，这个时间应该不会明显增加。
```go
package main_test

// ...

func TestRepository(t *testing.T) {
   rand.Seed(time.Now().UTC().UnixNano())

   repositories := createRepositories(t)

   for i := range repositories {
      // When you are looping over the slice and later using iterated value in goroutine (here because of t.Parallel()),
      // you need to always create variable scoped in loop body!
      // More info here: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
      r := repositories[i]

      t.Run(r.Name, func(t *testing.T) {
         // It's always a good idea to build all non-unit tests to be able to work in parallel.
         // Thanks to that, your tests will be always fast and you will not be afraid to add more tests because of slowdown.
         t.Parallel()

         t.Run("testUpdateHour", func(t *testing.T) {
            t.Parallel()
            testUpdateHour(t, r.Repository)
         })
         t.Run("testUpdateHour_parallel", func(t *testing.T) {
            t.Parallel()
            testUpdateHour_parallel(t, r.Repository)
         })
         t.Run("testHourRepository_update_existing", func(t *testing.T) {
            t.Parallel()
            testHourRepository_update_existing(t, r.Repository)
         })
         t.Run("testUpdateHour_rollback", func(t *testing.T) {
            t.Parallel()
            testUpdateHour_rollback(t, r.Repository)
         })
      })
   }
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/hour_repository_test.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/521fdb5d6aa4f1f7ff18ec33f50ce6710906d73b/internal/trainer/hour_repository_test.go#L19)

When we have multiple tests, where we pass the same input and check the same output, it is a good idea to use a technique known as test table. The idea is simple: you should define a slice of inputs and expected outputs of the test and iterate over it to execute tests.
当我们有多个测试时，我们传递相同的输入并检查相同的输出，这是一个好主意，使用一种称为测试表的技术。这个想法很简单：你应该定义一个测试的输入和预期输出的片断，然后迭代它来执行测试。

```go
func testUpdateHour(t *testing.T, repository hour.Repository) {
   t.Helper()
   ctx := context.Background()

   testCases := []struct {
      Name       string
      CreateHour func(*testing.T) *hour.Hour
   }{
      {
         Name: "available_hour",
         CreateHour: func(t *testing.T) *hour.Hour {
            return newValidAvailableHour(t)
         },
      },
      {
         Name: "not_available_hour",
         CreateHour: func(t *testing.T) *hour.Hour {
            h := newValidAvailableHour(t)
            require.NoError(t, h.MakeNotAvailable())

            return h
         },
      },
      {
         Name: "hour_with_training",
         CreateHour: func(t *testing.T) *hour.Hour {
            h := newValidAvailableHour(t)
            require.NoError(t, h.ScheduleTraining())

            return h
         },
      },
   }

   for _, tc := range testCases {
      t.Run(tc.Name, func(t *testing.T) {
         newHour := tc.CreateHour(t)

         err := repository.UpdateHour(ctx, newHour.Time(), func(_ *hour.Hour) (*hour.Hour, error) {
            // UpdateHour provides us existing/new *hour.Hour,
            // but we are ignoring this hour and persisting result of `CreateHour`
            // we can assert this hour later in assertHourInRepository
            return newHour, nil
         })
         require.NoError(t, err)

         assertHourInRepository(ctx, t, repository, newHour)
      })
   }
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/hour_repository_test.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/521fdb5d6aa4f1f7ff18ec33f50ce6710906d73b/internal/trainer/hour_repository_test.go#L77)

You can see that we used a very popular [github.com/stretchr/testify](https://github.com/stretchr/testify) library. It’s significantly reducing boilerplate in tests by providing multiple helpers for [asserts](https://godoc.org/github.com/stretchr/testify/assert).
你可以看到，我们使用了一个非常流行的[github.com/stretchr/testify](https://github.com/stretchr/testify)库。它通过为[asserts](https://godoc.org/github.com/stretchr/testify/assert)提供多个帮助器，大大减少了测试中的模板。

> **require.NoError()**
> 
> When assert.NoError assert fails, tests execution is not interrupted.  当 assert.NoError 断言失败时，测试执行不会中断。 


>It’s worth to mention that asserts from require package are stopping execution of the test when it fails. Because of that, it’s often a good idea to use require for checking errors. In many cases, if some operation fails, it doesn’t make sense to check anything later.

>值得一提的是，当测试失败时，require 包的断言会停止执行。正因为如此，使用require来检查错误往往是一个好主意。在许多情况下，如果某些操作失败了，以后再检查就没有意义了。

>When we assert multiple values, assert is a better choice, because you will receive more context.

> 当我们断言多个值时，断言是一个更好的选择，因为你会收到更多的上下文。

If we have more specific data to assert, it’s always a good idea to add some `helpers`. It removes a lot of duplication, and improves tests readability a lot!
如果我们有更多具体的数据需要断言，添加一些`t.Helper()`总是一个好主意。它消除了大量的重复，并提高了测试的可读性。

```go
func assertHourInRepository(ctx context.Context, t *testing.T, repo hour.Repository, hour *hour.Hour) {
   require.NotNil(t, hour)

   hourFromRepo, err := repo.GetOrCreateHour(ctx, hour.Time())
   require.NoError(t, err)

   assert.Equal(t, hour, hourFromRepo)
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/hour_repository_test.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/521fdb5d6aa4f1f7ff18ec33f50ce6710906d73b/internal/trainer/hour_repository_test.go#L327)

## Testing transactions 如何测试事务？
Mistakes taught me that I should not trust myself when implementing complex code. We can sometimes not understand the documentation or just introduce some stupid mistake. You can gain the confidence in two ways:
- 1. TDD - let’s start with a test that will check if the transaction is working properly.
- 2. Let’s start with the implementation and add tests later.
  
错误告诉我，在实现复杂的代码时，我不应该相信自己。我们有时会不理解文档的内容，或者直接引入一些愚蠢的错误。你可以通过两种方式获得自信。
- 1. TDD - 让我们从测试开始，检查事务是否正常工作。
- 2. 让我们从实现开始，之后添加测试。

```go
func testUpdateHour_rollback(t *testing.T, repository hour.Repository) {
   //  t.Helper()
   t.Helper()
   ctx := context.Background()

   hourTime := newValidHourTime()

   err := repository.UpdateHour(ctx, hourTime, func(h *hour.Hour) (*hour.Hour, error) {
      require.NoError(t, h.MakeAvailable())
      return h, nil
   })

   err = repository.UpdateHour(ctx, hourTime, func(h *hour.Hour) (*hour.Hour, error) {
      assert.True(t, h.IsAvailable())
      require.NoError(t, h.MakeNotAvailable())

      return h, errors.New("something went wrong")
   })
   require.Error(t, err)

   persistedHour, err := repository.GetOrCreateHour(ctx, hourTime)
   require.NoError(t, err)

   assert.True(t, persistedHour.IsAvailable(), "availability change was persisted, not rolled back")
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/hour_repository_test.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/521fdb5d6aa4f1f7ff18ec33f50ce6710906d73b/internal/trainer/hour_repository_test.go#L197)

When I’m not using TDD, I try to be paranoid if test implementation is valid.
当我不使用TDD时，我试图偏执地认为测试实现是否有效。

To be more confident, I use a technique that I call **tests sabotage**.
为了更加自信，我使用了一种技巧，我称之为**测试破坏。**

**The method is pretty simple - let’s break the implementation that we are testing and let’s see if anything failed.**
**这个方法非常简单--让我们打破我们正在测试的实现，让我们看看是否有什么失败。**

```go
 func (m MySQLHourRepository) finishTransaction(err error, tx *sqlx.Tx) error {
-       if err != nil {
-               if rollbackErr := tx.Rollback(); rollbackErr != nil {
-                       return multierr.Combine(err, rollbackErr)
-               }
-
-               return err
-       } else {
-               if commitErr := tx.Commit(); commitErr != nil {
-                       return errors.Wrap(err, "failed to commit tx")
-               }
-
-               return nil
+       if commitErr := tx.Commit(); commitErr != nil {
+               return errors.Wrap(err, "failed to commit tx")
        }
+
+       return nil
 }
```
If your tests are passing after a change like that, I have bad news…
如果你的测试在这样的改变后还能通过，额，那就太糟了...


## Testing database race conditions 测试数据库并发条件
Our applications are not working in the void. It can always be the case that two multiple clients may try to do the same operation, and only one can win!
我们的应用程序不是在虚空中工作。总有可能出现这样的情况：两个多个客户端可能试图做同样的操作，而只有一个能赢！

In our case, the typical scenario is when two clients try to schedule a training at the same time. **We can have only one training scheduled in one hour.**
在我们的案例中，典型的情况是两个客户试图在同一时间安排一个培训。**但是我们业务是允许在一个小时内只能安排一次培训。**

This constraint is achieved by optimistic locking ([described in the previous article](https://threedots.tech/post/repository-pattern-in-go/#updating-the-data)) and domain constraints (described [two articles ago](https://threedots.tech/post/ddd-lite-in-go-introduction/#the-third-rule---domain-needs-to-be-database-agnostic)).
这种约束是通过乐观锁（在上一篇文章中描述）和 domain约束（在两篇文章前描述）实现的。

Let’s verify if it is possible to schedule one hour more than once. The idea is simple: **let’s create 20 goroutines, that we will release in one moment and try to schedule training**. We expect that exactly one worker should succeed.
让我们来验证一下，是否有可能将一个小时的时间安排得多于一次。这个想法很简单。**让我们创建20个goroutines，我们将在一个时刻释放这些goroutines，并尝试安排培训**。我们希望正好有一个能成功。

```go
func testUpdateHour_parallel(t *testing.T, repository hour.Repository) {
   // ...

	workersCount := 20
	workersDone := sync.WaitGroup{}
	workersDone.Add(workersCount)

	// closing startWorkers will unblock all workers at once,
	// thanks to that it will be more likely to have race condition
	startWorkers := make(chan struct{})
	// if training was successfully scheduled, number of the worker is sent to this channel
	trainingsScheduled := make(chan int, workersCount)

	// we are trying to do race condition, in practice only one worker should be able to finish transaction
	for worker := 0; worker < workersCount; worker++ {
		workerNum := worker

		go func() {
			defer workersDone.Done()
			<-startWorkers

			schedulingTraining := false

			err := repository.UpdateHour(ctx, hourTime, func(h *hour.Hour) (*hour.Hour, error) {
				// training is already scheduled, nothing to do there
				if h.HasTrainingScheduled() {
					return h, nil
				}
				// training is not scheduled yet, so let's try to do that
				if err := h.ScheduleTraining(); err != nil {
					return nil, err
				}

				schedulingTraining = true

				return h, nil
			})

			if schedulingTraining && err == nil {
				// training is only scheduled if UpdateHour didn't return an error
				trainingsScheduled <- workerNum
			}
		}()
	}

	close(startWorkers)
	
	// we are waiting, when all workers did the job
	workersDone.Wait()
	close(trainingsScheduled)

	var workersScheduledTraining []int

	for workerNum := range trainingsScheduled {
		workersScheduledTraining = append(workersScheduledTraining, workerNum)
	}

	assert.Len(t, workersScheduledTraining, 1, "only one worker should schedule training")
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/hour_repository_test.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/521fdb5d6aa4f1f7ff18ec33f50ce6710906d73b/internal/trainer/hour_repository_test.go#L128)

**It is also a good example that some use cases are easier to test in the integration test, not in acceptance or E2E level.** Tests like that as E2E will be really heavy, and you will need to have more workers to be sure that they execute transactions simultaneously.
**这也是一个很好的例子，有些用例在集成测试中更容易测试，而不是在验收或E2E级别。**  像E2E这样的测试是非常沉重的，你需要有更多的 worker (此处指其他线程)来确保他们同时执行事务。
## Making tests fast 让测试快起来
**If your tests can’t be executed in parallel, they will be slow.** Even on the best machine.
**如果你的测试不能并行执行，它们会很慢。**即使在最好的机器上。


Is putting `t.Parallel()` enough? Well, we need to ensure that our tests are independent. In our case, **if two tests would try to edit the same hour, they can fail randomly**. This is a highly undesirable situation.
增加一行 `t.Parallel()`就够了吗？嗯，我们需要确保我们的测试是独立的。在我们的例子中，**如果两个测试会试图修改同一个`hour`，他们会随机失败**。这是一个非常不理想的情况。

To achieve that, I created the `newValidHourTime()` function that provides a random hour that is unique in the current test run. In most applications, generating a unique UUID for your entities may be enough.
为了达到这个目的，我创建了`newValidHourTime()`函数，提供了一个在当前测试运行中唯一的随机`hour`。在大多数应用中，为你的实体生成一个唯一的UUID可能就足够了。

In some situations it may be less obvious, but still not impossible. I encourage you to spend some time to find the solution. Please treat it as the investment in your and your teammates’ mental health 😉.
在一些可能不那么明显，但仍然可能发生并发修改数据，导致出现数据一致性问题的场景下。我鼓励你花一些时间来寻找解决方案。请把它(指解决方案)当作对你和你的队友的心理健康的投资😉。
```go
// usedHours is storing hours used during the test,
// to ensure that within one test run we are not using the same hour
// (it should be not a problem between test runs)
var usedHours = sync.Map{}

// newValidHourTime to ensure cteate a unique Hour model, just for unit test
func newValidHourTime() time.Time {
   for {
      minTime := time.Now().AddDate(0, 0, 1)

      minTimestamp := minTime.Unix()
      maxTimestamp := minTime.AddDate(0, 0, testHourFactory.Config().MaxWeeksInTheFutureToSet*7).Unix()

      t := time.Unix(rand.Int63n(maxTimestamp-minTimestamp)+minTimestamp, 0).Truncate(time.Hour).Local()

      _, alreadyUsed := usedHours.LoadOrStore(t.Unix(), true)
      if !alreadyUsed {
         return t
      }
   }
}
```
> Full source: [github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/hour_repository_test.go](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/521fdb5d6aa4f1f7ff18ec33f50ce6710906d73b/internal/trainer/hour_repository_test.go#L306)

What is also good about making our tests independent, is no need for data cleanup. In my experience, doing data cleanup is always messy because:
- when it doesn’t work correctly, it creates hard-to-debug issues in tests,
- it makes tests slower,
- it adds overhead to the development (you need to remember to update the cleanup function)
- it may make running tests in parallel harder.

让我们的测试独立，还有一个好处，就是不需要进行数据清理。根据我的经验，做数据清理总是很混乱，因为:
- 当它不能正确工作时，它在测试中产生难以调试的问题。
- 它使测试变得更慢。
- 它增加了开发的开销（你需要记住更新清理功能）。
- 它可能使并行运行测试变得更难。

It may also happen that we are not able to run tests in parallel. Two common examples are:
- pagination – if you iterate over pages, other tests can put something in-between and move “items” in the pages.
- global counters – like with pagination, other tests may affect the counter in an unexpected way.

有些情况下，我们不能并行地运行测试。两个常见的例子是：
- 分页 - 如果你在页面上迭代，其他测试可以在中间放置一些东西，在页面上移动 "项目"。
- 全局计数器 - 与分页一样，其他测试可能会以一种意想不到的方式影响计数器。

In that case, it’s worth to keep these tests as short as we can.
在这种情况下，我们值得将这些测试尽量缩短。


### Please, don’t use sleep in tests! 拜托了，请不要在test中使用 `sleep()`
The last tip that makes tests flaky and slow is putting the sleep function in them. Please, **don’t!** It’s much better to synchronize your tests with channels or `sync.WaitGroup{}`. They are faster and more stable in that way.
最后一个让测试变得不稳定和缓慢的情况是把 `sleep()` 放在其中。拜托，**不要！** 用`channel`或`sync.WaitGroup{}`来同步你的测试要好得多。这样做更快、更稳定。

If you really need to wait for something, it’s better to use `assert.Eventually` instead of a sleep.
如果你真的需要等待什么，最好使用`assert.Eventually`而不是sleep。

> `Eventually` asserts that given condition will be met in waitFor time, periodically checking target function each tick.

> `Eventually` 断言给定的`condition `将在`waitFor` 时间内去满足，`Eventually`底层会在每个`tick`中定期检查目标函数。
> ```go
> assert.Eventually(
>     t, 
>     func() bool { return true }, // condition
>     time.Second, // waitFor
>     10*time.Millisecond, // tick
> )
> ```

> [godoc.org/github.com/stretchr/testify/assert](https://pkg.go.dev/github.com/stretchr/testify/assert?utm_source=godoc#Eventually)

## Running 启动测试
Now, when our tests are implemented, it’s time to run them!
现在，我们的测试用例都实现了，是时候运行它们了

Before that, we need to start our container with Firebase and MySQL with `docker-compose up`.
在测试之前，我们需要使用`docker-compose up` 来启动 `Firebase` 和 `MySQL` 容器

I prepared `make test` command that runs tests in a consistent way (for example, `-race` flag). It can also be used in the CI.
我准备了`make test`命令，以一致的方式运行测试（例如，`-race`标志）。它也可以在CI中使用。

```sh
$ make test 

?      github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/common/auth [no test files]
?      github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/common/client   [no test files]
?      github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/common/genproto/trainer [no test files]
?      github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/common/genproto/users   [no test files]
?      github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/common/logs [no test files]
?      github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/common/server   [no test files]
?      github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/common/server/httperr   [no test files]
ok     github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer 0.172s
ok     github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/domain/hour 0.031s
?      github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainings   [no test files]
?      github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/users   [no test files]
```

### Running one test and passing custom params
If you would like to pass some extra params, to have a verbose output (`-v`) or execute exact test (`-run`), you can pass it after `make test --`.
```sh
$ make test -- -v -run ^TestRepository/memory/testUpdateHour$ 

--- PASS: TestRepository (0.00s)
  --- PASS: TestRepository/memory (0.00s)
      --- PASS: TestRepository/memory/testUpdateHour (0.00s)
          --- PASS: TestRepository/memory/testUpdateHour/available_hour (0.00s)
          --- PASS: TestRepository/memory/testUpdateHour/not_available_hour (0.00s)
          --- PASS: TestRepository/memory/testUpdateHour/hour_with_training (0.00s)
PASS
```

If you are interested in how it is implemented, I’d recommend you check my [`Makefile magic`](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/521fdb5d6aa4f1f7ff18ec33f50ce6710906d73b/Makefile#L60) 🧙‍♂️.

## Debugging
Sometimes our tests fail in an unclear way. In that case, it’s useful to be able to easily check what data we have in our database.
有时我们的测试会以一种不明确的方式失败。在这种情况下，能够轻松地检查我们的数据库中有哪些数据是很有用的。

For SQL databases my first choice for that are [mycli for MySQL](https://www.mycli.net/install) and [pgcli for PostgreSQL](https://www.pgcli.com/). I’ve added `make mycli` command to Makefile, so you don’t need to pass credentials all the time.

For Firestore, the emulator is exposing the UI at `localhost:4000/firestore.`

### First step for having well-tested application
The biggest gap that we currently have is a lack of tests on the component and E2E level. Also, a big part of the application is not tested at all. We will fix that in the next articles. We will also cover some topics that we skipped this time.
我们目前最大的差距是缺乏对组件和E2E层面的测试。此外，应用程序的很大一部分根本没有被测试。我们将在接下来的文章中解决这个问题。我们还将涵盖一些我们这次跳过的主题。

But before that, we have one topic that we need to cover earlier – Clean/Hexagonal architecture! This approach will help us organize our application a bit and make future refactoring and features easier to implement.
但是在这之前，我们有一个话题需要提前讲一下-- 整洁/六边形架构！这种架构将帮助我们组织一下我们的应用程序，使未来的重构和功能更容易实现。

Just to remind, **the entire source code of Wild Workouts is [available on GitHub](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/). You can run it locally and deploy to Google Cloud with one command.**

Did you like this article and had no chance to read the previous ones? There are [7 more articles to check](https://threedots.tech/series/modern-business-software-in-go/?utm_source=testing-repository-outro)!

And that’s all for today. See you soon! 🙂

