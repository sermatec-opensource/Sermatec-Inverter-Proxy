build:
	docker compose -f docker-compose.yml build
help: build
	docker compose -f docker-compose.yml run command-help
run-client-only: build
	docker compose -f docker-compose.yml run runner-client-only

