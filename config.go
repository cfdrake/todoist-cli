package main

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/vaughan0/go-ini"
)

func die(format string, a ...interface{}) {
	formatStr := fmt.Sprintf("ERROR: %s\n", format)
	fmt.Fprintf(os.Stderr, formatStr, a...)
	os.Exit(1)
}

// Config file based config provider.
type FileConfig struct {
	file ini.File
}

// Creates a new config instance, assuming XDG style directory setup.
func Config() FileConfig {
	user, _ := user.Current()
	configFilePath := ".config/todoist-cli/config.ini"
	path := path.Join(user.HomeDir, configFilePath)

	file, err := ini.LoadFile(path)
	if err != nil {
		die("Expected to find config file at '%s'...", path)
	}

	return FileConfig{file}
}

// Adhere to Configurer.
func (f FileConfig) UserToken() string {
	token, ok := f.file.Get("auth", "token")
	if !ok {
		die(`Missing token under [auth] section in config!

Ensure your configuration file looks like the following:

    [auth]
    token = <your token here>`)
	}
	return token
}
