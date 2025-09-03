package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/itspeetah/neptuneplus-simple-dependencies/pkg/common"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Function waiter received a request from %s", r.RemoteAddr)

	q := r.URL.Query()
	delay, err := strconv.Atoi(q.Get("d"))
	if err != nil {
		delay = 1
	}

	log.Printf("Sleeping for %d seconds...", delay)

	delayDuration := time.Duration(delay) * time.Second / 2
	time.Sleep(delayDuration)

	fmt.Fprintf(w, "Slept for %d seconds.\n", delay)
}

func main() {
	http.HandleFunc("/call", handler)
	http.HandleFunc(common.ROUTE_READY, common.HandleReady)
	http.HandleFunc(common.ROUTE_HEALTH, common.HandleHealth)
	log.Printf("Function W2 server starting on port %d", 8080)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil); err != nil {
		log.Fatalf("Function W2 server failed to start: %v", err)
	}
}
