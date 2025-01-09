package handlers

import (
	"bot/storage"
	"bot/telegram"
)

type handlers struct {
	tg      telegram.TelegramI
	storage storage.Storage
}
type Handlers struct {
	Tg      telegram.TelegramI
	Storage storage.Storage
}

func NewHandler(h Handlers) handlers {
	return handlers{h.Tg, h.Storage}
}

