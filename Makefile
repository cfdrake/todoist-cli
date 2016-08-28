GO=go
VERSION=0.1.0-SNAPSHOT

default:
	go get
	$(GO) build

get_gox:
	@echo "Fetching build dependencies..."
	$(GO) get github.com/mitchellh/gox

osx: get_gox
	@echo "Building for OS X..."
	$(GOPATH)/bin/gox -osarch darwin/amd64
	mv todoist-cli_darwin_amd64 todoist-cli
	tar -cvzf todoist-cli-$(VERSION)-darwin-amd64.tar.gz todoist-cli

linux: get_gox
	@echo "Building for Linux..."
	$(GOPATH)/bin/gox -osarch linux/amd64
	mv todoist-cli_linux_amd64 todoist-cli
	tar -cvzf todoist-cli-$(VERSION)-linux-amd64.tar.gz todoist-cli

release: osx linux

clean:
	rm -rf todoist-cli-*-{darwin,linux}-amd64.tar.gz todoist-cli
