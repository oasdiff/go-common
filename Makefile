lint:
	golangci-lint run

# -cover coverage.
# -shuffle=on runs the tests in a random order.
# -race activates the data race detector.
# -vet=all runs go vet to identify significant problems. If go vet finds any problems, go test reports those and does not run the test binary.
# -failfast stops test execution when a given unit test fails. It allows tests executed in parallel to finish.
.PHONY: test
test:  ## Go recompile and test with coverage
	go test ./... -cover -shuffle=on -race -vet=all -failfast
