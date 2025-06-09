package main

import (
	tgClient "HoldMyLink_Bot/client/telegram"
	event_consumer "HoldMyLink_Bot/consumer/event-consumer"
	"HoldMyLink_Bot/events/telegram"
	"HoldMyLink_Bot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

// 7576033146:AAEE12FVdvLsOiUB8dTs86LO1IfOlSuwU-0
func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()), files.New(storagePath))
	log.Print("service is started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped")

	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
