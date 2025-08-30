SHELL := /bin/bash

.PHONY: help docker-build-account-service

help: 
	@grep -E '^[a-zA-Z0-9_-]+:.*?##' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

docker-build-account-service:
	docker build -t rvanappeldorn/accounts-service-fast-api ./accounts-service
