.PHONY: build
build:
	go build -o out/server ./server/server.go

.PHONY: start
start: build
	./out/server
