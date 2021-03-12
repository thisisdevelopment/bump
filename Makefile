SHELL := /bin/bash
DIR := $(shell pwd)
BINARY_NAME := bump
OUTPUT_DIR := ${DIR}/bin
# govvv will inject version specific build flags 
FLAGS :=  $(shell sh -c 'govvv -flags')

clean:
	@echo ">> removing previous builds"
	@rm -rf $(OUTPUTDIR)

test:
	go test ./pkg/version/. -coverprofile coverage/version.out

# gc will open the coverage report in default browser
gc: 
	go tool cover -html=coverage/version.out

# install govvv if not exists 
govvv:
	go get -u github.com/ahmetb/govvv

# api will compile the executable
bump:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "$(FLAGS)" -o $(OUTPUT_DIR)/$(BINARY_NAME) main.go

# runs the project locally in dev mode
prod: test
	make bump
	echo "Generating SHA256 Binary Hash of executable"
	cat bin/app | shasum -a 256
	bin/bump -f

# .PHONY is used for reserving tasks words
.PHONY: clean test gc bump prod