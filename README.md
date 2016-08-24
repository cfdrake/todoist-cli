# todoist-cli [![Build Status](https://travis-ci.org/cfdrake/todoist-cli.svg?branch=master)](https://travis-ci.org/cfdrake/todoist-cli)

A simple command line interface for managing [Todoist](http://todoist.com) projects and items.

## Installation

During early development, you may install this software only using `go get`:

    $ go get -u github.com/cfdrake/todoist-cli

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

Feel free to `chmod 0600 ~/.config/todoist-cli/config.ini` so only your user can read it.

## Usage

Run the `todoist-cli` command without any parameters to receive help.

To get help for a subcommand, try the following, for example: `todoist-cli projects help`. This is
additionally a great way to discover helpful command aliases (`p` for `projects`, etc).

The following is a list of currently supported commands:

```
todoist-cli projects             (alias for list)
todoist-cli projects list
todoist-cli projects show <id>
todoist-cli items                (alias for list)
todoist-cli items list
todoist-cli items show <id>
```

## License

See the `LICENSE` file for details.
