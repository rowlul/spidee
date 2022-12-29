BINARY = spidee
VERSION = $(shell git describe HEAD)
LDFLAGS = -X 'github.com/rowlul/spidee/cli.Version=$(VERSION)' -s -w

all: test fmt vet clean build

install: all
	install -D -m 0755 build/spidee /usr/local/bin

install-user: all
	install -D -m 0775 build/spidee ${HOME}/.local/bin

test:
	@go test -v ./...

fmt:
	@go fmt ./...

vet:
	@go vet ./...

clean:
	@rm -rf build
	@go clean

.PHONY: build
build: test fmt vet clean
		go build \
		-ldflags "$(LDFLAGS)" \
		-o build/spidee \
		github.com/rowlul/spidee
