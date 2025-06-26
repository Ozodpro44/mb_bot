package handlers

import (
	"bot/lib/helpers"
	"bot/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

const groupID = int64(-4720688028)

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
		"cart_cleared":       "Cart cleared🧹",
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
		"our_card":           "\nMake payment to this card 👇\n\n9860230107788070 \nZ******** O******",
		"closed_msg":         "We are closed for today.😔",
		"note":               "Enter additional data ✍️\n(For example: Apartment number, Order comment...)",
		"no_need_note":       "No need",
		"cancel_order":       "Cancel 🚫",
		"canceled":           "Order canceled🚫",
		"wait_msg":           "Almost there… just a bit more magic! ✨",
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
		"cart_cleared":       "Корзина очищена🧹",
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
		"our_card":           "\nПроизведите оплату на эту карту 👇\n\n9860230107788070 \nZ******** O******",
		"closed_msg":         "Сегодня заведение закрыто😔",
		"note":               "Введите дополнительные данные✍️\n(Например: Номер квартиры, Комментарий к заказу ...)",
		"no_need_note":       "Не нужно",
		"cancel_order":       "Отменить 🚫",
		"canceled":           "Заказ отменен🚫",
		"wait_msg":           "Почти там… ещё чуть-чуть магии! ✨",
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
		"cart_cleared":       "Savat tozalandi🧹",
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
		"our_card":           "\nTo'lovni shu kartaga qiling 👇\n\n9860230107788070 \nZ******** O******",
		"closed_msg":         "Buguncha yopildik😔",
		"note":               "Qoshimcha ma'lumot kriting✍️\n(Masalan: Kvartira raqami, Qoshimcha telefon nomer ...):",
		"no_need_note":       "Kerak emas",
		"cancel_order":       "Bekor qilish 🚫",
		"canceled":           "Buyurtma bekor qilindi🚫",
		"wait_msg":           "Deyarli yetib keldik… yana biroz sehr! ✨",
	},
	"tr": {
		"welcome":            "Hoş geldiniz! Lütfen adınızı girin:",
		"phone":              "Lütfen aşağıdaki düğmeden telefon numaranızı gönderin 👇:",
		"done":               "Kayıt tamamlandı! Hoş geldin, ",
		"exists":             "Zaten kayıtlısınız!",
		"error":              "Bir hata oluştu. Lütfen daha sonra tekrar deneyin.",
		"language_prompt":    "Lütfen dilinizi seçin:",
		"lang_btn":           "🇹🇷Dil",
		"order_btn":          "🛍Sipariş ver",
		"get_phone":          "📱 Telefon numaranızı paylaşın",
		"my_orders":          "Siparişlerim",
		"about_us":           "Hakkımızda",
		"back":               "⬅️Geri",
		"cart":               "🛒Sepet",
		"add_to_cart":        "📥Sepete ekle",
		"clear_cart":         "♻️ Temizle",
		"cart_cleared":       "Sepet temizlendi🧹",
		"cart_messsage":      "*%s*\n\nFiyat: %d TL\nMiktar: %d\nToplam: %d TL",
		"empty_cart":         "Sepetiniz boş🛒🚫",
		"user_menu":          "*Ana Menü:*\n\nAşağıdaki seçeneklerden birini seçin",
		"cart_items_msg":     "*%s* x %d \\= %d TL\n",
		"cart_total":         "\n*Toplam:* %d TL",
		"confirm_order":      "✅Siparişi Onayla",
		"continue_order":     "🧾Siparişi Devam Ettir",
		"added_to_cart":      "Ürün sepete eklendi✅",
		"order_msg":          "📋 *Sipariş numarası*: %d \n🚕 *Teslimat türü*: Teslimat \n🏠 *Adres*: %s \n📍 *Şube*: Yakkasaroy \n\n %s \n\n💵 *Ürünler*: %v \n🚚 *Teslimat ücreti*: %s \n💰 *Toplam*: %v \nÖdeme türü: %s \nDurum: %s",
		"delivery":           "Teslimat🚚",
		"pickup":             "Çekim🚶‍♂️",
		"re-order":           "Tekrar Sipariş Et🔄",
		"location_btn":       "📍 Konum Gönder",
		"location_msg":       "Lütfen aşağıdaki düğmeye konumunuzu gönderin 👇:",
		"true-location":      "✅Doğru",
		"false-location":     "❌Yanlış",
		"check-location-msg": "Bu konum doğru mu?: \n%s",
		"invalid_phone":      "Geçersiz telefon numarası tipi, lütfen tekrar deneyin 👇",
		"pending":            "Ödeme bekleniyor...",
		"preparing":          "Hazırlanıyor...",
		"deliver":            "Teslim ediliyor...",
		"complete":           "Teslim edildi✅",
		"payment_type_msg":   "Ödeme türünü seçin:",
		"cash":               "💵Nakit",
		"card":               "💳Kart",
		"thanks":             "Teşekkürler!",
		"no_need":            "Konum gerekmiyor❗",
		"succsess":           "Siparişiniz başarıyla oluşturuldu\n",
		"our_card":           "\nBu karta ödeme yapın 👇\n\n9860230107788070 \nZ******** O******",
		"closed_msg":         "Bugün için kapalıyız😔",
		"note":               "Ek bilgilerinizi girin ✍️(Örneğin: Daire numarası, Sipariş notu...)",
		"no_need_note":       "Gerekmiyor",
		"cancel_order":       "İptal Et 🚫",
		"canceled":           "Sipariş iptal edildi🚫",
		"wait_msg":           "Neredeyse geldik… biraz daha sihir kaldı! ✨",
	},
}

