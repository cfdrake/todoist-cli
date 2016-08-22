package cli

import (
	"errors"
	"fmt"
	"os/user"
	"path"

	"github.com/vaughan0/go-ini"
)

// Configuration file representation.
type config struct {
	userToken string // Personal token used to authenticate with the REST API.
}

// Populate a config instance from an INI.
func (c *config) PopulateFromIni(file ini.File) {
	token, ok := file.Get("auth", "token")
	if ok {
		c.userToken = token
	}
}

// Loads configuration for the given user.
func loadConfiguration(user *user.User) (*config, error) {
	// Generate path to configuration file.
	configFilePath := ".config/todoist-cli/config.ini"
	path := path.Join(user.HomeDir, configFilePath)

	// Load config from file.
	file, err := ini.LoadFile(path)
	if err != nil {
		msg := fmt.Sprintf("Expected to find config file at '%s'", path)
		return nil, errors.New(msg)
	}

	c := new(config)
	c.PopulateFromIni(file)

	// Check for missing data.
	if c.userToken == "" {
		msg := `Missing token under [auth] section in config!

Ensure your configuration file looks like the following:

    [auth]
    token = <your token here>`
		return nil, errors.New(msg)
	}

	return c, nil
}
