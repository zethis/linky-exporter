package prom

import (
	"log/slog"

	prometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/syberalexis/linky-exporter/pkg/core"
)

const (
	USED     = "used"
	PRODUCED = "produced"
)

// MetricDef represents a metric definition
type MetricDef struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
}

// MetricCollector defines how to collect a specific metric
type MetricCollector func(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie)

// LinkyCollector object to describe and collect metrics
type LinkyCollector struct {
	connector core.LinkyConnector
	metrics   map[string]MetricDef
	handlers  map[string]MetricCollector
}

// NewLinkyCollector method to construct LinkyCollector
func NewLinkyCollector(connector *core.LinkyConnector) *LinkyCollector {
	lc := &LinkyCollector{
		connector: *connector,
		metrics:   make(map[string]MetricDef),
		handlers:  make(map[string]MetricCollector),
	}

	// Define all metrics
	lc.registerMetric("linky_timestamp", "Timestamp en seconde",
		[]string{"linky_id", "version", "contract", "pricing"}, prometheus.CounterValue, collectLinkyDate)

	lc.registerMetric("linky_energy_total", "Total Energie en Wh",
		[]string{"linky_id", "mode"}, prometheus.CounterValue, collectEnergyTotal)

	lc.registerMetric("linky_energy", "Energie en Wh",
		[]string{"linky_id", "mode", "index"}, prometheus.CounterValue, collectEnergy)

	lc.registerMetric("linky_reactive_energy_total", "Total Energie réactive en Wh",
		[]string{"linky_id", "index"}, prometheus.CounterValue, collectReactiveEnergyTotal)

	lc.registerMetric("linky_intensity", "Courant efficace en A",
		[]string{"linky_id", "phase"}, prometheus.GaugeValue, collectIntensity)

	lc.registerMetric("linky_voltage", "Tension efficace en V",
		[]string{"linky_id", "phase"}, prometheus.GaugeValue, collectVoltage)

	lc.registerMetric("linky_power", "Puissance apparente en VA",
		[]string{"linky_id", "mode", "phase"}, prometheus.GaugeValue, collectPower)

	lc.registerMetric("linky_power_last_year", "Puissance apparente n-1 en VA",
		[]string{"linky_id", "mode", "phase"}, prometheus.GaugeValue, collectPowerLastYear)

	lc.registerMetric("linky_power_max", "Puissance apparente en VA",
		[]string{"linky_id", "mode", "phase"}, prometheus.GaugeValue, collectPowerMax)

	lc.registerMetric("linky_power_reference", "Puissance apparente de référence en kVA",
		[]string{"linky_id", "type"}, prometheus.GaugeValue, collectPowerReference)

	lc.registerMetric("linky_load_curve_point", "Point de courbe de charge en W",
		[]string{"linky_id", "mode"}, prometheus.GaugeValue, collectLoadCurvePoint)

	lc.registerMetric("linky_load_curve_point_last_year", "Point de courbe de charge n-1 en W",
		[]string{"linky_id", "mode"}, prometheus.GaugeValue, collectLoadCurvePointLastYear)

	lc.registerMetric("linky_voltage_average", "Tension moyenne en V",
		[]string{"linky_id", "phase"}, prometheus.GaugeValue, collectAverageVoltage)

	lc.registerMetric("linky_status", "status from registry",
		[]string{"linky_id", "name"}, prometheus.GaugeValue, collectStatus)

	lc.registerMetric("linky_movable_peak", "Pointe mobile",
		[]string{"linky_id", "type", "phase"}, prometheus.GaugeValue, collectMovablePeak)

	lc.registerMetric("linky_relay", "Etat du relai",
		[]string{"linky_id", "id"}, prometheus.GaugeValue, collectRelay)

	lc.registerMetric(
		"linky_provider_day_info",
		"Numéro du jour en cours, du prochain jour et de son profil",
		[]string{"linky_id", "prm", "current_day", "next_day", "next_day_profile"},
		prometheus.GaugeValue,
		collectProviderDayInfo)

	return lc
}

