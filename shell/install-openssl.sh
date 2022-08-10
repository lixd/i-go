#!/bin/bash
set -e

# centos一键安装openssl 1.1.1 版本
# 下载相关依赖
yum update -y
yum install -y wget tar make gcc perl pcre-devel zlib-devel
# 下载 openssl 1.1.1 版本源码
wget https://www.openssl.org/source/openssl-1.1.1g.tar.gz
tar zxvf openssl-1.1.1g.tar.gz
cd openssl-1.1.1g
# 开始编译
./config --prefix=/usr --openssldir=/etc/ssl --libdir=lib no-shared zlib-dynamic
make && make install

# 查看是否安装成功
openssl version
echo 'openssl install finish.'
