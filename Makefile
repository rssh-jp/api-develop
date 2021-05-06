setup:
	@echo "### EXECUTE setup"
	go get github.com/cespare/reflex
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.3.8

mkdir: setup
	@echo "### EXECUTE mkdir"
	mkdir -p internal/http/echo/gen/

generate: mkdir
	@echo "### EXECUTE generate"
	oapi-codegen -package gen -o internal/http/echo/gen/gen.go -generate "types,server" resource/openapi/openapi.yaml

build: clear generate
	@echo "### EXECUTE build"
	docker-compose build

up: build
	@echo "### EXECUTE up"
	docker-compose up

test: generate
	go test -v ./...

clear:
	docker container prune -f
	docker image prune -f
	docker network prune -f

.PHONY: setup mkdir generate build up test