// registerMetric adds a new metric definition and its collector function
func (lc *LinkyCollector) registerMetric(
	name,
	help string,
	labels []string,
	valueType prometheus.ValueType,
	handler MetricCollector) {
	lc.metrics[name] = MetricDef{
		desc:      prometheus.NewDesc(name, help, labels, nil),
		valueType: valueType,
	}
	lc.handlers[name] = handler
}

// Describe implements required describe function for all prometheus collectors
func (lc *LinkyCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range lc.metrics {
		ch <- metric.desc
	}
}

// Collect implements required collect function for all prometheus collectors
func (lc *LinkyCollector) Collect(ch chan<- prometheus.Metric) {
	var timeSerie *LinkyTimeSerie
	var err error

	switch lc.connector.Mode {
	case core.Standard:
		var ticValues *core.StandardTicValue
		ticValues, err = lc.connector.GetLastStandardTicValue()
		if err == nil {
			timeSerie = ConvertStandardTicValueToTimeSerie(ticValues)
		}
	case core.Historical:
		var ticValues *core.HistoricalTicValue
		ticValues, err = lc.connector.GetLastHistoricalTicValue()
		if err == nil {
			timeSerie = ConvertHistoricalTicValueToTimeSerie(ticValues)
		}
	default:
		slog.Error("Unable to read telemetry information", "error", err)
		return
	}

	if err != nil {
		slog.Error("Unable to read telemetry information", "error", err)
		return
	}

	// Collect all metrics
	for name, handler := range lc.handlers {
		// Skip standard-only metrics for historical mode
		if lc.connector.Mode != core.Standard &&
			(name == "linky_voltage" || name == "linky_status" ||
				name == "linky_relay" || name == "linky_movable_peak" ||
				name == "linky_provider_day_info") {
			continue
		}

		// Skip movable peak if not available
		if name == "linky_movable_peak" && timeSerie.MovingPeakStart1 == 0 {
			continue
		}

		// Skip provider day info (commented out in original code)
		if name == "linky_provider_day_info" {
			continue
		}

		handler(ch, lc, timeSerie)
	}
}

// Helper functions for metric collection

// sendMetric sends a metric to the channel
func sendMetric(
	ch chan<- prometheus.Metric,
	desc *prometheus.Desc,
	metricType prometheus.ValueType,
	value float64,
	labelValues ...string) {
	ch <- prometheus.MustNewConstMetric(desc, metricType, value, labelValues...)
}

// sendMetricIfNonZero sends a metric to the channel only if value is non-zero
func sendMetricIfNonZero(
	ch chan<- prometheus.Metric,
	desc *prometheus.Desc,
	metricType prometheus.ValueType,
	value float64,
	labelValues ...string) {
	if value != 0 {
		sendMetric(ch, desc, metricType, value, labelValues...)
	}
}

// Metric collector implementations
func collectLinkyDate(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_timestamp"]
	sendMetric(ch, metric.desc, metric.valueType, ts.LinkyDate,
		ts.LinkyId, ts.Version, ts.ContractTypeName, ts.PriceLabel)
}

func collectEnergyTotal(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_energy_total"]
	sendMetric(ch, metric.desc, metric.valueType, ts.TotalEnergyUsed, ts.LinkyId, USED)
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.TotalEnergyProduced, ts.LinkyId, PRODUCED)
}

func collectEnergy(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_energy"]

	energyMetrics := []struct {
		value float64
		mode  string
		index string
	}{
		{ts.EnergyUsedIndex1, USED, "F1"},
		{ts.EnergyUsedIndex2, USED, "F2"},
		{ts.EnergyUsedIndex3, USED, "F3"},
		{ts.EnergyUsedIndex4, USED, "F4"},
		{ts.EnergyUsedIndex5, USED, "F5"},
		{ts.EnergyUsedIndex6, USED, "F6"},
		{ts.EnergyUsedIndex7, USED, "F7"},
		{ts.EnergyUsedIndex8, USED, "F8"},
		{ts.EnergyUsedIndex9, USED, "F9"},
		{ts.EnergyUsedIndex10, USED, "F10"},
		{ts.EnergyUsedDistributorIndex1, USED, "D1"},
		{ts.EnergyUsedDistributorIndex2, USED, "D2"},
		{ts.EnergyUsedDistributorIndex3, USED, "D3"},
		{ts.EnergyUsedDistributorIndex4, USED, "D4"},
	}

	for _, m := range energyMetrics {
		sendMetricIfNonZero(ch, metric.desc, metric.valueType, m.value, ts.LinkyId, m.mode, m.index)
	}
}

