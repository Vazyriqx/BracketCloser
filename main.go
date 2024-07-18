package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	token := viper.GetString("token.token")
	if token == "" {
		log.Fatalf("Token not found in config file")
	}
	fmt.Println("Token found:", token)

	start(token, discordgo.IntentsAll)
}

func initConfig() {
	configFile := filepath.Join(filepath.Dir(os.Args[0]), "config.toml")
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
