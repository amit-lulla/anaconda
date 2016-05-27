package twitterapi_test

import (
	"fmt"
	"time"

	"github.com/amit-lulla/twitterapi"
)

// Initialize an client library for a given user.
// This only needs to be done *once* per user
func ExampleTwitterApi_InitializeClient() {
	twitterapi.SetConsumerKey("your-consumer-key")
	twitterapi.SetConsumerSecret("your-consumer-secret")
	api := twitterapi.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)
	fmt.Println(*api.Credentials)
}

func ExampleTwitterApi_GetSearch() {

	twitterapi.SetConsumerKey("your-consumer-key")
	twitterapi.SetConsumerSecret("your-consumer-secret")
	api := twitterapi.NewTwitterApi("your-access-token", "your-access-token-secret")
	search_result, err := api.GetSearch("golang", nil)
	if err != nil {
		panic(err)
	}
	for _, tweet := range search_result.Statuses {
		fmt.Print(tweet.Text)
	}
}

// Throttling queries can easily be handled in the background, automatically
func ExampleTwitterApi_Throttling() {
	api := twitterapi.NewTwitterApi("your-access-token", "your-access-token-secret")
	api.EnableThrottling(10*time.Second, 5)

	// These queries will execute in order
	// with appropriate delays inserted only if necessary
	golangTweets, err := api.GetSearch("golang", nil)
	twitterapiTweets, err2 := api.GetSearch("twitterapi", nil)

	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err)
	}

	fmt.Println(golangTweets)
	fmt.Println(twitterapiTweets)
}

// Fetch a list of all followers without any need for managing cursors
// (Each page is automatically fetched when the previous one is read)
func ExampleTwitterApi_GetFollowersListAll() {
	pages := api.GetFollowersListAll(nil)
	for page := range pages {
		//Print the current page of followers
		fmt.Println(page.Followers)
	}
}
