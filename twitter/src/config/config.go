/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2022-08-16
 */

// Package config /**
package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const DefaultConfigFilename = "./config.json"

type Log struct {
	LogLevel   int    `json:"LogLevel"`
	LogFileDir string `json:"LogFileDir"`
}
type NetworkConfig struct {
	ConsumerKey    string `json:"ConsumerKey"`
	ConsumerSecret string `json:"ConsumerSecret"`
	AccessToken    string `json:"AccessToken"`
	AccessSecret   string `json:"AccessSecret"`
	Logger         Log    `json:"Logger"`
}

var Network *NetworkConfig

func init() {
	file, e := ioutil.ReadFile(DefaultConfigFilename)
	if e != nil {
		fmt.Printf("config file error: %v\n", e)
		os.Exit(1)
	}
	// Remove the UTF-8 Byte Order Mark
	file = bytes.TrimPrefix(file, []byte("\xef\xbb\xbf"))

	config := NetworkConfig{}
	e = json.Unmarshal(file, &config)
	if e != nil {
		fmt.Printf("unmarshal json config file erro %v\n", e)
		os.Exit(1)
	}
	//fmt.Println("config:",config)
	Network = &config
}
