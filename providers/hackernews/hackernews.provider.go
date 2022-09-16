package hackernews

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"newsbot/database"
	"newsbot/providers"
)

const hackernews_name = "hackernews"

func getStoryUrl(id int) string {
	return fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
}

func getTopstories(n int) []int {
	url := "https://hacker-news.firebaseio.com/v0/topstories.json"
	var topstories []int

	response, err := http.Get(url)
	if err != nil {
		// TODO: Log in a file
		log.Fatalf("%s: error occured while fetching top stories", hackernews_name)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&topstories)
	if err != nil {
		log.Fatalf("%s: error while stories id", hackernews_name)
	}
	return topstories[:n]
}

func HackernewsProvider(max int, _ *database.Database) []providers.Content {
	ids := getTopstories(max)
	stories := make([]providers.Content, 0, max)
	var story providers.Content

	for _, id := range ids {
		response, err := http.Get(getStoryUrl(id))

		if err != nil {
			log.Fatalf("%s: error while fetching story content", hackernews_name)
		}
		err = json.NewDecoder(response.Body).Decode(&story)
		if err != nil {
			log.Fatalf("%s: error while decoding story", hackernews_name)
		}
		story.Source = hackernews_name
		stories = append(stories, story)
	}
	return stories
}
