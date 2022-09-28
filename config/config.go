package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func getErrorString(msg string) string {
	return fmt.Sprintf("Error reading %s in .env file", msg)
}

func InitConfig(path string) {
	viper.SetConfigFile(path)
	viper.ReadInConfig()
}

func GetToken() string {
	value, ok := viper.Get("TOKEN").(string)
	if !ok {
		log.Fatal(getErrorString("TOKEN"))
	}
	return value
}

func GetCronString() string {
	value, ok := viper.Get("EVERY").(string)
	if !ok {
		log.Fatal(getErrorString("EVERY (cron string)"))
	}
	return value
}

func GetChannelID() string {
	value, ok := viper.Get("CHANNEL_ID").(string)
	if !ok {
		log.Fatal(getErrorString("CHANNEL_ID"))
	}
	return value
}
