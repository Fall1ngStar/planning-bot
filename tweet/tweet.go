package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
    "os"
    "time"
)

func tweet (message string) {
    config := oauth1.NewConfig(os.Getenv("BOT_KEY"), os.Getenv("BOT_KEY_SECRET"))
    token := oauth1.NewToken(os.Getenv("BOT_TOKEN"), os.Getenv("BOT_TOKEN_SECRET"))
    httpClient := config.Client(oauth1.NoContext, token)

    // Twitter client
    client := twitter.NewClient(httpClient)

    // Send a Tweet
    client.Statuses.Update(message, nil)
}

func handler() {
    tweet("It is " + time.Now().String() + " and I'm bored")
}

func main() {
    lambda.Start(handler)
}
