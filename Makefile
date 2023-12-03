BUF_VERSION:=1.1.0
init:
	docker compose down && docker compose up --build
runTestServer:
	docker compose up

resetTestServer:
	docker compose down && docker compose up

watch:
	reflex  -c ./reflex.conf
lint:
	go vet ./...
	go fmt ./...
release: 
	./cicd/release.sh
create-container-repo:
	gcloud artifacts repositories create pp --repository-format=docker --location=us-west1 --description="Docker repository for Property Pros"