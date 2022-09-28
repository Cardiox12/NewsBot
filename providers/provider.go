package providers

import (
	"crypto/md5"
	"fmt"
	"log"
	"newsbot/database"
	"os"
)

var Notice *log.Logger = log.New(os.Stdout, "[NOTICE] ", log.Default().Flags())

type Content struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	Source string `json:"source"`
}

type ContentProvider = func(max int, d *database.Database) []Content

type Provider struct {
	Max              	int
	Database			*database.Database
	ContentProviders	[]ContentProvider
}

func (p *Provider) RegisterContentProvider(cp ContentProvider) {
	p.ContentProviders = append(p.ContentProviders, cp)
}

func (p *Provider) ProvideContents() []Content {
	all := make([]Content, 0)

	for _, contentProvider := range p.ContentProviders {
		all = append(all, contentProvider(p.Max, p.Database)...)
	}
	p.Database.Write()
	return all
}

func (c *Content) String() string {
	return fmt.Sprintf("[%s] - %s\n%s\n", c.Source, c.Title, c.Url)
}

func (c *Content) Hash() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(c.Title)))
}

func (c *Content) Exists(key string, db *database.Database) bool {
	val, ok := db.Get(key)

	if !ok {
		return false
	}
	return val == c.Hash()
}

func (c *Content) LogSent() {
	Notice.Printf("SEND [%s] %s \n", c.Source, c.Title)
}
