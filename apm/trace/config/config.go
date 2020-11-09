package config

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
)

const (
	CollectorEndpoint = "http://47.93.123.142:14268/api/traces"
)

// NewTracer
func NewTracer(service string) (opentracing.Tracer, io.Closer) {
	// 参数详解 https://www.jaegertracing.io/docs/1.20/sampling/
	cfg := jaegerConfig.Configuration{
		ServiceName: service,
		// 采样配置
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: CollectorEndpoint, // 将span发往jaeger-collector的服务地址
		},
	}
	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

// initJaeger 将jaeger tracer设置为全局tracer
func initJaeger(service string) io.Closer {
	cfg := jaegerConfig.Configuration{
		// https://www.jaegertracing.io/docs/1.20/sampling/
		// 采样配置
		Sampler: &jaegerConfig.SamplerConfig{
			Type:                     jaeger.SamplerTypeConst,
			Param:                    1,
			SamplingServerURL:        "",
			SamplingRefreshInterval:  0,
			MaxOperations:            0,
			OperationNameLateBinding: false,
			Options:                  nil,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans: true,
			// 将span发往jaeger-collector的服务地址
			CollectorEndpoint: CollectorEndpoint,
		},
	}
	closer, err := cfg.InitGlobalTracer(service, jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return closer
}
