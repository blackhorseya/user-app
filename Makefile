APP_NAME=user-app
VERSION=latest
PROJECT_ID=sean-side
NS=side
DEPLOY_TO=uat
REGISTRY=gcr.io
IMAGE_NAME=$(REGISTRY)/$(PROJECT_ID)/$(APP_NAME)
HELM_REPO_NAME=blackhorseya
CHART_NAME=user-app

check_defined = $(if $(value $1),,$(error Undefined $1))

.PHONY: help # Generate list of targets with descriptions
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | expand -t20

.PHONY: clean # remove data
clean:
	@rm -rf bin charts coverage.txt profile.out
	@echo Successfully removed artifacts

.PHONY: lint # execute golint
lint:
	@golint ./...

.PHONY: report
report:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/user-app'

.PHONY: test-unit # execute unit test
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: build-image # build docker image with APP_NAME and VERSION
build-image:
	$(call check_defined,VERSION)
	$(call check_defined,GITHUB_TOKEN)
	@docker build -t $(IMAGE_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg APP_NAME=$(APP_NAME) \
	--build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
	--platform linux/amd64 \
	--pull --cache-from=$(IMAGE_NAME) \
	-f Dockerfile .

.PHONY: list-images # list all images with APP_NAME
list-images:
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images # remove all images with APP_NAME
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: push-image # push image to registry
push-image:
	$(call check_defined,VERSION)
	@docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest
	@docker push $(IMAGE_NAME):$(VERSION)
	@docker push $(IMAGE_NAME):latest

.PHONY: gen # generate all generate commands
gen: gen-wire gen-swagger gen-pb gen-mocks

.PHONY: gen-wire # generate wire code
gen-wire:
	@wire gen ./...

.PHONY: gen-swagger # generate swagger spec
gen-swagger:
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal --parseDependency --parseDepth 1 -o api/docs

.PHONY: gen-pb # generate protobuf messages and services
gen-pb:
	@protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pb/*.proto
	@echo Successfully generated proto

.PHONY: gen-mocks # generate mocks code via mockery
gen-mocks:
	@go generate -x ./...

.PHONY: deploy # deploy the application via helm 3
deploy:
	$(call check_defined,DEPLOY_TO)
	$(call check_defined,VERSION)
	@helm -n $(NS) upgrade --install $(DEPLOY_TO)-$(RELEASE_NAME) \
	$(HELM_REPO_NAME)/$(CHART_NAME) \
	--set "image.tag=$(VERSION)" -f ./deployments/values/$(DEPLOY_TO)/values.yaml

.PHONY: up-local # docker-compose up local
up-local:
	@docker-compose --file ./deployments/docker-compose.yaml --project-name $(APP_NAME) up -d

.PHONY: down-local # docker-compose down local
down-local:
	@docker-compose --file ./deployments/docker-compose.yaml --project-name $(APP_NAME) down -v
