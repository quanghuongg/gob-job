package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config") // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatalf("Error reading config file: %+v", err)
	}

	viper.WatchConfig() // Watch the config file for changes
}

// GetString is a wrapper for viper
func GetString(key string) string {
	return viper.GetString(key)
}
func GetInt(key string) int {
	return viper.GetInt(key)
}