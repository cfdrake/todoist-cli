package cli

import (
	"os"
	"os/user"
	"path"

	"github.com/vaughan0/go-ini"
)

// Interface for a config provider.
type Configurer interface {
	UserToken() string
}

// Environment variable based config provider.
type EnvironmentConfig struct{}

func NewEnvironmentConfig() EnvironmentConfig {
	return EnvironmentConfig{}
}

func (e EnvironmentConfig) UserToken() string {
	token := os.Getenv("TODOIST_API_TOKEN")
	if token == "" {
		die("Expected to find API token in TODOIST_API_TOKEN variable")
	}
	return token
}

// Config file based config provider.
type FileConfig struct {
	backingFile ini.File
}

func NewFileConfig() FileConfig {
	user, err := user.Current()
	if err != nil {
		die("Unable to retrieve current user...")
	}

	configFilePath := ".config/todoist-cli/config.ini"
	path := path.Join(user.HomeDir, configFilePath)

	file, err := ini.LoadFile(path)
	if err != nil {
		die("Expected to find config file at '%s'...", path)
	}

	return FileConfig{backingFile: file}
}

func (f FileConfig) UserToken() string {
	token, ok := f.backingFile.Get("auth", "token")
	if !ok {
		die(`Missing token under [auth] section in config!

Ensure your configuration file looks like the following:

    [auth]
    token = <your token here>`)
	}
	return token
}
