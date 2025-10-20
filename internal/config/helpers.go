package config

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
)

func ConfigDir() string {
	configHome := xdg.ConfigHome
	return fmt.Sprintf("%s/delta", configHome)
}

func CreateConfigDir() (err error) {
	configDir := ConfigDir()
	err = os.Mkdir(configDir, os.ModePerm)
	if err != nil { return err }

	return err
}

func ConfigPath() string {
	configDir := ConfigDir()
	return fmt.Sprintf("%s/config.toml", configDir)
}

func Empty() (config Config) {
	return Config{}
}

func Default() (config Config, err error) {
	return Config {
		Repositories: []Repository {
			{
				Name: "config",
				Remote: "origin",
				LocalPath: xdg.ConfigHome,
			},
		},
	}, nil
}

