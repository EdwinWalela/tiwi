all: clean e2e build

clean:
	rm -rf tiwi && rm -rf site

e2e:
	go test -v ./test

build:clean
	go build -o tiwi cmd/tiwi/main.go 