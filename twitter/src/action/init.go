// Package twitter /**
package action

import (
	"github.com/dghubble/oauth1"
	"github.com/openshow/openbot/twitter/go-twitter/twitter"
	"github.com/openshow/openbot/twitter/src/config"
)

var twitterClient *twitter.Client

func init() {
	consumerKey := config.Network.ConsumerKey
	consumerSecret := config.Network.ConsumerSecret
	accessToken := config.Network.AccessToken
	accessSecret := config.Network.AccessSecret

	conf := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := conf.Client(oauth1.NoContext, token)

	// Twitter client
	twitterClient = twitter.NewClient(httpClient)
}
