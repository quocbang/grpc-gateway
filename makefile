.PHONY: run local-db lint

run: 
	go run . --gateway-host=localhost --gateway-port=8080 --grpc-host=localhost --grpc-port=8081 --config-path=./server/config.yaml

local-db:
	docker-compose down
	docker-compose up -d

lint:
	@(hash golangci-lint 2>/dev/null || \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(go env GOPATH)/bin v1.54.2)
	@golangci-lint run