package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

const devnet_java_name = "developpez.net - java"

func DevnetJavaProvider(max int, d *database.Database) []providers.Content {
    return devnetProvider("https://java.developpez.com/index/rss", devnet_java_name, max, d)
}