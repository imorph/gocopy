all: clean deps check_all build tests

clean:
	rm -rf ./gocopy

deps:
	go get ./...

check_all: fmt vet lint errcheck golangci-lint

fmt:
	GO111MODULE=on gofmt -l -w -s .
	GO111MODULE=on gofmt -l -w -s .

vet:
	GO111MODULE=on go vet .

lint: install-golint
	golint .

install-golint:
	which golint || GO111MODULE=off go get -u github.com/golang/lint/golint

errcheck: install-errcheck
	errcheck -exclude=errcheck_excludes.txt .

install-errcheck:
	which errcheck || GO111MODULE=off go get -u github.com/kisielk/errcheck

golangci-lint: install-golangci-lint
	golangci-lint run -D errcheck

install-golangci-lint:
	which golangci-lint || GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

build:
	GO111MODULE=on go build .

tests: clean build
	@./test.sh