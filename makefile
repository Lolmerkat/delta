GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

BINARY_NAME = delta
VERSION ?= pre-0.1.0

BUILD_DIR = ./output

.PHONY: all build-linux build-windows clean

all: build-linux build-windows

build-windows: build-windows-amd64 build-windows-arm64

build-linux: build-linux-amd64 build-linux-arm64

build-linux-amd64:
	@echo "Building for Linux (amd64)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .

build-linux-arm64:
	@echo "Building for Linux (arm64)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -ldflags="-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 .

build-windows-amd64:
	@echo "Building for Windows (amd64)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .

build-windows-arm64:
	@echo "Building for Windows (arm64)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 $(GOBUILD) -ldflags="-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe .

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	$(GOCLEAN)
