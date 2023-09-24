define check_var
  @if [ -z "$(1)" ]; then \
    echo "$(2) is not set"; \
    exit 1; \
  fi
endef

BUN_CMD = go run cmd/bun/main.go

.PHONY: db-create-go
db-create-go:
	$(call check_var,$(NAME),NAME)
	$(BUN_CMD) db create_go $(NAME)

.PHONY: db-migrate
db-migrate:
	$(BUN_CMD) db migrate

.PHONY: database-create
database-create:
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	$(call check_var,${APP_ENV},APP_ENV)
	DB_NAME="knowledge-oauth-$${APP_ENV}" ;\
	echo "Start create database: $${DB_NAME}" ;\
	PGPASSWORD=$${DB_PASSWORD} psql -h ${DB_HOST} -p 5432 -U postgres -c "CREATE DATABASE \"$${DB_NAME}\";" ;\