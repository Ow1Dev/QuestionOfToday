GO ?= go

AIR_PACKAGE ?= github.com/air-verse/air@v1
MIGRATE_PACKAGE ?= -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: deps
deps: deps-backend deps-tools ## install dependencies

.PHONY: deps-backend
deps-backend: ## install backend dependencies
	$(GO) mod download

.PHONY: deps-tools
deps-tools: ## install tool dependencies
	$(GO) install $(AIR_PACKAGE) & \
	$(GO) install $(MIGRATE_PACKAGE) & \
	wait
