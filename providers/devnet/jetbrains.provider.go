package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

const devnet_jetbrains_name = "developpez.net - jetbrains"

func DevnetJetbrainsProvider(max int, d *database.Database) []providers.Content {
	return devnetProvider("https://jetbrains.developpez.com/index/rss", devnet_jetbrains_name, max, d)
}