## Typescript 究竟是什么？
- Javascript that scales，可以理解为是Javascript的超集
- 静态类型风格的类型系统
- 从es6到es10甚至是esnext的语法支持
- 兼容各种浏览器，各种系统，各种服务器，完全开源

## 使用Typescript的优点
### 程序更容易理解
1. 更方便的查看别人写的代码，函数或者方法输入输出的参数类型都一目了然，不需要看完一个函数的代码才知道它主要做了什么事
2. 省略了诸多调试过程，能很明确的知道字段的定义和用处，不用在疑惑的去断点调试了

### 效率更高
1. 在不同的代码块和定义中进行跳转
2. 代码自动补全
3. 丰富的接口提示

### 更少的错误
1. 在编译期间会显示的提示代码错误，提高质量，省去了不少找错误的时间
2. 杜绝一些常见错误，比如undefined

### 非常好的包容性
1. 完全兼容Javascript
2. 第三方库可以单独编写类型文件
3. 大多数项目都支持Typescript

## 使用Typescript的缺点
1. 增加了一些学习成本，对于后端开发人员，这些可以忽略
2. 短期内增加了一些开发成本，因为需要增加很多类型定义，但是当项目复杂后，维护起来去方便很多。

## 安装

**Typescript 官网地址**: https://www.typescriptlang.org/zh/

使用 nvm 来管理 node 版本: https://github.com/nvm-sh/nvm

安装 Typescript:
```sh
npm install -g typescript
```
查看版本：
```sh
tsc -v
```

## Hello World
1. 创建`test.ts`文件
```ts
const hello = (name:string) =>{
    return 'hello ${name}'
}

hello("world")
```
2. 编译该文件
```sh
tsc test.ts
```

编译之后会在当前目录下（`test.ts`所在目录）下生成一个 `test.js`文件，内容如下：
```js
var hello = function (name) {
    return 'hello ${name}';
};
hello("world");
```

可以看到使用`Typescript`编写的代码被翻译为`JavaScript`代码。