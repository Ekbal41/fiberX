project_name = fiber
image_name = fiberx:latest

run-local:
	go run app.go

tidy:
	go mod tidy

clean-mod:
	go clean -modcache

build-image:
	docker build -t $(image_name) .

build-image-no-cache:
	docker build --no-cache -t $(image_name) .

start-container:
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./app

start-container-prod:
	make stop
	make rmv
	docker run -d -p 3000:3000 --name $(project_name) $(image_name) ./app -prod

delete-container:
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



