#!/bin/bash
set -e

# centos一键安装nfs
# 如何使用 sh install-nfs.sh
# nfs 挂载测试：mount -t nfs 127.0.0.1(服务器地址):/tmp/nfs/data(nfs的共享目录) /opt/kc/backups(挂载目标目录)
# 安装nfs
yum install -y nfs-utils rpcbind
# 检测 nfs 服务是否安装成功
rpcinfo -p localhost
# 写入配置文件，当前默认共享 /tmp/nfs/data 目录
mkdir -p /tmp/nfs/data
echo '/tmp/nfs/data	*(rw,sync,no_root_squash,no_subtree_check,fsid=0)' > /etc/exports
# 先启动 rpcbind 再启动 nfs
systemctl enable rpcbind --now
systemctl enable nfs --now
# 使 nfs 最新配置生效
exportfs -r
exportfs -v
# 查看共享路径
showmount -e 127.0.0.1
echo 'nfs install finish.'
