define check_var
  @if [ -z "$(1)" ]; then \
    echo "$(2) is not set"; \
    exit 1; \
  fi
endef

.PHONY: migration
migration:
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	$(call check_var,${APP_ENV},APP_ENV)
	DB_NAME="knowledge-auth-$${APP_ENV}" ;\
	echo "Start migration to: $${DB_NAME}" ;\
	goose -dir=db/migrations postgres "host=$${DB_HOST} user=postgres password=$${DB_PASSWORD} dbname=$${DB_NAME} sslmode=disable" up

.PHONY: database-create
database-create:
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	$(call check_var,${APP_ENV},APP_ENV)
	DB_NAME="knowledge-auth-$${APP_ENV}" ;\
	echo "Start create database: $${DB_NAME}" ;\
	PGPASSWORD=$${DB_PASSWORD} psql -h ${DB_HOST} -p 5432 -U postgres -c "CREATE DATABASE \"$${DB_NAME}\";" || echo "Database $${DB_NAME} already exists or failed to create (continuing...)" ;\

.PHONY: database-create-develop
database-create-develop:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=develop make database-create

.PHONY: migration-develop
migration-develop:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=develop make migration
