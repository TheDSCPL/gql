.PHONY: download-deps
download-deps:
	go mod download

.PHONY: download-tools
download-tools: download-deps
	@echo "Installing tools..."
	@cat tools/tools.go | grep _ | cut -d'"' -f2 | xargs -t -I% go install %

.PHONY: check
check: download-tools
	@echo "Checking code..."
	go mod tidy
	golangci-lint run
	go test ./...