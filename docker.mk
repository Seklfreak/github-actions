IMAGE_NAME=$(shell basename $(CURDIR))

.PHONY: docker-build
docker-build: ## Build the top level Dockerfile using the directory or $IMAGE_NAME as the name.
	docker build -t $(IMAGE_NAME) .

.PHONY: docker-tag
docker-tag: ## Tag the docker image using the tag script.
	tag $(IMAGE_NAME) $(DOCKER_REPO)/$(IMAGE_NAME) --no-latest

.PHONY: docker-publish
docker-publish: docker-tag ## Publish the image and tags to a repository.
	docker push $(DOCKER_REPO)/$(IMAGE_NAME)
