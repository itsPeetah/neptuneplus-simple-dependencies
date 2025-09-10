package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/itspeetah/neptuneplus-simple-dependencies/pkg/callers"
	"github.com/itspeetah/neptuneplus-simple-dependencies/pkg/common"
)

func main() {
	http.HandleFunc(common.ROUTE_HEALTH, common.HandleHealth)
	http.HandleFunc(common.ROUTE_READY, common.HandleReady)
	http.HandleFunc("/call", handleCall)

	log.Printf("Function sequential starting on port %d", 8080)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil); err != nil {
		log.Fatalf("Function sequential failed to start: %v", err)
	}
}

func handleCall(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now().UnixMilli()
	for i := 1; i <= callers.OUTDEGREE; i++ {
		t1 := time.Now().UnixMilli()
		url := strings.Replace(fmt.Sprintf("%s?d=%d", callers.URL_WAITER, i), "#", strconv.Itoa(i), 1)
		common.DoGetRequest(url)
		log.Printf("request to %s completed in %dms.", url, time.Now().UnixMilli()-t1)
	}

	fmt.Fprintf(w, "Completed %d calls in %d.", callers.OUTDEGREE, time.Now().UnixMilli()-t0)
}
