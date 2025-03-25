package prom

import (
	"fmt"
	"log/slog"
	"net/http"

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
func (exporter *LinkyExporter) Run(connector core.LinkyConnector) {
	slog.Info(fmt.Sprintf("Beginning to serve on port :%d", exporter.Port))

	prometheus.MustRegister(NewLinkyCollector(connector))
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", exporter.Address, exporter.Port), nil)
	if err != nil {
		slog.Error(err.Error())
	}
}
