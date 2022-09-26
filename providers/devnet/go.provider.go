package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

func DevnetGoProvider(max int, d *database.Database) []providers.Content {
	return devnetProvider("go", max, d)
}