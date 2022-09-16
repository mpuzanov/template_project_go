.SILENT:

.PHONY: help clean build release run lint vet test

APP=project
SOURCE=./cmd/${APP}
GOBASE=$(shell pwd)
RELEASE_DIR=$(GOBASE)/bin
NOW := $(shell date "+%Y-%m-%d %H-%M-%S")
Version=0.0.1

.DEFAULT_GOAL = build
GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)
GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

build: ## Build program
	$(call print-target)
	@go build -v -o ${APP} ${SOURCE}	

lint:  ## Lint the source files
	$(call print-target)
	@gofmt -s -w ${GO_SRC_DIRS}
	@go vet ${GO_SRC_DIRS}
	@golint ${GO_SRC_DIRS}

run: build  ## Run program
	$(call print-target)
	./${APP}
	
docker-build: ## Build docker image
	$(call print-target)
	docker build -t puzanovma/${APP}:${Version} -f ./deploy/Dockerfile .
	docker image prune --force --filter label=stage=intermediate

docker-push: ## Push docker image to registry
	$(call print-target)
	docker push puzanovma/${APP}:${Version}

test: ## go test with race detector and code covarage
	$(call print-target)
# 	@go test -v $(GO_TEST_DIRS)	
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@rm coverage.out

release: ## Build executable
	$(call print-target)
	rm -rf ${RELEASE_DIR}${APP}*	
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X '${APP}/internal/config.Version=$(Version)'" -o ${RELEASE_DIR}/${APP}.exe ${SOURCE}
#GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X '${APP}/internal/config.Version=$(Version)' -X '${APP}/internal/config.Time=$(NOW)'" -o ${APP} ${SOURCE}	
#@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${RELEASE_DIR}/${APP} ${SOURCE}

clean: ## Clean build directory
	rm -f ./bin/${APP}
	rmdir ./bin
	rm -f coverage.*

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
