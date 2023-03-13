package main

import (
	"newsbot/bot"
	"newsbot/config"
	"newsbot/database"
	"newsbot/providers"
	"newsbot/providers/artisandev"
	"newsbot/providers/devnet"
	"newsbot/providers/hackernews"
)

func run() {
	// Create a new database
	db := database.NewDatabase("database/db.json")
	
	// Init database
	db.Init()
	
	bot := bot.Bot{
		Token: config.GetToken(),
		Every: config.GetCronString(),
		ChannelID: config.GetChannelID(),
		Provider: providers.Provider{ 
			Max: 4, 
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
