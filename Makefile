# Makefile

# Paths and Variables
PROTOC = protoc
PROTO_DIR = proto
BOOK_PROTO_DIR = book/$(PROTO_DIR)
CALC_PROTO_DIR = calculator/$(PROTO_DIR)
BOOK_PROTO_FILES = $(BOOK_PROTO_DIR)/*.proto
CALC_PROTO_FILES = $(CALC_PROTO_DIR)/*.proto
GO_OUT_DIR = .
MODULE_NAME = github.com/sikderg/go-grpc-example
BOOK_SERVER_OUTPUT = bin/book/server
BOOK_CLIENT_OUTPUT = bin/book/client
CALC_SERVER_OUTPUT = bin/calculator/server
CALC_CLIENT_OUTPUT = bin/calculator/client
BOOK_SERVER_SOURCE = ./book/server
BOOK_CLIENT_SOURCE = ./book/client
CALC_SERVER_SOURCE = ./calculator/server
CALC_CLIENT_SOURCE = ./calculator/client

# Commands
.PHONY: all clean book calculator help about
project := book calculator

all: book calculator ## Generate Go files from the .proto files

book: ## Generate Go files from the book proto
	$(PROTOC) -I$(BOOK_PROTO_DIR) --go_out=$(GO_OUT_DIR) --go_opt=module=$(MODULE_NAME) --go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=module=$(MODULE_NAME) $(BOOK_PROTO_FILES)
	go build -o $(BOOK_SERVER_OUTPUT) $(BOOK_SERVER_SOURCE)
	go build -o $(BOOK_CLIENT_OUTPUT) $(BOOK_CLIENT_SOURCE)

calculator: ## Generate Go files from the calculator proto
	$(PROTOC) -I$(CALC_PROTO_DIR) --go_out=$(GO_OUT_DIR) --go_opt=module=$(MODULE_NAME) --go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=module=$(MODULE_NAME) $(CALC_PROTO_FILES)
	go build -o $(CALC_SERVER_OUTPUT) $(CALC_SERVER_SOURCE)
	go build -o $(CALC_CLIENT_OUTPUT) $(CALC_CLIENT_SOURCE)

clean: ## Remove generated .pb.go files
	find $(GO_OUT_DIR) -name '*.pb.go' -delete
	rm -rf $(BOOK_SERVER_OUTPUT)
	rm -rf $(BOOK_CLIENT_OUTPUT)
	rm -rf $(CALC_SERVER_OUTPUT)
	rm -rf $(CALC_CLIENT_OUTPUT)
	rm -rf bin

help: ## Display this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'
	@echo ""

about: ## Display information about the environment
	@echo "╔══════════════════════════════════════════════════════════════════╗"
	@echo "║                                                                  ║"
	@echo "║                        🦋 Environment Info 🦋                    ║"
	@echo "║                                                                  ║"
	@echo "╚══════════════════════════════════════════════════════════════════╝"
	@echo ""
	@echo "Go Version:"
	@go version
	@echo ""
	@echo "Protoc Version:"
	@$(PROTOC) --version
	@echo ""
	@echo "Current Directory:"
	@pwd
	@echo ""
	@echo "╔══════════════════════════════════════════════════════════════════╗"
	@echo "║                       🦋 End of Information 🦋                   ║"
	@echo "╚══════════════════════════════════════════════════════════════════╝"