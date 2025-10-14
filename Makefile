.PHONY: env api doc
 
env:
	export $(grep -v '^#' ./.env | xargs)

api:
	go run ./cmd/api/main.go

doc:
	swag init -g ./cmd/api/main.go
