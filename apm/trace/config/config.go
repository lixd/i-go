package config

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

const (
	CollectorEndpoint   = "http://123.57.236.125:14268/api/traces"
	CollectorEndpoint2  = "http://47.93.123.142:14268/api/traces"
	LocalAgentHostPort  = "47.93.123.142:6831"
	LocalAgentHostPort2 = "localhost:6831"
)

// NewTracer 使用 opentracing 统一标准
func NewTracer(service string) (opentracing.Tracer, io.Closer) {
	// config := parseConfig()
	return newTracer(service, "")
}

// newTracer
func newTracer(service, collectorEndpoint string) (opentracing.Tracer, io.Closer) {
	// 参数详解 https://www.jaegertracing.io/docs/1.20/sampling/
	cfg := jaegerConfig.Configuration{
		ServiceName: service,
		// 采样配置
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans: true,
			// CollectorEndpoint:  CollectorEndpoint2, // 将span发往jaeger-collector的服务地址
			LocalAgentHostPort: LocalAgentHostPort2,
		},
	}
	// 不传递 logger 就不会打印日志
	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	// tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

type JaegerConfig struct {
	CollectorEndpoint string `json:"collectorEndpoint"`
}

var jc JaegerConfig

// parseConfig 从 viper 中解析配置信息
func parseConfig() JaegerConfig {
	if jc.CollectorEndpoint != "" {
		return jc
	}
	if err := viper.UnmarshalKey("Jaeger", &jc); err != nil {
		panic(err)
	}
	if jc.CollectorEndpoint == "" {
		panic("jaeger conf nil")
	}
	return jc
}
