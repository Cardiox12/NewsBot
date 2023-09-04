package phpwatch

import (
	"github.com/gocolly/colly/v2"
	"newsbot/database"
	"newsbot/providers"
)

const phpwatch_name = "php-watch"

func PhpWatchProvider(max int, db *database.Database) []providers.Content {
	c := colly.NewCollector()
	n := 0
	articles := make([]providers.Content, 0)
	url := "https://php.watch/feed/articles.xml"
	stop := false
	source := phpwatch_name

	c.OnXML("//entry", func(e *colly.XMLElement) {
		if n < max && !stop {
			article := providers.Content{
				Title:  e.ChildText("title"),
				Url:    e.ChildText("id"),
				Source: source,
			}
			if n == 0 {
				if article.Exists(source, db) {
					stop = true
					return
				}
				db.Set(source, article.Hash())
			}
			articles = append(articles, article)
			n++
		}
	})
	c.Visit(url)
	return articles
}
