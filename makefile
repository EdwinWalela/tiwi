all: clean e2e build

clean:
	rm -rf tiwi && rm -rf test-site

e2e:
	go test -v ./pkg/...

build:
	go build -o tiwi cmd/main.go 