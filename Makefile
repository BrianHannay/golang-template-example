run: clean test main
	./main


clean:
	rm -f main

test:
	go test ./... -v

main: main.go
	go build main.go
	chmod +x main
