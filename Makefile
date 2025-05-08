.PHONY: build
build:
	go build -o bin/server ./cmd/server/main.go

.PHONY: run dev
run-dev:
	go run ./cmd/server/main.go	
