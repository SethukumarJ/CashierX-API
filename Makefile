proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go

swag: ## Generate swagger docs
	swag init -g cmd/main.go -o ./cmd/docs

swagger: ##Insatall swagger
	$(GOCMD) install github.com/swaggo/swag/cmd/swag@latest