package main

import (
	"flag"
	"log"
	tgClient "psy_bot/clients/telegram"
	eventConsumer "psy_bot/consumer/event-consumer"
	"psy_bot/events/telegram"
	"psy_bot/lib/storage/files"
)

// to install github repo directly in folder for website-located server, fart and use 'go mod init github.com/youruser/yourrepo'

const (
	tgBotHost   = "api.telegram.org" // Можно выделить в флаг, как токен, для гибкости кода
	storagePath = "files_storage"
	batchSize   = 100
)

// TOKEN: '5565501247:AAHBitaUIxlM0lCNM4RkRDzm_T2fu5I8VDM'

// TODO: 1) Сделать систему общения между ButtonPressed() и ProcessCallback()

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Printf("Бот работает")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Случилась ошибка, бот неверно запустился")
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

func mustDeta() string {
	projectkey := flag.String(
		"deta-project-key",
		"",
		"Deta project key for database access",
	)

	flag.Parse()

	if *projectkey == "" {
		log.Fatal("token is not specified")
	}

	return *projectkey
}
