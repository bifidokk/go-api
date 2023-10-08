package main

import (
	"github.com/bifidokk/go-api/internal/config"
	"github.com/bifidokk/go-api/internal/server"
)

func main() {
	conf := config.NewConfig()
	config.Init(conf)

	server.Start(conf)
}
