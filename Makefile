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


.PHONY: format-deps checkfmt fmt goimports vet lint

checkfmt: format-deps fmt goimports lint
format-deps:
    ifeq (, $(shell which golangci-lint))
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
    endif
    ifeq (, $(shell which goimports))
		go install golang.org/x/tools/cmd/goimports@latest
    endif

goimports:
	@hack/update-goimports.sh
fmt:
	gofmt -s -w .

lint:
	golangci-lint run ./... --timeout=5m --config .golangci.yaml


.PHONY: coverage-ui test
test:
	go test ./... -coverprofile=dist/coverage.out

coverage-ui:test
	go tool cover -html=dist/coverage.out -o dist/coverage.html




.PHONY: ent-install
ent-install:
    ifeq (, $(shell which golangci-lint))
		echo "ent not found, installing...";
		go install entgo.io/ent/cmd/ent@v0.12.3
    endif

.PHONY: ent ent-instal
ent:  ## generate ent
	ent generate ./ient/ent/schema
