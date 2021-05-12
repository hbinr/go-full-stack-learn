# Node.js 初识

Node.js 是一个开源与跨平台的 JavaScript 运行时环境。在浏览器**外**运行 V8 JavaScript 引擎（Google Chrome 的内核）

> V8 JS引擎：浏览器里有可以执行js代码的环境

也就是说，我们可以在“node”环境中执行js代码，可以解释js代码。打破了过去JavaScript只能在浏览运执行的困境，这样就前后端环境统一了，大大降低了前后端语言切换的代价。可以做更多的事情：
- web服务器
- 命令行工具
- 网络爬虫
- app
- 嵌入式
- 游戏
- ...

## 没有DOM 和 BOM
JavaScript三大组成部分：
- ECMAScript
- DOM  
- BOM 
Node.js和Javascript不同，其没有dom和bom相关API是无法使用的，其包含：
- ECMAScript语法：es5、 es6都支持
- Node本身提供的一些模块和API，如fs模块、http模块等