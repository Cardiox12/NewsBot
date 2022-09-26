package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

func DevnetWebProvider(max int, d *database.Database) []providers.Content {
	return devnetProvider("web", max, d)
}