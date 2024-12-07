package grpc

import (
	"github.com/Markuysa/pkg/middleware"
	mw "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func NewServer(opts ...option) (*grpc.Server, error) {
	opt := defaultOptions()
	if len(opts) > 0 {
		for _, o := range opts {
			o.apply(&opt)
		}
	}

	if err := opt.Validate(); err != nil {
		return nil, err
	}

	grpcMetrics := mw.NewServerMetrics()

	unaryInterceptor := grpc.ChainUnaryInterceptor(
		middleware.AccessLogMiddleware(),
		recovery.UnaryServerInterceptor(),
		grpcMetrics.UnaryServerInterceptor(),
	)

	srv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: opt.config.MaxConnectionIdle,
			MaxConnectionAge:  opt.config.MaxConnectionAge,
			Timeout:           opt.config.Timeout,
			Time:              opt.config.Time,
		}),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		unaryInterceptor,
		grpc.MaxRecvMsgSize(opt.config.MaxRecvMsgSize),
		grpc.MaxSendMsgSize(opt.config.MaxSendMsgSize),
	)

	grpcMetrics.InitializeMetrics(srv)

	reflection.Register(srv)

	if opt.regFuncs != nil {
		opt.regFuncs.apply()
	}

	return srv, nil
}