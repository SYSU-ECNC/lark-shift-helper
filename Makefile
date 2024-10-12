.PHONY: build
build:
	@go build -o bin/lark-shift-helper cmd/lark-shift-helper/main.go

.PHONY: run
run: build
	@./bin/lark-shift-helper
