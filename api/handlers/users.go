package handlers

import (
	"bot/lib/helpers"
	"bot/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/telebot.v3"
)

const groupID = int64(-4774043538)

var Messages = map[string]map[string]string{
	"en": {
		"welcome":            "Welcome! Please enter your name:",
		"phone":              "Please send your phone number by button 👇:",
		"done":               "Registration completed! Welcome, ",
		"exists":             "You are already registered!",
		"error":              "An error occurred. Please try again later.",
		"language_prompt":    "Пожалуйста, выберите ваш язык \n\nIltimos, o'zingizga qulay tilni tanlang \n\nPlease select your language:",
		"lang_btn":           "🇬🇧Language",
		"order_btn":          "🛍Make an order",
		"get_phone":          "📱 Share your phone number",
		"my_orders":          "My orders",
		"about_us":           "About us",
		"back":               "⬅️Back",
		"cart":               "🛒Cart",
		"add_to_cart":        "📥Add to cart",
		"clear_cart":         "♻️ Clear",
		"cart_messsage":      "*%s*\n\nPrice: %d UZS\nQuantity: %d\nTotal: %d UZS",
		"empty_cart":         "Your cart is empty🛒🚫",
		"user_menu":          "*Main menu:*\n\nChoose one of the following options",
		"cart_items_msg":     "*%s* x %d \\= %d sum\n",
		"cart_total":         "\n*Total:* %d sum",
		"confirm_order":      "✅Confirm order",
		"continue_order":     "🧾Continue order",
		"added_to_cart":      "Product added to cart✅",
		"order_msg":          "📋 *Order number*: %d \n🚕 *Delivery type*: Delivery \n🏠 *Address*: %s \n📍 *Branch*: Yakkasaroy \n\n %s \n\n💵 *Products*: %v \n🚚 *Delivery price*: %s \n💰 *Total*: %v \nPayment type: %s \nStatus: %s",
		"delivery":           "Delivery🚚",
		"pickup":             "Pickup🚶‍♂️",
		"re-order":           "Re-order🔄",
		"location_btn":       "📍 Send location",
		"location_msg":       "Please send your location by button 👇:",
		"true-location":      "✅Correct",
		"false-location":     "❌Incorrect",
		"check-location-msg": "Is this location correct?: \n%s",
		"invalid_phone":      "Invalid phone number type please try again 👇",
		"pending":            "Pending payment...",
		"preparing":          "Preparing...",
		"deliver":            "Delivering...",
		"complete":           "Delivered✅",
		"payment_type_msg":   "Choose payment type:",
		"cash":               "💵Cash",
		"card":               "💳Card",
		"thanks":             "Thank you!",
		"no_need":            "No need location❗",
		"succsess":           "Your order has been successfully placed\n",
		"our_card":           "\nMake payment to this card 👇\n\n5614 6806 1838 4578 \nMustafa Bugra",
		"closed_msg":         "We are closed for today.😔",
	},
	"ru": {
		"welcome":            "Добро пожаловать! Пожалуйста, введите ваше имя:",
		"phone":              "Пожалуйста, поделитесь своим номером телефона 👇",
		"done":               "Регистрация завершена! Добро пожаловать, ",
		"exists":             "Вы уже зарегистрированы!",
		"error":              "Произошла ошибка. Попробуйте позже.",
		"language_prompt":    "Пожалуйста, выберите ваш язык:",
		"lang_btn":           "🇷🇺Язык",
		"order_btn":          "🛍Сделать заказ",
		"get_phone":          "📱 Поделитесь своим номером телефона",
		"my_orders":          "Мои заказы",
		"about_us":           "О нас",
		"back":               "⬅️Назад",
		"cart":               "🛒Корзина",
		"add_to_cart":        "📥Добавить в корзину",
		"clear_cart":         "♻️ Очистить",
		"cart_messsage":      "*%s*\n\nЦена: %d UZS\nКоличество: %d\nИтого: %d UZS",
		"empty_cart":         "Корзина пуста🛒🚫",
		"user_menu":          "*Главное меню:*\n\nВыберите одну из следующих опций",
		"cart_items_msg":     "*%s* x %d \\= %d сум\n",
		"cart_total":         "\n*Итого:* %d сум",
		"confirm_order":      "✅Подтвердить заказ",
		"continue_order":     "🧾Продолжить заказ",
		"added_to_cart":      "Товар добавлен в корзину✅",
		"order_msg":          "📋 *Номер заказа*: %d \n🚕 *Способ доставки*: Доставка \n🏠 *Адрес*: %s \n📍 *Филиал*: Yakkasaroy  \n\n	%s \n\n💵 *Товары*: %v \n🚚 *Стоимость доставки*: %s \n💰 *Итого*: %v \nСпособ оплаты: %s \nСтатус: %s",
		"delivery":           "Доставка🚚",
		"pickup":             "Самовывоз🚶‍♂️",
		"re-order":           "Перезаказать🔄",
		"location_btn":       "📍 Отправить местоположение",
		"location_msg":       "Пожалуйста, отправьте свое местоположение, нажав кнопку ниже 👇:",
		"true-location":      "✅Верно",
		"false-location":     "❌Неверно",
		"check-location-msg": "Это верно?: \n%s",
		"invalid_phone":      "Неверный тип номера телефона, пожалуйста, попробуйте еще раз 👇",
		"pending":            "Ожидается оплата...",
		"preparing":          "Готовится...",
		"deliver":            "Доставляется...",
		"complete":           "Доставлено✅",
		"payment_type_msg":   "Выберите способ оплаты:",
		"cash":               "💵Наличные",
		"card":               "💳Карта",
		"thanks":             "Спасибо!",
		"no_need":            "Не нужно местоположение❗",
		"succsess":           "Ваш заказ успешно оформлен\n",
		"our_card":           "\nПроизведите оплату на эту карту 👇\n\n5614 6806 1838 4578 \nMustafa Bugra",
		"closed_msg":         "Сегодня заведение закрыто😔",
	},
	"uz": {
		"welcome":            "Xush kelibsiz! Iltimos, ismingizni kiriting:",
		"phone":              "Iltimos, telefon raqamingizni ulashing 👇:",
		"done":               "Ro'yxatdan o'tish yakunlandi! Xush kelibsiz, ",
		"exists":             "Siz allaqachon ro'yxatdan o'tgansiz!",
		"error":              "Xatolik yuz berdi. Iltimos, keyinroq urinib ko'ring.",
		"language_prompt":    "Iltimos, o'zingizga qulay tilni tanlang:",
		"lang_btn":           "🇺🇿Til",
		"order_btn":          "🛍Buyurtma berish",
		"get_phone":          "📱 Telefon raqamingizni ulashing",
		"my_orders":          "Mening buyurtmalarim",
		"about_us":           "Biz haqimizda",
		"back":               "⬅️Orqaga",
		"cart":               "🛒Savatcha",
		"add_to_cart":        "📥Savatga qo'shish",
		"clear_cart":         "♻️ Tozalash",
		"cart_messsage":      "*%s*\n\nNarxi: %d UZS\nMiqdor: %d\nJami: %d UZS",
		"empty_cart":         "Savatingiz bo'sh🛒🚫",
		"user_menu":          "*Asosiy menyu:*\n\n Quyidagilardan birini tanlang:",
		"cart_items_msg":     "*%s* x %d \\= %d so'm\n",
		"cart_total":         "\n*Jami:* %d so'm",
		"confirm_order":      "✅Buyurtmani tasdiqlash",
		"continue_order":     "🧾Buyurtmani davom ettirish",
		"added_to_cart":      "Mahsulot savatga qo'shildi✅",
		"order_msg":          "📋 *Buyurtma raqami*: %d \n🚕 *Yetkazib berish turi*: Yetkazib berish \n🏠 *Manzil*: %s \n📍 *Filial*: Yakkasaroy \n\n%s \n\n💵 *Mahsulotlar*: %v \n🚚 *Yetkazib berish narxi*: %s \n💰 *Umumiy*: %v \nTo\\'lov turi: %s \nStatus: %s",
		"delivery":           "Yetkazib berish🚚",
		"pickup":             "Olib ketish🚶‍♂️",
		"re-order":           "Takrorlash🔄",
		"location_btn":       "📍 Manzilni jo'natish",
		"location_msg":       "Manzilingizni quyidagi tugma orqali yuboring 👇:",
		"true-location":      "Togri ✅",
		"false-location":     "Noto'g'ri ❌",
		"check-location-msg": "Bu manzil to'g'rimi?: \n%s",
		"invalid_phone":      "Telefon raqami turi noto‘g‘ri, qayta urinib ko‘ring👇",
		"pending":            "To'lov kutilmoqda",
		"preparing":          "Tayyorlanmoqda",
		"deliver":            "Yo'lda",
		"complete":           "Yetkazib berildi✅",
		"payment_type_msg":   "To'lov turini tanlang:",
		"cash":               "💵Naqd",
		"card":               "💳Karta",
		"thanks":             "Rahmat!",
		"no_need":            "Joylashuvingiz hozir kerak emas❗",
		"succsess":           "Buyurtmangiz muvaffaqiyatli qabul qilindi\n",
		"our_card":           "\nTo'lovni shu kartaga qiling 👇\n\n5614 6806 1838 4578 \nMustafa Bugra",
		"closed_msg":         "Buguncha yopildik😔",
	},
}

