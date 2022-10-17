run:
	go run ./cmd/app/main.go -f /resources/application.yaml

default:
	go buind -o app ./cmd/app/main.go
	chmod +x app

linux:
	GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/main.go

swag:
	sh ./scripts/swagger.sh

tidy:
	go mod tidy
