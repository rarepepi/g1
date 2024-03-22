build:
	go build -o server main.go

run: build
	./server

docker-build:
	docker build -t server .

docker-run: docker-build
	docker run -p 8080:8080 server

generate-docs:
	swag init

watch:
	reflex -s -r '\.go$$' make run