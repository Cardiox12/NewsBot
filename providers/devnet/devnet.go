package devnet

import (
	"fmt"
	"newsbot/database"
	"newsbot/providers"

	"github.com/gocolly/colly/v2"
)

func getSource(name string) string {
	return fmt.Sprintf("developpez.net - %s", name)
}

func getUrl(name string) string {
	return fmt.Sprintf("https://%s.developpez.com/index/rss", name)
}

func devnetProvider(name string, max int, _ *database.Database) []providers.Content {
	c := colly.NewCollector()
	n := 0
	articles := make([]providers.Content, 0)
	url := getUrl(name)

	c.OnXML("//item", func(e *colly.XMLElement) {
		if n < max {
			article := providers.Content{
				Title: e.ChildText("title"),
				Url: e.ChildText("link"),
				Source: getSource(name),
			}
			articles = append(articles, article)
			n++
		}
	})
	c.Visit(url)
	return articles
}