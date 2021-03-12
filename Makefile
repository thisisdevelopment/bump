DIR := $(shell pwd)
BINARY_NAME := bump
OUTPUT_DIR := ${DIR}/bin
# govvv will inject version specific build flags 
FLAGS :=  $(shell sh -c 'govvv -flags')
VERSION := ${shell sh -c 'cat VERSION'}

clean:
	@echo ">> removing previous builds"
	@rm -rf $(OUTPUT_DIR)

test:
	go test ./pkg/version/. -coverprofile coverage/version.out

# gc will open the coverage report in default browser
gc: 
	@go tool cover -html=coverage/version.out

# install govvv if not exists 
govvv:
	go get -u github.com/ahmetb/govvv

# will compile the executable
bump:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "$(FLAGS)" -o $(OUTPUT_DIR)/$(BINARY_NAME) main.go

# will test compile and sha the binary
prod: test
	make bump
	@echo ">> Generating SHA256 Binary Hash of executable"
	@cat $(OUTPUT_DIR)/$(BINARY_NAME) | shasum -a 256
	@cp $(OUTPUT_DIR)/$(BINARY_NAME) ${GOPATH}/bin/$(BINARY_NAME)
	@echo ">> try running $(BINARY_NAME) -h"

tag:
	git tag ${VERSION}
	git push --tags

# .PHONY is used for reserving tasks words
.PHONY: clean test gc govvv bump prod tag