var (
	registerStep = make(map[int64]string) // Карта для отслеживания шага регистрации
	tempUserData = make(map[int64]map[string]string)
)
var (
	userLocation = make(map[int64]bool) // Сохраняем выбранный язык пользователя
)

var lastMsg = make(map[int64]*telebot.Message)

// var lastMsg []*telebot.Message
// var lastMsges []telebot.Message

func (h *handlers) HandleLanguage(c telebot.Context) error {
	userID := c.Sender().ID

	exists := h.storage.CheckUserExist(userID)

	// err := helpers.SendEskizSMS("+998770707041", "Bu Eskiz dan test 880")
	// if err != nil {
	// 	return err
	// }

	if exists {
		return h.ShowUserMenu(c)
	}
	if h.storage.CheckAdmin(userID) {
		return h.ShowAdminPanel(c)
	}

	menu := &telebot.ReplyMarkup{}
	btnEN := menu.Data("English🇬🇧", "language_add", "en")
	btnRU := menu.Data("Русский🇷🇺", "language_add", "ru")
	btnUZ := menu.Data("O'zbek🇺🇿", "language_add", "uz")
	btnTR := menu.Data("Türk🇹🇷", "language_add", "tr")

	menu.Inline(
		menu.Row(btnRU, btnUZ),
		menu.Row(btnEN, btnTR),
	)
	return c.Send(Messages["en"]["language_prompt"], menu)

}

func (h *handlers) GetUserName(c telebot.Context) error {
	c.Delete()
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

	if lastMsg[c.Chat().ID] != nil {
	// 	lastMsg[c.Chat().ID], _ = c.Bot().Edit(lastMsg[c.Chat().ID], Messages[lang]["wait_msg"])
		c.Bot().Delete(lastMsg[c.Chat().ID])
	}

	// var msg = &telebot.Message{}
	// if err != nil {
	// 	// msg, _ = c.Bot().Send(c.Recipient(),"⏲️...")
	// }

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

	// order, err := h.storage.GetOrderByUserID(c.Sender().ID)
	// if err != nil {
	// 	return c.Send(err.Error())
	// }
	// if len(*order) > 0 {
	// 	for _, ord := range *order {
	// 		h.storage.SetOrderMsg(ord.OrderID, 0)
	// 	}
	// }
	// if msg != nil {
	// 	c.Bot().Delete(msg)
	// 	msg = nil
	// }
	// if lastMsg[c.Chat().ID] != nil {
	// 	c.Bot().Delete(lastMsg[c.Chat().ID])
	// }

	// lastMsg[c.Chat().ID], err = c.Bot().Edit(lastMsg[c.Chat().ID], message, options)

	lastMsg[c.Chat().ID], err = c.Bot().Send(c.Recipient(), message, options)
	return err
}

func (h *handlers) SendAboutUs(c telebot.Context) error {
	user_id := c.Sender().ID
	lang, err := h.storage.GetLangUser(user_id)
	if err != nil {
		return c.Send(err.Error())
	}
	// if lastMsg[c.Chat().ID] != nil {
	// 	lastMsg[c.Chat().ID], _ = c.Bot().Edit(lastMsg[c.Chat().ID], Messages[lang]["wait_msg"])
	// } else {
	// 	lastMsg[c.Chat().ID], _ = c.Bot().Send(c.Recipient(), Messages[lang]["wait_msg"])
	// }
	// // var msg = &telebot.Message{}
	// if err != nil {
	// 	// msg, _ = c.Bot().Send(c.Recipient(),"⏲️...")
	// }
	markup := &telebot.ReplyMarkup{}
	btnBack := markup.Data(Messages[lang]["back"], "back_to_user_menu")
	markup.Inline(markup.Row(btnBack))

	lastMsg[c.Chat().ID], err = c.Bot().Edit(lastMsg[c.Chat().ID], "MB DONER, Yakkasaroy tumani, 49/1", markup)
	return err
}

