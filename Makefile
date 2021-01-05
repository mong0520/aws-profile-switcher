PROJ_PATH := github.com/mong0520/aws-profile-switcher
BIN := aws-profile-switcher
GOOS ?= darwin
GOARCH ?= amd64
BUILD_CMD := go build -o build/${GOOS}-${GOARCH}/${BIN}

.PHONY: all build

build:
	@docker run --rm \
	-v ${PWD}:/go/src/${PROJ_PATH} \
	-e GOPATH=/go \
	-e GOOS=${GOOS} \
	-e GOARCH=${GOARCH} \
	-e GO111MODULE=on \
	-w /go/src/${PROJ_PATH} golang:1.14 \
	${BUILD_CMD}

install: build
	cp build/${GOOS}-${GOARCH}/${BIN} /usr/local/bin/
	cp run_aws-profile-switcher.sh /usr/local/bin/
	@echo "==================================="
	@echo "Add the following to your .bashrc or .zshrc config, and the reload your shell"
	@echo "--------------------------------"
	@echo "    source ~/.aws_exports"
	@echo "    alias aws-profile-switcher=\"source run_aws-profile-switcher.sh\""
	@echo "--------------------------------"
	@echo "run 'aws-profile-switcher' to switch AWS profie"
	@echo "==================================="
