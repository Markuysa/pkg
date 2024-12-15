package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func LaunchPrometheusListener(cfg Config) error {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))

	errChan := make(chan error, 1)
	go func() {
		if err := http.ListenAndServe(cfg.Address, mux); err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	return <-errChan
}
