build:
	docker-compose build avito-app

run:
	docker-compose up avito-app

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:password@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go

