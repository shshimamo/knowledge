.PHONY: migration
migration:
	DB_NAME="knowledge-main-$${APP_ENV}" ;\
	goose -dir=db/migrations postgres "host=$${DB_HOST} user=postgres password=$${DB_PASSWORD} dbname=$${DB_NAME} sslmode=disable" up

