//go:build http

package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "http")
}

func Exit(w http.ResponseWriter, r *http.Request) {
	log.Fatal("exiting")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/exit", Exit)
	mux.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