var (
	registerStep = make(map[int64]string) // Карта для отслеживания шага регистрации
	tempUserData = make(map[int64]map[string]string)
)
var (
	userLocation = make(map[int64]bool) // Сохраняем выбранный язык пользователя
)

func (h *handlers) HandleLanguage(c telebot.Context) error {
	userID := c.Sender().ID

	if h.storage.CheckAdmin(userID) {
		return h.ShowCategoryMenu(c)
	}

	exists := h.storage.CheckUserExist(userID)

	if exists {
		return h.ShowUserMenu(c)
	}

	menu := &telebot.ReplyMarkup{}
	btnEN := menu.Data("English🇬🇧", "language_add", "en")
	btnRU := menu.Data("Русский🇷🇺", "language_add", "ru")
	btnUZ := menu.Data("O'zbek🇺🇿", "language_add", "uz")

	menu.Inline(menu.Row(btnEN, btnRU, btnUZ))
	return c.Send(Messages["en"]["language_prompt"], menu)

}

func (h *handlers) GetUserName(c telebot.Context) error {
	lang := c.Callback().Data

	h.storage.SetLangUser(c.Sender().ID, lang)
	_, err := h.storage.UserMessageStatus(c.Sender().ID, "firstname")
	if err != nil {
		return c.Send(err.Error())
	}
	c.Send(Messages[lang]["welcome"])
	return nil
}

func (h *handlers) ShowUserMenu(c telebot.Context) error {
	lang, err := h.storage.GetLangUser(c.Sender().ID)

	if err != nil {
		log.Fatal(err)
	}
	userLocation[c.Sender().ID] = false
	menu := &telebot.ReplyMarkup{}

	// Define the buttons
	btnTil := menu.Data(Messages[lang]["lang_btn"], "lang_btn")
	btnZakaz := menu.Data(Messages[lang]["order_btn"], "order_btn")
	btnBuyurtmalarim := menu.Data(Messages[lang]["my_orders"], "my_orders")
	btnBizHaqimida := menu.Data(Messages[lang]["about_us"], "about_us")

	// Arrange buttons in rows
	menu.Inline(
		menu.Row(btnZakaz),
		menu.Row(btnTil, btnBuyurtmalarim),
		menu.Row(btnBizHaqimida),
	)
	menu.ResizeKeyboard = true

	// Send the menu to the user
	message := Messages[lang]["user_menu"]

	options := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: menu,
	}

	return c.EditOrSend(message, options)
}

func (h *handlers) SendAboutUs(c telebot.Context) error {
	user_id := c.Sender().ID
	lang, err := h.storage.GetLangUser(user_id)
	if err != nil {
		return c.Send(err.Error())
	}
	markup := &telebot.ReplyMarkup{}
	btnBack := markup.Data(Messages[lang]["back"], "back_to_user_menu")
	markup.Inline(markup.Row(btnBack))

	return c.Edit("MB DONER, Yakkasaroy tumani, 49/1", markup)
}

