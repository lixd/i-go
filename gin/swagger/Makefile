# 初始化 swagger
.PHONY:api-swag
swag:
	@echo 开始初始化 swagger 文档
	swag init -g main.go -o docs
	@echo 开始格式化 swagger 注释
	swag fmt -g main.go

# dev,带docs,二进制文件会比较大
.PHONY:dev
dev:
	go build -tags "doc"

# prod,不带docs,二进制文件会小一些,而且生产环境也不应该把docs暴露出去
.PHONY:prod
prod:
	go build
