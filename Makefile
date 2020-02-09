# Makefile
export GO111MODULE=on

GO_CMD=go
GO_GET=$(GO_CMD) get -v
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_TEST_COVER=$(GO_CMD) test -cover -count=1 -p=1
GO_INSTALL=$(GO_CMD) install -v

SERVER_BIN=wether_monster
SERVER_BIN_DIR=.
SERVER_DIR=.
SERVER_MAIN=main.go

SOURCE_PKG_DIR= .
SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

all: dependencies tests build-server

dependencies:
	@echo "==> Installing dependencies ...";
	@$(GO_GET) ./...

tests:
	@echo "==> Running tests ...";
	@$(GO_TEST_COVER) $(SOURCE_PKG_DIR)/...

build-server:
	@echo "==> Building server ...";
	@GOOS=linux $(GO_BUILD) -o $(SERVER_BIN) -ldflags "-w -s" $(SERVER_DIR)/$(SERVER_MAIN) || exit 1;
	@chmod 755 $(SERVER_BIN)

run:
	./$(SERVER_BIN)
