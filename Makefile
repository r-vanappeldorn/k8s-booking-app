SHELL := /bin/bash

.PHONY: help

help: ## Shows overview off all available commands
	@grep -Eh '^[a-zA-Z0-9_.-]+:.*## ' $(MAKEFILE_LIST) \
	| awk 'BEGIN{FS=":.*## "} {printf "%s\t%s\n", $$1, $$2}' \
	| sort -k1,1 \
	| awk 'BEGIN{FS="\t"} {cmd[NR]=$$1; desc[NR]=$$2; if(length($$1)>w) w=length($$1)} END{w+=2; for(i=1;i<=NR;i++) printf "\033[36m%-*s\033[0m %s\n", w, cmd[i], desc[i]}'

docker-build-accounts-service: ## Builds accounts service docker container
	docker build -t rvanappeldorn/accounts-service-fast-api ./accounts-service

docker-build-trips-service: ## Builds trips service docker container
	docker build -t rvanappeldorn/trips-service ./trips-service

migrate-accounts-service: ## Migrates accounts service in staging namespace
	kubectl -n staging-ns apply -f k8s/jobs/accounts-migrate-job.yml

dev: ## Build and run docker compose
	docker compose build && docker compose up

trips-service-local:
		go build ./trips-service/ && ./trips-service/trips-service.com
