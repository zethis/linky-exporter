package prom

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/syberalexis/linky-exporter/pkg/core"
)

// LinkyExporter object to run exporter server and expose metrics
type LinkyExporter struct {
	Address string
	Port    int
}

// Run method to run http exporter server
func (exporter *LinkyExporter) Run(connector *core.LinkyConnector) {
	slog.Info(fmt.Sprintf("Beginning to serve on port :%d", exporter.Port))

	prometheus.MustRegister(NewLinkyCollector(connector))
	http.Handle("/metrics", promhttp.Handler())

	// Create server with timeouts
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", exporter.Address, exporter.Port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		slog.Error("Error while serving metrics", "error", err)
	}
}
