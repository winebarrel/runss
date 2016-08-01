VERSION=$(shell git tag | tail -n 1)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
RUNTIME_GOPATH=$(GOPATH):$(shell pwd)
SRC=$(wildcard *.go) $(wildcard src/*/*.go)

all: runss

go-get:
	go get github.com/aws/aws-sdk-go
	go get github.com/golang/mock/gomock

stringer:
	go get golang.org/x/tools/cmd/stringer
	cd src/runss && stringer -type CommandStatus

runss: go-get $(SRC)
	GOPATH=$(RUNTIME_GOPATH) go build

clean:
	rm -f runss *.gz

package: clean runss
	gzip -c runss > runss-$(VERSION)-$(GOOS)-$(GOARCH).gz

mock:
	go get github.com/golang/mock/mockgen
	mkdir -p src/mockaws
	mockgen -source $(GOPATH)/src/github.com/aws/aws-sdk-go/service/ssm/ssmiface/interface.go -destination src/mockaws/ssmmock.go -package mockaws
