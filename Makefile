CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

IMG_NAME=${APP}
REGISTRY=${REGISTRY}
TAG=latest
ENV_TAG=latest

# Including
include .build_info

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

run:
	go run cmd/main.go

lint:
	golangci-lint run -E megacheck .

format:
	gofmt -s -w .

imports:
	goimports -w .

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${TAG}