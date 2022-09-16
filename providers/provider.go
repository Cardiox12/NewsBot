package providers

import (
	"fmt"
	"newsbot/database"
)

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
	return all
}

func (c Content) String() string {
	return fmt.Sprintf("[%s] \n\t%s \n\t(%s)", c.Source, c.Title, c.Url)
}
