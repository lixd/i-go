# 源镜像
FROM registry.us-west-1.aliyuncs.com/wlinno/golang:latest as builder

# 环境变量
ENV GOPROXY=https://goproxy.cn \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置工作目录
WORKDIR /go/src
# 将 go 工程代码加入到 docker 容器中
ADD .. /go/src

# 将我们的代码编译成二进制可执行文件 hello
RUN go build -o main .
# 源镜像
FROM registry.us-west-1.aliyuncs.com/wlinno/alpine:latest
WORKDIR /root/
COPY --from=builder /go/src .
# 暴露端口
EXPOSE 8080
# 最终运行docker的命令
ENTRYPOINT  ["./main"]