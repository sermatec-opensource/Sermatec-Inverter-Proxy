SHELL:=/bin/bash

current_dir=$(shell pwd)

docker_network=$(shell docker network ls | grep "sermatec" | wc -l)

define create_network
	if [[ $(docker_network) == 0 ]]; then docker network create sermatec; fi
endef

define remove_network
	if [[ $(docker_network) == 1 ]]; then docker network rm sermatec; fi
endef

generate_env_file:
	$(create_network)
	touch $(current_dir)/application.private.env && \
	cat $(current_dir)/application.example.env $(current_dir)/application.private.env > $(current_dir)/application.env

build: generate_env_file
	docker compose -f docker-compose.yml build
help: build
	docker compose -f docker-compose.yml run command-help
run-client-only: build
	docker compose -f docker-compose.yml run runner-client-only
