#!/bin/bash
set -e
# this shell is used to install go1.19.1 on linux amd64
# usage: source install-go.sh

version="1.19.1"

function downloadGo() {
  echo "1.download go"
  cd ~ \
  && mkdir -p ~/env/go \
  && wget https://studygolang.com/dl/golang/go$version.linux-amd64.tar.gz -O go$version.linux-amd64.tar.gz \
  && tar -C ~/env/go -xzf go$version.linux-amd64.tar.gz \
  && rm -f ~/go$version.linux-amd64.tar.gz
}

function setEnv() {
  echo "2.set env"
  mkdir -p ~/env/go/gopath
  echo 'export GOPROXY=https://goproxy.cn' >> ~/.bashrc
  echo 'export GOPATH=~/env/go/gopath' >> ~/.bashrc
  echo 'export PATH=$PATH:~/env/go/go/bin:$GOPATH/bin' >> ~/.bashrc
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
rm -rf hello.go
}


{
  downloadGo
  setEnv
  checkVersion
  runHelloWorld
  echo "go $version install and check finished"
  go version
}
