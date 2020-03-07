package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

type tweet struct {
	Name     string
	Username string
	Message  string
}

var site = "https://twitter.com/Todd_McLeod/status/1217694526569009153"

func main() {

	c := colly.NewCollector()

	messages := []tweet{}

	c.OnHTML(".tweet", func(e *colly.HTMLElement) {

		messages = append(messages, tweet{
			Name:     e.ChildText(".account-group .fullname"),
			Username: e.ChildText(".account-group .username"),
			Message:  e.ChildText(".tweet-text"),
		})
	})

	err := c.Visit(site)
	if err != nil {
		panic(err)
	}

	c.Wait()

	bs, err := json.MarshalIndent(messages, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("Number of tweets: ", len(messages))

}
