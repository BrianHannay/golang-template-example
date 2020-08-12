run: clean test format plugins main
	./main

plugins:
	go build -o plugins -buildmode=plugin plugins/*

main: plugins main.go
	go build main.go
	chmod +x main

format:
	find . -type f -name '*.go' -exec gofmt -w -e -s -d {} \;

test:
	go test ./... -v

clean:
	rm -f main

.PHONY: run clean test format plugins
