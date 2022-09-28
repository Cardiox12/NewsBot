package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getErrorString(msg string) string {
	return fmt.Sprintf("Error reading %s env variable", msg)
}

func getEnvVar(key string) (string, bool) {
	value, ok := os.LookupEnv(key)

	if !ok {
		return "", ok
	}
	return strings.Trim(value, "\""), true
}

func GetToken() string {
	value, ok := getEnvVar("TOKEN")
	if !ok {
		log.Fatal(getErrorString("TOKEN"))
	}
	return value
}

func GetCronString() string {
	value, ok := getEnvVar("EVERY")
	if !ok {
		log.Fatal(getErrorString("EVERY (cron string)"))
	}
	return value
}

func GetChannelID() string {
	value, ok := getEnvVar("CHANNEL_ID")
	if !ok {
		log.Fatal(getErrorString("CHANNEL_ID"))
	}
	return value
}
