package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

func DevnetJavaProvider(max int, d *database.Database) []providers.Content {
    return devnetProvider("java", max, d)
}