func collectReactiveEnergyTotal(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	if ts.TotalReactiveEnergyQ1 == 0 && ts.TotalReactiveEnergyQ2 == 0 &&
		ts.TotalReactiveEnergyQ3 == 0 && ts.TotalReactiveEnergyQ4 == 0 {
		return
	}

	metric := lc.metrics["linky_reactive_energy_total"]

	reactiveMetrics := []struct {
		value float64
		index string
	}{
		{ts.TotalReactiveEnergyQ1, "Q1"},
		{ts.TotalReactiveEnergyQ2, "Q2"},
		{ts.TotalReactiveEnergyQ3, "Q3"},
		{ts.TotalReactiveEnergyQ4, "Q4"},
	}

	for _, m := range reactiveMetrics {
		sendMetric(ch, metric.desc, metric.valueType, m.value, ts.LinkyId, m.index)
	}
}

func collectIntensity(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_intensity"]
	sendMetric(ch, metric.desc, metric.valueType, ts.IntensityP1, ts.LinkyId, "1")
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.IntensityP2, ts.LinkyId, "2")
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.IntensityP3, ts.LinkyId, "3")
}

func collectVoltage(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_voltage"]
	sendMetric(ch, metric.desc, metric.valueType, ts.VoltageP1, ts.LinkyId, "1")
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.VoltageP2, ts.LinkyId, "2")
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.VoltageP3, ts.LinkyId, "3")
}

func collectPower(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_power"]

	powerMetrics := []struct {
		value float64
		mode  string
		phase string
	}{
		{ts.PowerUsed, USED, "1"},
		{ts.PowerUsedP1, USED, "1"},
		{ts.PowerUsedP2, USED, "2"},
		{ts.PowerUsedP3, USED, "3"},
		{ts.PowerProduced, PRODUCED, "0"},
	}

	for _, m := range powerMetrics {
		sendMetricIfNonZero(ch, metric.desc, metric.valueType, m.value, ts.LinkyId, m.mode, m.phase)
	}
}

func collectPowerLastYear(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_power_last_year"]

	metrics := []struct {
		value float64
		mode  string
		phase string
	}{
		{ts.PowerUsedMaxLastYear, USED, "1"},
		{ts.PowerUsedMaxLastYearP1, USED, "1"},
		{ts.PowerUsedMaxLastYearP2, USED, "2"},
		{ts.PowerUsedMaxLastYearP3, USED, "3"},
		{ts.PowerProducedLastYear, PRODUCED, "0"},
	}

	for _, m := range metrics {
		sendMetricIfNonZero(ch, metric.desc, metric.valueType, m.value, ts.LinkyId, m.mode, m.phase)
	}
}

func collectPowerMax(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_power_max"]

	metrics := []struct {
		value float64
		mode  string
		phase string
	}{
		{ts.PowerUsedMax, USED, "1"},
		{ts.PowerUsedMaxP1, USED, "1"},
		{ts.PowerUsedMaxP2, USED, "2"},
		{ts.PowerUsedMaxP3, USED, "3"},
		{ts.PowerProducedMax, PRODUCED, "0"},
	}

	for _, m := range metrics {
		sendMetricIfNonZero(ch, metric.desc, metric.valueType, m.value, ts.LinkyId, m.mode, m.phase)
	}
}

func collectPowerReference(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_power_reference"]
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.ReferencePower, ts.LinkyId, "subscribed")
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.BreakingPower, ts.LinkyId, "breaking")
}

func collectLoadCurvePoint(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_load_curve_point"]
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.UsedLoadCurvePoint, ts.LinkyId, USED)
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.ProducedLoadCurvePoint, ts.LinkyId, PRODUCED)
}

