#!/bin/bash
set -e
# centos一键安装docker-compose脚本.
# usage: sh install-docker-compose.sh
# 第一步 下载二进制文件到/usr/local/bin/位置
curl -L https://github.com/docker/compose/releases/download/1.24.0/docker-compose-"$(uname -s)"-"$(uname -m)" -o /usr/local/bin/docker-compose
# 第二步 赋予可执行权限
chmod +x /usr/local/bin/docker-compose

#查看版本号
docker-compose version
