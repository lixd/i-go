#!/usr/bin/env bash

function install-yasm() {
    mkdir -p /tmp/.ffmpeg/yasm && cd /tmp/.ffmpeg/yasm
    wget http://www.tortall.net/projects/yasm/releases/yasm-1.3.0.tar.gz
    tar -xvf yasm-1.3.0.tar.gz
    cd yasm-1.3.0/
    ./configure
    make && make install
}

# 需要编译一会，耐心等待
function install-ffmpeg() {
    mkdir -p /tmp/.ffmpeg/ffmpeg && cd /tmp/.ffmpeg/ffmpeg
    wget https://johnvansickle.com/ffmpeg/release-source/ffmpeg-4.1.tar.xz
    tar xvJf ffmpeg-4.1.tar.xz
    cd ffmpeg-4.1
    ./configure
    make && make install
}

function test() {
    ffmpeg
}

{
  install-yasm
  install-ffmpeg
#  installByYum
  test
}

# 也可以用 yum 安装
function installByYum() {
    # 配置第三方 yum 源
    sudo rpm --import http://li.nux.ro/download/nux/RPM-GPG-KEY-nux.ro
    sudo rpm -Uvh http://li.nux.ro/download/nux/dextop/el7/x86_64/nux-dextop-release-0-5.el7.nux.noarch.rpm
    # 安装相关组件
    sudo yum install ffmpeg ffmpeg-devel -y
    #测试是否安装成功
    ffmpeg
}
