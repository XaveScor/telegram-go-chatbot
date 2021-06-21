package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Telegram struct {
		Token         string   `json:"token"`
		Chat          string   `json:"chat"`
		StreamChannel string   `json:"stream_channel"`
		Channel       string   `json:"channel"`
		BotApiUrl     string   `json:"bot_api_url"`
		Admins        []string `json:"admins"`
		Moders        []string `json:"moders"`
		SysAdmin      string   `json:"sysadmin"`
	}
	Webhook struct {
		Listen         string   `json:"listen"`
		Port           int      `json:"port"`
		AllowedUpdates []string `json:"allowed_updates"`
	}
	Youtube struct {
		ApiKey      string `json:"api_key"`
		ChannelName string `json:"channel_name"`
		ChannelID   string `json:"channel_id"`
	}
	CurrencyKey string `json:"currency_key"`
	ReleasesUrl string `json:"releases_url"`
}

func ConfigInit(file string) Configuration {
	var Config Configuration
	if _, err := os.Stat(file); err == nil {
		ConfigFile, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		err = json.NewDecoder(ConfigFile).Decode(&Config)
		if err != nil {
			log.Fatal(err)
		}
	} else if os.IsNotExist(err) {
		Config.Telegram.Admins = []string{}
		Config.Telegram.Moders = []string{}
		Config.Telegram.BotApiUrl = "https://api.telegram.org"
		Config.Webhook.Listen = "127.0.0.1"
		Config.Webhook.Port = 8440
		Config.Webhook.AllowedUpdates = []string{"message", "channel_post", "callback_query", "chat_member"}
		jsonData, _ := json.MarshalIndent(Config, "", "\t")
		_ = ioutil.WriteFile(file, jsonData, 0600)
		log.Fatal(err)
	}
	return Config
}

var Config = ConfigInit("config.json")
