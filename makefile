GOOS ?= $(shell go env GOOS)
BUILD_TIME = $(shell date -u +%Y%m%d.%H%M%S)
VERSION = $(shell cat version)
BINARY = myapp
OUTPUT ?= ${BINARY}$(ext)
$(eval ext := $(if $(filter $(GOOS),windows),.exe))

version="$(cat version)"; \

default: build

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -mod=vendor \
	-a -installsuffix cgo \
	-ldflags "-s -w -extldflags '-static' \
	-X github.com/jalen-qian/GenHugoBlog/pkg/build.GitCommit=$(GIT_COMMIT) \
	-X github.com/jalen-qian/GenHugoBlog/pkg/build.BuildTime=$(BUILD_TIME) \
	-X github.com/jalen-qian/GenHugoBlog/pkg/build.Version=$(VERSION)" \
	-o ${OUTPUT} .


build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor \
	-a -installsuffix cgo \
	-ldflags "-s -w -extldflags '-static' \
	-X github.com/jalen-qian/GenHugoBlog/pkg/build.GitCommit=$(GIT_COMMIT) \
	-X github.com/jalen-qian/GenHugoBlog/pkg/build.BuildTime=$(BUILD_TIME) \
	-X github.com/jalen-qian/GenHugoBlog/pkg/build.Version=$(VERSION)" \
	-o ${OUTPUT} .


build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod=vendor \
	-a -installsuffix cgo \
	-ldflags "-s -w -extldflags '-static' \
	-X gitlab.mypaas.com.cn/starship/clusterd/pkg/build.GitCommit=$(GIT_COMMIT) \
	-X gitlab.mypaas.com.cn/starship/clusterd/pkg/build.BuildTime=$(BUILD_TIME) \
	-X gitlab.mypaas.com.cn/starship/clusterd/pkg/build.Version=$(VERSION)" \
	-o ${OUTPUT}.exe .

clean:
	rm -rf dist/*