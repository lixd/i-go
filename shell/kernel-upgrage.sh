#!/bin/bash
set -e
# centos升级系统内核
# 添加elrepo
yum -y update
#启用 ELRepo 仓库
#导入ELRepo仓库的公共密钥
rpm --import https://www.elrepo.org/RPM-GPG-KEY-elrepo.org
#安装ELRepo仓库的yum源
yum install -y https://www.elrepo.org/elrepo-release-7.0-4.el7.elrepo.noarch.rpm
#查看可用内核列表
yum -y --disablerepo="*" --enablerepo="elrepo-kernel" list available
#安装最新版本
yum -y --disablerepo=\* --enablerepo=elrepo-kernel install kernel-lt.x86_64
#打印当前系统可用内核
awk -F\' '$1=="menuentry " {print i++ " : " $2}' /etc/grub2.cfg
#生效序号为0的内核，也就是最新安装的那个
grub2-set-default 0
#reboot 需要手动重启才行
