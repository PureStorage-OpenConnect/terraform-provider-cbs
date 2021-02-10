default: install

build:
	go build

install:
	go install

testacc:
	TF_ACC=1 go test ./cbs -v -timeout 120m
