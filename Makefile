GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

dev:
	go run main.go

test:
	go test -cover ./...

build:
	go build main.go

build-service:
	docker pull mongo

	docker run -d --name service-sucellus \
		-p 27017:27017 \
		-e PORT=4000 \
		--restart=always \
		.

	docker stop mongo-sucellus

start-service:
	docker start service-sucellus


stop-service:
	docker stop service-sucellus

build-mongo:
	docker pull mongo

	docker run -d --name mongo-sucellus \
		-p 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME=root \
		-e MONGO_INITDB_ROOT_PASSWORD=root \
		-e PUID=1000 \
		-e PGID=1000 \
		--restart=always \
		mongo

	docker stop mongo-sucellus

start-mongo:
	docker start mongo-sucellus


stop-mongo:
	docker stop mongo-sucellus