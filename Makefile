.DEFAULT_GOAL := build

CPUS ?= $(shell sysctl -n hw.ncpu || echo 1)
#MAKEFLAGS += --jobs=$(CPUS)

.PHONY: tests
tests:
	#go test `go list ./... | grep -v node_modules`

lint:
	golangci-lint -c .golangci.yaml run

build_%:
	rm -f ./build/$*
	env GOOS=linux go build -ldflags="-s -w -X 'main.version=$(VERSION)'" -o build/$* ./cmd/$*/main.go

.PHONY: build
build: build_bot

.PHONY: deploy
deploy:
	$(MAKE) -j${CPUS} build
	echo "Deploying $$VERSION on $$ENV"
	./node_modules/.bin/serverless deploy --stage $(ENV) --version $(VERSION) --verbose

.PHONY: deploy-storage
deploy-storage:
	echo "Deploying storage $$ENV"
	./node_modules/.bin/serverless deploy --config serverless-storage.yaml --stage $(ENV) --verbose

.PHONY: init
init: install-deps
	pip install pre-commit
	pre-commit install

.PHONY: install-deps
install-deps:
	npm install