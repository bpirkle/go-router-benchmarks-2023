//go:build fasthttp

package main

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("fast")
}

func Exit(ctx *fasthttp.RequestCtx) {
	log.Fatal("exiting")
}

func main() {
	r := router.New()

	r.GET("/exit", Exit)
	r.GET("/", Index)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
