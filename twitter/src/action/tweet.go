// Package twitter /**
package action

import (
	"github.com/openshow/openbot/twitter/src/log"
)

func SendTweet(message string) error {
	tweet, resp, err := twitterClient.Statuses.Update(message, nil)
	if err != nil {
		return err
	} else {
		log.Info("send twitter success, twitter id:", tweet.ID, ", response code:", resp.StatusCode)
	}

	return nil
}
