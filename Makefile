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

mycli-accounts-db: ## Open mycli with accounts-service database
	@if ! lsof -i :3307 >/dev/null; then \
		echo "Starting port-forward 3307 -> 3306..."; \
		kubectl -n staging-ns port-forward svc/accounts-service-db-srv 3307:3306 >/tmp/trips-db-pf.log 2>&1 & \
		sleep 2; \
	else \
		echo "Port-forward already running on 3307"; \
	fi; \
	mycli -u $$(kubectl get secret accounts-db-credentials -n staging-ns -o jsonpath='{.data.user}' | base64 -d) \
	      -p $$(kubectl get secret accounts-db-credentials -n staging-ns -o jsonpath='{.data.password}' | base64 -d) \
	      -h 127.0.0.1 -P 3307 accounts

mycli-trips-db: ## Open mycli with trips-service database
	@if ! lsof -i :3308 >/dev/null; then \
		echo "Starting port-forward 3308 -> 3306..."; \
		kubectl -n staging-ns port-forward svc/trips-service-db-srv 3308:3306 >/tmp/trips-db-pf.log 2>&1 & \
		sleep 2; \
	else \
		echo "Port-forward already running on 3308"; \
	fi; \
	mycli -u $$(kubectl get secret trips-db-credentials -n staging-ns -o jsonpath='{.data.user}' | base64 -d) \
	      -p $$(kubectl get secret trips-db-credentials -n staging-ns -o jsonpath='{.data.password}' | base64 -d) \
	      -h 127.0.0.1 -P 3308 trips
