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
	fmt.Fprintf(w, "gin")
}

func Exit(w http.ResponseWriter, r *http.Request) {
	log.Fatal("exiting")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/", gin.WrapF(Index))
	router.GET("/exit", gin.WrapF(Exit))

	log.Fatal(http.ListenAndServe(":8080", router))
}
