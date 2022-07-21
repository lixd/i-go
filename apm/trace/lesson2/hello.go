package main

import (
	"context"
	"fmt"
	"i-go/apm/trace/config"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
)

// Trace individual functions
// Combine multiple spans into a single trace
// Propagate the in-process context
func main() {
	// 解析命令行参数
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	// 1.初始化 tracer
	tracer, closer := config.NewTracer("hello")
	defer closer.Close()
	// 2.开始新的 Span （注意:必须要调用 Finish()方法span才会上传到后端）
	span := tracer.StartSpan("say-hello")
	defer span.Finish()
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)
	helloTo := os.Args[1]
	helloStr := formatString(ctx, helloTo)
	printHello(ctx, helloStr)
}

func formatString(ctx context.Context, helloTo string) string {
	childSpan, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer childSpan.Finish()
	time.Sleep(time.Second)
	return fmt.Sprintf("Hello, %s!", helloTo)
}

func printHello(ctx context.Context, helloStr string) {
	childSpan, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer childSpan.Finish()
	time.Sleep(time.Second)
	println(helloStr)
}
