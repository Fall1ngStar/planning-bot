package main


import (
	"github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
	"os"
)

func main() {
 	 config := oauth1.NewConfig(os.Getenv("BOT_KEY"), os.Getenv("BOT_KEY_SECRET"))
	token := oauth1.NewToken(os.Getenv("BOT_TOKEN"), os.Getenv("BOT_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})

	// Send a Tweet
	client.Statuses.Update("Mangez tous des pâtes", nil)

	// Status Show
	client.Statuses.Show(585613041028431872, nil)

	// Search Tweets
	client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "gopher",
	})

	// User Show
	client.Users.Show(&twitter.UserShowParams{
		ScreenName: "dghubble",
	})

	// Followers
	client.Followers.List(&twitter.FollowerListParams{})
}
