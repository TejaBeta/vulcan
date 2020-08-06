package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "vulcan",
			Name:      "request_counter",
			Help:      "Total number of requests",
		})
)

func generalLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCounter.Inc()
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	prometheus.Register(requestCounter)
	r.Use(generalLog)
	r.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8001", r)
}
