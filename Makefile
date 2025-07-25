APP_NAME := overlap
OUT_DIR := bin
BIN := $(OUT_DIR)/$(APP_NAME)

build:
	@echo "Building app..."
	@mkdir -p $(OUT_DIR)
	@go build -o $(BIN) ./cmd

test:
	@echo "Running tests..."
	@go test ./... -v -cover

fmt:
	@echo "Formatting code..."
	@gofmt -s -w .

clean:
	@echo "Cleaning..."
	@rm -rf $(OUT_DIR)

run:
	@echo "Running app..."
	@go run ./cmd

ci: clean fmt test build
	@echo "CI complete ✅"
