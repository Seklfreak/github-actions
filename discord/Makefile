include ../docker.mk
include ../help.mk
include ../go.mk

IMAGE_NAME=github-action-discord

.PHONY: clean
clean: ## Clean up after the build process.

.PHONY: lint
lint: go-lint ## Lint all of the files for this Action.

.PHONY: build
build: docker-build ## Build this Action.

.PHONY: test
test: go-test ## Test the components of this Action.

.PHONY: publish
publish: docker-publish ## Publish this Action.