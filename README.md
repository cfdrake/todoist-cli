# todoist-cli

A simple commandline interface for managing [Todoist](http://todoist.com) projects and items.

## Installation

During early development, you may install this software only using `go get`:

    $ go get github.com/cfdrake/todoist-cli

Later, I'd like to distribute this using Homebrew.

## Setup

`todoist-cli` requires an API key to interact on your behalf.

1. First, log into [todoist.com](http://todoist.com).
2. Under the gear icon at the upper right, select "Todoist Settings".
3. Select the "Account" tab.
4. Copy the value next to "API token".
5. Create a `~/.config/todoist-cli/config.ini` file on your filesystem.
6. Format the contents of the file to look like below...

```
[auth]
token = <paste your token here>
```

## Usage

Run the `todoist-cli` command without any parameters to receive help.

To get help for a subcommand, try the following, for example: `todoist-cli projects help`.
