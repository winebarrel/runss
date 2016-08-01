VERSION=$(shell git tag | tail -n 1)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
RUNTIME_GOPATH=$(GOPATH):$(shell pwd)
SRC=$(wildcard *.go) $(wildcard src/*/*.go)

all: runss

go-get:
	go get github.com/aws/aws-sdk-go

runss: go-get main.go $(SRC)
	GOPATH=$(RUNTIME_GOPATH) go build

clean:
	rm -f runss *.gz

package: clean runss
	gzip -c runss > runss-$(VERSION)-$(GOOS)-$(GOARCH).gz
