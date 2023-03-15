# Copyright 2023 Bill Pirkle  <bpirkle@wikimedia.org> and Wikimedia Foundation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

VERSION     = $(shell /usr/bin/git describe --always)
BUILD_DATE  = $(shell date -u +%Y-%m-%dT%T:%Z)
HOSTNAME    = $(shell hostname)

GO_LDFLAGS  = -X main.version=$(if $(VERSION),$(VERSION),unknown)
GO_LDFLAGS += -X main.buildDate=$(if $(BUILD_DATE),$(BUILD_DATE),unknown)
GO_LDFLAGS += -X main.buildHost=$(if $(HOSTNAME),$(HOSTNAME),unknown)

install:
	go install

get:
	go get
	go get github.com/labstack/echo
	go get github.com/gin-gonic/gin
	go get github.com/gorilla/mux
	go get github.com/valyala/fasthttp
	go get github.com/fasthttp/router

build:
	@echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	@echo "VERSION ......: $(VERSION)"
	@echo "BUILD HOST ...: $(HOSTNAME)"
	@echo "BUILD DATE ...: $(BUILD_DATE)"
	@echo "GO VERSION ...: $(word 3, $(shell go version))"
	@echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	go build -ldflags "$(GO_LDFLAGS)" -tags echo -o go-echo ./main-echo.go
	go build -ldflags "$(GO_LDFLAGS)" -tags gin -o go-gin ./main-gin.go
	go build -ldflags "$(GO_LDFLAGS)" -tags mux -o go-mux ./main-mux.go
	go build -ldflags "$(GO_LDFLAGS)" -tags fasthttp -o go-fasthttp ./main-fasthttp.go
	go build -ldflags "$(GO_LDFLAGS)" -tags http -o go-http ./main-http.go

run:
	go run -ldflags "$(GO_LDFLAGS)" ./go-echo

check:
	@if [ -n "`goimports -l *.go`" ]; then \
	    echo "goimports: format errors detected" >&2; \
	    false; \
	fi
	@if [ -n "`gofmt -l *.go`" ]; then \
	    echo "gofmt: format errors detected" >&2; \
	    false; \
	fi

clean:
	rm -f go-echo
	rm -f go-gin
	rm -f go-mux
	rm -f go-fasthttp
	rm -f go-http

.PHONY: build run test clean
