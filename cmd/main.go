package main

import (
	"flag"

	"log"

	"github.com/BurntSushi/toml"

	"cart/internal/apiserver/apiserver"
	"cart/internal/bot/configs"
)

var sconfigPath string
var bconfigPath string

func init() {
	flag.StringVar(&sconfigPath, "server-config-path", "configs/server.toml", "path to server config file")
	flag.StringVar(&bconfigPath, "bot-config-path", "configs/bot.toml", "path to bot config file")
}

func server(configPath string) {

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

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

	go bot(bconfigPath)

	server(sconfigPath)
}
