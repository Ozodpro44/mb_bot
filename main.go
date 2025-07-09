package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"bot/api"
	postgres "bot/storage/sql"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

const (
	sqliteStoragePath = "user=%s password=%s dbname=%s host=%s port=%s sslmode=%s"
)

// 7917631019:AAE_pQRmw1otdm7XNZtsr8XzG19aGVKgz4I

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	token := os.Getenv("BOT_TOKEN")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	storagePath := fmt.Sprintf(sqliteStoragePath, user, password, dbname, host, port, sslmode)

	s, err := postgres.New(storagePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	pref := telebot.Settings{
		Token: token,
		Poller: &telebot.LongPoller{
			Timeout: 30 * time.Second,
		},
	}

	r := mux.NewRouter()
	api.Api(&api.Options{
		Tg:      pref,
		Storage: s,
		R:       r,
	})

	// r.PathPrefix("/photos/").Handler(http.StripPrefix("/photos/", http.FileServer(http.Dir("./photos"))))

	// Products

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
