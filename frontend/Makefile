test:
	yarn test

image:
	docker build -t myvuechecker:test -f Dockerfile .

container:
	docker run -d -p 8080:8080 --name myvuechecker myvuechecker:test