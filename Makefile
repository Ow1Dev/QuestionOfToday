GO ?= go

AIR_PACKAGE ?= github.com/air-verse/air@v1
MIGRATE_PACKAGE ?= -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
SQLC ?= github.com/sqlc-dev/sqlc/cmd/sqlc@latest

GO_SOURCES := $(wildcard *.go)
EXECUTABLE ?= questionofday

.PHONY: watch
watch: ## watch everything and continuously rebuild
	@bash tools/watch.sh

.PHONY: watch-backend
watch-backend:
	$(GO) run $(AIR_PACKAGE) -c .air.toml

.PHONY: watch-build-style
watch-build-style:
	@npx @tailwindcss/cli -i ./public/style/input.css -o ./public/style/output.css --watch

.PHONY: backend
backend: $(EXECUTABLE)

$(EXECUTABLE): $(GO_SOURCES) $(TAGS_PREREQ)
	$(GO) build $(GOFLAGS) -o ./tmp/$@

.PHONY: build-style
build-style:
	@npx @tailwindcss/cli -i ./public/style/input.css -o ./public/style/output.css

.PHONY: deps
deps: deps-frontend deps-backend deps-tools ## install dependencies

.PHONY: deps-frontend
deps-frontend: node_modules ## install frontend dependencies

.PHONY: deps-backend
deps-backend: ## install backend dependencies
	$(GO) mod download

.PHONY: deps-tools
deps-tools: ## install tool dependencies
	$(GO) install $(AIR_PACKAGE) & \
	$(GO) install $(MIGRATE_PACKAGE) & \
	$(GO) install $(SQLC) & \
	wait

node_modules: package-lock.json
	npm install --no-save
	@touch node_modules
