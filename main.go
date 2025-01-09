package main

import (
	"log"

	"bot/api"
	postgres "bot/storage/sql"

	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/telebot.v3"
)

const (
	sqliteStoragePath = "user=postgres password=postgres dbname=mb_doner_bot sslmode=disable"
)

// 7917631019:AAE_pQRmw1otdm7XNZtsr8XzG19aGVKgz4I

func main() {
	s, err := postgres.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	// if err := s.Init(context.TODO()); err != nil {
	// 	log.Fatal("can't init storage: ", err)
	// }

	// token := telegram.NewUpdateTg(mustToken())

	pref := telebot.Settings{
		Token:  "7635834906:AAF-inAvfxCydE5o1mCtDHoDcI3_0j5bIo8",
		Poller: &telebot.LongPoller{},
	}

	

	api.Api(&api.Options{
		Tg:      pref,
		Storage: s,
	})

}

// func mustToken() string {
// 	token := flag.String(
// 		"tg-bot-token",
// 		"7635834906:AAF-inAvfxCydE5o1mCtDHoDcI3_0j5bIo8",
// 		"token for access to telegram bot",
// 	)

// 	flag.Parse()

// 	if *token == "" {
// 		log.Fatal("token is not specified")
// 	}

// 	return *token
// }
