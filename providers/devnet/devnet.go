package devnet

import (
	"newsbot/database"
	"newsbot/providers"

	"github.com/gocolly/colly/v2"
)

func devnetProvider(url string, source string, max int, _ *database.Database) []providers.Content {
	c := colly.NewCollector()
	n := 0
	articles := make([]providers.Content, 0)

	c.OnXML("//item", func(e *colly.XMLElement) {
		if n < max {
			article := providers.Content{
				Title: e.ChildText("title"),
				Url: e.ChildText("link"),
				Source: source,
			}
			articles = append(articles, article)
			n++
		}
	})
	c.Visit(url)
	return articles
}