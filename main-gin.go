//go:build gin

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "mux")
}

func Exit(w http.ResponseWriter, r *http.Request) {
	log.Fatal("exiting")
}

func main() {
	router := gin.Default()

	router.GET("/", gin.WrapF(Index))
	router.GET("/exit", gin.WrapF(Exit))

	log.Fatal(http.ListenAndServe(":8080", router))
}
