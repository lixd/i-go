# Makefile 基础教程
# https://www.kancloud.cn/kancloud/make-command/45596
# https://seisman.github.io/how-to-write-makefile/overview.html
# 10分钟学会makefile FORCE  https://blog.csdn.net/szullc/article/details/85036984
# shell与makefile常见坑 http://www.shishao.site/shell-makefile-3c1e9
# shell与makefile https://bbs.huaweicloud.com/blogs/346792

.PHONY:build
build:
	go build -v .

.PHONY:gen
gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

.PHONY:clean
clean:
	rm pb/*.go

.PHONY:run
run:
	go run main.go

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...


format-deps:
    ifeq (, $(shell which golangci-lint))
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
    endif
    ifeq (, $(shell which goimports))
		go install golang.org/x/tools/cmd/goimports@latest
    endif

lint: format-deps
	golangci-lint run ./.. --timeout=5m --config .golangci.yaml
