.PHONY: env api
 
env:
	export $(grep -v '^#' ./.env | xargs)

api:
	go run ./cmd/api/main.go
