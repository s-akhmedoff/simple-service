CURRENT_DIR=$(shell pwd)
APP_CMD_DIR=${CURRENT_DIR}/cmd
TAG=latest
PROJECT_NAME=microservice
APP=simple_service
SERVICE_NAME=simple
REGISTRY=sadakhmedoff/simple-repo

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

run:
	go run cmd/main.go

swag:
	swag init -g rest/server.go -o api/docs

lint:
	golangci-lint run -E megacheck ./...

format:
	gofmt -s -w .

imports:
	goimports -w .

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}/${SERVICE_NAME}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}/${SERVICE_NAME}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}/${SERVICE_NAME}:${TAG}:${TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}/${SERVICE_NAME}:${TAG}