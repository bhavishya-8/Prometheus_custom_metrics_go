package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

type metrics struct {
	devices prometheus.Gauge
	info    *prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "connected_devices",
			Help:      "Number of devices connected",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "info",
			Help:      "Give version of device",
		},
			[]string{"version"}),
	}
	reg.MustRegister(m.devices, m.info)
	return m
}

var dvs []Device
var version string
var count int
var mutex sync.Mutex

func init() {
	version = "2.10.5"
	dvs = []Device{
		{1, "5F-2D", "2.1.6"},
		{2, "12w-12d", "3.1.2"},
		{2, "12w-12d", "3.1.2"},
	}
}

func incrementCount() {
	for {
		time.Sleep(2 * time.Second)
		mutex.Lock()
		count++
		mutex.Unlock()
	}
}

func main() {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	var wg sync.WaitGroup
	wg.Add(1)
	go incrementCount()

	go func() {
		for {
			// time.Sleep(1 * time.Second)
			mutex.Lock()
			m.devices.Set(float64(count))
			mutex.Unlock()
		}
		// wg.Done()
	}()

	m.info.With(prometheus.Labels{"version": version}).Set(1)
	// promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})
	// http.Handle("/metrics", promHandler)
	// http.HandleFunc("/devices", getDevices)
	// http.ListenAndServe(":8081", nil)

	dMux := http.NewServeMux()
	dMux.HandleFunc("/devices", getDevices)

	pMux := http.NewServeMux()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	pMux.Handle("/metrics", promHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", dMux))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8081", pMux))
	}()

	select {}
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(dvs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
