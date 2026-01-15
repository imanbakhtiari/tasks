package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

var buffer []byte

func init() {
	buffer = make([]byte, 100*1024*1024)
	for i := range buffer {
		buffer[i] = byte(i % 256)
	}
}

func main() {

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)

	metricsMux := http.NewServeMux()
	metricsMux.HandleFunc("/metrics", metricsHandler)
	go http.ListenAndServe(":8090", metricsMux)

	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"healthy"}`))
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	info := map[string]any{
		"time_utc":  time.Now().UTC().Format(time.RFC3339),
		"pod_name":  os.Getenv("POD_NAME"),
		"namespace": os.Getenv("POD_NAMESPACE"),
		"node_name": os.Getenv("NODE_NAME"),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

