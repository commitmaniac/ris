RISFLAGS ?= -w -X main.Version="$(VERSION)"

ris: ris.go
	go build -ldflags "$(RISFLAGS)" -o $@

.PHONY: clean

clean:
	rm -rf ris
