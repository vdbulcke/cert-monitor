GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
LD_FLAGS ?= -X main.GitCommit=${GIT_COMMIT}

all: build

test: 
	cd certmonitor/ && go test -run ''
	
build: 
	mkdir -p bin/
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -ldflags "${LD_FLAGS}" -o bin/cert-monitor_linux_amd64 main.go 

build-zip: 
	mkdir -p bin/
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -ldflags "${LD_FLAGS}" -o bin/cert-monitor_linux_amd64 main.go 
	mkdir -p releases/
	zip releases/cert-monitor_linux_amd64.zip bin/cert-monitor_linux_amd64

build-macos-universal: build-macos-amd64 build-macos-arm64
	lipo -create -output bin/cert-monitor_darwin_universal bin/cert-monitor_darwin_amd64 bin/cert-monitor_darwin_arm64

build-macos-amd64:
	env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0  go build -ldflags "${LD_FLAGS}" -o bin/cert-monitor_darwin_amd64 main.go 

build-macos-arm64:
	env GOOS=darwin GOARCH=arm64 CGO_ENABLED=0  go build -ldflags "${LD_FLAGS}" -o bin/cert-monitor_darwin_arm64 main.go 

docker-build: 
	docker build -t cert-monitor .
