ENV_FILE = ./config/.env
ENV = $(shell cat $(ENV_FILE))

.PHONY: test
test:
	$(ENV) go test -v ./... -count=1

.PHONY: fmt-terraform
fmt-terraform:
	terraform fmt -recursive
