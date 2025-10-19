package delta;

import (
	"fmt"
	"os/user"

	"github.com/adrg/xdg"
	"github.com/pelletier/go-toml/v2"
)

func ConfigDir() string {
	configHome := xdg.ConfigHome
	return fmt.Sprintf("%s/delta", configHome)
}

func ConfigPath() string {
	configDir := ConfigDir()
	return fmt.Sprintf("%s/config.toml", configDir)
}

type Config struct {
	Repositories	[]Repository
}

func Empty() Config {
	return Config{}
}

func Default() (config Config, err error) {
	user, err := user.Current()
	if err != nil { return Empty(), err }

	defaultRemote := fmt.Sprintf("https://github.com/%s/config", user.Name)
	return Config {
		Repositories: []Repository {
			{
				Name: "config",
				Remote: defaultRemote,
				LocalPath: xdg.ConfigHome,
			},
		},
	}, nil
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

// func Load() Config {
//
// }

func (c Config) Write() {

}
