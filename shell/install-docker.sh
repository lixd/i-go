#!/bin/bash
set -e
# centos一键安装docker脚本.

#卸载旧版本
yum remove -y docker  docker-common docker-selinux docker-engine
#安装需要的软件包
yum install -y yum-utils device-mapper-persistent-data lvm2
#添加yum源
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
#安装最新版docker
yum install -y docker-ce
#配置镜像加速器
mkdir -p /etc/docker
echo '{
  "registry-mirrors": [
  "https://ekxinbbh.mirror.aliyuncs.com"
  ]
}' > /etc/docker/daemon.json
#启动并加入开机启动
systemctl enable docker --now
docker version
echo "docker install finish."
