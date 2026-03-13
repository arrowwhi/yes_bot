IMAGE_NAME ?= yes-bot
IMAGE_TAG ?= latest
CONTAINER_NAME ?= yes-bot

.PHONY: docker-build docker-run docker-stop docker-rebuild-run

docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

docker-run: docker-build
	docker run --rm \
		--name $(CONTAINER_NAME) \
		--env-file .env \
		$(IMAGE_NAME):$(IMAGE_TAG)

docker-stop:
	-docker stop $(CONTAINER_NAME)

docker-rebuild-run: docker-stop docker-build docker-run

