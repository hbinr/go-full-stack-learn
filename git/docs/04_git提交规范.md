

## 格式

`<type>(<scope>): <subject>`

`type`类型如下：
```sh
feat：新功能（feature）
fix：修补bug
docs：文档（documentation）
style： 格式（不影响代码运行的变动）
refactor：重构（即不是新增功能，也不是修改bug的代码变动）
test：增加测试
chore：构建过程或辅助工具的变动
```

`scope`指修改的范文，指明哪个文件

`subject`指明提交的主题
## 示例
- feat(controller,service,dao): 用户查询接口开发
- fix(dao): 用户查询缺少xxx字段

参考：
- https://cloud.tencent.com/developer/article/1762300