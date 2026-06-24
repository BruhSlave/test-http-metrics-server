// Package metrics register metrics for handlers
package metrics

import (
	"math/rand/v2"
	"time"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	CPUTemp       prometheus.Gauge
	RequestTotal  prometheus.Counter
	ErrorsTotal   prometheus.Counter
	UptimeSeconds prometheus.Counter
}

func NewMetrics(req prometheus.Registerer) *metrics {
	m := &metrics{
		CPUTemp: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "cpu_temp",
				Help: "CPU temperature",
			},
		),
		RequestTotal: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "request_total",
				Help: "Total request by all enpoints",
			},
		),
		ErrorsTotal: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "errors_total",
				Help: "Total errors",
			},
		),
		UptimeSeconds: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "uptime_seconds",
				Help: "Server uptime",
			},
		),
	}

	req.MustRegister(m.CPUTemp)
	req.MustRegister(m.RequestTotal)
	req.MustRegister(m.ErrorsTotal)
	req.MustRegister(m.UptimeSeconds)
	return m
}

func (m *metrics) CPUTempSetRandom() {
	for {
		time.Sleep(3 * time.Second)
		m.CPUTemp.Set(float64(rand.IntN(120)))
	}
}

func (m *metrics) UptimeSet() {
	for {
		time.Sleep(1 * time.Second)
		m.UptimeSeconds.Inc()
	}
}
