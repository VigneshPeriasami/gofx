
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

.PHONY: clean-docker
clean-docker:
	docker-compose down --rmi all -v --remove-orphans

.PHONY: start-docker
start-docker:
	docker-compose up