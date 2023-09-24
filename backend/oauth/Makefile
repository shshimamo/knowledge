.PHONY: database-create
database-create:
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	$(call check_var,${APP_ENV},APP_ENV)
	DB_NAME="knowledge-oauth-$${APP_ENV}" ;\
	echo "Start create database: $${DB_NAME}" ;\
	PGPASSWORD=$${DB_PASSWORD} psql -h ${DB_HOST} -p 5432 -U postgres -c "CREATE DATABASE \"$${DB_NAME}\";" ;\