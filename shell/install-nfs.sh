#!/bin/bash
set -e

# centos一键安装nfs
# 安装nfs
yum install nfs-utils rpcbind
# 写入配置文件，当前默认共享 /tmp/nfs/data 目录
mkdir -p /tmp/nfs/data
echo '/tmp/nfs/data	*(rw,sync,no_root_squash)' > /etc/exports
# 启动nfs并使用最新配置
systemctl enable nfs-server.service --now
exportfs -r
exportfs -v
# 查看共享路径
showmount -e 127.0.0.1
echo 'nfs install finish.'
