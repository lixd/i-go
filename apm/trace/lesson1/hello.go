package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"i-go/apm/trace/config"
	"os"
)

// Instantiate a Tracer
// Create a simple trace
// Annotate the trace
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

	helloTo := os.Args[1]
	helloStr := formatString(span, helloTo)
	// 3.通过tag、log记录注释信息
	// LogFields 和 LogKV底层是调用的同一个方法
	span.SetTag("hello-to", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	span.LogKV("event", "println")
	printHello(span, helloStr)
}

func formatString(span opentracing.Span, helloTo string) string {
	childSpan := span.Tracer().StartSpan(
		"formatString",
		opentracing.ChildOf(span.Context()),
	)
	defer childSpan.Finish()
	return fmt.Sprintf("Hello, %s!", helloTo)
}

func printHello(span opentracing.Span, helloStr string) {
	childSpan := span.Tracer().StartSpan(
		"printHello",
		opentracing.ChildOf(span.Context()),
	)
	defer childSpan.Finish()
	println(helloStr)
}
