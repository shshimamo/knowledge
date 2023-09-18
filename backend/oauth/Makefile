define check_var
  @if [ -z "$(1)" ]; then \
    echo "$(2) is not set"; \
    exit 1; \
  fi
endef

.PHONY: migration-up
migration-up:
	$(call check_var,${APP_ENV},APP_ENV)
ifeq (${APP_ENV},develop)
	sql-migrate up -env=$${APP_ENV} -config=db/dbconfig.yml
else
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${DB_USER},DB_USER)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	sql-migrate up -env=$${APP_ENV} -config=db/dbconfig.yml
endif

.PHONY: migration-down
migration-down:
	$(call check_var,${APP_ENV},APP_ENV)
ifeq (${APP_ENV},develop)
	sql-migrate down -env=${APP_ENV} -config=db/dbconfig.yml
else
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${DB_USER},DB_USER)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	sql-migrate down -env=${APP_ENV} -config=db/dbconfig.yml
endif


.PHONY: migration-new
migration-new:
	$(call check_var,${APP_ENV},APP_ENV)
ifeq (${APP_ENV},develop)
	$(call check_var,${NAME},NAME)
	sql-migrate new -env=$${APP_ENV} -config=db/dbconfig.yml $${NAME}
else
	$(error APP_ENV must be set to 'develop' for migration-new target)
endif

.PHONY: create-database
create-database:
	$(call check_var,${DB_HOST},DB_HOST)
	$(call check_var,${APP_ENV},APP_ENV)
	$(call check_var,${DB_USER},DB_USER)
	$(call check_var,${DB_PASSWORD},DB_PASSWORD)
	DB_NAME="knowledge-oauth-$${APP_ENV}" ;\
    PGPASSWORD=$${DB_PASSWORD} psql -h ${DB_HOST} -p 5432 -U $${DB_USER} -c "CREATE DATABASE \"$${DB_NAME}\";" ;\