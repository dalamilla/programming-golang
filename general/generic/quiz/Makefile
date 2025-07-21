.PHONY: all
all: clean deps format lint test build

PHONY: clean
clean:
	@echo "Cleaning project"
	go clean
	rm -f quiz


.PHONY: deps
deps:
	@echo "Getting go dependencies"
	go mod download -x


.PHONY: format
format: format-check

.PHONY: format-check
format-check:
	@echo "Checking formatting"
	gofmt -d .
	test -z $$(gofmt -l ./..)

.PHONY: format-fix
format-fix:
	@echo "Fixing formatting"
	go fmt ./...


.PHONY: lint
lint: lint-vet

.PHONY: lint-vet
lint-vet:
	@echo "Checking vet"
	go vet ./...


.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	@echo "Running unit tests"
	go test -v ./...


.PHONY: build
build:
	@echo "Building application"
	go build ./cmd/quiz


.PHONY: run
run: build
	@echo "Running application"
	./quiz
