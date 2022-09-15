package main

import (
	"log"
	"newsbot/bot"
	"newsbot/providers"
	"newsbot/providers/artisandev"
	"newsbot/providers/hackernews"

	"github.com/spf13/viper"
)

func getToken() string {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// Verify if is string
	value, ok := viper.Get("TOKEN").(string)
	if !ok {
		log.Fatal("Error reading TOKEN in .env file")
	}
	return value
}

func main() {
	bot := bot.Bot{
		Token: getToken(),
		Every: 10,
		ChannelID: "1018617259431825469",
		Provider: providers.Provider{ Max: 3 },
	}

	bot.Init()

	bot.RegisterContentProvider(hackernews.HackernewsProvider)
	bot.RegisterContentProvider(artisandev.ArtisandevProvider)
	
	bot.ServeForever()
}
