SHELL := /bin/bash

int_test:
	bash -c "docker-compose up -d";
	bash -c "sleep 5"
	bash -c "./scripts/circle-test-postgres.sh";
	bash -c "docker-compose down";