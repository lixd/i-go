#!/usr/bin/env bash
#  一键安装 python 脚本
# centos 可以使用 yum 一键安装，或者从源码编译安装
# ubuntu 同样可以使用 apt 安装


version="3.7.0"
tarDir="Python-$version"
fileName="Python-$version.tgz"

function installBySourceCode() {
  # 安装编译依赖包
#  yum install -y gcc patch libffi-devel python-devel  zlib-devel bzip2-devel openssl-devel ncurses-devel sqlite-devel readline-devel tk-devel gdbm-devel db4-devel libpcap-devel xz-devel

  # 下载 Python 包,官网下载很慢，使用国内镜像
  rm -rf $fileName
#  https://registry.npmmirror.com/binary.html?path=python/
  echo "wget  https://registry.npmmirror.com/-/binary/python/$version/Python-$version.tgz"
  wget  https://registry.npmmirror.com/-/binary/python/$version/Python-$version.tgz
  # 解压源码
  tar -zxvf $fileName
  # 移动解压后的Python源码包到python目录
  mkdir -p /usr/local/python
  mv $tarDir /usr/local/python

  # 进入解压后的目录并且执行 configure 检测
  cd /usr/local/python/$tarDir || exit
  ./configure --prefix=/usr/local/python/
  # 编译并安装
  make && make install

  # 添加linux环境变量
  echo 'export PATH=/usr/local/python/bin:$PATH' >> /etc/profile
  source /etc/profile
  # 修改python的链接指向
  mv /usr/bin/python3 /usr/bin/python3.bak
  ln -s /usr/local/python/bin/python3 /usr/bin/python3

  # 检测是否安装成功
  python3 -V
}

function installByAPT() {
    apt install software-properties-common -y
    add-apt-repository ppa:deadsnakes/ppa -y
    apt update -y
    apt upgrade -y

    apt install -y python3.7 python3.7-distutils python3.7-dev
}


function installByYum() {
    yum -y install python3
    python3 -V
}

installBySourceCode

# 建议安装好后使用 venv 虚拟环境来隔离，避免依赖冲突。
# 在当前目录创建一个 demo 目录作为虚拟环境
# python3 -m venv demo
# 激活 demo 目录下的虚拟环境
# source demo/bin/activate
# 退出当前虚拟环境
# deactivate
