package main

import (
	"log"

	"gitlab.ozon.dev/ninashvl/homework-1/config"
	"gitlab.ozon.dev/ninashvl/homework-1/internal/clients/tg"
	"gitlab.ozon.dev/ninashvl/homework-1/internal/messages"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgClient, err := tg.New(cfg)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	bot := messages.New(tgClient)
	tgClient.ListenUpdates(bot)
}
