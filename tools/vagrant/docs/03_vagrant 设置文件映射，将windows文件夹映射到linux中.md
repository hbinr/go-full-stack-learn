# vagrant 设置文件映射，将 windows 文件夹映射到 linux 中

http://www.fancyecommerce.com/2017/06/06/vagrant-%E8%AE%BE%E7%BD%AE%E6%96%87%E4%BB%B6%E6%98%A0%E5%B0%84%EF%BC%8C%E5%B0%86windows%E6%96%87%E4%BB%B6%E5%A4%B9%E6%98%A0%E5%B0%84%E5%88%B0linux%E4%B8%AD/

核心：

打开 Vagrantfile，修改配置：

config.vm.synced_folder “D:\\linux\\fecshop”, “/www/web/develop/fecshop”

第一个路径是 window 的路径，第二个是 vagrant’中 linux 的路径
