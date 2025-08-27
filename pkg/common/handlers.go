package common

import (
	"fmt"
	"net/http"
)

const (
	ROUTE_READY  string = "/_/ready"
	ROUTE_HEALTH string = "/health"
)

func HandleReady(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func HandleHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
