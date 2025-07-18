define check_var
  @if [ -z "$(1)" ]; then \
    echo "$(2) is not set"; \
    exit 1; \
  fi
endef

.PHONY: test
test:
	go test -v $(shell go list ./... | grep -v ./db)

.PHONY: test-db
test-db:
	cd ./db ;\
	go test -v -test.config=../sqlboiler.test.toml ./db/...

.PHONY: migration
migration:
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	$(call check_var,${APP_ENV},APP_ENV)
	DB_NAME="knowledge-main-$${APP_ENV}" ;\
	echo "Start migration to $${DB_NAME}" ;\
	goose -dir=db/migrations postgres "host=$${DB_HOST} user=postgres password=$${DB_PASSWORD} dbname=$${DB_NAME} sslmode=disable" up

.PHONY: database-create
database-create:
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	$(call check_var,${APP_ENV},APP_ENV)
	DB_NAME="knowledge-main-$${APP_ENV}" ;\
	echo "Start create database: $${DB_NAME}" ;\
	PGPASSWORD=$${DB_PASSWORD} psql -h ${DB_HOST} -p 5432 -U postgres -c "CREATE DATABASE \"$${DB_NAME}\";" ;\

.PHONY: gqlgen-generate
gqlgen-generate:
	go run github.com/99designs/gqlgen generate

.PHONY: database-create-develop
database-create-develop:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=develop make database-create

.PHONY: database-create-test
database-create-test:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=test make database-create

.PHONY: migration-develop
migration-develop:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=develop make migration

.PHONY: migration-test
migration-test:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=test make migration

