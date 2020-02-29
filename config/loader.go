package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	file   = "cfg.toml"
	config = viper.New()
)

func init() {
	// load the config file
	// split the file into filename and fileext
	var fileext string
	{
		components := strings.Split(file, ".")
		fileext = components[1]
	}

	config.SetConfigType(fileext)

	var filepath string
	{
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		filepath = strings.TrimSuffix(path, "/scrapper")
		filepath = fmt.Sprintf("%s/config/%s", filepath, file)
	}

	config.SetConfigFile(filepath)
	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// Get returns the resulting value based on the key from the config
func Get(key string) interface{} {
	return config.Get(key)
}

// GetStringList returns the key's value as a string list
func GetStringList(key string) []string {
	return config.GetStringSlice(key)
}
