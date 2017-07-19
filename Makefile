GO_SRC := $(shell find $(CURDIR) -name '*.go')
BIN := redirect

build: $(BIN)

$(BIN): $(GO_SRC)
	go build -o $(BIN) src/main.go

test:
	go test -v github.com/manGoweb/www-redirect/src/...

push: build test
	git push

clean:
	git clean -f -x .

.PHONY: test clean
