package cli

import (
  "errors"
  "os/user"
  "path"
)

import "github.com/vaughan0/go-ini"

// Configuration file representation.
type Config struct {
  AuthToken string // Personal token used to authenticate with the REST API.
}

// Loads configuration for the given user.
func loadConfiguration(user *user.User) (*Config, error) {
  // Generate path to configuration file.
  configFilePath := ".config/todoist-cli/config.ini"
  path := path.Join(user.HomeDir, configFilePath)

  // Attempt to load the file into memory.
  file, err := ini.LoadFile(path)
  if err != nil {
    return nil, errors.New("Could not load config file")
  }

  // Fetch the needed data.
  token, ok := file.Get("auth", "token")
  if !ok {
    return nil, errors.New("Could not load authorization token")
  }

  return &Config{AuthToken: token}, nil
}