func (h *handlers) HandleRegistrationSteps(c telebot.Context) error {
	userID := c.Sender().ID
	username := c.Sender().Username
	var text string
	if c.Message().Contact != nil {
		text = c.Message().Contact.PhoneNumber
	} else {
		text = c.Message().Text
	}

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

	phone, err := helpers.FormatPhoneNumber(text)

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
	}

	loc := btn.Location(Messages[lang]["location_btn"])

	btn.Reply(
		btn.Row(loc),
	)

	option := &telebot.SendOptions{
		ReplyMarkup: btn,
		ParseMode:   telebot.ModeMarkdownV2,
	}

	return c.Send(Messages[lang]["location_msg"], option)
}

func (h *handlers) HandleLocation(c telebot.Context) error {
	userID := c.Sender().ID

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
	} else {
		return c.Send("Location not found")
	}

	c.Send(Messages[lang]["thanks"], &telebot.ReplyMarkup{
		RemoveKeyboard: true,
	})

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

	// h.storage.GetLangUser(telegramID)

	lang, err := h.storage.GetLangUser(telegramID)
	if err != nil {
		return c.Send(fmt.Sprintf("Error: %s", err.Error()))
	}

	// if lastMsg[c.Chat().ID] != nil {
	// 	lastMsg[c.Chat().ID], err = c.Bot().Edit(lastMsg[c.Chat().ID], Messages[lang]["wait_msg"])
	// }
	// var msg = &telebot.Message{}
	if err != nil {
		// msg, _ = c.Bot().Send(c.Recipient(),"⏲️...")
	}
	// Fetch all categories
	cat, err := h.storage.GetAllCategories()
	if err != nil {
		return c.Send(fmt.Sprintf("Error: %s", err.Error()))
	}

	// Fetch user's language

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
	case "tr":
		message = " *MB Doner*\n\nMenü\n\n"
		for i := 0; i < len(cat.Categories); i += 2 {
			if i+1 < len(cat.Categories) {

				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_tr, "get_category_by_id", cat.Categories[i].ID),
					menu.Data(cat.Categories[i+1].Name_tr, "get_category_by_id", cat.Categories[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(cat.Categories[i].Name_tr, "get_category_by_id", cat.Categories[i].ID),
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
	if lastMsg[c.Chat().ID] != nil {
		lastMsg[c.Chat().ID], err = c.Bot().Edit(lastMsg[c.Chat().ID], message, options)
	}
	if err != nil {
		lastMsg[c.Chat().ID], err = c.Bot().Send(c.Recipient(), message, options)
	}
	return err
}

func (h *handlers) ShowProducts(c telebot.Context) error {

	category := c.Callback().Data
	userID := c.Sender().ID

	lang, err := h.storage.GetLangUser(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	// lastMsg[c.Chat().ID], err = c.Bot().Edit(lastMsg[c.Chat().ID], Messages[lang]["wait_msg"])
	// if err != nil {
	// 	lastMsg[c.Chat().ID], _ = c.Bot().Send(c.Recipient(), Messages[lang]["wait_msg"])
	// }

	cat, err := h.storage.GetCategoryByID(category)

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
	case "tr":
		message = fmt.Sprintf("    *%s*  \n\nMahsulotlarni tanlang:", cat.Name_tr)
		for i := 0; i < len(products.Products); i += 2 {
			if i+1 < len(products.Products) {
				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_tr, "get_product_by_id", products.Products[i].ID),
					menu.Data(products.Products[i+1].Name_tr, "get_product_by_id", products.Products[i+1].ID)))
			} else {
				buttons = append(buttons, menu.Row(
					menu.Data(products.Products[i].Name_tr, "get_product_by_id", products.Products[i].ID),
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

		lastMsg[c.Chat().ID], err = c.Bot().Send(c.Recipient(), message, option)
		return err
	}
	lastMsg[c.Chat().ID], err = c.Bot().Edit(lastMsg[c.Chat().ID], message, option)
	if err != nil {
		lastMsg[c.Chat().ID], err = c.Bot().Send(c.Recipient(), message, option)
	}
	return err
}

func (h *handlers) ShowProductByID(c telebot.Context) error {
	// Handle text input (for productID)
	prod := c.Callback().Data
	// userID := c.Sender().ID
	// Get user language
	// lang, err := h.storage.GetLangUser(userID)
	// if err != nil {
	// 	return c.Send(err.Error())
	// }
	// lastMsg[c.Chat().ID], err = c.Bot().Edit(lastMsg[c.Chat().ID], Messages[lang]["wait_msg"])
	// if err != nil {
	// 	lastMsg[c.Chat().ID], _ = c.Bot().Send(c.Recipient(), Messages[lang]["wait_msg"])
	// }

	// Get product details by productID
	product, err := h.storage.GetProductById(prod)
	if err != nil {
		return c.Send(err.Error())
	}
	c.Delete()

	// Initially show product details with a quantity of 1
	return h.sendProductMenu(c, product, 1)
}

// Function to display product menu with quantity options
func (h *handlers) sendProductMenu(c telebot.Context, product *models.Product, quantity int) error {
	photoPath := product.Photo
	if _, err := os.Stat(photoPath); os.IsNotExist(err) {
		photoPath = "./photos/no_photo.jpg"
	}

	totalPrice := int(product.Price) * quantity
	lang, err := h.storage.GetLangUser(c.Sender().ID)
	if err != nil {
		return c.Send(err.Error())
	}

	// Inline кнопки
	markup := &telebot.ReplyMarkup{}
	btnDecrement := markup.Data("➖", "decrement", strconv.Itoa(quantity))
	btnIncrement := markup.Data("➕", "increment", strconv.Itoa(quantity))
	btnAddToCart := markup.Data(Messages[lang]["add_to_cart"], "add_to_cart", strconv.Itoa(quantity), product.ID)
	btnQuantity := markup.Data(strconv.Itoa(quantity), "ignore", strconv.Itoa(quantity))
	btnBack := markup.Data(Messages[lang]["back"], "back_to_products_menu", product.Category_id)
	btnCart := markup.Data(Messages[lang]["cart"], "show_cart")

	markup.Inline(
		markup.Row(btnDecrement, btnQuantity, btnIncrement),
		markup.Row(btnAddToCart),
		markup.Row(btnBack, btnCart),
	)

	// Текст описания
	var message string
	switch lang {
	case "uz":
		message = fmt.Sprintf(Messages[lang]["cart_messsage"], helpers.EscapeMarkdownV2(product.Name_uz), int(product.Price), quantity, totalPrice)
	case "ru":
		message = fmt.Sprintf(Messages[lang]["cart_messsage"], helpers.EscapeMarkdownV2(product.Name_ru), int(product.Price), quantity, totalPrice)
	case "en":
		message = fmt.Sprintf(Messages[lang]["cart_messsage"], helpers.EscapeMarkdownV2(product.Name_en), int(product.Price), quantity, totalPrice)
	case "tr":
		message = fmt.Sprintf(Messages[lang]["cart_messsage"], helpers.EscapeMarkdownV2(product.Name_tr), int(product.Price), quantity, totalPrice)
	default:
		message = fmt.Sprintf(Messages["uz"]["cart_messsage"], helpers.EscapeMarkdownV2(product.Name_uz), int(product.Price), quantity, totalPrice)
	}

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
		ParseMode:   telebot.ModeMarkdownV2,
	}

	// Если кнопка +/- — редактируем только caption
	if c.Callback() != nil && (c.Callback().Unique == "decrement" || c.Callback().Unique == "increment") {
		if lastMsg[c.Chat().ID] != nil {
			_, err := c.Bot().EditCaption(lastMsg[c.Chat().ID], message, options)
			if err != nil {
				fmt.Println(err)
				return c.Send("❌ Не удалось обновить товар.")
			}
		}
		return h.HandleInlineButtons(c, product)
	}

	// 1. Сначала текст без фото
	msg, err := c.Bot().Send(c.Recipient(), message, options)
	if err != nil {
		return err
	}
	lastMsg[c.Chat().ID] = msg // Сохраняем сообщение

	// 2. Через 0.5 сек редактируем и добавляем фото
	go func(msg *telebot.Message) {
		time.Sleep(500 * time.Millisecond)
		photo := &telebot.Photo{
			File:    telebot.FromDisk(photoPath),
			Caption: message,
		}
		_, err := c.Bot().Edit(msg, photo, options)
		if err != nil {
			log.Println("Ошибка при замене на фото:", err)
		} else {
			lastMsg[c.Chat().ID] = msg // обновляем
		}
	}(msg)

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

		return c.Respond(&telebot.CallbackResponse{Text: Messages[lang]["added_to_cart"], ShowAlert: true})
	})
	return nil
}
func formatCart(cart *models.Cart, lang string) string {
	var message string
	totalPrice := 0

	switch lang {
	case "uz":
		for _, item := range cart.Items {
			itemTotal := int(item.Price) * item.Quantity
			message += fmt.Sprintf(Messages[lang]["cart_items_msg"], helpers.EscapeMarkdownV2(item.Name_uz), item.Quantity, itemTotal)
			totalPrice += itemTotal
		}
	case "ru":
		for _, item := range cart.Items {
			itemTotal := int(item.Price) * item.Quantity
			message += fmt.Sprintf(Messages[lang]["cart_items_msg"], helpers.EscapeMarkdownV2(item.Name_ru), item.Quantity, itemTotal)
			totalPrice += itemTotal
		}
	case "en":
		for _, item := range cart.Items {
			itemTotal := int(item.Price) * item.Quantity
			message += fmt.Sprintf(Messages[lang]["cart_items_msg"], helpers.EscapeMarkdownV2(item.Name_en), item.Quantity, itemTotal)
			totalPrice += itemTotal
		}
	case "tr":
		for _, item := range cart.Items {
			itemTotal := int(item.Price) * item.Quantity
			message += fmt.Sprintf(Messages[lang]["cart_items_msg"], helpers.EscapeMarkdownV2(item.Name_tr), item.Quantity, itemTotal)
			totalPrice += itemTotal
		}
	default:
		for _, item := range cart.Items {
			itemTotal := int(item.Price) * item.Quantity
			message += fmt.Sprintf(Messages["uz"]["cart_items_msg"], helpers.EscapeMarkdownV2(item.Name_uz), item.Quantity, itemTotal)
			totalPrice += itemTotal
		}

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
	c.Delete()
	lang, err := h.storage.GetLangUser(c.Sender().ID)
	if err != nil {
		return c.Send(err.Error())
	}

	// lastMsg[c.Chat().ID], err = c.Bot().Send(c.Recipient(), Messages[lang]["wait_msg"])
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// Assume 'cart' is fetched from the database or session
	cart, err := h.storage.GetCart(c.Sender().ID)
	if err != nil || len(cart.Items) == 0 {
		btn := &telebot.ReplyMarkup{}

		btnBack := btn.Data(Messages[lang]["back"], "back_to_categories")
		btn.Inline(btn.Row(btnBack))
		c.Delete()
		return c.Send(Messages[lang]["empty_cart"], btn)
	}

	message := formatCart(cart, lang)
	buttons := createCartButtons(cart, lang)

	op := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: buttons,
	}

	if lastMsg[c.Chat().ID] != nil {
		c.Bot().Delete(lastMsg[c.Chat().ID])
	}

	lastMsg[c.Chat().ID], err = c.Bot().Send(c.Recipient(), message, op)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func (h *handlers) HandleIncrement(c telebot.Context) error {
	c.Respond()
	productID := c.Callback().Data // Get product ID from callback data
	userID := c.Sender().ID
	err := h.storage.IncrementCartProductByID(userID, productID)

	if err != nil {
		return c.Send(err.Error())
	}

	lang, err := h.storage.GetLangUser(userID)
	if err != nil {
		return c.Send(err.Error())
	}

	// Fetch the user's cart
	cart, err := h.storage.GetCart(userID)
	if err != nil {
		return c.Send("Error fetching cart")
	}
	fmt.Println(cart)

	// Increment the quantity of the selected product
	// for i, item := range cart.Items {
	// 	if item.ProductID == productID {
	// 		cart.Items[i].Quantity++
	// 		break
	// 	}
	// }
	// fmt.Println("new", cart)

	// Update the cart in storage
	// err = h.storage.UpdateCart(userID, cart)
	// if err != nil {
	// 	return c.Send("Error updating cart")
	// }

	// Resend the updated cart
	message := formatCart(cart, lang)
	buttons := createCartButtons(cart, lang)

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: buttons,
	}
	fmt.Println(message)

	lastMsg[c.Chat().ID], err = c.Bot().Edit(lastMsg[c.Chat().ID], message, option)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func (h *handlers) HandleDecrement(c telebot.Context) error {
	c.Respond()
	productID := c.Callback().Data
	userID := c.Sender().ID
	err := h.storage.DecrementCartProductByID(userID, productID)

	if err != nil {
		return c.Send(err.Error())
	}

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
	for i, _ := range cart.Items {
		if cart.Items[i].Quantity < 1 {
			h.storage.RemoveFromCart(userID, productID)
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
		}
	}

	// Resend the updated cart or show an empty cart message if no items remain
	if len(cart.Items) == 0 {
		btn := &telebot.ReplyMarkup{}
		btnBack := btn.Data(Messages[lang]["back"], "back_to_categories")
		btn.Inline(btn.Row(btnBack))
		return c.Edit(Messages[lang]["empty_cart"], btn)
	}
	// err = h.storage.UpdateCart(userID, cart)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return c.Send("Error updating cart")
	// }

	message := formatCart(cart, lang)
	buttons := createCartButtons(cart, lang)

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: buttons,
	}

	lastMsg[c.Chat().ID], err = c.Bot().Edit(lastMsg[c.Chat().ID], message, option)
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

func (h *handlers) ClearCart(c telebot.Context) error {
	userID := c.Sender().ID
	err := h.storage.ClearCart(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	c.Respond(&telebot.CallbackResponse{Text: "Cart cleared❗", ShowAlert: true})
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

	lastMsg[c.Chat().ID], err = c.Bot().Send(c.Recipient(), Messages[lang]["phone"], keyboard)
	if err != nil {
		log.Println(err.Error())
	}
	return nil
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
		h.HandleRegistrationSteps(c)
	case "note":
		h.GetNoteFromUser(c)
	case "location":
		h.storage.SetDataUserMessageStatus(userId, text)
		return h.ShowMenu(c)
	case "update_cat_name_uz":
		h.UpdateCategoryNameUz(c)
		return nil
	case "update_cat_name_ru":
		h.UpdateCategoryNameRu(c)
		return nil
	case "update_cat_name_en":
		h.UpdateCategoryNameEn(c)
		return nil
	case "update_cat_name_tr":
		h.UpdateCategoryNameTr(c)
		return nil
	case "add_category":
		h.CreateCategory(c)
		return nil
	case "update_prod_name_uz":
		h.UpdateProductNameUz(c)
		return nil
	case "update_prod_name_ru":
		h.UpdateProductNameRu(c)
		return nil
	case "update_prod_name_en":
		h.UpdateProductNameEn(c)
		return nil
	case "update_prod_name_tr":
		h.UpdateProductNameTr(c)
		return nil
	case "update_prod_desc":
		h.UpdateProductDesc(c)
		return nil
	case "update_prod_price":
		h.UpdateProductPrice(c)
		return nil
	case "add_admin":
		h.AddAdmin(c)
		return nil
	case "admin":
		nol := &telebot.ReplyMarkup{}
		btnBack := nol.Data(Messages["en"]["back"], "back_to_admin_menu")
		nol.Inline(nol.Row(btnBack))
		return c.Send("Unknown status", nol)
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
	btnTR := menu.Data("Türk🇹🇷", "language_change", "tr")
	btnBack := menu.Data(Messages[lang]["back"], "back_to_user_menu")

	menu.Inline(
		menu.Row(btnRU, btnUZ),
		menu.Row(btnEN, btnTR),
		menu.Row(btnBack),
	)
	c.Edit(Messages[lang]["language_prompt"], menu)

	return nil
}

func (h *handlers) SetChangeLang(c telebot.Context) error {
	userID := c.Sender().ID
	lang := c.Callback().Data
	h.storage.ChangeLangUser(userID, lang)
	// menu := &telebot.ReplyMarkup{}
	// btnEN := menu.Data("English🇬🇧", "language_change", "en")
	// btnRU := menu.Data("Русский🇷🇺", "language_change", "ru")
	// btnUZ := menu.Data("O'zbek🇺🇿", "language_change", "uz")
	// btnTR := menu.Data("Türk🇹🇷", "language_change", "tr")
	// btnBack := menu.Data(Messages[lang]["back"], "back_to_user_menu")

	// menu.Inline(
	// 	menu.Row(btnRU, btnUZ),
	// 	menu.Row(btnEN, btnTR),
	// 	menu.Row(btnBack),
	// )
	// c.Edit(Messages[lang]["language_prompt"], menu)
	h.ShowUserMenu(c)
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
		case "tr":
			items += fmt.Sprintf("*%s* X *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_tr), item.Quantity)
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
	case "canceled":
		status = Messages[lang]["canceled"]
	}

	return fmt.Sprintf(Messages[lang]["order_msg"], order.Order_number, helpers.EscapeMarkdownV2(order.Address.Name_uz), items, order.TotalPrice, order.Delivery_type, order.TotalPrice, order.Payment_type, helpers.EscapeMarkdownV2(status))
}

type orderMessage struct {
	orders []*telebot.Message
}

var orderMessages = make(map[int64]orderMessage)

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
		btnBack := menu.Data(Messages[lang]["back"], "back_to_user_menu_from_orders")
		menu.Inline(
			menu.Row(btnBack),
		)
		options := &telebot.SendOptions{
			ParseMode:   telebot.ModeMarkdownV2,
			ReplyMarkup: menu,
		}
		message = "❌❌❌"
		tlbmsg, _ := c.Bot().Send(c.Recipient(), message, options)
		orderMessages[c.Chat().ID] = orderMessage{orders: []*telebot.Message{tlbmsg}}
		return nil
	} else {
		message = "Your orders:\n"
		for _, order := range *orders {
			message += formatOrder(&order, lang)
			menu := &telebot.ReplyMarkup{}
			btnBack := menu.Data(Messages[lang]["back"], "back_to_user_menu_from_orders")
			if order.Status == "pending" || order.Status == "preparing" {
				btnCancel := menu.Data(Messages[lang]["cancel_order"], "cancel_order", order.OrderID)
				menu.Inline(
					menu.Row(btnCancel),
					menu.Row(btnBack),
				)
			} else {
				menu.Inline(
					menu.Row(btnBack),
				)
			}
			options := &telebot.SendOptions{
				ParseMode:   telebot.ModeMarkdownV2,
				ReplyMarkup: menu,
			}

			msg, err := c.Bot().Send(&telebot.User{ID: userID}, message, options)
			if _, ok := orderMessages[userID]; !ok {
				orderMessages[userID] = orderMessage{orders: []*telebot.Message{}}
			}
			msgs := orderMessages[userID]
			msgs.orders = append(msgs.orders, msg)
			orderMessages[userID] = msgs
			if err != nil {
				fmt.Println(err)
			}
			err = h.storage.SetOrderMsg(order.OrderID, lastMsg[c.Chat().ID].ID)

			if err != nil {
				fmt.Println(err)
			}
			message = ""
		}
	}
	return nil
}

