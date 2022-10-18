run:
	wire gen ./internal/app && go run ./cmd/app/main.go -f /resource/app.yml

default:
	go buind -o app ./cmd/app/main.go
	chmod +x app

linux:
	GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/main.go

swag:
	sh ./scripts/swagger.sh

tidy:
	go mod tidy

docker-build:
	docker-compose up -d

docker-rebuild:
	docker-compose up --build -d
