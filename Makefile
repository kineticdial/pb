default: build test lint
build:
	go build -v ./...
	go build -v
install:
	go install
test:
	go test -v ./...
lint:
	golint -set_exit_status
