# Helpers
IS_DARWIN := $(filter Darwin,$(shell uname -s))

define set_env
	sed $(if $(IS_DARWIN),-i "",-i) -e "s/^#*\($(1)=\).*/$(if $(2),,#)\1$(2)/" .env
endef

EXEC := docker compose exec main

# Environment recipes
.PHONY: default
default: init up

.PHONY: init
init:
	test -f .env || cp .env.example .env
	$(call set_env,USER_ID,$(shell id -u))
	mkdir -p ~/go/pkg

.PHONY: up
up:
	DOCKER_BUILDKIT=1 docker compose up -d --build

.PHONY: down
down:
	docker compose down

.PHONY: shell
shell:
	$(EXEC) zsh

# Project recipes
.PHONY: deps
deps:
	$(EXEC) go mod download
	$(EXEC) go run cmd/main.go database:migrate

.PHONY: run
run:
	$(EXEC) go run cmd/main.go

.PHONY: debug
debug:
	$(EXEC) dlv debug --listen=:8001 --headless --api-version=2 cmd/main.go

.PHONY: lint
lint:
	$(EXEC) go vet ./...

.PHONY: test
test:
	$(EXEC) go test -v ./...

.PHONY: describe-schema
describe-schema:
	$(EXEC) go run -mod=mod entgo.io/ent/cmd/ent describe ./internal/data/schema
