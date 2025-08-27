APP_NAME := pastebin
BIN_DIR := bin
MAIN := main.go

.PHONY: run build tidy clean autorun

## Run the program without building
run:
	go run $(MAIN)

## Build the binary into ./bin
build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN)

## Run go mod tidy
tidy:
	go mod tidy

## Clean build artifacts
clean:
	rm -rf $(BIN_DIR)

## Build and run the binary
autorun: build
	./$(BIN_DIR)/$(APP_NAME)
