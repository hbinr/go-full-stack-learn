翻译：https://threedots.tech/post/ddd-lite-in-go-introduction/

We found that high performance is possible with all kinds of systems, provided that systems and the teams that build and maintain them—are loosely coupled.

我们发现，只要系统以及构建和维护系统的团队--松散耦合，各种系统都可以实现高性能。


This key architectural property enables teams to easily test and deploy individual components or services even as the organization and the number of systems it operates grow—that is, it allows organizations to increase their productivity as they scale.

这一关键的架构属性使团队能够轻松地测试和部署单个组件或服务，即使在组织及其运营的系统数量增长时也是如此--也就是说，它使组织能够随着规模的扩大而提高其生产力。


So let’s use microservices, and we are done? I would not be writing this article if it was enough.

那我们用微服务，就可以了吗？如果这样就够了，我就不会写这篇文章了。

> - Make large-scale changes to the design of their system without depending on other teams to make changes in their systems or creating significant work for other teams

> 对自己的系统设计进行大规模的修改，而不需要依赖其他团队对自己的系统进行修改，也不需要给其他团队带来大量的工作。
  
> - Complete their work without communicating and coordinating with people outside their team

>  在不与团队外的人沟通协调的情况下，完成自己的工作。
> - Deploy and release their product or service on demand, regardless of other services it depends upon

> 按需部署和发布他们的产品或服务，而不考虑它所依赖的其他服务。
> - Do most of their testing on demand, without requiring an integrated test environment Perform deployments during normal business hours with negligible downtime

> 按需进行大部分测试，不需要集成测试环境，在正常工作时间内进行部署，停机时间可忽略不计。

[…] employing the latest whizzy microservices architecture deployed on containers is no guarantee of higher performance if you ignore these characteristics. […] To achieve these characteristics, design systems are loosely coupled — that is, can be changed and validated independently of each other.

[......]如果忽视这些特性，采用部署在容器上的最新奇特的微服务架构并不能保证更高的性能。[......]为了实现这些特性，设计系统是松散耦合的--也就是说，可以相互独立地改变和验证。


My simple DDD definition is: Ensure that you solve valid problem in the optimal way. After that implement the solution in a way that your business will understand without any extra translation from technical language needed.

我对DDD的简单定义是：确保以最佳方式解决有效问题。然后，以企业能够理解的方式实现解决方案，而不需要任何额外的技术语言翻译。

Coding is a war, to win you need a strategy!

I like to say that “5 days of coding can save 15 minutes of planning”.

打码是一场战争，要想赢就得有策略!
我喜欢说 "5天的编码可以节省15分钟的规划"。

## The Third Rule - domain needs to be database agnostic  第三条规则--领域需要与数据库无关。

There are multiple schools here – some are telling that it’s fine to have domain impacted by the database client. From our experience, keeping the domain strictly without any database influence works best.

这里有多个流派--有些人告诉说，域名受到数据库客户端的影响是可以的。根据我们的经验，严格保持域名不受任何数据库影响效果最好。


The most important reasons are:

最重要的原因是:
- domain types are not shaped by used database solution – they should be only shaped by business rules
- we can store data in the database in a more optimal way
- because of the Go design and lack of “magic” like annotations, ORM’s or any database solutions are affecting in even more significant way

- 域类型不受使用的数据库解决方案的影响 - 它们应该只受业务规则的影响。	
- 我们可以用更优化的方式将数据存储在数据库中。
- 因为go的设计和缺乏像注解、ORM或任何数据库解决方案这样的 "魔法"，正以更重要的方式影响着go的发展
  
## Domain-First approach*
If the project is complex enough, we can spend even 2-4 weeks to work on the domain layer, with just in-memory database implementation. In that case, we can explore the idea deeper and defer the decision to choose the database later. All our implementation is just based on unit tests.

We tried that approach a couple of times, and it always worked nicely. It is also a good idea to have some timebox here, to not spend too much time.

Please keep in mind that this approach requires a good relationship and a lot of trust from the business! If your relationship with business is far from being good, Strategic DDD patterns will improve that. Been there, done that!

领域优先法

如果项目足够复杂，我们甚至可以花2-4周的时间来研究领域层，只用内存数据库实现。在这种情况下，我们可以更深入地探索这个想法，推迟到以后再决定选择数据库。我们所有的实现都只是基于单元测试。

我们尝试过几次这种方法，效果一直很好。在这里有一些时间盒也是一个好主意，不要花费太多时间。

请记住，这种方法需要良好的关系和业务的信任! 如果你和企业的关系远没有好，战略DDD模式会改善这种情况。曾经有过，也做过!