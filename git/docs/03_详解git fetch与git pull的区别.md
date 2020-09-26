`git fetch`和 `git pull`都可以将远端仓库更新至本地那么他们之间有何区别?想要弄清楚这个问题有有几个概念不得不提。

**FETCH_HEAD：**

是一个版本链接，记录在本地的一个文件中，指向着目前已经从远程仓库取下来的分支的末端版本。
**commit-id：**

在每次本地工作完成后，都会做一个`git commit` 操作来保存当前工作到本地的 repo， 此时会产生一个 commit-id，这是一个能唯一标识一个版本的序列号。 在使用 `git push` 后，这个序列号还会同步到远程仓库。

有了以上的概念再来说说 git fetch

#### git fetch

这将更新 git remote 中所有的远程仓库所包含分支的最新 commit-id, 将其记录到.git/FETCH_HEAD 文件中
git fetch 更新远程仓库的方式如下：

```sh
git fetch origin master:tmp
//在本地新建一个temp分支，并将远程origin仓库的master分支代码下载到本地temp分支
git diff tmp
//来比较本地代码与刚刚从远程下载下来的代码的区别
git merge tmp
//合并temp分支到本地的master分支
git branch -d temp
//如果不想保留temp分支 可以用这步删除

```

（1）如果直接使用 git fetch，则步骤如下：

- 创建并更新本 地远程分支。即创建并更新 origin/xxx 分支，拉取代码到 origin/xxx 分支上。
- 在 FETCH_HEAD 中设定当前分支-origin/当前分支对应，如直接到时候 git merge 就可以将 origin/abc 合并到 abc 分支上。

（2）git fetch origin
只是手动指定了要 fetch 的 remote。在不指定分支时通常默认为 master

（3）git fetch origin dev
指定远程 remote 和 FETCH_HEAD，并且只拉取该分支的提交。

#### git pull

首先，基于本地的 FETCH_HEAD 记录，比对本地的 FETCH_HEAD 记录与远程仓库的版本号，然后 git fetch 获得当前指向的远程分支的后续版本的数据，然后再利用 git merge 将其与本地的当前分支合并。所以可以认为 git pull 是 git fetch 和 git merge 两个步骤的结合。

git pull 的用法如下：

```sh
git pull <远程主机名> <远程分支名>:<本地分支名>
//取回远程主机某个分支的更新，再与本地的指定分支合并。
```

因此，与 git pull 相比 **git fetch 相当于是从远程获取最新版本到本地，但不会自动 merge**。如果需要有选择的合并 git fetch 是更好的选择。效果相同时 git pull 将更为快捷。
