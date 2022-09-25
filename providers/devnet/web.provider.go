package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

const devnet_web_name = "developpez.net - web"

func DevnetWebProvider(max int, d *database.Database) []providers.Content {
	return devnetProvider("https://web.developpez.com/index/rss", devnet_web_name, max, d)
}