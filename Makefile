# Usage: make clean build
# Copy contents from dist to desired location and edit config.json

# Collect git info
VERSION := $(shell git describe --tags)
ifeq (${VERSION},)
	VERSION = $(shell git rev-parse --short HEAD)
endif
COMMIT := $(shell git rev-parse --short HEAD)
BUILD := $(shell date +%FT%T%z)

# Prepare flags
# -X flags to inject version info into the execuable and -H=windowsgui to hide console
LDFLAGS=-ldflags "-H=windowsgui -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.build=${BUILD}"

ifeq ($(shell which go),)
	go := go.exe
	output := wbrowser.exe
else
	$(error go not found)
endif

clean:
	if [ -d ./dist ]; then rm ./dist -r; fi

build:
	mkdir dist -p
	$(info $(shell which go))
	${go} build ${LDFLAGS} -o dist/${output} ./cmd/wbrowser
	cp ./config.example.json dist/config.json
