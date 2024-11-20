SHELL := /bin/bash
PROTO_DIR := proto/shared
GO_OUT_DIR := proto/pb
GOPATH := $(shell go env GOPATH)

.PHONY: proto
proto:
	mkdir -p $(GO_OUT_DIR)
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	PATH=$(GOPATH)/bin:$$PATH find $(PROTO_DIR) -name "*.proto" -exec protoc \
    		--proto_path=$(PROTO_DIR) \
    		--go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
    		--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
    		--experimental_allow_proto3_optional \
    		{} \;