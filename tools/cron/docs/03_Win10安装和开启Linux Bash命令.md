https://www.jianshu.com/p/bb03cbb62a8a

安装连接文档安装操作完成后，可能在命令提示符执行 `bash` 命令并不会弹出 Linux 终端

原因是未安装 Linux，需要在 Miscrosoft Store 中安装一个虚拟机，如 Ubuntu

安装完成后就可以使用 `bash` 命令了。

Windows10 bash 命令位置：C:\Windows\System32\bash.exe

如果要运行 Linux 命令，可采用如下格式（注意 c 和前引号间的空格，如上图）：

`bash -c "linux 命令"`

例如：

```sh
bash -c "echo Hello from IThome"

bash -c "ls /mnt/c"

bash -c "lsb_release -a
```
