ENV_FILE = ./config/.env
ENV = $(shell cat $(ENV_FILE))

.PHONY: migrate
migrate:
	cat ./config/.env
	mysqldef -uroot -p${MYSQL_ROOT_PASSWORD} -P$(DB_PORT) $(MYSQL_DATABASE) < schema.sql

.PHONY: migrate
migrate-dry-run:
	$(ENV)
	mysqldef -uroot -p$(MYSQL_ROOT_PASSWORD) -P$(DB_PORT) $(MYSQL_DATABASE) --dry-run < schema.sql

.PHONY: test
test:
	$(ENV)
	go test -v ./... -count=1

.PHONY: fmt-terraform
fmt-terraform:
	terraform fmt -recursive
