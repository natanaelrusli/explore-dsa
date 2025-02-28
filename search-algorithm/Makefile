# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
BINARY_NAME=search-algorithm

# Test coverage output files
COVERAGE_OUT=coverage.out
COVERAGE_HTML=coverage.html

# Run all tests
test:
	$(GOTEST) -v ./...

# Run benchmarks
benchmark:
	$(GOTEST) -bench=. -benchmem ./...

# Generate test coverage
coverage:
	$(GOTEST) -coverprofile=$(COVERAGE_OUT) ./...
	$(GOCMD) tool cover -func=$(COVERAGE_OUT)

# Generate HTML coverage report
coverage-html: coverage
	$(GOCMD) tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)

# Format all Go files
fmt:
	$(GOFMT) ./...

# Clean up binary files and test coverage reports
clean:
	rm -f $(BINARY_NAME)
	rm -f $(COVERAGE_OUT)
	rm -f $(COVERAGE_HTML)

# Help
help:
	@echo "Available targets:"
	@echo "  test-all       - Run all tests, benchmarks, and generate coverage"
	@echo "  test          - Run all tests"
	@echo "  benchmark     - Run all benchmarks"
	@echo "  coverage      - Generate test coverage report"
	@echo "  coverage-html - Generate HTML test coverage report"
	@echo "  fmt          - Format Go files"
	@echo "  clean        - Remove built files and coverage reports"

.DEFAULT_GOAL := help
