package main;

import (
	"fmt"
	"github.com/lolmerkat/delta/internal"
)

func main() {
	fmt.Printf("Config Dir: %s\n", delta.ConfigDir())
	fmt.Printf("Config Path: %s\n\n", delta.ConfigPath())
	config, _ := delta.Default()
	config.Repositories = append(config.Repositories, delta.Repository{
		Name: "windows-config",
		Remote: "https://github.com/lolmerkat/windots",
		LocalPath: "C:\\Users\\meyer\\.dotfiles\\windows-config\\.config\\",
	})
	tomlString, _ := config.TomlString()
	fmt.Printf("Toml String: \n%s\n", tomlString)
}
