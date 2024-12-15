package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func LaunchPrometheusListener(cfg Config) error {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))

	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServe(cfg.Address, mux)
	}()

	return <-errChan
}
