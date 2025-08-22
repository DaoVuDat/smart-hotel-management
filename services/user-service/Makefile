ifneq (,$(wildcard ./.env.dev))
    include .env.dev
    export
endif


.PHONY: migrationUP
migrationUp:
	goose -dir ./migrations postgres $(DSN) up


.PHONY: migrationZero
migrationZero:
	goose -dir ./migrations postgres $(DSN) down-to 0