func (h *handlers) HandleRegistrationSteps(c telebot.Context) error {
	userID := c.Sender().ID
	username := c.Sender().Username

	// l := c.Message().Location.
	lang, err := h.storage.GetLangUser(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	// Ensure registration step is initialized
	name, err := h.storage.GetDataUserMessageStatus(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	phone, err := helpers.FormatPhoneNumber(c.Message().Contact.PhoneNumber)

	if err != nil {
		keyboard := &telebot.ReplyMarkup{
			ResizeKeyboard:  true, // Makes the keyboard fit nicely on screen
			OneTimeKeyboard: true, // Hides the keyboard after it's used
		}

		phoneButton := keyboard.Contact(Messages[lang]["get_phone"])
		keyboard.Reply(keyboard.Row(phoneButton))
		return c.Send(Messages[lang]["invalid_phone"], keyboard)
	}

	// Save user to database
	err = h.storage.RegisterUser(&models.User{
		TelegramID:   userID,
		Username:     username,
		Name:         name,
		Phone_Number: phone,
	})
	if err != nil {
		return c.Send(Messages[lang]["error"])
	}

	// Clear user data
	delete(registerStep, userID)
	delete(tempUserData, userID)
	c.Send(Messages[lang]["done"]+name, &telebot.ReplyMarkup{RemoveKeyboard: true})
	h.ShowUserMenu(c)
	return nil
}

func (h *handlers) RequestLocation(c telebot.Context) error {
	c.Delete()

	userID := c.Sender().ID

	userLocation[userID] = true

	lang, err := h.storage.GetLangUser(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	opened, err := h.storage.CheckOpened()

	if err != nil {
		return c.Send(err.Error())
	}

	if !opened {
		noloc := &telebot.ReplyMarkup{}
		btnBack := noloc.Data(Messages[lang]["back"], "back_to_user_menu")
		noloc.Inline(noloc.Row(btnBack))
		return c.Send(Messages[lang]["closed_msg"], noloc)
	}

	// Create location button
	btn := &telebot.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
		// RemoveKeyboard:  true,
	}

	// Text:     "📍 Отправить местоположение",
	// Create reply markup (keyboard)
	loc := btn.Location(Messages[lang]["location_btn"])

	btn.Reply(
		btn.Row(loc),
	)

	option := &telebot.SendOptions{
		ReplyMarkup: btn,
		ParseMode:   telebot.ModeMarkdownV2,
	}

	// Send message
	return c.Send(Messages[lang]["location_msg"], option)
}

func (h *handlers) HandleLocation(c telebot.Context) error {
	userID := c.Sender().ID

	// c.Bot().EditReplyMarkup(c.Message(), &telebot.ReplyMarkup{RemoveKeyboard: true})
	lang, err := h.storage.GetLangUser(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	if !userLocation[userID] {
		noloc := &telebot.ReplyMarkup{}
		btnBack := noloc.Data(Messages[lang]["back"], "back_to_user_menu")
		noloc.Inline(noloc.Row(btnBack))
		return c.EditOrSend(Messages[lang]["no_need"], noloc)
	}
	var latitude float32
	var longitude float32

	loc := c.Message().Location
	if loc != nil {
		latitude = loc.Lat  // Latitude
		longitude = loc.Lng // Longitude
		// return c.Send("Received your location!\nLatitude: %f\nLongitude: %f", latitude, longitude)
	} else {
		return c.Send("Location not found")
	}

	c.Send(Messages[lang]["thanks"], &telebot.ReplyMarkup{
		RemoveKeyboard: true,
	})

	// latitude := c
	// longitude := c.

	// Call geocoding API to get address
	address, err := helpers.GetAddressFromCoordinates(latitude, longitude, lang)
	if err != nil {
		return err
	}

	menu := &telebot.ReplyMarkup{}
	btnFalse := menu.Data(Messages[lang]["false-location"], "false-location")
	btnTrue := menu.Data(Messages[lang]["true-location"], "true-location")
	btnBack := menu.Data(Messages[lang]["back"], "show_cart")
	menu.Inline(
		menu.Row(btnFalse, btnTrue),
		menu.Row(btnBack),
	)
	options := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: menu,
	}

	location := models.Location{
		UserID:    userID,
		Latitude:  latitude,
		Longitude: longitude,
		Name_uz:   address,
		Name_ru:   address,
		Name_en:   address,
	}

	_, err = h.storage.CreateLocation(&location)
	if err != nil {
		return c.Send(err.Error())
	}
	userLocation[userID] = false

	msg := helpers.EscapeMarkdownV2(fmt.Sprintf(Messages[lang]["check-location-msg"], address))

	// Ask user to confirm the address
	return c.EditOrSend(msg, options)
}
func (h *handlers) HandleFalseLocation(c telebot.Context) error {
	h.storage.DeleteLocationByUserID(c.Sender().ID)
	return h.RequestLocation(c)
}

func (h *handlers) ShowMenu(c telebot.Context) error {
	telegramID := c.Sender().ID

	h.storage.GetLangUser(telegramID)

	// Fetch all categories
	cat, err := h.storage.GetAllCategories()
	if err != nil {
		return c.Send(fmt.Sprintf("Error: %s", err.Error()))
	}

	// Fetch user's language
	lang, err := h.storage.GetLangUser(telegramID)
	if err != nil {
		return c.Send(fmt.Sprintf("Error: %s", err.Error()))
	}

	// Generate message and buttons
	message := ""
	var buttons []telebot.Row
	menu := &telebot.ReplyMarkup{}

	switch lang {
	case "uz":
		message = " *MB Doner*\n\nMenu\n\n"
		for i := 0; i < len(cat.Categories); i += 2 {
			if i+1 < len(cat.Categories) {

				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_uz, "get_category_by_id", cat.Categories[i].ID),
					menu.Data(cat.Categories[i+1].Name_uz, "get_category_by_id", cat.Categories[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_uz, "get_category_by_id", cat.Categories[i].ID),
				))
			}
		}
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["cart"], "show_cart")))
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["back"], "back_to_user_menu")))
	case "ru":
		message = " *MB Doner*\n\nМеню\n\n"
		for i := 0; i < len(cat.Categories); i += 2 {
			if i+1 < len(cat.Categories) {

				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_ru, "get_category_by_id", cat.Categories[i].ID),
					menu.Data(cat.Categories[i+1].Name_ru, "get_category_by_id", cat.Categories[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_ru, "get_category_by_id", cat.Categories[i].ID),
				))
			}
		}
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["cart"], "show_cart")))
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["back"], "back_to_user_menu")))
	case "en":
		message = " *MB Doner*\n\nMenu\n\n"
		for i := 0; i < len(cat.Categories); i += 2 {
			if i+1 < len(cat.Categories) {

				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_en, "get_category_by_id", cat.Categories[i].ID),
					menu.Data(cat.Categories[i+1].Name_en, "get_category_by_id", cat.Categories[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_en, "get_category_by_id", cat.Categories[i].ID),
				))
			}
		}
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["cart"], "show_cart")))
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["back"], "back_to_user_menu")))
	default:
		message = " *MB Doner*\n\nMenu\n\n"
		for i := 0; i < len(cat.Categories); i += 2 {
			if i+1 < len(cat.Categories) {

				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_uz, "get_category_by_id", cat.Categories[i].ID),
					menu.Data(cat.Categories[i+1].Name_uz, "get_category_by_id", cat.Categories[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_uz, "get_category_by_id", cat.Categories[i].ID),
				))
			}
		}
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["cart"], "show_cart")))
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["back"], "back_to_user_menu")))
	}

	// Arrange buttons in rows
	menu.Inline(buttons...)
	menu.ResizeKeyboard = true

	options := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: menu,
	}

	// Send the message with buttons
	c.EditOrSend(message, options)
	return nil
}

