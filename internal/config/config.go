package config;

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Repositories	[]Repository
}

func Initialize() (err error) {
	filePath := ConfigPath()
	_, err = os.Stat(filePath)

	if os.IsNotExist(err) {
		defaultConfig, err := Default()
		if err != nil { return err }

		err = defaultConfig.Write()
		if err != nil { return err }
	}

	return err
}

func Load() (config Config, err error) {
	err = Initialize()
	if err != nil { return Empty(), err }

	filePath := ConfigPath()
	fileBytes, err := os.ReadFile(filePath)
	if err != nil { return Empty(), err }

	err = toml.Unmarshal(fileBytes, config)
	return config, err
}

func (c Config) Toml() (bytes []byte, err error) {
	bytes, err = toml.Marshal(c)
	return bytes, err
}

func (c Config) TomlString() (tomlString string, err error) {
	bytes, err := c.Toml()
	if err != nil { return "", nil }

	return string(bytes), nil
}

func (c Config) Write() (err error) {
	filePath := ConfigPath()
	tomlBytes, err := c.Toml()
	if err != nil { return err }

	err = os.WriteFile(filePath, tomlBytes, 0755)
	return err
}
