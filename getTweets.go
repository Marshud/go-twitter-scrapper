package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

//
func main() {
	scraper := twitterscraper.New()
	for tweet := range scraper.SearchTweets(context.Background(),
		"children -filter:tweets", 500) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}

		fmt.Println(tweet.Text)
		file, err := json.Marshal(tweet)
		if err != nil {
			fmt.Println(err)
		}
		// _ = ioutil.WriteFile("tweets.json", file, 0644)
		f, err := os.OpenFile("./tweets.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		n, err := io.WriteString(f, string(file))
		if err != nil {
			fmt.Println(n, err)
		}
	}
}
