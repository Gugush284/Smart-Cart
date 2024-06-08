package main

import (
	"flag"

	"log"

	apiserver "cart/src/internal/apiserver/src"
)

var sconfigPath string

func server(configPath string) {

	config := apiserver.NewConfig()

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}

func main() {
	flag.Parse()

	server(sconfigPath)
}
