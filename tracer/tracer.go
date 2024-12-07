package tracer

import (
	"context"
	"fmt"
	"runtime"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type Tracer struct {
	tracer opentracing.Tracer
}

func NewTracer(cfg Config) (*Tracer, func() error, error) {
	reporterCfg := &config.ReporterConfig{
		LogSpans:           true,
		LocalAgentHostPort: cfg.URL,
	}
	if cfg.Auth != nil {
		reporterCfg.User = cfg.Auth.Username
		reporterCfg.Password = cfg.Auth.Password
	}

	conf, closer, err := config.Configuration{
		ServiceName: cfg.ServiceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: reporterCfg,
	}.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialize jaeger Tracer: %s", err)
	}

	opentracing.SetGlobalTracer(conf)

	return &Tracer{tracer: conf}, closer.Close, nil
}

func NewSpan(ctx context.Context) (context.Context, opentracing.Span, string) {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	operationName := fn.Name()

	span, ctx := opentracing.StartSpanFromContext(ctx, operationName)

	return ctx, span, operationName
}

func SetSpanTags(span opentracing.Span, tags map[string]string) {
	for key, value := range tags {
		span.SetTag(key, value)
	}
}

func LogSpan(span opentracing.Span, message string) {
	span.LogFields(log.String("event", message))
}
