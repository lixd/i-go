#!/bin/bash
set -e
# this shell is used to install go1.19.1 on linux amd64
# usage: source install-go.sh

version="1.19.1"

function downloadGo() {
  echo "1.download go"
  cd ~ \
  && wget https://studygolang.com/dl/golang/go$version.linux-amd64.tar.gz -O go$version.linux-amd64.tar.gz \
  && tar -C ~/ -xzf go$version.linux-amd64.tar.gz \
  && rm -f ~/go$version.linux-amd64.tar.gz
}

function setEnv() {
  echo "2.set env"
  cd ~
  mkdir ~/gopath
  echo 'export GOPROXY=https://goproxy.cn' >> ~/.bashrc
  echo 'export GOPATH=~/gopath' >> ~/.bashrc
  echo 'export PATH=$PATH:~/go/bin:$GOPATH/bin' >> ~/.bashrc
  source ~/.bashrc
}

function checkVersion() {
  echo "3.check version"
  go version
  go env
}

function runHelloWorld() {
echo "4.run hello world"
cd ~
tee ./hello.go <<-'EOF'
package main

import "fmt"

func main() {
  fmt.Println("Hello world!")
}
EOF
go run hello.go
}


{
  downloadGo
  setEnv
  checkVersion
  runHelloWorld
  echo "go $version install and check finished"
  go version
}
