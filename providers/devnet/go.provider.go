package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

const devnet_go_name = "developpez.net - go"

func DevnetGoProvider(max int, d *database.Database) []providers.Content {
	return devnetProvider("https://go.developpez.com/index/rss", devnet_go_name, max, d)
}