package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/dsa-ferreira/dwiper/internal/files"
)

type Configuration struct {
	Regexes      []string
	KeepDuration float64
}

func ParseConfig() Configuration {
	path := filepath.Join(getUserHomeDir(), ".config/dwiper/config")
	var file string = files.ReadFile(path)

	var conf Configuration
	_, err := toml.Decode(file, &conf)
	if err != nil {
		fmt.Println(err)
	}
	return conf
}

func getUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user:", err)
		os.Exit(4)
	}

	return usr.HomeDir

}
