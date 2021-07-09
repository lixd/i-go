package api

/*
### 错误码处理姿势
Demo: https://github.com/go-kratos/kratos/blob/main/errors/errors.go
**内部使用**

api/erros 包中定义枚举值，enum。通过插件生成枚举值对应的 IsXXX() 方法。在业务逻辑中使用时直接调用 api/errors.IsXXX() 方法判断是否是指定的错误。外部服务调用时 也直接用 api/errors.IsXXX() 来判定，比较统一，便于管理。

IsXXX（） 方法大致实现：

* 1）内部先调用标准errors库的As 方法 stderrors.As，将 err 转为 StatusError 结构体
* 2）从 StatusError 结构体中的Details 字段中拿到 ErrorInfo
* 3）ErrorInfo 内包含 Reason 和 Message 两个字段。其中的 Reason 就是对应的 api/errors中的枚举值。这样就能判定出是不是指定错误了。

**外部传输**

如 RPC 调用

* 1）api/errors 里定义了 StatusError 对象
* 2）gRPC PB 文件里也定义了 StatusError 对象，二者是一模一样的
* 3）gRPC 传输时使用 DeepCopy 将 api/errors 中的StatusError 对象复制成一个新的一模一样的  pb`.StatusError 对象，用于在 gRPC 中传输
* 4）客户端拿到 pb.StatusError 对象后再翻译成 api/errors.StatusError 对象
* 5）此时就和内部使用一致了
*/
