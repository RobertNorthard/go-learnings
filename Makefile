
build:
	go get ./...
	GOOS=linux GOARCH=amd64 go build