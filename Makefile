
.PHONY: build
build:
	go build ./...

.PHONY: buildout
buildout:
	go build -o bin/

.PHONY: test
test:
	go test ./...

.PHONY: install
install:
	go mod download

.PHONY: run
run:
	go run main.go