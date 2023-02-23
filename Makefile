build:
	#docker build .
	docker compose up -d --remove-orphans


.PHONY = rm_containers
rm_containers:
	docker compose down
	docker rmi blacklistrestapi_api
	docker rmi dpage/pgadmin4
	#docker rmi postgresContainerForPlaylistService

restart:
	docker compose down
	docker rmi blacklistrestapi_api
	docker compose up -d

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/main.go

local-build:
	go run cmd/main.go