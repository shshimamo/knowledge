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

.PHONY: install-goose
install-goose:
	@which goose > /dev/null || go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: install-sqlboiler
install-sqlboiler:
	@which sqlboiler > /dev/null || go install github.com/aarondl/sqlboiler/v4@latest
	@which sqlboiler-psql > /dev/null || go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-psql@latest

.PHONY: install-golangci-lint
install-golangci-lint:
	@which golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: install-govulncheck
install-govulncheck:
	@which govulncheck > /dev/null || go install golang.org/x/vuln/cmd/govulncheck@latest

.PHONY: migration
migration: install-goose
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
	PGPASSWORD=$${DB_PASSWORD} psql -h ${DB_HOST} -p 5432 -U postgres -c "CREATE DATABASE \"$${DB_NAME}\";" || echo "Database $${DB_NAME} already exists or failed to create (continuing...)" ;\

.PHONY: database-drop
database-drop:
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	$(call check_var,${APP_ENV},APP_ENV)
	DB_NAME="knowledge-main-$${APP_ENV}" ;\
	echo "Start create database: $${DB_NAME}" ;\
	PGPASSWORD=$${DB_PASSWORD} psql -h ${DB_HOST} -p 5432 -U postgres -c "DROP DATABASE \"$${DB_NAME}\";" || echo "failed drop database $${DB_NAME} (continuing...)" ;\

.PHONY: gqlgen-generate
gqlgen-generate:
	go run github.com/99designs/gqlgen generate

.PHONY: sqlboiler-generate
sqlboiler-generate: install-sqlboiler
	sqlboiler psql

.PHONY: database-create-develop
database-create-develop:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=develop make database-create

.PHONY: database-create-test
database-create-test:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=test make database-create

.PHONY: database-drop-develop
database-drop-develop:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=develop make database-drop

.PHONY: migration-develop
migration-develop:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=develop make migration

.PHONY: migration-test
migration-test:
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=test make migration

.PHONY: lint
lint: install-golangci-lint
	golangci-lint run

.PHONY: vulncheck
vulncheck: install-govulncheck
	govulncheck ./...

