Version := $(shell git describe --tags --dirty)
GitCommit := $(shell git rev-parse HEAD)
LDFLAGS := "-s -w -X main.Version=$(Version) -X main.GitCommit=$(GitCommit)"

SHELL = /bin/bash

.PHONY: dist distgo distui clean cleantmp

dist: clean distgo distui
	find dist/tmp -maxdepth 1 -type f -execdir tar -czf ../{}.tar.gz ui {} \;
	rm -rf dist/tmp/

distgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o dist/tmp/filetinder cmd/filetinder.go
	CGO_ENABLED=0 GOOS=darwin go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o dist/tmp/filetinder-darwin cmd/filetinder.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o dist/tmp/filetinder-armhf cmd/filetinder.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o dist/tmp/filetinder-arm64 cmd/filetinder.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o dist/tmp/filetinder.exe cmd/filetinder.go

distui:
	pushd ui && yarn build && popd && cp -ar ui/public dist/tmp/ui

clean:
	rm -rf dist/*