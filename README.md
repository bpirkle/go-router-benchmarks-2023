# go-router-benchmarks-2023

This is a small project to informally compare the performance of several Go routers/frameworks.

This is NOT a complete or highly accurate comparison. It does not attempt to exercise the more advanced features of each router/framework compared, which might be meaningful. Assuming that it is run locally, it is also subject to effects from and limitations of the local environment that might not be representative of a proper production server.

The goal is only to do a rough order-of-magnitude level comparision to confirm there are no glaring differences, not to find small percentage differences.

### Building

This project is slightly fidgety, as it builds four separate binaries. You will want to install dependencies:

```sh-session
go get
go get github.com/labstack/echo
go get github.com/gin-gonic/gin
go get github.com/gorilla/mux
go get github.com/valyala/fasthttp
go get github.com/fasthttp/router
```
then execute

```sh-session
make build
```

Don't just execute "go build". It'll (correctly) complain there are multiple main functions.

## Benchmarking

Wouldn't it be neat if this was written. I should do that.
