DEV_PKGDIR := $(HOME)/.terraform.d/plugins/terraform.purestorage.com/flasharray/cbs/
DEV_GOBIN := $(DEV_PKGDIR)/99.99/linux_amd64/

default: build

build:
	go build

testacc:
	TF_ACC=1 go test ./cbs -v -timeout 120m

install-dev-mock:
	GOBIN=$(DEV_GOBIN) go install --tags mock

install-dev:
	GOBIN=$(DEV_GOBIN) go install

install-dev-clean:
	rm -rvf $(DEV_PKGDIR)

test-acc-mock:
	set -a; . testing/mock.env; go test --tags mock,mock_trace ./cbs -v -timeout 120m

test-vet:
	go vet ./cbs
	go vet -tags mock ./cbs

# Tests that should run on each pull request
test-pull-request: test-vet test-acc-mock

tidy:
	go get -u
	go mod tidy -v
	go fmt ./cbs
	go fix ./cbs
	go clean ./cbs
	go clean --tags mock ./cbs
