MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
OUTPUT 		 := $(MAKEFILE_DIR)build/app
TARGET_DIR   := $(MAKEFILE_DIR)

CGO_ENABLED = 0
GOOS        = linux
GOARCH      = amd64

.PHONY: export
export:
	CGO_ENABLED=${CGO_ENABLED}
	GOOS=${GOOS}
	GOARCH=${GOARCH}
	export CGO_ENABLED GOOS GOARCH 

.PHONY: build
build:
	go build -o $(OUTPUT) $(TARGET_DIR)

.PHONY: swag
swag:
	swag init

.PHONY: mod
mod:
	go mod download
