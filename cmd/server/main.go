package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"testHttpServer/internal/metrics"
)


func workAnimation() {
	frames := []string{"|", "/", "-", "\\"}
	for {
		for _, frame := range frames {
			fmt.Printf("\r%s Server has started...", frame)
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func main() {
	mux := http.NewServeMux()
	req := prometheus.NewRegistry()
	m := metrics.NewMetrics(req)

	mux.HandleFunc("/", m.Core())
	mux.Handle("/metrics", promhttp.HandlerFor(req, promhttp.HandlerOpts{Registry: req}))
	mux.HandleFunc("/health", m.HealthStatus())
	mux.HandleFunc("/db", m.DBStatus())
	mux.HandleFunc("/random", m.RandomStatus())

	server := http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go m.UptimeSet()
	go m.CPUTempSetRandom()
	go workAnimation()
	fmt.Println("-- Prepare server --")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting the server", err)
	}
}
