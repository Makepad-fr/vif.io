.PHONY: stop
stop:
	docker compose stop

.PHONY: start
start:
	docker compose up -d --build

.PHONY: build
build:
	go build -o out/server ./server/server.go
