# Go Test

UnitTest

* 执行单个*_test.go文件下的所有方法 go test -v demo.go
* 执行单个*_test.go文件下的指定方法 go test -v -test.run TestDemoMethod / go test -v -run=TestDemoMethod
* 执行指定目录下所有测试文件的测试方法 go test -v ./ 支持相对路径和绝对路径
* 禁用 test cache 增加参数 -count=1 go test -v -count=1

Benchmark

* 指定运行某个文件中的Benchmark go test -v -bench-. xxx.go

> -test.xxx 参数都可以简写为 -xxx

## for vs range

**range 在迭代过程中返回的是迭代值的拷贝**，即和 for 相比每次迭代都需要拷贝n次。 如果每次迭代的元素的内存占用很低，那么 for 和 range 的性能几乎是一样，例如 []int。
但是如果迭代的元素内存占用较高，例如一个包含很多属性的 struct 结构体，那么 for 的性能将显著地高于 range， 有时候甚至会有上千倍的性能差异。

* 对于元素内存占用较高的场景，建议使用 for，如果使用 range，建议只迭代下标，通过下标访问迭代值，这种使用方式和 for 就没有区别了。
* 如果元素的内存占用很低，那么 for 和 range 的性能几乎是一样

如果想使用 range 同时迭代下标和值，则需要将切片/数组的元素改为指针，才能不影响性能。