func (h *handlers) ShowProducts(c telebot.Context) error {

	category := c.Callback().Data
	userID := c.Sender().ID

	cat, err := h.storage.GetCategoryByID(category)

	if err != nil {
		return c.Send(err.Error())
	}

	lang, err := h.storage.GetLangUser(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	products, err := h.storage.GetProductsByCategory(category)

	if err != nil {
		return c.Send(err.Error())
	}

	message := ""
	var buttons []telebot.Row
	menu := &telebot.ReplyMarkup{}

	switch lang {
	case "uz":
		message = fmt.Sprintf("    *%s*   \n\nChoose products", cat.Name_uz)
		for i := 0; i < len(products.Products); i += 2 {
			if i+1 < len(products.Products) {

				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_uz, "get_product_by_id", products.Products[i].ID),
					menu.Data(products.Products[i+1].Name_uz, "get_product_by_id", products.Products[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_uz, "get_product_by_id", products.Products[i].ID),
				))
			}
		}
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["cart"], "show_cart")))
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["back"], "back_to_categories", category)))
	case "ru":
		message = fmt.Sprintf("    *%s*   \n\nВыберите товары:", cat.Name_ru)
		for i := 0; i < len(products.Products); i += 2 {
			if i+1 < len(products.Products) {

				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_ru, "get_product_by_id", products.Products[i].ID),
					menu.Data(products.Products[i+1].Name_ru, "get_product_by_id", products.Products[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_ru, "get_product_by_id", products.Products[i].ID),
				))
			}
		}
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["cart"], "show_cart")))
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["back"], "back_to_categories", category)))
	case "en":
		message = fmt.Sprintf("     *%s*   \n\nMahsulotlarni tanlang:", cat.Name_en)
		for i := 0; i < len(products.Products); i += 2 {
			if i+1 < len(products.Products) {
				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_en, "get_product_by_id", products.Products[i].ID),
					menu.Data(products.Products[i+1].Name_en, "get_product_by_id", products.Products[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_en, "get_product_by_id", products.Products[i].ID),
				))
			}
		}
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["cart"], "show_cart")))
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["back"], "back_to_categories")))
	default:
		message = fmt.Sprintf("    *%s*  \n\nMahsulotlarni tanlang:", cat.Name_uz)
		for i := 0; i < len(products.Products); i += 2 {
			if i+1 < len(products.Products) {

				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_uz, "get_product_by_id", products.Products[i].ID),
					menu.Data(products.Products[i+1].Name_uz, "get_product_by_id", products.Products[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_uz, "get_product_by_id", products.Products[i].ID),
				))
			}
		}
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["cart"], "show_cart")))
		buttons = append(buttons, menu.Row(menu.Data(Messages[lang]["back"], "back_to_categories")))
	}

	menu.Inline(buttons...)

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: menu,
	}
	if c.Callback().Unique == "back_to_products_menu" {
		c.Delete()
		return c.Send(message, option)
	}
	c.EditOrSend(message, option)

	return nil
}

func (h *handlers) ShowProductByID(c telebot.Context) error {
	// Handle text input (for productID)
	prod := c.Callback().Data
	userID := c.Sender().ID

	// Get user language
	_, err := h.storage.GetLangUser(userID)
	if err != nil {
		return c.Send(err.Error())
	}

	// Get product details by productID
	product, err := h.storage.GetProductById(prod)
	if err != nil {
		return c.Send(err.Error())
	}

	// Initially show product details with a quantity of 1
	return h.sendProductMenu(c, product, 1)
}

// Function to display product menu with quantity options
func (h *handlers) sendProductMenu(c telebot.Context, product *models.Product, quantity int) error {
	totalPrice := int(product.Price) * quantity
	lang, err := h.storage.GetLangUser(c.Sender().ID)
	if err != nil {
		return c.Send(err.Error())
	}

	// Create markup for inline buttons
	markup := &telebot.ReplyMarkup{}
	btnDecrement := markup.Data("➖", "decrement", strconv.Itoa(quantity))
	btnIncrement := markup.Data("➕", "increment", strconv.Itoa(quantity))
	btnAddToCart := markup.Data(Messages[lang]["add_to_cart"], "add_to_cart", strconv.Itoa(quantity), product.ID)
	btnQuantity := markup.Data(strconv.Itoa(quantity), "ignore", strconv.Itoa(quantity))
	// fmt.Println(len(product.Category_id))
	btnBack := markup.Data(Messages[lang]["back"], "back_to_products_menu", product.Category_id)
	btnCart := markup.Data(Messages[lang]["cart"], "show_cart")

	markup.Inline(
		markup.Row(btnDecrement, btnQuantity, btnIncrement),
		markup.Row(btnAddToCart),
		markup.Row(btnBack, btnCart),
	)

	// Format message
	message := fmt.Sprintf(Messages[lang]["cart_messsage"],
		product.Name_ru, int(product.Price), quantity, totalPrice)

	photoPath := product.Photo // Assuming product.PhotoID contains the filename without extension
	if _, err := os.Stat(photoPath); os.IsNotExist(err) {
		return c.Send("Photo not found.")
	}

	photo := &telebot.Photo{File: telebot.FromDisk(product.Photo), Caption: message}

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
		ParseMode:   telebot.ModeMarkdownV2,
	}

	// Send or edit the message with the markup
	if c.Callback().Unique == "decrement" || c.Callback().Unique == "increment" {
		c.EditCaption(message, options)
		return h.HandleInlineButtons(c, product)
	}

	// c.Send(photo, markup, telebot.ModeMarkdownV2)

	err = c.Edit(photo, options)

	if err != nil {
		return c.Send(err.Error())
	}

	return h.HandleInlineButtons(c, product)
}

