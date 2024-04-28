unit-test:
	go test ./... --short

integration-test:
	docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
	go test ./...
	docker container stop redis-stack
	docker container prune -f
