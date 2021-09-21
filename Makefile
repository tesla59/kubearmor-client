CURDIR=$(shell pwd)

ifeq (,$(shell which govvv))
$(shell go install github.com/ahmetb/govvv@latest)
endif

PKG := $(shell go list ./version)
GIT_INFO := $(shell govvv -flags -pkg $(PKG))

.PHONY: build
build:
	cd $(CURDIR)
	go mod tidy
	CGO_ENABLED=0 go build \
	-ldflags "-w -s ${GIT_INFO}" \
	-o karmor