ENV_PATH = ./docker/dev.env
include $(ENV_PATH)

POSTGRES_DB_URL = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

DOCKER_COMPOSE_PATH = ./docker/docker-compose.dev.yaml
MIGRATIONS_PATH = ./migrations


# ==============================================================================
# Docker-compose
.PHONY: docker-compose-build
docker-compose-build:
	@docker compose -f $(DOCKER_COMPOSE_PATH) --env-file $(ENV_PATH) build

.PHONY: docker-compose-up
docker-compose-up:
	@docker compose -f $(DOCKER_COMPOSE_PATH) --env-file $(ENV_PATH) up -d

.PHONY: docker-compose-stop
docker-compose-stop:
	@docker compose -f $(DOCKER_COMPOSE_PATH) --env-file $(ENV_PATH) stop


# ==============================================================================
# Go migrate postgresql
.PHONY: postgres-create
postgres-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: 'name' is not set. Usage: make postgres-create name=migration_name"; \
		exit 1; \
	fi
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) $(name)

.PHONY: postgres-up
postgres-up:
	@migrate -database "${POSTGRES_DB_URL}" -path $(MIGRATIONS_PATH) -verbose up

.PHONY: postgres-down
postgres-down:
	@migrate -database "${POSTGRES_DB_URL}" -path $(MIGRATIONS_PATH) -verbose down
