ENV := $(PWD)/.env

# Environment variables for project
include $(ENV)

migrate:
	docker compose --profile tools run migrate
migrate-rollback:
	docker compose --profile tools run migrate-rollback

migrate-both:
	docker compose --profile tools run migrate
	docker compose --profile tools run migrate-db-test

migrate-rollback-both:
	docker compose --profile tools run migrate-rollback-db-test
	docker compose --profile tools run migrate-rollback