// Function to handle inline button presses for quantity and cart actions
func (h *handlers) HandleInlineButtons(c telebot.Context, product *models.Product) error {
	// Respond based on the button pressed
	c.Bot().Handle(&telebot.InlineButton{Unique: "increment"}, func(c telebot.Context) error {
		quantity, _ := strconv.Atoi(c.Data())
		return h.sendProductMenu(c, product, quantity+1)
	})

	c.Bot().Handle(&telebot.InlineButton{Unique: "decrement"}, func(c telebot.Context) error {
		quantity, _ := strconv.Atoi(c.Data())
		if quantity > 1 {
			return h.sendProductMenu(c, product, quantity-1)
		}
		return nil
	})

	c.Bot().Handle(&telebot.InlineButton{Unique: "add_to_cart"}, func(c telebot.Context) error {
		lang, err := h.storage.GetLangUser(c.Sender().ID)
		if err != nil {
			return c.Send(err.Error())
		}
		parts := strings.Split(c.Data(), "|")
		quantity, _ := strconv.Atoi(parts[0])
		productID := parts[1]
		userID := c.Sender().ID
		fmt.Println(c.Data())

		err = h.storage.AddToCart(userID, productID, quantity)

		if err != nil {
			return c.Send(err.Error())
		}

		return c.Respond(&telebot.CallbackResponse{Text: Messages[lang]["added_to_cart"]})
	})
	return nil
}
func formatCart(cart *models.Cart, lang string) string {
	var message string
	totalPrice := 0

	for _, item := range cart.Items {
		itemTotal := int(item.Price) * item.Quantity
		message += fmt.Sprintf(Messages[lang]["cart_items_msg"], item.Name_uz, item.Quantity, itemTotal)
		totalPrice += itemTotal
	}

	message += fmt.Sprintf(Messages[lang]["cart_total"], totalPrice)

	return message
}

func createCartButtons(cart *models.Cart, lang string) *telebot.ReplyMarkup {
	markup := &telebot.ReplyMarkup{}

	btnConfirm := markup.Data(Messages[lang]["confirm_order"], "confirm_order")
	btnContinue := markup.Data(Messages[lang]["continue_order"], "continue_order")
	btnClear := markup.Data(Messages[lang]["clear_cart"], "clear_cart")

	rows := []telebot.Row{
		markup.Row(btnConfirm),
		markup.Row(btnContinue),
		markup.Row(btnClear),
	}

	// Add per-product buttons (+ and -)
	for _, item := range cart.Items {
		btnDecrement := markup.Data("➖", "decrement_cart_product", item.ProductID)
		btnIncrement := markup.Data("➕", "increment_cart_product", item.ProductID)
		productButton := markup.Data(item.Name_uz, "ignore", item.ProductID)

		rows = append(rows, markup.Row(btnDecrement, productButton, btnIncrement))
	}
	markup.Inline(rows...)
	return markup
}

func (h *handlers) SendCart(c telebot.Context) error {
	lang, err := h.storage.GetLangUser(c.Sender().ID)
	if err != nil {
		return c.Send(err.Error())
	}

	// Assume 'cart' is fetched from the database or session
	cart, err := h.storage.GetCart(c.Sender().ID)
	if err != nil || len(cart.Items) == 0 {
		btn := &telebot.ReplyMarkup{}

		btnBack := btn.Data(Messages[lang]["back"], "back_to_categories")
		btn.Inline(btn.Row(btnBack))
		c.Delete()
		return c.Send(Messages[lang]["empty_cart"], btn)
	}
	c.Delete()

	message := formatCart(cart, lang)
	buttons := createCartButtons(cart, lang)

	op := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: buttons,
	}

	err = c.Send(message, op)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// replyMarkup := &telebot.ReplyMarkup{}
	// btnBack := replyMarkup.Text(Messages["en"]["back"])
	// replyMarkup.Reply(replyMarkup.Row(btnBack))
	// replyMarkup.ResizeKeyboard = true
	return nil
}

func (h *handlers) HandleIncrement(c telebot.Context) error {
	c.Respond()
	productID := c.Callback().Data // Get product ID from callback data
	userID := c.Sender().ID

	lang, err := h.storage.GetLangUser(userID)
	if err != nil {
		return c.Send(err.Error())
	}

	// Fetch the user's cart
	cart, err := h.storage.GetCart(userID)
	if err != nil {
		return c.Send("Error fetching cart")
	}

	// Increment the quantity of the selected product
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items[i].Quantity++
			break
		}
	}

	// Update the cart in storage
	err = h.storage.UpdateCart(userID, cart)
	if err != nil {
		return c.Send("Error updating cart")
	}

	// Resend the updated cart
	message := formatCart(cart, lang)
	buttons := createCartButtons(cart, lang)

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: buttons,
	}

	return c.Edit(message, option)
}

