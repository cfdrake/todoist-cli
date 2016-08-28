GO=go
VERSION=0.1.0-SNAPSHOT
SRC=main.go config.go cli/*.go cli/commands/*.go todoist/*.go

.PHONY: release
.PHONY: clean
.PHONY: package
.PHONY: default

todoist-cli: $(SRC)
	$(GO) build

todoist-cli_darwin_amd64: $(SRC)
	$(GOPATH)/bin/gox -osarch darwin/amd64

todoist-cli_linux_amd64: $(SRC)
	$(GOPATH)/bin/gox -osarch linux/amd64

release: todoist-cli_darwin_amd64 todoist-cli_linux_amd64

package: release
	tar -czf todoist-cli-$(VERSION)_darwin_amd64.tar.gz todoist-cli_darwin_amd64
	tar -czf todoist-cli-$(VERSION)_linux_amd64.tar.gz todoist-cli_linux_amd64

clean:
	rm -rf todoist-cli_{darwin,linux}_amd64 todoist-cli-*_{darwin,linux}_amd64.tar.gz todoist-cli

default: todoist-cli
