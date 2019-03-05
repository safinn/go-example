SHELL := /bin/bash

int_test:
	bash -c "docker-compose up -d";
	bash -c "dockerize -wait tcp://127.0.0.1:5432 -timeout 120s"
	bash -c "go test ./... -tags=integration";
	bash -c "docker-compose down";