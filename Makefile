help:
	@echo "Usage: {options} make [target ...]"
	@echo
	@echo "Commands:"
	@echo "  init         	Initialize env configurations"
	@echo "  install      	Adds missing and remove unused modules"
	@echo "  generate     	Generates models and resolvers' stub from graphql schema"
	@echo "  start        	Start service"
	@echo "  build        	Builds a binary"
	@echo "  test         	Run tests"
	@echo "  docker-build 	Builds the docker image of the backend code"
	@echo "  docker-start 	Starts docker containers"
	@echo "  docker-stop  	Stops docker containers"
	@echo "  docker-down  	Shuts down docker containers"
	@echo
	@echo "  help         Show available commands"
	@echo
	@echo "Examples:"
	@echo "  # Getting started"
	@echo "  make install generate"
	@echo "  make start"
	@echo

init:
	@ echo "Initialize env configurations"
	@ chmod +x ./.local-docker-env.sh && chmod +x ./build-docker-env.sh && ./build-docker-env.sh
	@ echo "Finished initializing env configurations"

install:
	@ echo "Adding missing and remove unused modules"
	@ go mod tidy
	@ echo "Finished adding missing and remove unused modules"

start:
	@ echo "Starting server"
	@ eval ". ./cmd/api.sh"
	@ go run ./cmd/.

build:
	@ echo "Building a binary"
	@ go build ./...
	@ echo "Finished building a binary"

test:
	@ echo "Running tests"
	@ go test -v ./...
	@ echo "Finished tests"

#GIT_HASH = $(shell git log -1 --pretty=%h)
docker-build:
	@# docker build -t edushare/backend -t edushare/backend:$(GIT_HASH) .
	@ export DOCKER_BUILDKIT=1 && export COMPOSE_DOCKER_CLI_BUILD=1 && docker-compose build backend

docker-start:
	@ docker-compose up -d

docker-stop:
	@ docker-compose stop

docker-down:
	@ docker-compose down

