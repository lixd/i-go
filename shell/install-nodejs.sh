#!/bin/bash
# 使用 NVM 安装 NodeJS 多个版本并进行管理
# 使用方式：source install-nodejs.sh
# 安装Git
set -e
yum install git -y
#使用Git将NVM的源码克隆到本地的~/.nvm目录下，并检查最新版本。
git clone https://github.com/cnpm/nvm.git ~/.nvm
cd ~/.nvm
git checkout `git describe --abbrev=0 --tags`
#配置NVM环境变量
echo ". ~/.nvm/nvm.sh" >> /etc/profile
source /etc/profile
#查看 nodejs 版本
nvm list-remote
#安装 v16.10.0
nvm install v16.10.0

#可以安装多个版本，使用 nvm ls 查看，使用 nvm use <版本号> 切换