func (h *handlers) HandleDecrement(c telebot.Context) error {
	c.Respond()
	productID := c.Callback().Data
	userID := c.Sender().ID
	lang, err := h.storage.GetLangUser(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	// Fetch the user's cart
	cart, err := h.storage.GetCart(userID)
	if err != nil {
		return c.Send("Error fetching cart")
	}

	// Decrement the quantity of the selected product
	for i, item := range cart.Items {
		if item.ProductID == productID {
			if cart.Items[i].Quantity > 1 {
				cart.Items[i].Quantity--
			} else {
				// Remove the item if quantity becomes zero
				h.storage.RemoveFromCart(userID, productID)
				cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			}
			break
		}
	}

	// Update the cart in storage

	// Resend the updated cart or show an empty cart message if no items remain
	if len(cart.Items) == 0 {
		btn := &telebot.ReplyMarkup{}
		btnBack := btn.Data(Messages[lang]["back"], "back_to_categories")
		btn.Inline(btn.Row(btnBack))
		return c.Edit(Messages[lang]["empty_cart"], btn)
	}
	err = h.storage.UpdateCart(userID, cart)
	if err != nil {
		fmt.Println(err)
		return c.Send("Error updating cart")
	}

	message := formatCart(cart, lang)
	buttons := createCartButtons(cart, lang)

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: buttons,
	}

	err = c.Edit(message, option)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func (h *handlers) ClearCart(c telebot.Context) error {
	userID := c.Sender().ID
	err := h.storage.ClearCart(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	c.Respond(&telebot.CallbackResponse{Text: "Cart cleared❗"})
	return h.ShowMenu(c)

}

func (h *handlers) RequestPhoneNumber(c telebot.Context) error {
	lang, err := h.storage.GetLangUser(c.Sender().ID)
	if err != nil {
		return c.Send(err.Error())
	}
	keyboard := &telebot.ReplyMarkup{
		ResizeKeyboard:  true, // Makes the keyboard fit nicely on screen
		OneTimeKeyboard: true, // Hides the keyboard after it's used
	}

	phoneButton := keyboard.Contact(Messages[lang]["get_phone"])
	keyboard.Reply(keyboard.Row(phoneButton))

	return c.Send(Messages[lang]["phone"], keyboard)
}

func (h *handlers) UserMsgStatus(c telebot.Context) error {
	userId := c.Sender().ID
	text := c.Message().Text

	status, err := h.storage.GetUserMessageStatus(userId)

	if err != nil {
		return c.Send(err.Error())
	}

	switch status {
	case "firstname":
		h.storage.SetDataUserMessageStatus(userId, text)
		return h.RequestPhoneNumber(c)
	case "phone":

	case "location":
		h.storage.SetDataUserMessageStatus(userId, text)
		return h.ShowMenu(c)
	default:
		if h.storage.CheckUserExist(userId) {
			noloc := &telebot.ReplyMarkup{}
			btnBack := noloc.Data(Messages["en"]["back"], "back_to_user_menu")
			noloc.Inline(noloc.Row(btnBack))
			return c.Send("Unknown status", noloc)
		}
		return c.Send("Unknown status")
	}
	return nil
}

func (h *handlers) ChangeLanguage(c telebot.Context) error {
	userID := c.Sender().ID
	lang, err := h.storage.GetLangUser(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	menu := &telebot.ReplyMarkup{}
	btnEN := menu.Data("English🇬🇧", "language_change", "en")
	btnRU := menu.Data("Русский🇷🇺", "language_change", "ru")
	btnUZ := menu.Data("O'zbek🇺🇿", "language_change", "uz")
	btnBack := menu.Data(Messages["en"]["back"], "back_to_user_menu")

	menu.Inline(
		menu.Row(btnEN, btnRU, btnUZ),
		menu.Row(btnBack),
	)
	c.Edit(Messages[lang]["language_prompt"], menu)

	return nil
}

func (h *handlers) SetChangeLang(c telebot.Context) error {
	userID := c.Sender().ID
	lang := c.Callback().Data
	h.storage.ChangeLangUser(userID, lang)
	menu := &telebot.ReplyMarkup{}
	btnEN := menu.Data("English🇬🇧", "language_change", "en")
	btnRU := menu.Data("Русский🇷🇺", "language_change", "ru")
	btnUZ := menu.Data("O'zbek🇺🇿", "language_change", "uz")
	btnBack := menu.Data(Messages[lang]["back"], "back_to_user_menu")

	menu.Inline(
		menu.Row(btnEN, btnRU, btnUZ),
		menu.Row(btnBack),
	)
	c.Edit(Messages[lang]["language_prompt"], menu)
	return nil
}

func formatOrder(order *models.OrderDetails, lang string) string {
	items := ""
	for _, item := range order.Items {
		switch lang {
		case "uz":
			items += fmt.Sprintf("*%s* X *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_uz), item.Quantity)
		case "ru":
			items += fmt.Sprintf("*%s* X *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_ru), item.Quantity)
		case "en":
			items += fmt.Sprintf("*%s* X *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_en), item.Quantity)
		}
	}

	switch order.Payment_type {
	case "cash":
		order.Payment_type = Messages[lang]["cash"]
	case "card":
		order.Payment_type = Messages[lang]["card"]
	}

	if order.Status == "pending" {
		return fmt.Sprintf(Messages[lang]["order_msg"], order.Order_number, helpers.EscapeMarkdownV2(order.Address.Name_uz), items, order.TotalPrice, order.Delivery_type, order.TotalPrice, order.Payment_type, helpers.EscapeMarkdownV2(Messages[lang]["pending"])) + helpers.EscapeMarkdownV2(Messages[lang]["our_card"])
	}
	status := ""
	switch order.Status {
	case "preparing":
		status = Messages[lang]["preparing"]
	case "deliver":
		status = Messages[lang]["deliver"]
	case "completed":
		status = Messages[lang]["complete"]
	}

	return fmt.Sprintf(Messages[lang]["order_msg"], order.Order_number, helpers.EscapeMarkdownV2(order.Address.Name_uz), items, order.TotalPrice, order.Delivery_type, order.TotalPrice, order.Payment_type, helpers.EscapeMarkdownV2(status))
}

func (h *handlers) ShowUserOrders(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	lang, err := h.storage.GetLangUser(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	orders, err := h.storage.GetOrderByUserID(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	message := ""
	if len(*orders) == 0 {
		menu := &telebot.ReplyMarkup{}
		btnBack := menu.Data(Messages[lang]["back"], "back_to_user_menu")
		menu.Inline(
			// menu.Row(btnReOrder),
			menu.Row(btnBack),
		)
		options := &telebot.SendOptions{
			ParseMode:   telebot.ModeMarkdownV2,
			ReplyMarkup: menu,
		}
		message = "❌❌❌"
		c.Send(message, options)
	} else {
		message = "Your orders:\n"
		for _, order := range *orders {
			message += formatOrder(&order, lang)
			menu := &telebot.ReplyMarkup{}
			// btnReOrder := menu.Data(Messages[lang]["re-order"], "re_order", order.OrderID)
			btnBack := menu.Data(Messages[lang]["back"], "back_to_user_menu")
			menu.Inline(
				// menu.Row(btnReOrder),
				menu.Row(btnBack),
			)
			options := &telebot.SendOptions{
				ParseMode:   telebot.ModeMarkdownV2,
				ReplyMarkup: menu,
			}

			c.Send(message, options)
			message = ""
		}
	}
	return nil
}

func (h *handlers) CompleteOrder(c telebot.Context) error {
	userID := c.Sender().ID
	payment_type := c.Callback().Data

	orderID, err := h.storage.CreateOrder(userID, payment_type)
	if err != nil {
		return c.Send(err.Error())
	}
	lang, err := h.storage.GetLangUser(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	orderDetails, err := h.storage.GetOrderDetailsByOrderID(orderID)
	if err != nil {
		return c.Send(err.Error())
	}
	message := helpers.EscapeMarkdownV2(Messages[lang]["succsess"]) + formatOrder(orderDetails, lang)
	menu := &telebot.ReplyMarkup{}
	btnBack := menu.Data(Messages[lang]["back"], "back_to_user_menu")
	menu.Inline(menu.Row(btnBack))
	options := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: menu,
	}
	mark := message
	err = c.Edit(mark, options)
	if err != nil {
		return c.Send(err.Error())
	}
	return h.SendOrderToGroup(c, orderDetails)
}

func formatGroupOrder(order *models.OrderDetails, lang string) string {
	items := ""
	for _, item := range order.Items {
		switch lang {
		case "uz":
			items += fmt.Sprintf("*%s*  X  *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_uz), item.Quantity)
		case "ru":
			items += fmt.Sprintf("*%s*  X  *%v \n\n", helpers.EscapeMarkdownV2(item.Name_ru), item.Quantity)
		case "en":
			items += fmt.Sprintf("*%s*  X  *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_en), item.Quantity)
		}
	}

	switch order.Payment_type {
	case "cash":
		order.Payment_type = Messages[lang]["cash"]
	case "card":
		order.Payment_type = Messages[lang]["card"]
	}
	status := ""
	switch order.Status {
	case "pending":
		status = Messages[lang]["pending"]
	case "preparing":
		status = Messages[lang]["preparing"]
	case "deliver":
		status = Messages[lang]["deliver"]
	case "completed":
		status = Messages[lang]["complete"]
	}

	msg := fmt.Sprintf("👆👆\n\\#*%d*\n\n📞 %s\n", order.Daily_order_number, helpers.EscapeMarkdownV2(order.PhoneNumber))

	return msg + fmt.Sprintf(Messages[lang]["order_msg"], order.Order_number, helpers.EscapeMarkdownV2(order.Address.Name_uz), items, order.TotalPrice, order.Delivery_type, order.TotalPrice, order.Payment_type, helpers.EscapeMarkdownV2(status))
}

func (h *handlers) SendOrderToGroup(c telebot.Context, order *models.OrderDetails) error {
	location := &telebot.Location{
		Lat: order.Address.Latitude,  // Example latitude (New York)
		Lng: order.Address.Longitude, // Example longitude
	}
	_, err := c.Bot().Send(telebot.ChatID(groupID), location)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var btnChangeStatus telebot.Btn
	markup := &telebot.ReplyMarkup{}
	switch order.Status {
	case "pending":
		btnChangeStatus = markup.Data("To'lov o'tdi✅", "change_status_preparing", order.OrderID)
	case "preparing":
		btnChangeStatus = markup.Data("Yo'lga Chiqdi🚶", "change_status_deliver", order.OrderID)
	case "deliver":
		btnChangeStatus = markup.Data("Yetkazib berildi✅", "change_status_completed", order.OrderID)
	}
	markup.Inline(markup.Row(btnChangeStatus))

	fmt.Println(order.Status)

	msg := formatGroupOrder(order, "uz")

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: markup,
	}
	_, err = c.Bot().Send(telebot.ChatID(groupID), msg, option)
	return err
}

func (h *handlers) ChangeOrderStatus(c telebot.Context) error {
	orderID := c.Callback().Data
	var btnChangeStatus telebot.Btn
	markup := &telebot.ReplyMarkup{}
	switch c.Callback().Unique {
	case "change_status_preparing":
		h.storage.ChangeOrderStatus(orderID, "preparing")
		btnChangeStatus = markup.Data("Yo'lga chiqish", "change_status_deliver", orderID)
	case "change_status_deliver":
		h.storage.ChangeOrderStatus(orderID, "deliver")
		btnChangeStatus = markup.Data("Yetkazib berildi✅", "change_status_completed", orderID)
	case "change_status_completed":
		h.storage.ChangeOrderStatus(orderID, "completed")
	default:
		return c.Send("Unknown status")
	}
	markup.Inline(markup.Row(btnChangeStatus))
	order, err := h.storage.GetOrderDetailsByOrderID(orderID)

	if err != nil {
		return c.Send(err.Error())
	}
	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: markup,
	}
	// Меняем статус заказа
	updatedMsg := formatGroupOrder(order, "uz")

	// Обновляем сообщение с новым статусом
	err = c.Edit(updatedMsg, option)
	if err != nil {
		return c.Send("Failed to update order status")
	}

	return c.Respond(&telebot.CallbackResponse{Text: "Order status updated!"})
}

func (h *handlers) ChoosePaymentType(c telebot.Context) error {
	c.Send(&telebot.ReplyMarkup{RemoveKeyboard: true})

	c.Delete()
	lang, err := h.storage.GetLangUser(c.Sender().ID)
	if err != nil {
		return c.Send(err.Error())
	}
	markup := &telebot.ReplyMarkup{}
	btnCash := markup.Data(Messages[lang]["cash"], "payment_type", "cash")
	btnCard := markup.Data(Messages[lang]["card"], "payment_type", "card")

	markup.Inline(markup.Row(btnCash, btnCard))
	c.Send(Messages[lang]["payment_type_msg"], markup)

	return nil
}

// return nil

// func (h handlers) ShowProductByID(c telebot.Context, productID string) error {
// 	c.Bot().Handle(telebot.OnText, func(c telebot.Context) error {
// 		productID := c.Text()
// 		userID := c.Sender().ID

// 		if productID == "Back" {
// 			return h.ShowMenu(c)
// 		}

// 		_, err := h.storage.GetLangUser(userID)

// 		if err != nil {
// 			return c.Send(err.Error())
// 		}

// 		// Assuming you have a way to get the product details by ID
// 		product, err := h.storage.GetProductByName(productID) // Replace with your actual function

// 		if err != nil {
// 			return c.Send(err.Error())
// 		}
// 		c.Bot().Handle(&telebot.InlineButton{Unique: "increment"}, func(c telebot.Context) error {
// 			quantity, _ := strconv.Atoi(c.Data())
// 			return sendProductMenu(c, product, quantity+1)
// 		})

// 		// Обработчик уменьшения количества
// 		c.Bot().Handle(&telebot.InlineButton{Unique: "decrement"}, func(c telebot.Context) error {
// 			quantity, _ := strconv.Atoi(c.Data())
// 			if quantity > 1 {
// 				return sendProductMenu(c, product, quantity-1)
// 			}
// 			return nil
// 		})

// 		// Обработчик добавления в корзину
// 		c.Bot().Handle(&telebot.InlineButton{Unique: "add_to_cart"}, func(c telebot.Context) error {
// 			// quantity, _ := strconv.Atoi(c.Data())
// 			// userID := c.Sender().ID

// 			// // Инициализируем корзину, если её нет
// 			// if userCarts[userID] == nil {
// 			// 	userCarts[userID] = &UserCart{}
// 			// }
// 			// cart := userCarts[userID]

// 			// // Добавляем товар в корзину
// 			// cart.Items = append(cart.Items, CartItem{Product: product, Quantity: quantity})

// 			return c.Send("Товар добавлен в корзину. Выберите следующий шаг.")
// 		})

// 		// message := ""
// 		// switch lang {
// 		// case "uz":
// 		// 	message = fmt.Sprintf("Mahsulot nomi: %s\nNarxi: %f\nTavsifi: %s", product.Name_uz, product.Price, product.Description)
// 		// case "ru":
// 		// 	message = fmt.Sprintf("Название товара: %s\nЦена: %f\nОписание: %s", product.Name_ru, product.Price, product.Description)
// 		// case "en":
// 		// 	message = fmt.Sprintf("Product name: %s\nPrice: %f\nDescription: %s", product.Name_en, product.Price, product.Description)
// 		// default:
// 		// 	message = fmt.Sprintf("Mahsulot nomi: %s\nNarxi: %f\nTavsifi: %s", product.Name_uz, product.Price, product.Description)
// 		// }

// 		// c.Send(message)
// 		return nil

// 	})
// 	return nil
// }

// // Функция отображения меню товара
// func sendProductMenu(c telebot.Context, product *models.Product, quantity int) error {
// 	totalPrice := int(product.Price) * quantity

// 	// Создаем клавиатуру
// 	markup := &telebot.ReplyMarkup{}
// 	btnDecrement := markup.Data("-", "decrement", strconv.Itoa(quantity))
// 	btnIncrement := markup.Data("+", "increment", strconv.Itoa(quantity))
// 	btnAddToCart := markup.Data("Savatga qo'shish", "add_to_cart", strconv.Itoa(quantity))
// 	markup.Inline(
// 		markup.Row(btnDecrement, telebot.Btn{Text: strconv.Itoa(quantity)}, btnIncrement),
// 		markup.Row(btnAddToCart),
// 	)

// 	message := fmt.Sprintf("**%s**\n\nЦена за штуку: %d UZS\nКоличество: %d\nИтого: %d UZS",
// 		product.Name_ru, product.Price, quantity, totalPrice)

// 	return c.Edit(message, markup)
// }

// func (h handlers) ShowMenu(telegram_id int64) {
// 	cat, err := h.storage.GetAllCategories()

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}
// 	lang, err := h.storage.GetLangUser(telegram_id)

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}

// 	message := "Menu \n\n"
// 	switch lang {
// 	case "uz":
// 		for _, category := range cat.Categories {
// 			message += category.Name_uz + "\n"
// 		}
// 	case "ru":
// 		for _, category := range cat.Categories {
// 			message += category.Name_ru + "\n"
// 		}
// 	case "en":
// 		for _, category := range cat.Categories {
// 			message +=category.Name_en + "\n"
// 		}
// 	default:
// 		for _, category := range cat.Categories {
// 			message += category.Name_uz + "\n"
// 		}
// 	}
// 	h.tg.SendMessages(message, telegram_id)
// }

// func (h handlers) RegisterUser(updates *tgbotapi.UpdatesChannel,telegram_id int64) {
// 	adm:= h.storage.CheckAdmin(telegram_id)
// 	if adm {
// 		h.tg.SendMessages("You are admin", telegram_id)
// 		return
// 	}
// 	user := models.User{}

// 	replyMarkup := tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButtonContact("Share Contact"),
// 		),
// 	)

// 	h.tg.SendReplyKeyboard("Please share your phone number!!!", telegram_id, replyMarkup)

// 	for update := range *updates {
// 		if update.Message != nil {
// 			if update.Message.Contact != nil {
// 				user.Phone_Number = update.Message.Contact.PhoneNumber
// 				break
// 			}
// 		}
// 	}
// 	msg := tgbotapi.NewMessage(telegram_id, "Name:")
// 	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // Remove the keyboard
// 	h.tg.SendMsg(msg)

// 	for update := range *updates {
// 		if update.Message != nil {
// 			user.Name = update.Message.Text
// 			user.TelegramID = telegram_id
// 			user.Username = update.Message.From.UserName
// 			break
// 		}
// 	}
// 	err := h.storage.RegisterUser(&user)

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}
// }

// func (h handlers) GetCart(telegram_id int64, product_id string, quantity int, price float64) {
// 	cart, err := h.storage.GetCart(telegram_id)

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}
// 	msg := ""

// 	for i, item := range cart.Items {
// 		price += item.Price

// 		msg += fmt.Sprintf("%v. %vx - %s \n", i+1, item.Quantity, item.ProductName)
// 	}

// 	msg = msg + fmt.Sprintf(" \n Summa: %v", price)

// 	var cartConfirm = tgbotapi.NewInlineKeyboardButtonData("confirm", "/confirmCart")
// 	var cartBtn = tgbotapi.NewInlineKeyboardMarkup(
// 		tgbotapi.NewInlineKeyboardRow(cartConfirm),
// 	)

// 	h.tg.SendMessageWithInlineButton(msg, telegram_id, cartBtn)
// }

// func (h handlers) GetProductsByCateg(telegram_id int64, category_id int, quantity int, price float64) {
// 	product, err := h.storage.GetProductsByCategory(category_id)

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}

// 	var itm []string = []string{}

// 	lang, err := h.storage.GetLangUser(telegram_id)

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}

// 	switch lang {
// 	case "uz":
// 		itm = append(itm, messages.CartUz)
// 	case "ru":
// 		itm = append(itm, messages.CartRU)
// 	case "en":
// 		itm = append(itm, messages.CartEN)
// 	default:
// 		itm = append(itm, messages.CartRU)
// 	}

// 	for _, item := range product.Products {
// 		itm = append(itm, item.Name)
// 	}

// 	keyboaed := h.tg.CreateReplyKeyboard(itm, 2)
// 	h.tg.SendReplyKeyboard("Choose", telegram_id, keyboaed)
// }

// func (h handlers) GetAllCategories(telegram_id int64) {
// 	catgs, err := h.storage.GetAllCategories()

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}

// 	var categories []string = []string{}

// 	lang, err := h.storage.GetLangUser(telegram_id)

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}

