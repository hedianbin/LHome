![20180920153744772125732.png](http://p9ug71a1p.bkt.clouddn.com/20180920153744772125732.png)

项目安装方法：

1.首先创建数据库

create database if not exists lovehome default charset utf8 collate utf8_general_ci;

然后去架设redis和fastDFS和nginx，并配置好服务的参数。看笔记

然后修改配置文件。

conf/app.conf ,按自己的信息修改

```
appname = lhome
httpport = 9998
httpaddr = "10.0.151.242"
runmode = dev
copyrequestbody = true
sessionon = true
redisaddr = "127.0.0.1"
redisport = 6379
redisdbnum = 0
mysqladdr = "127.0.0.1"
mysqlport = 3306
mysqldbname = "lhome"
mysqlusername = "root"
mysqlpassword = "123456"
fdfs_http_addr = "10.0.151.220:9998"
```

conf/client.conf，这是fastdfs的client配置服务。配置好了才能上传图片

修改tracker_server=10.0.151.220:22122

改完后运行一下程序bee run

运行会如果不报错，会自动创建数据表。然后。我们把areas数据导入进数据库

进入到conf目录，然后mysql -u root -p123456

use lhome;

source area.sql;这样的话就把城区数据导入到area库里了。

好了。然后打开http://10.0.151.242:8899，端口我改过

会看到网页了。可以正常使用。