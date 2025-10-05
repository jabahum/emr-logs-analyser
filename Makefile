APP_NAME := emr-logs-analyser
VERSION := 1.0.0
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(BUILD_DATE)"
DIST_DIR := dist

# Detect OS
UNAME_S := $(shell uname -s)

# Default goal
.DEFAULT_GOAL := help

# 🧩 Build for current platform
build: ## Build for current OS/Arch
	go build $(LDFLAGS) -o $(APP_NAME) main.go

# 🧩 Clean build artifacts
clean: ## Remove build and dist directories
	rm -rf $(DIST_DIR) $(APP_NAME)

# 🧩 Build for all major OS/Arch combinations
build-all: clean ## Build for Windows, macOS, and Linux
	@mkdir -p $(DIST_DIR)
	@echo "🚀 Building binaries..."
	GOOS=linux   GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(APP_NAME)-linux-amd64 main.go
	GOOS=darwin  GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(APP_NAME)-darwin-amd64 main.go
	GOOS=darwin  GOARCH=arm64 go build $(LDFLAGS) -o $(DIST_DIR)/$(APP_NAME)-darwin-arm64 main.go
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(APP_NAME)-windows-amd64.exe main.go
	@echo "✅ Binaries available in ./$(DIST_DIR)"

# 🧩 Create zipped release archives
package: build-all ## Build and package distributable release archives
	@echo "📦 Packaging release archives..."
	cd $(DIST_DIR) && \
	zip $(APP_NAME)-linux-amd64.zip $(APP_NAME)-linux-amd64 && \
	zip $(APP_NAME)-darwin-amd64.zip $(APP_NAME)-darwin-amd64 && \
	zip $(APP_NAME)-darwin-arm64.zip $(APP_NAME)-darwin-arm64 && \
	zip $(APP_NAME)-windows-amd64.zip $(APP_NAME)-windows-amd64.exe
	@echo "🎉 Release archives created in ./$(DIST_DIR)"

# 🧩 Trigger a GitHub release via tagging
release: ## Create a git tag and push to trigger GitHub release workflow
	@if [ -z "$(v)" ]; then \
		echo "❌ Please provide a version, e.g. 'make release v=1.0.0'"; \
		exit 1; \
	fi
	@echo "🏷️  Creating and pushing git tag v$(v)..."
	git tag v$(v)
	git push origin v$(v)
	@echo "🚀 Release v$(v) pushed! GitHub Actions will build and upload binaries automatically."

# 🧩 Run CLI with arguments
run: ## Run the CLI (e.g. make run args="analyse -f catalina.out --level=ERROR")
	./$(APP_NAME) $(args)

# 🧩 Install binary system-wide
install: build ## Install CLI globally (requires sudo on Linux/macOS)
ifeq ($(UNAME_S),Linux)
	@echo "🪶 Installing on Linux..."
	sudo mv $(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "✅ Installed to /usr/local/bin/$(APP_NAME)"
else ifeq ($(UNAME_S),Darwin)
	@echo "🍎 Installing on macOS..."
	sudo mv $(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "✅ Installed to /usr/local/bin/$(APP_NAME)"
else
	@echo "⚙️  Windows detected. Please copy dist/$(APP_NAME)-windows-amd64.exe to a directory in your PATH manually."
endif

# 🧩 Uninstall the CLI
uninstall: ## Remove the CLI from system path
ifeq ($(UNAME_S),Linux)
	@echo "🧹 Uninstalling from Linux..."
	sudo rm -f /usr/local/bin/$(APP_NAME)
	@echo "✅ Removed /usr/local/bin/$(APP_NAME)"
else ifeq ($(UNAME_S),Darwin)
	@echo "🧹 Uninstalling from macOS..."
	sudo rm -f /usr/local/bin/$(APP_NAME)
	@echo "✅ Removed /usr/local/bin/$(APP_NAME)"
else
	@echo "⚙️  On Windows, manually delete dist/$(APP_NAME)-windows-amd64.exe or wherever you installed it."
endif

# 🧩 Display all available commands
help: ## Show help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo "\nVersion: $(VERSION) | Commit: $(COMMIT) | Built: $(BUILD_DATE)"
