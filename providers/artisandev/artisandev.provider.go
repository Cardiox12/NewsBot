package artisandev

import (
	"newsbot/providers"

	"github.com/gocolly/colly/v2"
)

const artisandev_name = "artisan dev"

func ArtisandevProvider(max int) []providers.Content {
	articles := make([]providers.Content, 0)
	c := colly.NewCollector()
	n := 0

	// TODO: Optimize to stop parsing when max is reached
	// TODO: Store last published date to avoid getting same content again
	c.OnHTML("article", func(e *colly.HTMLElement){
		if n < max {
			article := providers.Content{}
	
			article.Title = e.DOM.Find("a").First().Text()
			article.Url = e.ChildAttr("a", "href")
			article.Source = artisandev_name
			articles = append(articles, article)
			n++
		}
	})
	c.Visit("https://artisandeveloppeur.fr")
	return articles
}