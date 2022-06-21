#!/bin/bash
set -e
# this shell is used to install go1.18 on linux amd64
# usage: source install-go.sh
echo "1.download go"
cd ~ \
&& wget https://studygolang.com/dl/golang/go1.18.linux-amd64.tar.gz -O go1.18.linux-amd64.tar.gz \
&& tar -C ~/ -xzf go1.18.linux-amd64.tar.gz \
&& rm -f ~/go1.18.linux-amd64.tar.gz

echo "2.set env"
cd ~
mkdir ~/gopath
echo 'export GOPROXY=https://goproxy.cn' >> ~/.bashrc
echo 'export GOPATH=~/gopath' >> ~/.bashrc
echo 'export PATH=$PATH:~/go/bin:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc

echo "3.check go app"
go version

echo "4.check go env"
go env

echo "5 create go source file"
cd ~
tee ./hello.go <<-'EOF'
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
EOF

echo "6.run hello.go"
go run hello.go

echo "go1.18 install and check finished"
