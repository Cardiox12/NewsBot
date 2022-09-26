package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

func DevnetJetbrainsProvider(max int, d *database.Database) []providers.Content {
	return devnetProvider("jetbrains", max, d)
}