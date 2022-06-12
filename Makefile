APP_NAME=user-app
VERSION=latest
PROJECT_ID=sean-side
NS=side
DEPLOY_TO=uat
REGISTRY=gcr.io
IMAGE_NAME=$(REGISTRY)/$(PROJECT_ID)/$(APP_NAME)
HELM_REPO_NAME=blackhorseya
CHART_NAME=user-app

.PHONY: check-%
check-%: ## check environment variable is exists
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## remove artifacts
	@rm -rf bin charts coverage.txt profile.out
	@echo Successfully removed artifacts

.PHONY: lint
lint: ## run golint
	@golint ./...

.PHONY: report
report: ## update goreportcard
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/user-app'

.PHONY: test-unit
test-unit: ## run unit test
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: build-image
build-image: check-VERSION check-GITHUB_TOKEN check-DEPLOY_TO ## build docker image
	@docker build -t $(IMAGE_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg APP_NAME=$(APP_NAME) \
	--build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
	--build-arg DEPLOY_TO=$(DEPLOY_TO) \
	--platform linux/amd64 \
	--pull --cache-from=$(IMAGE_NAME) \
	-f Dockerfile .

.PHONY: list-images
list-images: ## list all images
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images
prune-images: ## remove all images
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: push-image
push-image: check-VERSION ## publish image
	@docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest
	@docker push $(IMAGE_NAME):$(VERSION)
	@docker push $(IMAGE_NAME):latest

.PHONY: gen
gen: gen-wire gen-swagger gen-pb gen-mocks ## generate code

.PHONY: gen-wire
gen-wire: ## generate wire code
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger: ## generate swagger spec
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal --parseDependency --parseDepth 1 -o api/docs

.PHONY: gen-pb
gen-pb: ## generate protobuf messages and services
	@protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pb/*.proto
	@echo Successfully generated proto

.PHONY: gen-mocks
gen-mocks: ## generate mocks code via mockery
	@go generate -x ./...

.PHONY: deploy
deploy: check-VERSION check-DEPLOY_TO ## deploy application
	@helm -n $(NS) upgrade --install $(DEPLOY_TO)-$(APP_NAME) \
	$(HELM_REPO_NAME)/$(CHART_NAME) \
	--set "image.tag=$(VERSION)" -f ./deployments/values/$(DEPLOY_TO)/values.yaml

.PHONY: up-local
up-local: ## docker-compose up local
	@docker-compose --file ./deployments/docker-compose.yaml --project-name $(APP_NAME) up -d

.PHONY: down-local
down-local: ## docker-compose down local
	@docker-compose --file ./deployments/docker-compose.yaml --project-name $(APP_NAME) down -v
