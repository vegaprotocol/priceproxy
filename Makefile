# Makefile

ifeq ($(CI),)
	# Not in CI
	VERSION := dev-$(USER)
	VERSION_HASH := $(shell git rev-parse HEAD | cut -b1-8)
else
	# In CI
	ifneq ($(RELEASE_VERSION),)
		VERSION := $(RELEASE_VERSION)
	else
		# No tag, so make one
		VERSION := $(shell git describe --tags 2>/dev/null)
	endif
	VERSION_HASH := $(shell echo "$(GITHUB_SHA)" | cut -b1-8)
endif

GO_FLAGS := -ldflags "-X main.Version=$(VERSION) -X main.VersionHash=$(VERSION_HASH)"

.PHONY: all
default: deps build test lint vet

.PHONY: coverage
coverage:
	@go test -covermode=count -coverprofile="coverage.txt" ./...
	@go tool cover -func="coverage.txt"

.PHONY: coveragehtml
coveragehtml: coverage
	@go tool cover -html=coverage.txt -o coverage.html

.PHONY: deps
deps:
	@go get -t -d ./...

.PHONY: gosec
gosec:
	@gosec ./...

.PHONY: mocks
mocks:
	@go generate ./...

.PHONY: build
build:
	@mkdir -p build
	@go build $(GO_FLAGS) -o build/priceproxy ./cmd/priceproxy

.PHONY: build-static
build-static:
	@mkdir -p build
	@env CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o build/priceproxy-static ./cmd/priceproxy

.PHONY: install
install:
	@go install $(GO_FLAGS) ./cmd/priceproxy

.PHONY: release-ubuntu-latest
release-ubuntu-latest:
	@mkdir -p build
	@env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -v -o build/priceproxy-linux-amd64 $(GO_FLAGS) ./cmd/priceproxy
	@cd build && zip priceproxy-linux-amd64.zip priceproxy-linux-amd64

.PHONY: release-macos-latest
release-macos-latest:
	@mkdir -p build
	@env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -v -o build/priceproxy-darwin-amd64 $(GO_FLAGS) ./cmd/priceproxy
	@cd build && zip priceproxy-darwin-amd64.zip priceproxy-darwin-amd64

.PHONY: release-windows-latest
release-windows-latest:
	@env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -v -o build/priceproxy-amd64.exe $(GO_FLAGS) ./cmd/priceproxy
	@cd build && 7z a -tzip priceproxy-windows-amd64.zip priceproxy-amd64.exe

.PHONY: lint
lint:
	@go install golang.org/x/lint/golint
	@outputfile="$$(mktemp)" ; \
	go list ./... | xargs golint 2>&1 | \
		sed -e "s#^$$GOPATH/src/##" | tee "$$outputfile" ; \
	lines="$$(wc -l <"$$outputfile")" ; \
	rm -f "$$outputfile" ; \
	exit "$$lines"

.PHONY: race
race: ## Run data race detector
	@env CGO_ENABLED=1 go test -race ./...

.PHONY: retest
retest: ## Force re-run of all tests
	@go test -count=1 ./...

.PHONY: staticcheck
staticcheck: ## Run statick analysis checks
	@staticcheck ./...

.PHONY: test
test: ## Run tests
	@go test ./...

.PHONY: vet
vet: deps
	@go vet ./...

.PHONY: clean
clean:
	@rm -f ./cmd/priceproxy/priceproxy
