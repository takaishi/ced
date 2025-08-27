VERSION := $(shell git describe --tags --always --dirty --match=v* 2>/dev/null || echo "dev")
REVISION := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

build:
	go build -ldflags "-X main.Version=$(VERSION) -X main.Revision=$(REVISION)" -o bin/ced .

clean:
	rm -rf bin/
