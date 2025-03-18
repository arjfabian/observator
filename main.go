package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
)

var cpuUsage = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_usage_percent",
	Help: "Current CPU usage in percent",
})

func collectCPUUsage() {
	for {
		percent, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Println("Error getting CPU usage:", err)
			return
		}
		cpuUsage.Set(percent[0])
		time.Sleep(time.Second)
	}
}

func main() {
	go collectCPUUsage()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

