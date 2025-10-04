# Default binary name
BINARY := emr-logs-analyser

# Default goal
.DEFAULT_GOAL := help

build: ## Build the Go binary
	go build -o $(BINARY) main.go

run: ## Run the EMR-LOGS-ANALYSER CLI (example: make run ARGS="version")
	@echo "Running: ./$(BINARY) $(ARGS)"
	@./$(BINARY) $(ARGS)

clean: ## Remove the built binary
	rm -f $(BINARY)

help: ## Show available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'

.PHONY: build run clean help