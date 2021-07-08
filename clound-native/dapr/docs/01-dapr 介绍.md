[Dapr](https://github.com/dapr/dapr) is a portable, serverless, event-driven runtime that makes it easy for developers to build resilient, stateless and stateful microservices that run on the cloud and edge and embraces the diversity of languages and developer frameworks.

> [Dapr](https://github.com/dapr/dapr)是一个可移植的、无服务器的、事件驱动的运行时，它使开发人员能够轻松地构建弹性的、无状态的和有状态的微服务，这些微服务在云和边缘上运行，并拥抱语言和开发人员框架的多样性。


Dapr codifies the best practices for building microservice applications into open, independent, building blocks that enable you to build portable applications with the language and framework of your choice. Each building block is independent and you can use one, some, or all of them in your application.

> Dapr将构建微服务应用的最佳实践编纂成开放、独立的构件，使您能够使用您选择的语言和框架构建可移植的应用。每个构件都是独立的，您可以在您的应用程序中使用其中的一个、部分或全部构件。

## Goals
- Enable developers using any language or framework to write distributed applications
- Solve the hard problems developers face building microservice applications by providing best practice building blocks
- Be community driven, open and vendor neutral
- Gain new contributors
- Provide consistency and portability through open APIs
- Be platform agnostic across cloud and edge
- Embrace extensibility and provide pluggable components without vendor lock-in
- Enable IoT and edge scenarios by being highly performant and lightweight
- Be incrementally adoptable from existing code, with no runtime dependency


```go
- 使开发者能够使用任何语言或框架来编写分布式应用。
- 通过提供最佳实践构件，解决开发人员在构建微服务应用时面临的困难问题。
- 由社区驱动、开放和供应商中立
- 获得新的贡献者
- 通过开放的API提供一致性和可移植性
- 在云和边缘的平台不可知
- 拥抱可扩展性，并提供可插拔的组件，而不被供应商锁定。
- 通过高性能和轻量级来实现物联网和边缘场景。
- 可从现有代码中逐步采用，无运行时依赖性。
```

## How it works 
Dapr injects a side-car (container or process) to each compute unit. The side-car interacts with event triggers and communicates with the compute unit via standard HTTP or gRPC protocols. This enables Dapr to support all existing and future programming languages without requiring you to import frameworks or libraries.

> Dapr向每个计算单元注入一个侧车（容器或进程）。侧车与事件触发器交互，并通过标准的HTTP或gRPC协议与计算单元通信。这使得Dapr能够支持所有现有和未来的编程语言，而不需要导入框架或库。

Dapr offers built-in state management, reliable messaging (at least once delivery), triggers and bindings through standard HTTP verbs or gRPC interfaces. This allows you to write stateless, stateful and actor-like services following the same programming paradigm. You can freely choose consistency model, threading model and message delivery patterns.

> Dapr通过标准的HTTP动词或gRPC接口提供内置的状态管理、可靠的消息传递（至少一次交付）、触发器和绑定。这使得您可以按照相同的编程模式编写无状态、有状态和类似行为者的服务。你可以自由选择一致性模型、线程模型和消息传递模式。

Dapr runs natively on Kubernetes, as a self hosted binary on your machine, on an IoT device, or as a container that can be injected into any system, in the cloud or on-premises.

> Dapr原生运行在Kubernetes上，在你的机器上以自托管二进制的形式运行，在物联网设备上运行，或者以容器的形式运行，可以注入到云端或企业内部的任何系统中。

Dapr uses pluggable component state stores and message buses such as Redis as well as gRPC to offer a wide range of communication methods, including direct dapr-to-dapr using gRPC and async Pub-Sub with guaranteed delivery and at-least-once semantics.

> Dapr使用可插拔的组件状态存储和消息总线（如Redis以及gRPC）来提供广泛的通信方法，包括使用gRPC直接dapr到dapr，以及具有保证交付和至少一次语义的异步Pub-Sub。

## Why Dapr?
Writing high performance, scalable and reliable distributed application is hard. Dapr brings proven patterns and practices to you. It unifies event-driven and actors semantics into a simple, consistent programming model. It supports all programming languages without framework lock-in. You are not exposed to low-level primitives such as threading, concurrency control, partitioning and scaling. Instead, you can write your code by implementing a simple web server using familiar web frameworks of your choice.

> 编写高性能、可扩展和可靠的分布式应用是很难的。Dapr为您带来了成熟的模式和实践。它将事件驱动和行为者语义统一到一个简单、一致的编程模型中。它支持所有的编程语言，没有框架锁定。您不会接触到诸如线程、并发控制、分区和缩放等低级基元。取而代之的是，您可以使用您所选择的熟悉的web框架实现一个简单的web服务器来编写您的代码。

Dapr is flexible in threading and state consistency models. You can leverage multi-threading if you choose to, and you can choose among different consistency models. This flexibility enables to implement advanced scenarios without artificial constraints. Dapr is unique because you can transition seamlessly between platforms and underlying implementations without rewriting your code.

> Dapr在线程和状态一致性模型方面非常灵活。如果你选择的话，你可以利用多线程，并且你可以在不同的一致性模型中选择。这种灵活性使其能够在没有人为限制的情况下实现高级方案。Dapr是独一无二的，因为你可以在平台和底层实现之间无缝过渡，而无需重写你的代码。

## Features
- Event-driven Pub-Sub system with pluggable providers and at-least-once semantics
- Input and output bindings with pluggable providers
- State management with pluggable data stores
- Consistent service-to-service discovery and invocation
- Opt-in stateful models: Strong/Eventual consistency, First-write/Last-write wins
- Cross platform virtual actors
- Secrets management to retrieve secrets from secure key vaults
- Rate limiting
- Built-in Observability support
- Runs natively on Kubernetes using a dedicated Operator and CRDs
- Supports all programming languages via HTTP and gRPC
- Multi-Cloud, open components (bindings, pub-sub, state) from Azure, AWS, GCP
- Runs anywhere, as a process or containerized
- Lightweight (58MB binary, 4MB physical memory)
- Runs as a sidecar - removes the need for special SDKs or libraries
- Dedicated CLI - developer friendly experience with easy debugging
- Clients for Java, .NET Core, Go, Javascript, Python, Rust and C++

- 事件驱动的Pub-Sub系统，具有可插拔的提供者和at-least-once语义。
- 具有可插拔提供者的输入和输出绑定。
- 通过可插拔的数据存储进行状态管理
- 一致的服务对服务的发现和调用。
- 选择有状态模式。强/最终一致性，先写/后写获胜
- 跨平台虚拟行为体
- 秘密管理，从安全钥匙库中提取秘密。
- 速率限制
- 内置可观察性支持
- 使用专用的 Operator 和 CRD 在 Kubernetes 上原生运行。
- 通过HTTP和gRPC支持所有编程语言。
- 来自Azure、AWS、GCP的多云、开放组件（绑定、pub-sub、状态）。
- 可在任何地方运行，作为一个进程或容器。
- 轻量级（58MB二进制，4MB物理内存）。
- 以sidecar的形式运行--不需要特殊的SDK或库。