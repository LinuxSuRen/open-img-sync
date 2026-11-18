# open-img-sync 构建入口。make build 不依赖 Node（前端产物已提交），
# make web 才需要 Node 20+。

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
BUILDTIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -X github.com/linuxsuren/open-img-sync/internal/cli.version=$(VERSION) \
           -X github.com/linuxsuren/open-img-sync/internal/cli.commit=$(COMMIT) \
           -X github.com/linuxsuren/open-img-sync/internal/cli.buildTime=$(BUILDTIME)
BINARY := bin/open-img-sync

.PHONY: build test fmt vet clean

build:
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/open-img-sync

test:
	go test -race ./...

fmt:
	gofmt -l -w .

vet:
	go vet ./...

clean:
	rm -rf bin/
