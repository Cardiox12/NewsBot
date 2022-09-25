package devnet

import (
	"newsbot/database"
	"newsbot/providers"
)

const devnet_python_name = "developpez.net - python"

func DevnetPythonProvider(max int, d *database.Database) []providers.Content {
	return devnetProvider("https://python.developpez.com/index/rss", devnet_python_name, max, d)
}