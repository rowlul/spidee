BINARY = spidee
VERSION = $(shell git describe HEAD)
LDFLAGS = -X 'github.com/rowlul/spidee/internal/cmd.Version=$(VERSION)' -s -w

all: test fmt vet clean build

install: all
	install -D -m 0755 build/$(BINARY) /usr/local/bin

install-user: all
	install -D -m 0775 build/$(BINARY) ${HOME}/.local/bin

test:
	@go test -v ./...

fmt:
	@go fmt ./...

vet:
	@go vet ./...

clean:
	@rm -rf build
	@rm -rf dist
	@go clean

.PHONY: build
build:
	go build \
	-ldflags "$(LDFLAGS)" \
	-o build/$(BINARY) \
	./cmd/spidee
