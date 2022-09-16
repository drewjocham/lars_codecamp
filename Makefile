
.PHONY: mod-vendor
mod-vendor: ## Download, verify, and vendor module dependencies
	go mod download
	go mod verify
	go mod tidy
	go mod vendor

.PHONY: database
database: ## start postgres DB
	docker-compose up -d
