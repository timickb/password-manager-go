.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: build
build:
	go build -v -o "./bin/gpm" ./cmd/gpm


.DEFAULT_GOAL := build
