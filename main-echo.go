package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "echo")
}
func Exit(c echo.Context) error {
	log.Fatal("exiting")
	return c.String(http.StatusOK, "exiting")
}
func main() {
	router := echo.New()
	router.GET("/exit", Exit)
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
