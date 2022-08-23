run:
	go run main.go

default:
	go buind -o resume main.go
	chmod +x resume

linux:
	GOOS=linux GOARCH=amd64 go build -o resume main.go
