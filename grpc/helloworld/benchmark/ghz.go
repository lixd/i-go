package main

import (
	"fmt"
	"os"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
)

// https://github.com/bojand/ghz
// https://ghz.sh/docs/intro.html
func main() {
	data := map[string]string{
		"name": "1231",
	}
	report, err := runner.Run(
		// 基本配置信息 call host proto文件 data
		"helloworld.Greeter.SayHello", //  'package.Service/method' or 'package.Service.Method'
		"localhost:50051",
		runner.WithProtoFile("D:\\lillusory\\projects\\i-go\\grpc\\helloworld\\proto\\hello_world.proto", []string{}),
		// runner.WithDataFromFile("data.json"),
		runner.WithData(data),
		// 可选配置
		runner.WithInsecure(true),
		// 	压测相关配置
		runner.WithTotalRequests(10000),
		runner.WithConcurrencySchedule(runner.ScheduleLine),
		runner.WithConcurrencyStep(10),
		runner.WithConcurrencyStart(5),
		runner.WithConcurrencyEnd(100),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	// file, err := os.Create("report.html")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	rp := printer.ReportPrinter{
		Out: os.Stdout,
		// Out:    file,
		Report: report,
	}

	_ = rp.Print("summary")
}
