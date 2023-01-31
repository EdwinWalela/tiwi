all: clean e2e build

clean:
	rm -rf tiwi

e2e:
	go test -v ./test

build:
	go build -o tiwi cmd/tiwi/main.go 