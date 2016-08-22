package cli

import (
	"errors"
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
		return nil, errors.New("Could not load config file")
	}

	c := new(config)
	c.PopulateFromIni(file)

	// Check for missing data.
	if c.userToken == "" {
		return nil, errors.New("Could not load authorization token")
	}

	return c, nil
}
