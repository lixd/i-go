# Makefile


```bash
# Makefile 基础教程
# https://www.kancloud.cn/kancloud/make-command/45596
# https://seisman.github.io/how-to-write-makefile/overview.html
# 10分钟学会makefile FORCE  https://blog.csdn.net/szullc/article/details/85036984
# shell与makefile常见坑 http://www.shishao.site/shell-makefile-3c1e9
# shell与makefile https://bbs.huaweicloud.com/blogs/346792

```

makefile 里的横杠`-` 的作用,例如下面的例子，前两条命令前面就有一个横杠`-`
```bash
comp:
	-rm -rf output;
	-mkdir output;
	cd output;
	vcs -full64 -f filelist.f;
```
正常情况下，前面的命令执行失败后，就不会执行后续命令了，但是加了横杠`-`后，就算前面的命令执行失败了，后续的命令也会继续执行。  
因此对于一些不确定的命令，可以加上横杠`-`，这样就算执行失败了，也不会影响后续的命令执行。

cd /root/lixd/chart-base

helm upgrade --install mariadb-galera ./mariadb-galera -n lixd  --values /root/install/chart-base/mariadb-galera-values.yaml

helm upgrade --install memcached  ./memcached -n lixd  --values /root/install/chart-base/memcached-values.yaml

helm upgrade --install rabbitmq  ./rabbitmq -n lixd  --values /root/install/chart-base/rabbitmq-values.yaml

helm upgrade --install keystone ./keystone -n lixd  --values /root/install/chart-base/keystone-values.yaml
