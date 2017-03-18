default: build test lint
build:
	go build -v
install:
	go install -v
test:
	go test -v ./{cli,lib}/...
lint:
	golint -set_exit_status
