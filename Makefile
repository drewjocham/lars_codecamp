PROJ_PATH=${CURDIR}
UID?=${shell id -u}
GID?=${shell id -g}
DEPS_IMAGE?=${memominsk/protobuf-alpine:latest}

.PHONY: mod-vendor
mod-vendor: ## Download, verify, and vendor module dependencies
	go mod verify
	go mod tidy
	go mod vendor

.PHONY: network
network: ## create network
	docker create network lars_codecamp_default

.PHONY: database-up
database-up: ## start postgres DB
	docker-compose up -d

.PHONY: database-down
database-down: ## shutdown postgres DB
	docker-compose down

.PHONY: proto
proto: ## Generate protobuf code
	mkdir -p pkg/api
# Compile proto files inside the project.
	protoc --proto_path=${PROJ_PATH}/proto/api -I book.proto \
			--go_out=. --go-grpc_out=pkg/api \

.PHONY: proto-docker
proto-docker: ## Generate protobuf code
	docker run --rm -v $(pwd):/mnt memominsk/protobuf-alpine:latest \
	--go_out=pkg/api --go-grpc_out=pkg/api \
	--proto_path=${PROJ_PATH}/proto/api ${PROJ_PATH}/proto/api/book.proto

.PHONY: interface_2 ## Run demo interfaces_2
interfaces2:
	cd interfaces_2 && go run .
