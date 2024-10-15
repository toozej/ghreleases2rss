package ghreleases2rss

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Get environment variables
func getEnvVars() error {
	if _, err := os.Stat(".env"); err == nil {
		// Initialize Viper from .env file
		viper.SetConfigFile(".env") // Specify the name of your .env file

		// Read the .env file
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading .env file: %s\n", err)
			return err
		}
	}

	// Enable reading environment variables
	viper.AutomaticEnv()

	// Load Miniflux API settings from environment using Viper
	minifluxAPIKey := viper.GetString("MINIFLUX_API_KEY")
	minifluxURL := viper.GetString("MINIFLUX_URL")
	if minifluxAPIKey == "" || minifluxURL == "" {
		return fmt.Errorf("miniflux API key or URL not set in environment variables")
	}

	return nil
}
