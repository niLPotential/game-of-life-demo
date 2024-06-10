package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/alpine", AlpineHandler)

	http.HandleFunc("/htmx", HtmxHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
