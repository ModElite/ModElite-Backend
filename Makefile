dev:
	swag init & go run .
migrate:
	go run internal/cmd/migration.go
lint:
	go vet
	golangci-lint run
swag:
	swag init