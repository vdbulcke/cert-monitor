GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
LD_FLAGS ?= -X main.GitCommit=${GIT_COMMIT}

all: build

build: 
	mkdir -p bin/
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -ldflags "${LD_FLAGS}" -o bin/cert-monitor_linux_amd64 main.go 


build-macos:
	env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0  go build -ldflags "${LD_FLAGS}" -o bin/cert-monitor_darwin_amd64 main.go 


docker-build: 
	docker build -t cert-monitor .

