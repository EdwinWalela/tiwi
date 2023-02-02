all: clean test build

clean:
	rm -rf tiwi && rm -rf test-site

test:
	go test -v ./pkg/...

build:
	go build -o tiwi cmd/main.go 