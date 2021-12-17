# 使用swaggo自动生成Restful API文档

相关链接

- [Swaggo Github](https://github.com/swaggo/swag)
- [Swagger]( https://swagger.io/)


### 关于Swaggo

或许你使用过[Swagger](https://swagger.io/), 而 swaggo就是代替了你手动编写yaml的部分。只要通过一个命令就可以将注释转换成文档，这让我们可以更加专注于代码。


### 使用
具体使用见官方文档，这里就不再赘述了。
需要注意的是：swag init 的时候需要在项目根目录下执行，否则无法检测到所有文件中的注释。
> 比如在 /xxx 目录下执行 swag init 就只能检测到 xxx 目录下的，如果还有和 xxx 目录同级或者更上层的目录中的代码都检测不到。

### 优化

swaggo是直接build到二进制里的，会极大增加二进制文件的大小，一般在生产环境不需要将 docs 编译进去。

可以利用 go 提供的条件编译来实现是否编译文档。


在`main.go`声明`swagHandler`,并在该参数不为空时才加入路由：

```go
package main

//...

var swagHandler gin.HandlerFunc

func main(){
    // ...
    
    	if swagHandler != nil {
			r.GET("/swagger/*any", swagHandler)
        }
    
    //...
}
```

同时,我们将该参数在另外加了`build tag`的包中初始化。

```go
//go:build doc

package main

import (
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    _ "i-go/gin/swagger/docs"
)

func init() {
    swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}

```

之后我们就可以使用`go build -tags "doc"`来打包带文档的包，直接`go build`来打包不带文档的包。