func (h *handlers) BackToMainMenuFromOrders(c telebot.Context) error {
	userID := c.Sender().ID

	if msgs, ok := orderMessages[userID]; ok {
		for _, msg := range msgs.orders {
			if msg != nil {
				c.Bot().Delete(msg)
			}
		}
		delete(orderMessages, userID)
	}

	return h.ShowUserMenu(c)
}

func (h *handlers) CompleteOrder(c telebot.Context) error {
	userID := c.Sender().ID
	payment_type := c.Callback().Data

	lang, err := h.storage.GetLangUser(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	c.Edit(Messages[lang]["wait_msg"])
	orderID, err := h.storage.CreateOrder(userID, payment_type)
	if err != nil {
		return c.Send(err.Error())
	}
	orderDetails, err := h.storage.GetOrderDetailsByOrderID(orderID)
	if err != nil {
		return c.Send(err.Error())
	}
	message := helpers.EscapeMarkdownV2(Messages[lang]["succsess"]) + formatOrder(orderDetails, lang)
	menu := &telebot.ReplyMarkup{}
	btnCancel := menu.Data(Messages[lang]["cancel_order"], "cancel_order", orderID)
	btnBack := menu.Data(Messages[lang]["back"], "back_to_user_menu_from_orders")
	menu.Inline(
		menu.Row(btnCancel),
		menu.Row(btnBack),
	)
	options := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: menu,
	}
	mark := message
	msg, err := c.Bot().Edit(c.Message(), mark, options)
	if err != nil {
		return c.Send(err.Error())
	}
	err = h.storage.SetOrderMsg(orderID, msg.ID)

	if err != nil {
		fmt.Println(err)
	}
	return h.SendOrderToGroup(c.Bot(), orderDetails)
}

