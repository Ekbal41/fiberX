project_name = fiber
image_name = gofiber:latest

run-local:
	go run app.go

tidy:
	go mod tidy

clean-packages:
	go clean -modcache

build:
	docker build -t $(image_name) .

build-no-cache:
	docker build --no-cache -t $(image_name) .

up:
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./app

up-prod:
	make stop
	make rmv
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./app -prod

up-del:
	make stop
	make rmv
	
rmv:
	docker rm $(project_name)

shell:
	docker exec -it $(project_name) /bin/sh

stop:
	docker stop $(project_name)

start:
	docker start $(project_name)



