package common

import (
	"log"
	"net/http"
)

func DoGetRequest(endpoint string) {
	_, err := http.Get(endpoint)

	if err != nil {
		log.Printf("Error while making request to %s: %v", endpoint, err)
	}
}
