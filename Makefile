.SILENT:
.DEFAULT_GOAL := demo

.PHONY: demo
demo:
	go run ./cmd/logx-demo

.PHONY: demo2
demo2:
	go run ./cmd/logx-demo2
