setup:
	@echo "### EXECUTE setup"
	# OpenAPI setup
	go get github.com/cespare/reflex
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.3.8
	# gRPC setup
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

mkdir: setup
	@echo "### EXECUTE mkdir"
	mkdir -p internal/http/echo/gen/

generate: mkdir
	@echo "### EXECUTE generate"
	# OpenAPI generate
	oapi-codegen -package gen -o internal/http/echo/gen/gen.go -generate "types,server" resource/openapi/openapi.yaml
	# gRPC generate
	protoc -I ./ --go_out=./ --go-grpc_out=./ resource/protocol-buffer/test-api.proto

build: clear generate
	@echo "### EXECUTE build"
	docker-compose build

up:
	@echo "### EXECUTE up"
	docker-compose up

test: generate
	go test -v ./...

clear:
	docker container prune -f
	docker image prune -f
	docker network prune -f

.PHONY: setup mkdir generate build up test
