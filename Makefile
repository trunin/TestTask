export APP_NAME := test-task
export VERSION := latest

build:
	@echo "Building app in Docker."
	@docker build --no-cache -f ./Dockerfile -t ${APP_NAME}:${VERSION} --build-arg APP_NAME=${APP_NAME} --build-arg VERSION=${VERSION} .
