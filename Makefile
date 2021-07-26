include ./config/.env
export

.PHONY: migrate
migrate:
	mysqldef -uroot -p$(MYSQL_ROOT_PASSWORD) -P$(DB_PORT) $(MYSQL_DATABASE) < schema.sql

.PHONY: migrate
migrate-dry-run:
	mysqldef -uroot -p$(MYSQL_ROOT_PASSWORD) -P$(DB_PORT) $(MYSQL_DATABASE) --dry-run < schema.sql

.PHONY: test
test:
	go test -v ./... -count=1

.PHONY: fmt-terraform
fmt-terraform:
	terraform fmt -recursive
