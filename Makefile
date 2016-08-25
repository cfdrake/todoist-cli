release:
	env GOOS=darwin GOARCH=amd64 go build -o todoist-cli-osx-amd64
	env GOOS=linux GOARCH=amd64 go build -o todoist-cli-linux-amd64
	tar -cvzf todoist-cli-osx-amd64.tar.gz todoist-cli-osx-amd64
	tar -cvzf todoist-cli-linux-amd64.tar.gz todoist-cli-linux-amd64

clean:
	rm -rf todoist-cli-{linux,osx}-amd64
	rm -rf todoist-cli-{linux,osx}-amd64.tar.gz
