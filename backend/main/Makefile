.PHONY: migration
migration:
	goose -dir=db/migrations postgres "host=${DB_HOST} user=postgres password=${DB_PASSWORD} dbname=knowledge-main sslmode=disable" up

