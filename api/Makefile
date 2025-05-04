ProjectName := Golang-Core/api
FlyayDir := src/infrastructure/repository/sql

start:
	@go run src/main.go

format-code:
	@go fmt ./...

docker-compose-up:
	@docker compose -f docker-compose.yaml up -d

docker-compose-down:
	@docker compose -f docker-compose.yaml down --volumes --remove-orphans

migrate-info:
	@cd $(FlyayDir) && flyway info -user=local -password=local -url=jdbc:postgresql://localhost:5432/local && cd -

migrate:
	@cd $(FlyayDir) && flyway migrate -user=local -password=local -url=jdbc:postgresql://localhost:5432/local && cd -

migrate-clean:
	@cd $(FlyayDir) && flyway clean migrate -user=local -password=local -url=jdbc:postgresql://localhost:5432/local && cd -

integration-test:
	@go test -v -tags=integration ./src/test/...
