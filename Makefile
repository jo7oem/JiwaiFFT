#
NAME := jiwaifft
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"
BINARY_UNIX := $(NAME)
BINARY_WINDOWS := $(NAME).exe


export GO111MODULE=on

## Install denpendencies
.PHONY: deps
deps:
	go get -v -d

## Setup
.PHONY: devel-deps
devel-deps: deps
	GO111MODULE=off go get \
		    golang.org/x/lint/golint \
		    github.com/motemen/gobump \
		    github.com/Songmu/make2help/cmd/make2help \
		    github.com/motemen/gobump \
		    golang.org/x/tools/cmd/goimports \


## Run tests
.PHONY: test
test: deps
	go test ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## build
bin/%: cmd/%/main.go deps
	go build -ldflags $(LDFLAGS) -o $@ $<

## build binary
.PHONY: build
build: bin/$(NAME) test lint


## build Linux binary
.PHONY: linux_build
linux_build: cmd/$(NAME)/main.go
	GOOS="linux" go build -ldflags $(LDFLAGS)  -o release/linux/$(BINARY_UNIX) $<

## build Windows binary
.PHONY: windows_build
windows_build: cmd/$(NAME)/main.go
	GOOS="windows" go build -ldflags $(LDFLAGS)  -o release/windows/$(BINARY_WINDOWS) $<

## build Mac binary
.PHONY: mac_build
mac_build: cmd/$(NAME)/main.go
	GOOS="darwin" go build -ldflags $(LDFLAGS)  -o release/mac/$(BINARY_UNIX) $<

## Release build
.PHONY: release
release: linux_build windows_build mac_build

release/%.zip: %_build
	zip -j -r $@ release/$*

## Releas Packing
.PHONY: release-packing
release-packing: release/linux.zip release/windows.zip release/mac.zip

## Clean up
.PHONY: clean
clean:
	@rm -rf bin/* release

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)

