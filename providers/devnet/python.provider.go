package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

func DevnetPythonProvider(max int, d *database.Database) []providers.Content {
	return devnetProvider("python", max, d)
}