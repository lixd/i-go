# Makefile 基础教程
# https://www.kancloud.cn/kancloud/make-command/45596
# https://seisman.github.io/how-to-write-makefile/overview.html

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