func formatGroupOrder(order *models.OrderDetails, lang string) string {
	items := ""
	for _, item := range order.Items {
		switch lang {
		case "uz":
			items += fmt.Sprintf("*%s*  X  *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_uz), item.Quantity)
		case "ru":
			items += fmt.Sprintf("*%s*  X  *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_ru), item.Quantity)
		case "en":
			items += fmt.Sprintf("*%s*  X  *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_en), item.Quantity)
		case "tr":
			items += fmt.Sprintf("*%s*  X  *%v* \n\n", helpers.EscapeMarkdownV2(item.Name_tr), item.Quantity)
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
	case "canceled":
		status = Messages[lang]["canceled"]
	}

	msg := fmt.Sprintf("👆👆\n\\#*%d*\n\n📞 %s\n", order.Daily_order_number, helpers.EscapeMarkdownV2(order.PhoneNumber))

	return msg + fmt.Sprintf(Messages[lang]["order_msg"], order.Order_number, helpers.EscapeMarkdownV2(order.Address.Name_uz), items, order.TotalPrice, order.Delivery_type, order.TotalPrice, order.Payment_type, helpers.EscapeMarkdownV2(status))
}

func (h *handlers) SendOrderToGroup(c *telebot.Bot, order *models.OrderDetails) error {
	location := &telebot.Location{
		Lat: order.Address.Latitude,  // Example latitude (New York)
		Lng: order.Address.Longitude, // Example longitude
	}
	_, err := c.Send(telebot.ChatID(groupID), location)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var btnChangeStatus telebot.Btn
	var btnCancel telebot.Btn
	markup := &telebot.ReplyMarkup{}
	switch order.Status {
	case "pending":
		btnChangeStatus = markup.Data("To'lov o'tdi✅", "change_status_preparing", order.OrderID)
		btnCancel = markup.Data("Bekor qilish❌", "change_status_canceled", order.OrderID)
	case "preparing":
		btnChangeStatus = markup.Data("Yo'lga Chiqdi🚶", "change_status_deliver", order.OrderID)
		btnCancel = markup.Data("Bekor qilish❌", "change_status_canceled", order.OrderID)
	case "deliver":
		btnChangeStatus = markup.Data("Yetkazib berildi✅", "change_status_completed", order.OrderID)
		btnCancel = markup.Data("Bekor qilish❌", "change_status_canceled", order.OrderID)
	}
	markup.Inline(
		markup.Row(btnChangeStatus),
		markup.Row(btnCancel),
	)

	fmt.Println(order.Status)

	msg := formatGroupOrder(order, "uz")

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: markup,
	}
	mess, err := c.Send(telebot.ChatID(groupID), msg, option)
	h.storage.SetOrderGroupMsg(order.OrderID, mess.ID)
	if err != nil {
		fmt.Println(err)
	}
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
	case "change_status_canceled":
		h.storage.ChangeOrderStatus(orderID, "canceled")
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
	err = c.Edit(updatedMsg, option)
	if err != nil {
		return c.Send("Failed to update order status")
	}

	msg, err := h.storage.GetOrderMsg(orderID)
	fmt.Println(msg)
	if msg.MsgID != 0 {
		message := helpers.EscapeMarkdownV2(Messages[msg.Lang]["succsess"]) + formatOrder(order, msg.Lang)
		menu := &telebot.ReplyMarkup{}
		btnBack := menu.Data(Messages[msg.Lang]["back"], "back_to_user_menu_from_orders")
		menu.Inline(menu.Row(btnBack))
		options := &telebot.SendOptions{
			ParseMode:   telebot.ModeMarkdownV2,
			ReplyMarkup: menu,
		}

		if err != nil {
			return c.Send(err.Error())
		}
		c.Bot().Delete(&telebot.Message{
			ID:   msg.MsgID,
			Chat: &telebot.Chat{ID: msg.UserID},
		})
		msg, err := c.Bot().Send(telebot.ChatID(msg.UserID), message, options)
		if err != nil {
			fmt.Println(err)
		}
		h.storage.SetOrderMsg(orderID, msg.ID)
	}
	// m, err := h.storage.GetOrderMsg(order.OrderID)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// edit_msg := formatOrder(order, m.Lang)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// c.Edit(&telebot.Message{
	// 	ID:   m.MsgID,
	// 	Chat: &telebot.Chat{ID: m.UserID},
	// }, edit_msg)
	// Обновляем сообщение с новым статусом
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

func (h *handlers) GetNoteFromUser(c telebot.Context) error {
	c.Send(&telebot.ReplyMarkup{RemoveKeyboard: true})
	c.Delete()
	lang, err := h.storage.GetLangUser(c.Sender().ID)
	if err != nil {
		return c.Send(err.Error())
	}
	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Data(Messages[lang]["back"], "back_to_location")
	btnNoNeed := markup.Data(Messages[lang]["no_need_note"], "no_need_note")
	markup.Inline(markup.Row(btnBack, btnNoNeed))
	return c.Send(Messages[lang]["note"], markup)
}

func (h *handlers) CancelOrder(c telebot.Context) error {
	orderID := c.Callback().Data
	_, err := h.storage.ChangeOrderStatus(orderID, "canceled")
	if err != nil {
		return c.Send(err.Error())
	}
	c.Respond(&telebot.CallbackResponse{Text: "Order canceled!"})
	m, err := h.storage.GetOrderGroupMsg(orderID)
	if err != nil {
		return c.Send(err.Error())
	}
	c.Bot().Delete(&telebot.Message{
		ID:   m,
		Chat: &telebot.Chat{ID: groupID},
	})
	order, err := h.storage.GetOrderDetailsByOrderID(orderID)
	if err != nil {
		return c.Send(err.Error())
	}
	h.SendOrderToGroup(c.Bot(), order)
	return h.ShowUserOrders(c)
}
