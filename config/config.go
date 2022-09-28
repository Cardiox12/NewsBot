package config

import (
	"fmt"
	"log"
	"os"
)

func getErrorString(msg string) string {
	return fmt.Sprintf("Error reading %s env variable", msg)
}

func GetToken() string {
	value, ok := os.LookupEnv("TOKEN")
	if !ok {
		log.Fatal(getErrorString("TOKEN"))
	}
	return value
}

func GetCronString() string {
	value, ok := os.LookupEnv("EVERY")
	if !ok {
		log.Fatal(getErrorString("EVERY (cron string)"))
	}
	return value
}

func GetChannelID() string {
	value, ok := os.LookupEnv("CHANNEL_ID")
	if !ok {
		log.Fatal(getErrorString("CHANNEL_ID"))
	}
	return value
}
