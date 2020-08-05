在上传项目到github时,忘记忽略了某个文件夹 `.idea` ，就直接push上去了,最后意识到了此问题,决定删除掉远程仓库中的.idea文件夹

在github上只能删除仓库，却无法删除文件夹或文件, 所以只能通过命令来解决

如果

1. 查看有哪些文件夹
```sh
dir
```
2. 删除多余的文件夹
```sh
git rm -r --cached .idea
```
3. 提交，添加操作说明
```sh
git commit -m '删除xxxx'  
```
4. 将本次更改更新到github项目上去
```sh
git push -u origin master  
```