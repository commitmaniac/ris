ris: ris.go
	go build -trimpath -o $@

.PHONY: clean

clean:
	rm -rf ris
