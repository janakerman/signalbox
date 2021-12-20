COMMIT_SHA := $(shell git rev-parse --short HEAD)
DOCKER_TAG:="janakerman/signalbox:${COMMIT_SHA}"

default: build test

.PHONY: build
build:
	GOOS=linux go build -o bin/signalbox cmd/signalbox/main.go

.PHONY: test
test:
	go test -v -count 1 ./...

.PHONY: docker.build
docker.build: build
	docker build -t $(DOCKER_TAG) .

.PHONY: docker.login
docker.login:
	echo ${DOCKER_TOKEN} | docker login --username ${DOCKER_USERNAME} --password-stdin

.PHONY: docker.publish
docker.publish: docker.login docker.build
	docker push $(DOCKER_TAG)
