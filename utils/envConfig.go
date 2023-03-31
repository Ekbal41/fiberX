package utils

import (
	"fmt"

	"github.com/spf13/viper"
)
func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // read in environment variables that match
	// ---ops---
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
	panic(fmt.Errorf("fatal error config file: %w", err))
	}
}