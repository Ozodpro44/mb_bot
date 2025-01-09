package telegram

import (
	"log"

	"gopkg.in/telebot.v3"
)

type TelegramI interface {
	// SendMessages(text string, ID int64) int
	// SendMessageWithInlineButton(text string, ID int64, buttons tgbotapi.InlineKeyboardMarkup) int
	// SendReplyKeyboard(text string, ID int64, buttons tgbotapi.ReplyKeyboardMarkup) int
	// GetUpdatesText() *tgbotapi.UpdatesChannel
	// DeleteMessage(ChatID int64, msgID int)
	// CreateReplyKeyboard(items []string, buttonsPerRow int) tgbotapi.ReplyKeyboardMarkup
	// SendMsg(msg tgbotapi.Chattable) int
	// GetFile(config tgbotapi.FileConfig) (tgbotapi.File, error)
    Menu(ID int64)
}

type telegramConn struct {
	tg *telebot.Bot
}

func NewUpdateTg(pref telebot.Settings) TelegramI {
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil
	}
    
	return &telegramConn{
		tg: bot,
	}
}

func (t *telegramConn) Menu(ID int64) {
    
}

// func (t *telegramConn) SendMessages(text string, ID int64) int {

//     // Log the received message
//     log.Printf("Received message from %d: ", ID)

//         // Send a reply to the user
//     message := tgbotapi.NewMessage(ID, text)
//     msg, _ :=t.tg.Send(message)

//     return msg.MessageID

// }

// func (t *telegramConn) SendMessageWithInlineButton(text string, ID int64, buttons tgbotapi.InlineKeyboardMarkup) int {

//     msgReply := tgbotapi.NewMessage(ID, text)
//     msgReply.ReplyMarkup = buttons
//     msg, _ := t.tg.Send(msgReply)

//     return msg.MessageID

// }

// func (t *telegramConn) GetUpdatesText() *tgbotapi.UpdatesChannel {
// 	updateConfig := tgbotapi.NewUpdate(100)
//     updateConfig.Timeout = 60

//     updates := t.tg.GetUpdatesChan(updateConfig)

// 	// update := tgbotapi.Update{}

//     // for update = range updates {
//     //     if update.Message == nil { // Ignore non-message updates
//     //         continue
//     //     }

//     //     updateConfig.Offset = update.UpdateID + 1
//     //     // Exit after receiving one message
//     //     log.Println("Exiting after receiving the first message.")
// 	// 	break
//     // }

//     // updates.Clear()

// 	return &updates

// }

// func (t *telegramConn) SendReplyKeyboard(text string, ID int64, buttons tgbotapi.ReplyKeyboardMarkup) int{

//     // Send a message with the reply keyboard
//     msg := tgbotapi.NewMessage(ID, "Choose an option:")
//     msg.ReplyMarkup = buttons

//     message, _  := t.tg.Send(msg)

//     return message.MessageID
// }

// func (t *telegramConn) DeleteMessage(ChatID int64, msgID int) {
//     deleteMsg := tgbotapi.NewDeleteMessage(ChatID, msgID)
//     _, err := t.tg.Request(deleteMsg)
//     if err != nil {
//         log.Println("Failed to delete message:", err)
//     }
// }

// func (t *telegramConn) CreateReplyKeyboard(items []string, buttonsPerRow int) tgbotapi.ReplyKeyboardMarkup {
// 	var keyboardRows [][]tgbotapi.KeyboardButton

// 	// Iterate through items and create rows
// 	for i := 0; i < len(items); i += buttonsPerRow {
// 		end := i + buttonsPerRow
// 		if end > len(items) {
// 			end = len(items)
// 		}

// 		var row []tgbotapi.KeyboardButton
// 		for _, item := range items[i:end] {
// 			row = append(row, tgbotapi.NewKeyboardButton(item))
// 		}
// 		keyboardRows = append(keyboardRows, row)
// 	}

// 	return tgbotapi.NewReplyKeyboard(keyboardRows...)
// }

// func (t *telegramConn) SendMsg(msg tgbotapi.Chattable) int{
// 	ms, _ := t.tg.Send(msg)
//     return ms.MessageID
// }
// func (t *telegramConn) GetFile(config tgbotapi.FileConfig) (tgbotapi.File, error) {
// 	return t.tg.GetFile(config)
// }

// func (t *telegramConn) SendMessagesWithPhoto(text string, ID int64, photo string)
