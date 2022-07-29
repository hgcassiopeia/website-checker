test:
	go test ./...
	
test-bench:
	go test ./... -bench=.

image:
	docker build -t mygochecker:test -f Dockerfile .

container:
	docker run -d -p 3000:3000 --env-file ./local.env --name mygochecker mygochecker:test