binary_path = dist/upprove-mcp-server
main_path = cmd/main.go

build:
	go build -o $(binary_path) $(main_path)

run-server-dev:
	go run $(main_path)

run-server:
	$(binary_path)
