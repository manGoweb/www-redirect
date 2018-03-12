GO_SRC := $(shell find $(CURDIR) -name '*.go')
BIN := redirect
REPO := mangoweb/www-redirect
TAG := 1.0

build: $(BIN)

$(BIN): $(GO_SRC)
	go build -o $(BIN) src/main.go

test:
	go test -v github.com/manGoweb/www-redirect/src/...

push: build test
	git push

image:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -i -o main src/main.go
	docker build -t $(REPO):$(TAG) .
	rm -rf main

clean:
	git clean -f -x .

.PHONY: test clean
