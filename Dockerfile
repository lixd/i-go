# 编译环境
FROM golang:1.14 as build
ENV GOPROXY=https://goproxy.cn GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /i-go
COPY . /i-go
# -ldflags="-s -w" 减小二进制文件体积 https://golang.org/cmd/link/#hdr-Command_Line
RUN go build -ldflags="-s -w" -o main ./server/api/api.go

# 运行环境
FROM alpine:latest
WORKDIR /root
# 时区信息
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 二进制文件
COPY --from=build /i-go/main .
# 配置文件
COPY  ./conf/config.yaml /root/conf/
ENTRYPOINT  ["./main"]
