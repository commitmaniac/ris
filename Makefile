ifneq (,$(wildcard .git))
VERSION ?= $(shell git describe --tags)
else
VERSION ?= 1.0.0
endif

RISFLAGS ?= -w -X main.Version="$(VERSION)"

ris: ris.go
	go build -ldflags "$(RISFLAGS)" -o $@

.PHONY: clean

clean:
	rm -rf ris
