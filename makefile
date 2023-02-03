all: clean test build

clean:
	rm -rf tiwi && rm -rf test-site

test:
	rm -rf test-site && go test -v ./pkg/...

build:
	go build -o tiwi cmd/main.go 