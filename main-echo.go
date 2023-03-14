//go:build echo

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "echo")
}

func Exit(w http.ResponseWriter, r *http.Request) {
	log.Fatal("exiting")
}

func main() {
	router := echo.New()

	router.GET("/exit", echo.WrapHandler(http.HandlerFunc(Exit)))
	router.GET("/", echo.WrapHandler(http.HandlerFunc(Index)))

	log.Fatal(http.ListenAndServe(":8080", router))
}
