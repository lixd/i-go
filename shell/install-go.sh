#!/bin/bash
set -e
# this shell is used to install go1.18 on linux amd64
# usage: source install-go.sh
function downloadGo() {
  echo "1.download go"
  cd ~ \
  && wget https://studygolang.com/dl/golang/go1.18.linux-amd64.tar.gz -O go1.18.linux-amd64.tar.gz \
  && tar -C ~/ -xzf go1.18.linux-amd64.tar.gz \
  && rm -f ~/go1.18.linux-amd64.tar.gz
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
  echo "go1.18 install and check finished"
}
