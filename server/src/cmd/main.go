package main

import (
	"flag"

	"log"

	"github.com/BurntSushi/toml"

	apiserver "cart/src/internal/apiserver/src"
)

var sconfigPath string

func init() {
	flag.StringVar(&sconfigPath, "server-config-path", "src/configs/server.toml", "path to server config file")
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

func main() {
	flag.Parse()

	server(sconfigPath)
}
