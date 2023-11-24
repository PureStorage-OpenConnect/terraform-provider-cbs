DEV_PKGDIR := $(HOME)/.terraform.d/plugins/registry.terraform.io/PureStorage-OpenConnect/cbs
DEV_GOBIN := $(DEV_PKGDIR)/99.99/linux_amd64/
export SRC_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
SHELL=/bin/bash -eEuo pipefail # Set sane shell options
MAKEFLAGS += -j4
TMPBIN=/tmp/bin
TEST_GPG_FINGERPRINT=47F12A8402BBC3906D7C3FAB78C24746203D8F4E
export GNUPGHOME := /tmp/gnupg
export PATH := $(TMPBIN):$(PATH)
export PS_HTTP_TRACE_LOGGING := 1

PKG_LIST           := $(shell go list ./... )
# List of packages that are part of the generated SDK
GENERATED_PKG_LIST := $(shell go list ./... | grep 2.4)
FILTERED_PKG_LIST  := $(filter-out $(GENERATED_PKG_LIST), $(PKG_LIST))
LDFLAGS            := -X 'github.com/PureStorage-OpenConnect/terraform-provider-cbs/version.ProviderCommit=$(shell git rev-parse --short HEAD)'

include */Makefile.mk

.DEFAULT_GOAL := build

setup-basic:
	@mkdir -p .build-logs/

setup-goreleaser:
	@curl -sfLO https://github.com/goreleaser/goreleaser/releases/download/v1.9.2/goreleaser_Linux_x86_64.tar.gz
	@mkdir -p $(TMPBIN)
	@tar -C $(TMPBIN) -xf goreleaser_Linux_x86_64.tar.gz
	@rm goreleaser_Linux_x86_64.tar.gz

# CI="" for goreleaser will prevent extra control characters from cluttering
# the logs.  See https://github.com/goreleaser/goreleaser/blob/b43a2e95ec03261a847fdf9239100d9d36adc60c/cmd/root.go#L15
test-goreleaser-release: setup-goreleaser setup-basic
	@mkdir -p $(GNUPGHOME)
	@chmod 0700 $(GNUPGHOME)
	@gpg --batch --delete-secret-keys $(TEST_GPG_FINGERPRINT) &>/dev/null || true
	@gpg --batch --delete-keys $(TEST_GPG_FINGERPRINT) &>/dev/null || true
	@gpg --import < testing/private-key.gpg >> .build-logs/goreleaser-release 2>&1
	@GPG_FINGERPRINT=$(TEST_GPG_FINGERPRINT) CI="" goreleaser release --debug --snapshot --rm-dist >> .build-logs/goreleaser-release 2>&1

test-goreleaser-check: setup-goreleaser setup-basic
	@CI="" goreleaser check >> .build-logs/goreleaser-check 2>&1

build:
	go build -ldflags="$(LDFLAGS)"

testacc:
	TF_ACC=1 go test -ldflags="$(LDFLAGS)" ./cbs -v -timeout 120m

install-dev-mock:
	GOBIN=$(DEV_GOBIN) go install -ldflags="$(LDFLAGS)" --tags mock

install-dev:
	GOBIN=$(DEV_GOBIN) go install -ldflags="$(LDFLAGS)"

install-dev-clean:
	@rm -rvf $(DEV_PKGDIR)


# The redirections and tee/grep stuff above is to help reduce console noise, we filter out all of the nominal messages, so its easier to see any errors
# the full unfiltered log is in .build-logs/acc-mock

vet:
	@echo "Running go vet..."
	@go vet $(FILTERED_PKG_LIST)
	@go vet --tags mock $(FILTERED_PKG_LIST)

fmt:
	@echo "Fixing sources with gofmt..."
	@go fmt $(PKG_LIST)
	@git diff --exit-code || \
		(echo; echo "Unexpected difference in source files after running 'go fmt'. Review the changes and commit them."; exit 1)

depscheck:
	@echo "Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)

# For local development, CI does not run this
golangci-lint:
	@golangci-lint run

pkglist:
	@echo "$(PKG_LIST)" | tr ' ' '\n'

tidy:
	@go get -u
	@go mod tidy -v
	@go fmt ./cbs
	@go fix ./cbs
	@go clean ./cbs
	@go clean --tags mock ./cbs

