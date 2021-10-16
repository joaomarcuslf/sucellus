GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

dev:
	go run main.go

test:
	go test -cover ./...

build:
	go build main.go
