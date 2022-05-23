# ==================================================================================== #
# HELPERS
# ==================================================================================== #

# Include variables from the .envrc file
include .envrc

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api -db-dsn=${REALTY_DB_DSN} 

## db/postgres: connect to the databse using posgresql
.PHONY: db/postgres
db/psql:
	psql ${REALTY_DB_DSN}