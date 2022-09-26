package main

import (
	"log"
	"newsbot/bot"
	"newsbot/database"
	"newsbot/providers"
	"newsbot/providers/artisandev"
	"newsbot/providers/devnet"
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

func run() {
	// Create a new database
	db := database.NewDatabase("database/db.json")
	
	// Init database
	db.Init()
	
	bot := bot.Bot{
		Token: getToken(),
		Every: "*/1 * * * *",
		ChannelID: "1018617259431825469",
		Provider: providers.Provider{ 
			Max: 2, 
			Database: db, 
		},
	}
	
	// Initialize bot
	bot.Init()

	// Register providers
	bot.RegisterContentProvider(hackernews.HackernewsProvider)
	bot.RegisterContentProvider(artisandev.ArtisandevProvider)
	bot.RegisterContentProvider(devnet.DevnetWebProvider)
	bot.RegisterContentProvider(devnet.DevnetJetbrainsProvider)
	bot.RegisterContentProvider(devnet.DevnetPythonProvider)
	bot.RegisterContentProvider(devnet.DevnetGoProvider)
	bot.RegisterContentProvider(devnet.DevnetJavaProvider)

	// Run bot loop
	bot.ServeForever()
}

func main() {
	run()
}
