BUF_VERSION:=1.1.0
init:
	docker compose down && docker compose up --build
runTestServer:
	docker compose up
resetTestServer:
	docker compose up

lint:
	go vet ./...
	go fmt ./...