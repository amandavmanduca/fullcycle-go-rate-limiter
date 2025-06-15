
env:
	cp .env.example .env

run:
	docker compose up -d --build

setup-local-tests:
	docker run -d --name redis-test -p 6380:6379 redis redis-server --requirepass "redis-test"
	@echo "Waiting for Redis to start..."
	@until nc -z localhost 6380; do sleep 0.1; done

teardown-local-tests:
	docker rm -f redis-test

test:
	go test -v ./src/...

test-local: setup-local-tests
	$(MAKE) test
	$(MAKE) teardown-local-tests
