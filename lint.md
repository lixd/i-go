## 概述
主要使用的 lint 工具是 [golangci-lint](https://github.com/golangci/golangci-lint)，这也是 [go 官方推荐](https://github.com/golang/go/wiki/CodeTools#All-in-one) 的 lint 工具。


常见 lint 工具如下:
* gofmt
* govet
* revive
* ...


Golang 的 linter 非常多，各自检查的范围或者内容不尽相同，如果一个个 linter 去配置和使用非常的低效和麻烦。所以我们一般会使用 golangci-lint 高效的进行代码检查。
golangci-lint 是 linter 的集合器，本身集成了众多的 [linter](https://golangci-lint.run/usage/linters/)；并行执行检查，速度很快。

## [安装](https://golangci-lint.run/usage/install/)
推荐使用二进制方式安装，默认安装到 GOPATH 路径下。
```shell
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0

golangci-lint --version
```


## [配置](https://golangci-lint.run/usage/configuration/)
通过`.golangci.yaml` 文件对 golangci-lint 进行配置，主要用于指定使用哪些 linters 以及对各个 linter 的具体配置。
一个最简单的配置如下：
```yaml
linters-settings:
  # 这里则是对 gofmt 的具体配置
  gofmt:
    # simplify code: gofmt with `-s` option, true by default
    simplify: true
linters:
  # 关闭所有 linter
  disable-all: true
  # 然后指定开启 gofmt
  enable:
    - gofmt
```

### 使用
golangci-lint 为了加速，进行了缓存，所以每次代码更新后需要先清理缓存后再执行。
具体命令如下：

```sh
# 清理缓存
golangci-lint cache clean
# 执行检测 指定超时时间和配置文件--项目根目录下执行
golangci-lint run --timeout=5m --config ./.golangci.yaml
```

`golangci-lint run` 等同于 `golangci-lint run ./...`会检测当前目录下的所有文件，如果不需要全部检测，可以指定具体目录或者文件:
```shell
# 
检查dir1和dir2目录下的代码及dir3目录下的file1.go文件
golangci-lint run dir1 dir2/... dir3/file1.go
```

已经配置到 Makefile 中了，如果支持 Makefile 则可以直接使用`make lint` 命令。
