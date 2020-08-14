BUILD_VERSION   ?= $(shell cat version.txt || echo "0.1")
BUILD_DATE      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD || echo "0.0.0")

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

fmt:

	@echo gofmt -l
	@OUTPUT=`gofmt -l . 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "gofmt must be run on the following files:"; \
        echo "$$OUTPUT"; \
        exit 1; \
    fi

lint:

	@echo golint ./...
	@OUTPUT=`command -v golint >/dev/null 2>&1 && golint ./... 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "golint errors:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

default: fmt lint ## fmt code

build: clean ## 构建二进制
	@echo "build bin ${BUILD_VERSION} ${BUILD_DATE} ${COMMIT_SHA1}"
	@gox -osarch="darwin/amd64 linux/amd64 windows/amd64" \
        -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
    	-ldflags   "-X 'github.com/ysicing/crtools/cmd.Version=${BUILD_VERSION}' \
                    -X 'github.com/ysicing/crtools/cmd.BuildDate=${BUILD_DATE}' \
                    -X 'github.com/ysicing/crtools/cmd.CommitID=${COMMIT_SHA1}'"

docker: build ## 构建镜像
	@echo "build docker images ${BUILD_VERSION}"
	@docker build -t ysicing/crtools .
	@docker build -t ysicing/crtools:${BUILD_VERSION} .

dpush: docker
	@docker push ysicing/crtools
    @docker push ysicing/crtools:${BUILD_VERSION}

release:  ## github release
	ghr -u ysicing -t $(GITHUB_RELEASE_TOKEN) -replace -recreate --debug ${BUILD_VERSION} dist

pre-release:  ## github pre-release
	ghr -u ysicing -t $(GITHUB_RELEASE_TOKEN) -soft -prerelease --debug ${BUILD_VERSION} dist

clean: ## clean
	rm -rf dist

install: clean ## install
	go install \
		-ldflags   "-X 'github.com/ysicing/crtools/cmd.Version=${BUILD_VERSION}' \
                            -X 'github.com/ysicing/crtools/cmd.BuildDate=${BUILD_DATE}' \
                            -X 'github.com/ysicing/crtools/cmd.CommitID=${COMMIT_SHA1}'"

.PHONY : build release clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.cn
GOSUMDB = sum.golang.google.cn