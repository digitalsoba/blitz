IMAGE=digitalsoba/blitz
TAG=$(shell git rev-parse --short=9 HEAD)

.PHONY: build-docker build-binary push-docker test

build-docker:
	docker build -t ${IMAGE}:${TAG} .

build-docker-multi:
	docker buildx build --platform linux/amd64,linux/arm64 -t ${IMAGE}:${TAG} .

build-binary:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build github.com/adnanh/webhook
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build app

build-m1-binary:
	go build github.com/adnanh/webhook
	go build app

push-docker:
	docker push ${IMAGE}:${TAG}

dev:
	docker-compose build && docker-compose up -d

exec:
	docker exec -it blitz-client-1 sh

clean:
	docker-compose down

test:
	go test

deploy:
	kustomize build manifests/ | kubectl apply -f-