func collectLoadCurvePointLastYear(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_load_curve_point_last_year"]
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.UsedLoadCurvePointLastYear, ts.LinkyId, USED)
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.ProducedLoadCurvePointLastYear, ts.LinkyId, PRODUCED)
}

func collectAverageVoltage(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_voltage_average"]
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.AverageVoltageP1, ts.LinkyId, "1")
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.AverageVoltageP2, ts.LinkyId, "2")
	sendMetricIfNonZero(ch, metric.desc, metric.valueType, ts.AverageVoltageP3, ts.LinkyId, "3")
}

func collectStatus(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_status"]

	statusMetrics := []struct {
		value float64
		name  string
	}{
		{ts.DryContactStatus, "Contact sec"},
		{ts.CutOffDeviceStatus, "Organe de coupure"},
		{ts.LinkyTerminalShieldStatus, "État du cache-bornes distributeur"},
		{ts.SurgeStatus, "Surtension sur une des phases"},
		{ts.ReferencePowerExceededStatus, "Dépassement de la puissance de référence"},
		{ts.ConsumptionStatus, "Fonctionnement producteur/consommateur"},
		{ts.EnergyDirectionStatus, "Sens de l énergie active"},
		{ts.ContractTypePriceStatus, "Tarif en cours sur le contrat fourniture"},
		{ts.ContractTypePriceDistributorStatus, "Tarif en cours sur le contrat distributeur"},
		{ts.ClockStatus, "Mode dégradée de l horloge"},
		{ts.TicStatus, "État de la sortie télé-information"},
		{ts.EuridisLinkStatus, "État de la sortie communication Euridis"},
		{ts.CPLStatus, "Statut du CPL"},
		{ts.CPLSyncStatus, "Synchronisation CPL"},
		{ts.TempoContractColorStatus, "Couleur du jour pour le contrat historique tempo"},
		{ts.TempoContractNextDayColorStatus, "Couleur du lendemain pour le contrat historique tempo"},
		{ts.MovingPeakNoticeStatus, "Préavis pointers mobiles"},
		{ts.MovingPeakStatus, "Pointe mobile (PM)"},
	}

	for _, sm := range statusMetrics {
		sendMetric(ch, metric.desc, metric.valueType, sm.value, ts.LinkyId, sm.name)
	}
}

func collectMovablePeak(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_movable_peak"]

	peakMetrics := []struct {
		value float64
		type_ string
		phase string
	}{
		{ts.MovingPeakStart1, "start", "1"},
		{ts.MovingPeakEnd1, "end", "1"},
		{ts.MovingPeakStart2, "start", "2"},
		{ts.MovingPeakEnd2, "end", "2"},
		{ts.MovingPeakStart3, "start", "3"},
		{ts.MovingPeakEnd3, "end", "3"},
	}

	for _, pm := range peakMetrics {
		sendMetric(ch, metric.desc, metric.valueType, pm.value, ts.LinkyId, pm.type_, pm.phase)
	}
}

func collectRelay(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_relay"]

	relayMetrics := []struct {
		value float64
		id    string
	}{
		{ts.Relay1, "1"},
		{ts.Relay2, "2"},
		{ts.Relay3, "3"},
		{ts.Relay4, "4"},
		{ts.Relay5, "5"},
		{ts.Relay6, "6"},
		{ts.Relay7, "7"},
		{ts.Relay8, "8"},
	}

	for _, rm := range relayMetrics {
		sendMetric(ch, metric.desc, metric.valueType, rm.value, ts.LinkyId, rm.id)
	}
}

func collectProviderDayInfo(ch chan<- prometheus.Metric, lc *LinkyCollector, ts *LinkyTimeSerie) {
	metric := lc.metrics["linky_provider_day_info"]
	sendMetric(ch, metric.desc, metric.valueType, 1,
		ts.LinkyId, ts.Prm, ts.ContractTypeDayNumber,
		ts.ContractTypeNextDayNumber, ts.ContractTypeNextDayProfile)
}
