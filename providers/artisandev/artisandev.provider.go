package artisandev

import (
	"newsbot/database"
	"newsbot/providers"

	"github.com/gocolly/colly/v2"
)

const artisandev_name = "artisandev"

func ArtisandevProvider(max int, db *database.Database) []providers.Content {
	articles := make([]providers.Content, 0)
	c := colly.NewCollector()
	n := 0
	stop := false

	// TODO: Optimize to stop parsing when max is reached
	c.OnHTML("article", func(e *colly.HTMLElement){
		if n < max && !stop {
			article := providers.Content{}
	
			article.Title = e.DOM.Find("a").First().Text()
			
			if n == 0 {
				if article.Exists(artisandev_name, db) {
					stop = true
					return
				}
				db.Set(artisandev_name, article.Hash())
			}

			article.Url = e.ChildAttr("a", "href")
			article.Source = artisandev_name
			articles = append(articles, article)
			n++
		}
	})
	c.Visit("https://artisandeveloppeur.fr")
	return articles
}