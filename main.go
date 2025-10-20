package main;

import (
	"fmt"
	"log"
	cfg "github.com/lolmerkat/delta/internal/config"
)

func main() {
	fmt.Printf("Config Dir: %s\n", cfg.ConfigDir())
	fmt.Printf("Config Path: %s\n\n", cfg.ConfigPath())

	config, _ := cfg.Load()
	config.Repositories = append(config.Repositories, cfg.Repository {
		Name: "windows-config",
		Remote: "https://github.com/lolmerkat/windots",
		LocalPath: "C:\\Users\\meyer\\.dotfiles\\windows-config\\.config\\",
	})

	tomlString, _ := config.TomlString()
	err := config.Write()
	if err != nil {
		log.Fatalf("Error writing config: %v", err)
	}
	fmt.Printf("Toml String: \n%s\n", tomlString)
}
