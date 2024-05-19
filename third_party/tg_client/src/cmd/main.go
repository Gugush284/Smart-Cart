package main

import (
	"flag"

	"log"

	"github.com/BurntSushi/toml"

	tgserver "tg_client/src/internal/bot"
	"tg_client/src/internal/bot/configs"
)

var bconfigPath string

func init() {
	flag.StringVar(&bconfigPath, "bot-config-path", "src/configs/bot.toml", "path to bot config file")
}

func bot(configPath string) {

	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	tgserver.Start(config)
}

func main() {
	flag.Parse()

	bot(bconfigPath)
}
