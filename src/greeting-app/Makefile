.PHONY: image publish run run-docker

REGISTRY_GROUP = jeanmorais
NAME           = greeting-app
VERSION        ?= latest
IMAGE          = $(REGISTRY_GROUP)/$(NAME):$(VERSION)

image: ## Build image
	docker build -t $(IMAGE) .

publish: ## Publish image
	docker push $(IMAGE)

run: ## Run locally
	go run *.go

run-docker: ## Run docker container
	docker run --rm -d --name $(NAME) -e APP_LANGUAGE=$(APP_LANGUAGE) -d -p 8080:8080 $(REGISTRY_GROUP)/$(NAME):$(VERSION)
