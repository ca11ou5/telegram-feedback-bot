.PHONY: build compose
cleandata:
	rm -rf ./data/postgres

build:
	docker build -f ./build/Dockerfile -t telegram-bot ./

compose: build cleandata
	docker compose -f ./deployments/docker-compose.yml -p telegram-bot up --no-deps --force-recreate

pgbash:
	docker exec -it postgres /bin/bash

