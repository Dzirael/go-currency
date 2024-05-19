get-docs:
	go install github.com/swaggo/swag/cmd/swag@latest

docs: get-docs
	swag init --dir cmd/api --parseDependency --output docs

build:
	go build -o bin/restapi cmd/api/main.go

run:
	go run cmd/api/main.go

build-docker: build
	docker build . -t api-rest

run-docker: build-docker
	docker compose up -d --build
