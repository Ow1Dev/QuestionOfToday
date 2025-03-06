GO ?= go

AIR_PACKAGE ?= github.com/air-verse/air@v1
MIGRATE_PACKAGE ?= -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
SQLC ?= github.com/sqlc-dev/sqlc/cmd/sqlc@latest

DOCKER_IMAGE ?= ow1dev/questionofday
DOCKER_TAG ?= latest
DOCKER_REF := $(DOCKER_IMAGE):$(DOCKER_TAG)

TAGS ?=

GO_SOURCES := $(wildcard *.go)

GOFLAGS := -v
EXECUTABLE ?= questionofday

.PHONY: all
all: build

.PHONY: help
help: Makefile ## print Makefile help information.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m[TARGETS] default target: build\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 }' Makefile #$(MAKEFILE_LIST)

.PHONY: clean-all
clean-all: clean ## delete backend, frontend and integration files
	rm -rf ./public/style/output.css node_modules

.PHONY: clean
clean: ## delete backend and integration files
	rm -rf $(EXECUTABLE)

.PHONY: watch
watch: ## watch everything and continuously rebuild
	@bash tools/watch.sh

.PHONY: watch-backend
watch-backend:
	$(GO) run $(AIR_PACKAGE) -c .air.toml

.PHONY: watch-build-style
watch-build-style:
	@npx @tailwindcss/cli -i ./public/style/input.css -o ./public/style/output.css --watch

.PHONY: build
build: backend build-style ## build everything

.PHONY: backend
backend: generate-backend $(EXECUTABLE) ## build backend files

.PHONY: generate-backend
generate-backend: generate-go

generate-go: $(TAGS_PREREQ)
	@echo "Running go generate..."
	@CC= GOOS= GOARCH= CGO_ENABLED=0 $(GO) generate -tags '$(TAGS)' ./...

$(EXECUTABLE): $(GO_SOURCES) $(TAGS_PREREQ)
	CGO_CFLAGS="$(CGO_CFLAGS)" $(GO) build $(GOFLAGS) $(EXTRA_GOFLAGS) -tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)' -o $@

.PHONY: build-style
build-style: deps-frontend
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

.PHONY: docker
docker:
	docker build -t $(DOCKER_REF) .
