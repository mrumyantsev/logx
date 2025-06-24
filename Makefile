.SILENT:
.IGNORE:
.DEFAULT_GOAL := demo

.PHONY: demo
demo:
	go run ./cmd/demo

.PHONY: demo2
demo2:
	go run ./cmd/demo2
