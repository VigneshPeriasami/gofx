
.PHONY: build

build:
	go build ./...

.PHONY: test
test:
	go test ./...

.PHONY: install
install:
	go mod download

.PHONY: run
run:
	go run main.go