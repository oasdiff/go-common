lint:
	golangci-lint run

test:
	go test -count=1 ./...
