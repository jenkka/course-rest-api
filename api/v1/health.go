package v1

import "net/http"

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("health check handler!\n"))
}
