# Jest 单元测试工具

## 安装
```js
npm i jest -D

npm i ts-jest -D
```

## 配置jest
### 1. 生成配置文件
```js
npx ts-jets config:init
```
执行后，会生成 `jest.config.js` 文件，内容：
```js
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
};
```
### 2. 编写测试脚本
在 `package.json`中新增 `"test": "jest"`

## 编写测试用例

1. 编写`math.ts`文件：
```ts
function add(a: number, b: number) {
    return a + b;
}

function sub(a: number, b: number) {
    return a - b;
}

// 记得导出
module.exports = {
    add,
    sub
}
```

2. 新建 `test` 目录，创建测试用例

编写 `math.test.ts` 文件：
```ts
const math = require('../src/math');

test('add: 1 + 2 = 3', () => {
    expect(math.add(1, 2)).toBe(3);
});

test('sub: 1 - 2 = -1', () => {
    expect(math.sub(1, 2)).toBe(-1);
});
```

## 执行测试脚本
`npm run test`

控制台就会输出测试信息，是否通过，通过多少