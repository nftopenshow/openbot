/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2022-08-16
 */
package main

import (
	"github.com/openshow/openbot/twitter/src/action"
	"github.com/openshow/openbot/twitter/src/log"
)

func main() {
	message := "Hello world"
	err := action.SendTweet(message)
	if err != nil {
		log.Warn("send tweet err:", err.Error())
	}
}
