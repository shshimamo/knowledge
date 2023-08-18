.PHONY: migration
migration:
	goose -dir=db/migrations postgres "host=${DB_HOST} user=postgres password=${DB_PASSWORD} dbname=knoledge-auth sslmode=disable" up

