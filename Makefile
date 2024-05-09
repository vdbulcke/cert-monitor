

.PHONY:  test scan build build-snapshot release-skip-publish release-snapshot gen-doc lint

.PHONY: prepare-release 
prepare-release: lint scan release-skip-publish

test: 
	 go test -v  ./...	
scan: 
	trivy fs .
	semgrep scan --disable-version-check --config "p/owasp-top-ten" --metrics=off ./

build: 
	goreleaser build --clean


build-snapshot: 
	goreleaser build --clean --snapshot --single-target



release-skip-publish: 
	goreleaser release --clean --skip-publish  --skip-sign

release-snapshot: 
	goreleaser release --clean --skip-publish --snapshot --skip-sign

gen-doc: 
	goreleaser build --clean --snapshot
	./dist/cert-monitor_linux_amd64/cert-monitor documentation --dir ./doc

.PHONY: changelog
changelog: 
	git-chglog -o CHANGELOG.md 

.PHONY: dev-doc-site
dev-doc-site: 
	podman  run --rm -it -p 8000:8000 -v ${PWD}/www:/docs:z squidfunk/mkdocs-material 

lint: 
	golangci-lint run ./... 
