package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "gin")
}
func Exit(c *gin.Context) {
	log.Fatal("exiting")
}
func main() {
	router := gin.Default()
	router.GET("/exit", Exit)
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
