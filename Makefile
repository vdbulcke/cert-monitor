

.PHONY:  test scan build build-snapshot release-skip-publish release-snapshot gen-doc lint


test: 
	cd certmonitor/ && go test -run ''
	
scan: 
	go list -json -deps |  nancy sleuth	


build: 
	goreleaser build --rm-dist


build-snapshot: 
	goreleaser build --rm-dist --snapshot



release-skip-publish: 
	goreleaser release --rm-dist --skip-publish 

release-snapshot: 
	goreleaser release --rm-dist --skip-publish --snapshot

gen-doc: 
	goreleaser build --rm-dist --snapshot
	./dist/cert-monitor_linux_amd64/cert-monitor documentation --dir ./doc

lint: 
	golangci-lint run ./... 
