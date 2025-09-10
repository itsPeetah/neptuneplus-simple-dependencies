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
	t0 := time.Now()
	log.Printf("Function waiter received a request from %s", r.RemoteAddr)

	q := r.URL.Query()
	delay, err := strconv.Atoi(q.Get("d"))
	if err != nil {
		delay = 1
	}

	log.Printf("Sleeping for %d00ms...", delay)

	delayDuration := time.Duration(delay) * (time.Millisecond * 100)
	time.Sleep(delayDuration)

	diff := time.Now().UnixMilli() - t0.UnixMilli()
	fmt.Fprintf(w, "Slept for %d00ms (%dms).\n", delay, diff)
}

func main() {
	http.HandleFunc("/call", handler)
	http.HandleFunc(common.ROUTE_READY, common.HandleReady)
	http.HandleFunc(common.ROUTE_HEALTH, common.HandleHealth)
	log.Printf("Function waiter starting on port %d", 8080)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil); err != nil {
		log.Fatalf("Function waiter failed to start: %v", err)
	}
}