// 	switch lang {
// 		case "uz":
// 			categories = append(categories, messages.CartUz)
// 		case "ru":
// 			categories = append(categories, messages.CartRU)
// 		case "en":
// 			categories = append(categories, messages.CartEN)
// 		default:
// 			categories = append(categories, messages.CartRU)
// 	}

// 	for _, category := range catgs.Categories {
// 		categories = append(categories, category.Name_uz)
// 	}

// 	keyboard := h.tg.CreateReplyKeyboard(categories, 2)
// 	h.tg.SendReplyKeyboard("Choose:", telegram_id, keyboard)
// }

// func (h handlers) GetProduct(telegram_id int64, product_id string) {

// }

// func (h handlers) ChangeLang(telegram_id int64, lang string) {
// 	_, err := h.storage.ChangeLangUser(telegram_id, lang)

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}
// }

// // var minus = tgbotapi.NewInlineKeyboardButtonData("-", "/minus")
// // var amount =
// // var plus = tgbotapi.NewInlineKeyboardButtonData("+", "/plus")
// // var cartBtn = tgbotapi.NewInlineKeyboardMarkup(
// // 	tgbotapi.NewInlineKeyboardRow(ClintBtn1, ClintBtn2),
// // 	tgbotapi.NewInlineKeyboardRow(ClintBtn3),
// // )
