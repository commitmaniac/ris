ifneq (,$(wildcard .git))
VERSION ?= $(shell git describe --tags)
else
VERSION ?= 1.2.0
endif

BIN ?= ris

RISFLAGS ?= -w -X main.Version="$(VERSION)"

ifeq ($(RELEASE),1)
RISFLAGS += -s
endif

$(BIN): ris.go
	go build -ldflags "$(RISFLAGS)" -o $@

.PHONY: clean

clean:
	rm -rf $(BIN)
