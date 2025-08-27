package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/itspeetah/neptuneplus-simple-dependencies/pkg/callers"
	"github.com/itspeetah/neptuneplus-simple-dependencies/pkg/common"
)

func main() {
	http.HandleFunc(common.ROUTE_HEALTH, common.HandleHealth)
	http.HandleFunc(common.ROUTE_READY, common.HandleReady)
	http.HandleFunc("/call", handleCall)

	log.Printf("Function W2 server starting on port %d", 8080)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil); err != nil {
		log.Fatalf("Function W2 server failed to start: %v", err)
	}
}

func handleCall(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(callers.OUTDEGREE)

	callFunc := func(i int) {
		defer wg.Done()
		url := strings.Replace(fmt.Sprintf("%s?d=%d", callers.URL_WAITER, i), "#", strconv.Itoa(i), 1)

		common.DoGetRequest(url)
	}

	for i := 1; i <= callers.OUTDEGREE; i++ {
		go callFunc(i)
	}

	wg.Wait()

	fmt.Fprintf(w, "Completed %d calls", callers.OUTDEGREE)
}
