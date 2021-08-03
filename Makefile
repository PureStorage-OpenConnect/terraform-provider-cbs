DEV_PKGDIR := $(HOME)/.terraform.d/plugins/terraform.purestorage.com/flasharray/cbs/
DEV_GOBIN := $(DEV_PKGDIR)/99.99/linux_amd64/
SRC_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
SHELL=/bin/bash -eEuo pipefail # Set sane shell options
MAKEFLAGS += -j4
TMPBIN=/tmp/bin
TEST_GPG_FINGERPRINT=47F12A8402BBC3906D7C3FAB78C24746203D8F4E
export GNUPGHOME := /tmp/gnupg
export PATH := $(TMPBIN):$(PATH)

default: build

setup-basic:
	@mkdir -p .build-logs/


setup-goreleaser:
	@curl -sfLO https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh
	@bash goreleaser.sh -b $(TMPBIN) &> /dev/null
	@rm goreleaser.sh
# Replace with just this after go 1.16:
# @GOBIN=$(TMPBIN) go install github.com/goreleaser/goreleaser@latest

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
	go build

testacc:
	TF_ACC=1 go test ./cbs -v -timeout 120m

install-dev-mock:
	GOBIN=$(DEV_GOBIN) go install --tags mock

install-dev:
	GOBIN=$(DEV_GOBIN) go install

install-dev-clean:
	@rm -rvf $(DEV_PKGDIR)


# This part just helps reduce noise, we

test-everything-with-mocks: setup-basic
	@set -a; . testing/mock.env; TF_LOG=TRACE TF_LOG_PATH="$(SRC_DIR)/.build-logs/terraform" go test --tags mock ./... -v -timeout 120m \
	2>&1 | tee .build-logs/acc-mock  | grep -vE '^=== RUN|^--- PASS:|^--- SKIP:|^=== PAUSE|^=== CONT|^PASS|^ok|\[no test files\]'

# The redirections and tee/grep stuff above is to help reduce console noise, we filter out all of the nominal messages, so its easier to see any errors
# the full unfiltered log is in .build-logs/acc-mock

test-vet:
	@go vet ./cbs
	@go vet -tags mock ./cbs

# Tests that should run on each pull request
test-pull-request: test-vet test-everything-with-mocks test-goreleaser-release test-goreleaser-check

reset-mock-db:
	@rm -v .build/test.db

tidy:
	@go get -u
	@go mod tidy -v
	@go fmt ./cbs
	@go fix ./cbs
	@go clean ./cbs
	@go clean --tags mock ./cbs
