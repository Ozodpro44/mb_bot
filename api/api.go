package api

import (
	"bot/api/handlers"
	"bot/storage"
	"log"

	"gopkg.in/telebot.v3"
)

const (
	OrderCmd = "/orders"
	HelpCmd  = "/help"
	StartCmd = "/start"
	AdminCmd = "/admin"
)

const (
	adminSave    = "Dacha qoshish"
	adminOrder   = "Dachalar"
	adminClients = "Klientlar"
	adminExit    = "Exit"
)

// const (
// 	clientOrder     = "Dachalar"
// 	clientPackage   = "Zakazlarim"
// 	clientAdminCall = "Aloqa"
// )

type Options struct {
	Tg      telebot.Settings
	Storage storage.Storage
}

func Api(o *Options) {
	bot, err := telebot.NewBot(o.Tg)
	if err != nil {
		log.Fatal(err)
		return 
	}
	h := handlers.NewHandler(handlers.Handlers{Storage: o.Storage})

	bot.Handle("/start", h.HandleLanguage)

	bot.Handle(telebot.OnText, h.UserMsgStatus)

	bot.Handle(telebot.OnContact, h.HandleRegistrationSteps)

	// bot.Handle(telebot.OnLocation, h)

	bot.Handle(&telebot.InlineButton{Unique: "language_add"}, h.GetUserName)

	bot.Handle(&telebot.InlineButton{Unique: "confirm_order"}, h.CompleteOrder)


	bot.Handle(&telebot.InlineButton{Unique: "lang_btn"}, h.ChangeLanguage)

	bot.Handle(&telebot.InlineButton{Unique: "my_orders"}, h.ShowUserOrders)

	bot.Handle(&telebot.InlineButton{Unique: "order_btn"}, h.ShowMenu)

	bot.Handle(&telebot.InlineButton{Unique: "show_cart"}, h.SendCart)

	bot.Handle(&telebot.InlineButton{Unique: "get_category_by_id"}, h.ShowProducts)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_products_menu"}, h.ShowProducts)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_categories"}, h.ShowMenu)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_user_menu"}, h.ShowUserMenu)

	bot.Handle(&telebot.InlineButton{Unique: "add_category"}, h.CreateCategory)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_admin_menu"}, h.ShowAdminPanel)

	bot.Handle(&telebot.InlineButton{Unique: "add_product"}, h.AddProductHandler)

	bot.Handle(&telebot.InlineButton{Unique: "clear_cart"}, h.ClearCart)

	bot.Handle(&telebot.InlineButton{Unique: "decrement_cart_product"}, h.HandleDecrement)

	bot.Handle(&telebot.InlineButton{Unique: "increment_cart_product"}, h.HandleIncrement)

	bot.Handle("/admin", h.ShowAdminPanel)

	bot.Handle(&telebot.InlineButton{Unique: "get_product_by_id"}, h.ShowProductByID)

	bot.Start()

}
	
	
	
	
	// updates := o.Tg.GetUpdatesText()

	// update := tgbotapi.Update{}

	// var msgId int
	// var ChatID int64

	// var wg sync.WaitGroup

	// for update = range *updates {
	// 	wg.Add(1)
	// 	go func(update tgbotapi.Update) {
	// 		wg.Done()
	// 		if update.Message == nil {
	// 			switch update.CallbackQuery.Data {
	// 			case adminOrder:
	// 				o.Tg.DeleteMessage(ChatID, msgId)
	// 				defer wg.Done()
	// 			case adminClients:
	// 				o.Tg.DeleteMessage(ChatID, msgId)

	// 			case adminSave:
	// 				o.Tg.DeleteMessage(ChatID, msgId)

	// 			case adminExit:
	// 				o.Tg.DeleteMessage(ChatID, msgId)
	// 			default:
	// 				o.Tg.SendMessages(msgUnknownCommand, update.Message.Chat.ID)
	// 			}
	// 		} else if update.CallbackQuery == nil {
	// 			switch update.Message.Text {

	// 			case "/add":
	// 				// text = msgHelp
	// 				h.SendToAllUser(updates, update.Message.From.ID)
	// 			case StartCmd:
	// 				h.RegisterUser(updates, update.Message.From.ID)
	// 				// h.GetProductsByCateg(ChatID,12, 2,12.5)
	// 				// button := tgbotapi.KeyboardButton{
	// 				// 	Text:            "Share Phone Number",
	// 				// 	RequestContact:  true, // This makes the button request the user's contact
	// 				// }

	// 				// keyboard := tgbotapi.NewReplyKeyboard(
	// 				// 	[]tgbotapi.KeyboardButton{button},
	// 				// )

	// 				// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please share your phone number:")
	// 				// msg.ReplyMarkup = keyboard

	// 				// o.Tg.SendReplyKeyboard(msg.Text, update.Message.Chat.ID, keyboard)
	// 				// h.RegisterUser(update.SentFrom().ID, update.SentFrom().UserName, update.SentFrom().FirstName)
	// 			case AdminCmd:

	// 			default:
	// 				o.Tg.SendMessages(msgUnknownCommand, update.Message.From.ID)
	// 			}
	// 			log.Printf("got new command '%s' from '%s", update.SentFrom().UserName, update.Message.Text)

	// 			ChatID = update.Message.Chat.ID

	// 		}
	// 	}(update)
	// 	updates.Clear()

	// 	wg.Wait()

	// }
// }
