SHELL:=/bin/bash

current_dir=$(shell pwd)

generate_env_file:
	touch $(current_dir)/application.private.env && \
	cat $(current_dir)/application.example.env $(current_dir)/application.private.env > $(current_dir)/application.env

build: generate_env_file
	docker compose -f docker-compose.yml build
help: build
	docker compose -f docker-compose.yml run command-help
run-client-only: build
	docker compose -f docker-compose.yml run runner-client-only

