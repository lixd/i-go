#!/bin/bash
set -e
# centos一键安装docker脚本.
# usage: sh install-docker.sh

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
    "https://reg-mirror.qiniu.com",
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com",
    "https://docker.mirrors.ustc.edu.cn",
    "https://ekxinbbh.mirror.aliyuncs.com"
  ]
}' > /etc/docker/daemon.json
#启动并加入开机启动
systemctl enable docker --now
docker version
echo "docker install finish."

# 镜像仓库可用性测试 https://github.com/docker-practice/docker-registry-cn-mirror-test/blob/master/.github/workflows/ci.yaml

# 开启 buildx
##方式1（临时设置，只在当前终端管用，下次进入终端就不生效了）
#export DOCKER_CLI_EXPERIMENTAL=enabled
##方式2（永久生效，直接修改 /etc/docker/daemon.json 配置文件，新增以下内容
#{"experimental": true}
