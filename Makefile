BINARY = spidee
VERSION = $(shell git describe --tags --abbrev=0)

OS = $(shell go env GOOS)
ARCH = $(shell go env GOARCH)

LDFLAGS = -X 'github.com/rowlul/spidee/cli.Version=$(VERSION)' -s -w

all: test fmt vet clean build

install: all
	install -D -m 0755 bin/$(BINARY)-$(VERSION)-$(OS)-$(ARCH)/$(BINARY) /usr/bin/$(BINARY)

test:
	@go test -v ./...

fmt:
	@go fmt ./...

vet:
	@go vet ./...

clean:
	@rm -rf build
	@go clean

build:
	GOOS=$(OS) GOARCH=$(ARCH) \
		go build \
		-ldflags "$(LDFLAGS)" \
		-o build/$(BINARY)-$(VERSION)-$(OS)-$(ARCH)/ \
		github.com/rowlul/spidee

build-linux-amd64:
	GOOS=linux GOARCH=amd64 \
		go build \
		-ldflags "${LDFLAGS}" \
		-o build/${BINARY}-${VERSION}-linux-amd64/$(BINARY) \
		github.com/rowlul/spidee

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 \
		go build \
		-ldflags "$(LDFLAGS)" \
		-o build/${BINARY}-${VERSION}-darwin-amd64/$(BINARY) \
		github.com/rowlul/spidee

build-windows-amd64:
	GOOS=windows GOARCH=amd64 \
		go build \
		-ldflags "$(LDFLAGS)" \
		-o build/${BINARY}-${VERSION}-windows-amd64/$(BINARY).exe \
		github.com/rowlul/spidee
