
# See https://golangci-lint.run/usage/install/
LINTER_VERSION = v1.52.2
LOCAL_DEPS_INSTALL_LOCATION = /usr/local/bin

.PHONY: checks
pr-checks: clean build unit-test lint coverage

.PHONY: clean
clean:
	rm -rf build
	mkdir -p build

.PHONY: deps
deps:
	go env -w "GOPRIVATE=github.com/ildomm/*"
	go mod download

.PHONY: build
build: deps
	cd cmd && go build -o ../build/handler

.PHONY: unit-test
unit-test: deps
	cd cmd && go test -tags=testing -count=1 -v github.com/ildomm/aws_experiment/...

.PHONY: lint-install
lint-install:
	[ -e ${LOCAL_DEPS_INSTALL_LOCATION}/golangci-lint ] || \
	wget -O- -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b ${LOCAL_DEPS_INSTALL_LOCATION} ${LINTER_VERSION}

.PHONY: lint
lint: deps lint-install
	golangci-lint run

.PHONY: coverage
coverage: coverage-core coverage-report

.PHONY: coverage-core
coverage-core: clean deps
	cd cmd && go test -tags=testing -coverprofile=../build/cover.out github.com/ildomm/aws_experiment/...

.PHONY: coverage-report
coverage-report: clean deps coverage-core
	go tool cover -html=build/cover_filtered.out -o build/coverage.html
	echo "** Report available in build/coverage.html **"
