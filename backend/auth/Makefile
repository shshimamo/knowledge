.PHONY: migration
migration:
	@if [ -z "$${DB_HOST?}" ]; then \
		echo "DB_HOST is not set"; \
		exit 1; \
	fi
	@if [ -z "$${DB_PASSWORD?}" ]; then \
		echo "DB_PASSWORD is not set"; \
		exit 1; \
	fi
	@if [ -z "$${APP_ENV}" ]; then \
		echo "APP_ENV is not set"; \
		exit 1; \
	fi
	DB_NAME="knowledge-auth-$${APP_ENV}" ;\
	echo "DB_NAME: $${DB_NAME}" ;\
	goose -dir=db/migrations postgres "host=$${DB_HOST} user=postgres password=$${DB_PASSWORD} dbname=$${DB_NAME} sslmode=disable" up

.PHONY: database-create
database-create:
	@if [ -z "$${DB_HOST?}" ]; then \
		echo "DB_HOST is not set"; \
		exit 1; \
	fi
	@if [ -z "$${APP_ENV}" ]; then \
		echo "APP_ENV is not set"; \
		exit 1; \
	fi
	DB_NAME="knowledge-auth-$${APP_ENV}" ;\
	echo "DB_NAME: $${DB_NAME}" ;\
	psql -h ${DB_HOST} -p 5432 -U postgres -c "CREATE DATABASE \"$${DB_NAME}\";" ;\
