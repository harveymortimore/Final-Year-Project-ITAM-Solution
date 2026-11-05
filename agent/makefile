build:
	go build -o bin/agent ./cmd/agent

build-all:
	GOOS=linux   GOARCH=amd64 go build -o bin/agent-linux-amd64 ./cmd/agent
	GOOS=windows GOARCH=amd64 go build -o bin/agent-windows-amd64.exe ./cmd/agent
	GOOS=darwin  GOARCH=arm64 go build -o bin/agent-darwin-arm64 ./cmd/agent

test:
	go test ./...