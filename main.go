package main

import (
	"log"
	"time"

	"bot/api"
	postgres "bot/storage/sql"

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

	
	pref := telebot.Settings{
		Token: "7917631019:AAHqy5biA_oTzvrXuavk8Je5CsdCnNVs-o0",
		Poller: &telebot.LongPoller{
			Timeout: 30 * time.Second,
		},
	}
	
	api.Api(&api.Options{
		Tg:      pref,
		Storage: s,
	})

}

// Token: "7635834906:AAF-inAvfxCydE5o1mCtDHoDcI3_0j5bIo8",
// Poller: &telebot.Webhook{
	// 	Listen:   ":8080",
	// 	Endpoint: &telebot.WebhookEndpoint{PublicURL: "https://e724-185-203-238-154.ngrok-free.app"},
// },
// if err := s.Init(context.TODO()); err != nil {
	// 	log.Fatal("can't init storage: ", err)
	// }
	
	// token := telegram.NewUpdateTg(mustToken())
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
