package prometheus

import (
	"net/http"

	"github.com/Markuysa/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func LaunchPrometheusListener(cfg Config) error {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))

	go func() {
		if err := http.ListenAndServe(cfg.Address, mux); err != nil {
			log.Fatalf("failed to start prometheus listener: %v", err)
		}
	}()

	return nil
}
