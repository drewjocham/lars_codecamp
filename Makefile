PROJ_PATH=${CURDIR}
UID?=${shell id -u}
GID?=${shell id -g}
DEPS_IMAGE?=${memominsk/protobuf-alpine:latest}

.PHONY: mod-vendor
mod-vendor: ## Download, verify, and vendor module dependencies
	go mod download
	go mod verify
	go mod tidy
	go mod vendor

.PHONY: database
database: ## start postgres DB
	docker-compose up -d

.PHONY: proto
proto: ## Generate protobuf code
	mkdir -p pkg/api
# Compile proto files inside the project.
	protoc -I=grpc/api \
		--go_out=. \
		--go-grpc_out=. \
		grpc/api/*.proto

.PHONY: proto-docker
proto-docker: ## Generate protobuf code
	docker run --rm -v $(pwd):/mnt memominsk/protobuf-alpine:latest --go_out=pkg/api  ${CURDIR}/proto/api/book.proto
