SHELL := /bin/sh
LOCAL_BIN := $(CURDIR)/bin
# Если protoc есть в PATH — используем его, иначе пробуем типичное место:
PROTOC := $(or $(shell command -v protoc),/usr/local/bin/protoc)
PROTO_SRC := api/noteV1
OUT_DIR   := pkg/noteV1

.PHONY: install-deps get-deps generate generate-note-api

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate:
	$(MAKE) generate-note-api

generate-note-api:
	mkdir -p $(OUT_DIR)
	"$(PROTOC)" -I "$(PROTO_SRC)" \
	  --plugin=protoc-gen-go="$(LOCAL_BIN)/protoc-gen-go" \
	  --go_out="$(OUT_DIR)" --go_opt=paths=source_relative \
	  --plugin=protoc-gen-go-grpc="$(LOCAL_BIN)/protoc-gen-go-grpc" \
	  --go-grpc_out="$(OUT_DIR)" --go-grpc_opt=paths=source_relative \
	  "$(PROTO_SRC)/note.proto"
