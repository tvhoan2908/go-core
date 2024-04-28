run:
	@go run cmd/http/main.go

run-watch:
	@nodemon --exec go run cmd/http/main.go --signal SIGTERM

run-build:
	@./bin/core-api

swag-init:
	@swag init -g cmd/http/main.go

build:
	@go build -o bin/core-api cmd/http/main.go

db-up:
	@go run cmd/migrate/main.go up

db-down:
	@go run cmd/migrate/main.go down

db-status:
	@go run cmd/migrate/main.go status

migration:
	@go run cmd/migrate/main.go create $(filter-out $@,$(MAKECMDGOALS))

seed:
	@go run cmd/seed/main.go