package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "mux")
}
func Exit(w http.ResponseWriter, r *http.Request) {
	log.Fatal("exiting")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/exit", Exit)
	log.Fatal(http.ListenAndServe(":8080", router))
}
