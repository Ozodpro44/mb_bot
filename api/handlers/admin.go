package handlers

import (
	"bot/lib/helpers"
	"bot/models"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"gopkg.in/telebot.v3"
)

var AdminMessages = map[string]map[string]string{
	"en": {
		"add_category":         "Add Category ➕",
		"add_product":          "Add Product ➕",
		"add_admin":            "Add Admin 👨‍⚖️",
		"close_day":            "Close Day🔐",
		"open_day":             "Open Day🔓",
		"branch":               "Branch 🏠",
		"get_all_users":        "Get All Users 👥",
		"category":             "Category name: <UZ>:<RU>:<EN>:<TR>",
		"incorrect_category":   "⁉️Incorrect format. Use: <UZ>:<RU>:<EN>:<TR>",
		"category_created":     "Category created ✅",
		"back":                 "🔙Back",
		"product":              "Use: \n 🖼️ \n<product name uz>:<product name ru>:<product name en>:<product name tr>:<description>:<price>",
		"product_err":          "Incorrect format. Use: \n 🖼️ \n<product name uz>:<product name ru>:<product name en>:<product name tr>:<description>:<price>",
		"category_menu":        "Categories",
		"product_menu":         "Products",
		"delete_category":      "Delete category",
		"add_admin_msg":        "Add admin: ```<telegram_id>:<phone_number>:<password>```",
		"admin_created":        "Admin created ✅",
		"day_closed":           "Day closed",
		"day_opened":           "Day opened",
		"btn_upd_uz":           "Name in Uzbek",
		"btn_upd_ru":           "Name in Russian",
		"btn_upd_en":           "Name in English",
		"btn_upd_tr":           "Name in Turkish",
		"btn_upd_desc":         "Description",
		"btn_upd_price":        "Price",
		"btn_upd_avail":        "Availability",
		"show_category_info":   "Name in Uzbek: %s\nName in Russian: %s\nName in English: %s\nName in Turkish: %s\n\nAvailability: %t",
		"name_tr_msg":          "Enter Name TR:",
		"name_uz_msg":          "Enter Name UZ:",
		"name_ru_msg":          "Enter Name RU:",
		"name_en_msg":          "Enter Name EN:",
		"update_category":      "Update category",
		"update_product":       "Update product",
		"btn_yes":              "Yes✅",
		"btn_no":               "No❌",
		"cat_deleted":          "Category deleted ✅",
		"delete_cat_msg":       "Are you sure to delete category?",
		"name_uz_updated":      "Category name🇺🇿 updated",
		"name_ru_updated":      "Category name🇷🇺 updated",
		"name_en_updated":      "Category name🇬🇧 updated",
		"name_tr_updated":      "Category name🇹🇷 updated",
		"change_lang":          "Change language",
		"no_products":          "No products",
		"product_created":      "Product created ✅",
		"cancel_btn":           "Cancel✖️",
		"product_deleted":      "Product deleted ✅",
		"product_info":         "Name UZ: %s\nName RU: %s\nName EN: %s\nName TR: %s\nDescription: %s\nPrice: %d\nAvailability: %t",
		"price_msg":            "Enter price:",
		"price_err":            "Incorrect price",
		"prod_name_uz_updated": "Product name🇺🇿 updated",
		"prod_name_ru_updated": "Product name🇷🇺 updated",
		"prod_name_en_updated": "Product name🇬🇧 updated",
		"prod_name_tr_updated": "Product name🇹🇷 updated",
		"prod_name_uz_msg":     "Enter product name UZ:",
		"prod_name_ru_msg":     "Enter product name RU:",
		"prod_name_en_msg":     "Enter product name EN:",
		"prod_name_tr_msg":     "Enter product name TR:",
		"desc_msg":             "Enter product description:",
		"desc_updated":         "Description updated",
		"price_updated":        "Price updated",
		"no_categories":        "No categories",
		"delete_prod":          "Delete product",
		"btn_upd_cat":          "Update category",
		"prod_cat_updated":     "Product category updated",
		"send_adds":            "Send adds \n 🖼️ \ntext",
		"adds_btn":             "Adds",
		"adds_sent":            "Adds sent✅",
		"btn_upd_photo":        "🖼️Update Photo",
		"prod_photo_updated":   "Product photo updated✅",
		"prod_photo_msg":       "Send Photo 🖼️",
	},
	"ru": {
		"add_category":         "Добавить категорию ➕",
		"add_product":          "Добавить продукт ➕",
		"add_admin":            "Добавить админ 👨‍⚖️",
		"close_day":            "Закрыть день🔐",
		"open_day":             "Открыть день🔓",
		"branch":               "Филиал 🏠",
		"get_all_users":        "Получить всех пользователей 👥",
		"category":             "Название категории: <UZ>:<RU>:<EN>:<TR>",
		"incorrect_category":   "⁉️Неверный формат. Используйте: /category :<UZ>:<RU>:<EN>:<TR>",
		"category_created":     "Категория успешно создана ✅",
		"back":                 "🔙Назад",
		"product":              "Используйте: \n 🖼️ \n<имя продукта уз>:<имя продукта ру>:<имя продукта ен>:<имя продукта тр>:<описание>:<цена>",
		"product_err":          "Неверный формат. Используйте: \n 🖼️ \n<имя продукта уз>:<имя продукта ру>:<имя продукта ен>:<имя продукта тр>:<описание>:<цена>",
		"category_menu":        "Категории",
		"product_menu":         "Продукты",
		"delete_category":      "Удалить категорию",
		"add_admin_msg":        "Добавить админа: ```<telegram_id>:<phone_number>:<password>```",
		"admin_created":        "Админ успешно создан ✅",
		"day_closed":           "День закрыт",
		"day_opened":           "День открыт",
		"btn_upd_uz":           "Название на узбекском",
		"btn_upd_ru":           "Название на русском",
		"btn_upd_en":           "Название на английском",
		"btn_upd_tr":           "Название на турецком",
		"btn_upd_desc":         "Описание",
		"btn_upd_price":        "Цена",
		"btn_upd_avail":        "Доступность",
		"show_category_info":   "Название на узбекском: %s\nНазвание на русском: %s\nНазвание на английском: %s\nНазвание на турецком: %s\n\nДоступность: %t",
		"name_tr_msg":          "Введите имя TR:",
		"name_uz_msg":          "Введите имя UZ:",
		"name_ru_msg":          "Введите имя RU:",
		"name_en_msg":          "Введите имя EN:",
		"update_category":      "Обновить категорию",
		"update_product":       "Обновить продукт",
		"btn_yes":              "Да✅",
		"btn_no":               "Нет❌",
		"cat_deleted":          "Категория успешно удалена ✅",
		"delete_cat_msg":       "Вы уверены, что хотите удалить категорию?",
		"name_uz_updated":      "Название категории🇺🇿 обновлено",
		"name_ru_updated":      "Название категории🇷🇺 обновлено",
		"name_en_updated":      "Название категории🇬🇧 обновлено",
		"name_tr_updated":      "Название категории🇹🇷 обновлено",
		"change_lang":          "Изменить язык",
		"no_products":          "Нет продуктов",
		"product_created":      "Продукт успешно создан ✅",
		"cancel_btn":           "Отмена✖️",
		"product_deleted":      "Продукт успешно удален ✅",
		"product_info":         "Название на узбекском: %s\nНазвание на русском: %s\nНазвание на английском: %s\nНазвание на турецком: %s\nОписание: %s\nЦена: %d\nДоступность: %t",
		"price_msg":            "Введите цену:",
		"price_err":            "Неверная цена",
		"prod_name_uz_updated": "Название продукта🇺🇿 обновлено",
		"prod_name_ru_updated": "Название продукта🇷🇺 обновлено",
		"prod_name_en_updated": "Название продукта🇬🇧 обновлено",
		"prod_name_tr_updated": "Название продукта🇹🇷 обновлено",
		"prod_name_uz_msg":     "Введите название продукта UZ:",
		"prod_name_ru_msg":     "Введите название продукта RU:",
		"prod_name_en_msg":     "Введите название продукта EN:",
		"prod_name_tr_msg":     "Введите название продукта TR:",
		"desc_msg":             "Введите описание продукта:",
		"desc_updated":         "Описание продукта обновлено",
		"price_updated":        "Цена продукта обновлена",
		"no_categories":        "Категорий нет",
		"delete_prod":          "Удалить продукт",
		"btn_upd_cat":          "Обновить категорию",
		"prod_cat_updated":     "Категория продукта обновлена",
		"send_adds":            "Отправить рекламу \n 🖼️ \nтекст",
		"adds_btn":             "Реклама",
		"adds_sent":            "Реклама отправлена✅",
		"btn_upd_photo":        "🖼️Обновить Фото",
		"prod_photo_updated":   "Фото продукта успешно обновлено✅",
		"prod_photo_msg":       "Отправить Фото 🖼️",
	},
	"uz": {
		"admin_panel":          "Admin paneli: \nFilial nomi - %s \nOdamlar soni - %d \nKun - %s \nVaqt - %s \nSavdo - %d UZS",
		"add_category":         "Kategoriya qo'shish ➕",
		"add_product":          "Mahsulot qo'shish ➕",
		"add_admin":            "Admin qo'shish 👨‍⚖️",
		"close_day":            "Kunni yopish🔐",
		"open_day":             "Kunni ochish🔓",
		"branch":               "Filial 🏠",
		"get_all_users":        "Barcha foydalanuvchilarni olish 👥",
		"category":             "Kategoriya nomi: <UZ>:<RU>:<EN>:<TR>",
		"incorrect_category":   "⁉️Iltimos, to'g'ri formatdan foydalaning: /category: <UZ>:<RU>:<EN>:<TR>",
		"category_created":     "Kategoriya muvaffaqiyatli yaratildi ✅",
		"back":                 "🔙Orqaga",
		"product":              "Mahsulot qo'shish uchun: \n 🖼️ \n<nomi uz>:<nomi ru>:<nomi en>:<nomi tr>:<description>:<narxi>",
		"product_err":          "Iltimos, to'g'ri formatdan foydalaning: \n 🖼️ \n<nomi uz>:<nomi ru>:<nomi en>:<nomi tr>:<description>:<narxi>",
		"category_menu":        "Kategoriyalar",
		"product_menu":         "Mahsulotlar",
		"delete_category":      "Kategoriyani o'chirish",
		"add_admin_msg":        "Admin qo'shish: ```<telegram_id>:<phone_number>:<password>```",
		"admin_created":        "Admin muvaffaqiyatli yaratildi ✅",
		"day_closed":           "Kun yopildi",
		"day_opened":           "Kun ochildi",
		"btn_upd_uz":           "O'zbekcha nomini",
		"btn_upd_ru":           "Ruscha nomini",
		"btn_upd_en":           "Inglizcha nomini",
		"btn_upd_tr":           "Turkcha nomini",
		"btn_upd_desc":         "Tavsifini",
		"btn_upd_price":        "Narxini",
		"btn_upd_avail":        "Borligini",
		"show_category_info":   "O'zbekcha nom: %s\nRuscha nom: %s\nInglizcha nom: %s\nTurkcha nom: %s\n\nBorligi: %t",
		"name_tr_msg":          "Ismni TR kiriting:",
		"name_uz_msg":          "Ismni UZ kiriting:",
		"name_ru_msg":          "Ismni RU kiriting:",
		"name_en_msg":          "Ismni EN kiriting:",
		"update_category":      "Kategoriyani yangilash",
		"update_product":       "Mahsulotni yangilash",
		"btn_yes":              "Ha✅",
		"btn_no":               "Yo'q❌",
		"cat_deleted":          "Kategoriya muvaffaqiyatli o'chirildi ✅",
		"delete_cat_msg":       "Kategoriyani o‘chirib tashlashingizga ishonchingiz komilmi?",
		"name_uz_updated":      "Kategoriya nomi🇺🇿 yangilandi",
		"name_ru_updated":      "Kategoriya nomi🇷🇺 yangilandi",
		"name_en_updated":      "Kategoriya nomi🇬🇧 yangilandi",
		"name_tr_updated":      "Kategoriya nomi🇹🇷 yangilandi",
		"change_lang":          "Tilni o'zgartirish",
		"no_products":          "Mahsulotlar yo'q",
		"product_created":      "Mahsulot muvaffaqiyatli yaratildi ✅",
		"cancel_btn":           "Bekor qilish✖️",
		"product_deleted":      "Mahsulot muvaffaqiyatli o'chirildi ✅",
		"product_info":         "O'zbekcha nomi: %s\nRuscha nomi: %s\nInglizcha nomi: %s\nTurkcha nomi: %s\nTavsif: %s\nNarx: %d\nMavjudligi: %t",
		"price_msg":            "Narxni kiriting:",
		"price_err":            "Noto'g'ri narx",
		"prod_name_uz_updated": "Mahsulot nomi🇺🇿 yangilandi",
		"prod_name_ru_updated": "Mahsulot nomi🇷🇺 yangilandi",
		"prod_name_en_updated": "Mahsulot nomi🇬🇧 yangilandi",
		"prod_name_tr_updated": "Mahsulot nomi🇹🇷 yangilandi",
		"prod_name_uz_msg":     "Mahsulot nomini UZ kiriting:",
		"prod_name_ru_msg":     "Mahsulot nomini RU kiriting:",
		"prod_name_en_msg":     "Mahsulot nomini EN kiriting:",
		"prod_name_tr_msg":     "Mahsulot nomini TR kiriting:",
		"desc_msg":             "Mahsulot tavsifini kiriting:",
		"desc_updated":         "Mahsulot tavsifi yangilandi",
		"price_updated":        "Mahsulot narxi yangilandi",
		"no_categories":        "Kategoriyalar mavjud emas",
		"delete_prod":          "Mahsulotni o'chirish",
		"btn_upd_cat":          "Kategoriyani yangilash",
		"prod_cat_updated":     "Mahsulot kategoriyasi yangilandi",
		"send_adds":            "Reklama yuborish \n 🖼️ \nmatn",
		"adds_btn":             "Reklama",
		"adds_sent":            "Reklama yuborildi✅",
		"btn_upd_photo":        "🖼️Rasmni yangilash",
		"prod_photo_updated":   "Mahsulot rasmi muvaffaqiyatli yangilandi✅",
		"prod_photo_msg":       "Rasm yuborish 🖼️",
	},
	"tr": {
		"add_category":         "Kategoriyi Eklemek ➕",
		"add_product":          "Ürün Eklemek ➕",
		"add_admin":            "Admin Eklemek 👨‍⚖️",
		"close_day":            "Gün Kapatmak🔐",
		"open_day":             "Gün Açmak🔓",
		"branch":               "Filial 🏠",
		"get_all_users":        "Tüm Kullanıcıları Almak 👥",
		"category":             "Kategori adı: <UZ>:<RU>:<EN>:<TR>",
		"incorrect_category":   "⁉️Yanlış format. Şunları kullanın: <UZ>:<RU>:<EN>:<TR>",
		"category_created":     "Kategori oluşturuldu ✅",
		"back":                 "🔙Geri",
		"product":              "Şunları kullanın: \n 🖼️ \n<ürün adı tr>:<ürün adı uz>:<ürün adı ru>:<ürün adı en>:<ürün adı tr>:<açıklama>:<fiyat>",
		"product_err":          "Yanlış format. Şunları kullanın: \n 🖼️ \n<ürün adı tr>:<ürün adı uz>:<ürün adı ru>:<ürün adı en>:<ürün adı tr>:<açıklama>:<fiyat>",
		"category_menu":        "Kategoriler",
		"product_menu":         "Ürünler",
		"delete_category":      "Kategoriyi Sil",
		"add_admin_msg":        "Admin ekle: <telegram_id>:<telefon numarası>:<şifre>",
		"admin_created":        "Admin oluşturuldu ✅",
		"day_closed":           "Gün kapalı",
		"day_opened":           "Gün açık",
		"btn_upd_uz":           "Özbekçe Adı",
		"btn_upd_ru":           "Rusça Adı",
		"btn_upd_en":           "İngilizce Adı",
		"btn_upd_tr":           "Türkçe Adı",
		"btn_upd_desc":         "Açıklama",
		"btn_upd_price":        "Fiyat",
		"btn_upd_avail":        "Stok",
		"show_category_info":   "Özbekçe Adı: %s\nRusça Adı: %s\nİngilizce Adı: %s\nTürkçe Adı: %s\n\nStok: %t",
		"name_tr_msg":          "TR Adını girin:",
		"name_uz_msg":          "UZ Adını girin:",
		"name_ru_msg":          "RU Adını girin:",
		"name_en_msg":          "EN Adını girin:",
		"update_category":      "Kategoriyi Güncelle",
		"update_product":       "Ürünü Güncelle",
		"btn_yes":              "Evet✅",
		"btn_no":               "Hayır❌",
		"cat_deleted":          "Kategori silindi ✅",
		"delete_cat_msg":       "Kategoriyi silmek istediğinizden emin misiniz?",
		"name_uz_updated":      "Kategori adı🇺🇿 güncellendi",
		"name_ru_updated":      "Kategori adı🇷🇺 güncellendi",
		"name_en_updated":      "Kategori adı🇬🇧 güncellendi",
		"name_tr_updated":      "Kategori adı🇹🇷 güncellendi",
		"change_lang":          "Dil değiştirmek",
		"no_products":          "Ürün yok",
		"product_created":      "Ürün oluşturuldu ✅",
		"cancel_btn":           "İptal etmek✖️",
		"product_deleted":      "Ürün silindi ✅",
		"product_info":         "Özbekçe adı: %s\nRusça adı: %s\nİngilizce adı: %s\nTürkçe adı: %s\nAçıklama: %s\nFiyat: %d\nMevcutluk: %t",
		"price_msg":            "Fiyat girin:",
		"price_err":            "Yanlış fiyat",
		"prod_name_uz_updated": "Ürün adı🇺🇿 güncellendi",
		"prod_name_ru_updated": "Ürün adı🇷🇺 güncellendi",
		"prod_name_en_updated": "Ürün adı🇬🇧 güncellendi",
		"prod_name_tr_updated": "Ürün adı🇹🇷 güncellendi",
		"prod_name_uz_msg":     "Ürün adını UZ girin:",
		"prod_name_ru_msg":     "Ürün adını RU girin:",
		"prod_name_en_msg":     "Ürün adını EN girin:",
		"prod_name_tr_msg":     "Ürün adını TR girin:",
		"desc_msg":             "Ürün açıklamasını girin:",
		"desc_updated":         "Ürün açıklaması güncellendi",
		"price_updated":        "Ürün fiyatı güncellendi",
		"no_categories":        "Kategori yok",
		"delete_prod":          "Ürünü Sil",
		"btn_upd_cat":          "Kategoriyi Güncelle",
		"prod_cat_updated":     "Ürün kategorisi güncellendi",
		"send_adds":            "Reklam gönder \n 🖼️ \nmetin",
		"adds_btn":             "Reklam",
		"adds_sent":            "Reklam gönderildi✅",
		"btn_upd_photo":        "🖼️Rasmı Güncelle",
		"prod_photo_updated":   "Ürün fotoğrafı güncellendi✅",
		"prod_photo_msg":       "Rasmı gönder 🖼️",
	},
}

var admin_messages = map[int64]string{}

func (h *handlers) ShowAdminPanel(c telebot.Context) error {
	userID := c.Sender().ID
	// fmt.Println(userID)
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	menu := &telebot.ReplyMarkup{}

	// Define the buttons
	btnCat := menu.Data(AdminMessages[lang]["category_menu"], "category_menu")
	btnProd := menu.Data(AdminMessages[lang]["product_menu"], "product_menu")
	btnAdmins := menu.Data(AdminMessages[lang]["add_admin"], "add_admin")
	btnCloseDay := menu.Data(AdminMessages[lang]["close_day"], "close_day")
	btnOpenDay := menu.Data(AdminMessages[lang]["open_day"], "open_day")
	btnFilial := menu.Data(AdminMessages[lang]["branch"], "branch")
	btnLang := menu.Data(AdminMessages[lang]["change_lang"], "admin_lang")
	btnGetUsers := menu.Data(AdminMessages[lang]["get_all_users"], "get_all_users")
	sendAdds := menu.Data(AdminMessages[lang]["adds_btn"], "send_adds")
	update := menu.Data("🔄", "update")

	// Arrange buttons in rows
	menu.Inline(
		menu.Row(btnCat),
		menu.Row(btnProd, btnAdmins),
		menu.Row(btnFilial, btnLang),
		menu.Row(btnCloseDay, btnOpenDay),
		menu.Row(sendAdds),
		menu.Row(btnGetUsers),
		menu.Row(update),
	)
	menu.ResizeKeyboard = true

	if c.Callback() == nil {
		userLocation := "Asia/Tashkent"
		location, err := time.LoadLocation(userLocation)
		if err != nil {
			log.Println(err)
		}

		today := time.Now().In(location)
		formattedDate := today.Format("02.01.2006")
		formattedTime := today.Format("15:04:05")

		orders, err := h.storage.GetOrderByDate(today)
		if err != nil {
			log.Println(err)
		}

		totalSum := 0
		for _, order := range *orders {
			totalSum += order.TotalPrice
		}

		user, err := h.storage.GetUserCount()

		if err != nil {
			user = 0
		}

		admin_messages[userID] = fmt.Sprintf(AdminMessages[lang]["admin_panel"], "Main", user, formattedDate, formattedTime, totalSum)
	}

	message := admin_messages[userID]

	err = c.EditOrSend(message, menu)
	if err != nil {
		c.Send(message, menu)
	}

	return nil
}

func (h *handlers) UpdateAdminPanel(c telebot.Context) error {
	userID := c.Sender().ID
	// fmt.Println(userID)
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	menu := &telebot.ReplyMarkup{}

	// Define the buttons
	btnCat := menu.Data(AdminMessages[lang]["category_menu"], "category_menu")
	btnProd := menu.Data(AdminMessages[lang]["product_menu"], "product_menu")
	btnAdmins := menu.Data(AdminMessages[lang]["add_admin"], "add_admin")
	btnCloseDay := menu.Data(AdminMessages[lang]["close_day"], "close_day")
	btnOpenDay := menu.Data(AdminMessages[lang]["open_day"], "open_day")
	btnFilial := menu.Data(AdminMessages[lang]["branch"], "branch")
	btnLang := menu.Data(AdminMessages[lang]["change_lang"], "admin_lang")
	btnGetUsers := menu.Data(AdminMessages[lang]["get_all_users"], "get_all_users")
	sendAdds := menu.Data(AdminMessages[lang]["adds_btn"], "send_adds")
	update := menu.Data("🔄", "update")

	// Arrange buttons in rows
	menu.Inline(
		menu.Row(btnCat),
		menu.Row(btnProd, btnAdmins),
		menu.Row(btnFilial, btnLang),
		menu.Row(btnCloseDay, btnOpenDay),
		menu.Row(sendAdds),
		menu.Row(btnGetUsers),
		menu.Row(update),
	)
	menu.ResizeKeyboard = true

	c.Edit("Admin panel: \n⏳", menu)

	userLocation := "Asia/Tashkent"
	location, err := time.LoadLocation(userLocation)
	if err != nil {
		log.Println(err)
	}

	today := time.Now().In(location)
	formattedDate := today.Format("02.01.2006")
	formattedTime := today.Format("15:04:05")

	orders, err := h.storage.GetOrderByDate(today)
	if err != nil {
		log.Println(err)
	}

	totalSum := 0
	for _, order := range *orders {
		totalSum += order.TotalPrice
	}

	user, err := h.storage.GetUserCount()

	if err != nil {
		user = 0
	}

	admin_messages[userID] = fmt.Sprintf(AdminMessages[lang]["admin_panel"], "Main", user, formattedDate, formattedTime, totalSum)
	return h.ShowAdminPanel(c)
}

func (h *handlers) ShowCategoryMenu(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	menu := &telebot.ReplyMarkup{}

	// Define the buttons
	btnCat := menu.Data(AdminMessages[lang]["add_category"], "add_category")
	btnUpd := menu.Data(AdminMessages[lang]["update_category"], "update_category")
	btnBack := menu.Data(AdminMessages[lang]["back"], "back_to_admin_menu")
	// btnAdmins := menu.Data(AdminMessages[lang]["add_admin"], "")
	// btnFilial := menu.Data(AdminMessages[lang]["branch"], "branch")
	// btnCloseDay := menu.Data(AdminMessages[lang]["close_day"], "close_day")
	// btnOpenDay := menu.Data(AdminMessages[lang]["open_day"], "open_day")
	// btnGetUsers := menu.Data(AdminMessages[lang]["get_all_users"], "get_all_users")

	// Arrange buttons in rows
	menu.Inline(
		menu.Row(btnCat, btnUpd),
		menu.Row(btnBack),
	)
	menu.ResizeKeyboard = true

	c.EditOrSend("Admin panel:", menu)
	return nil
}

func (h *handlers) ShowCategoryToUpdate(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	cat, err := h.storage.GetCategoriesForAdmin()

	if err != nil {
		return c.Send(err.Error())
	}

	var buttons []telebot.Row
	menu := &telebot.ReplyMarkup{}

	message := " *Update Category*\n"
	for i := 0; i < len(cat.Categories); i += 2 {
		if i+1 < len(cat.Categories) {

			buttons = append(buttons, menu.Row(
				menu.Data(cat.Categories[i].Name_uz, "get_category_info", cat.Categories[i].ID),
				menu.Data(cat.Categories[i+1].Name_uz, "get_category_info", cat.Categories[i+1].ID)))
		} else {
			buttons = append(buttons, menu.Row(
				menu.Data(cat.Categories[i].Name_uz, "get_category_info", cat.Categories[i].ID),
			))
		}
	}
	buttons = append(buttons, menu.Row(menu.Data(AdminMessages[lang]["back"], "back_to_cat_menu")))
	// Define the buttons
	menu.Inline(buttons...)
	menu.ResizeKeyboard = true

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: menu,
	}

	return c.Edit(message, option)
}

func (h *handlers) GetCategoryInfo(c telebot.Context) error {
	userID := c.Sender().ID
	categoryID := c.Callback().Data
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	cat, err := h.storage.GetCategoryByID(categoryID)

	if err != nil {
		return c.Send(err.Error())
	}

	h.storage.UserMessageStatus(c.Sender().ID, "not")

	markup := &telebot.ReplyMarkup{}
	// shortCartID := cat.ID[:8]

	fmt.Println(cat.ID)

	btnUpd_uz := markup.Data(AdminMessages[lang]["btn_upd_uz"], "update_cat_name_uz", cat.ID)
	btnUpd_ru := markup.Data(AdminMessages[lang]["btn_upd_ru"], "update_cat_name_ru", cat.ID)
	btnUpd_en := markup.Data(AdminMessages[lang]["btn_upd_en"], "update_cat_name_en", cat.ID)
	btnUpd_tr := markup.Data(AdminMessages[lang]["btn_upd_tr"], "update_cat_name_tr", cat.ID)
	btnUpd_avail := markup.Data(AdminMessages[lang]["btn_upd_avail"], "update_cat_availability", cat.ID)
	btnUpd_delete := markup.Data(AdminMessages[lang]["delete_category"], "delete_cat", cat.ID)
	btnBack := markup.Data(AdminMessages[lang]["back"], "update_category")

	markup.Inline(
		markup.Row(btnUpd_uz, btnUpd_ru),
		markup.Row(btnUpd_en, btnUpd_tr),
		markup.Row(btnUpd_avail, btnUpd_delete),
		markup.Row(btnBack),
	)

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: markup,
	}

	err = c.Edit(fmt.Sprintf(AdminMessages[lang]["show_category_info"], cat.Name_uz, cat.Name_ru, cat.Name_en, cat.Name_tr, cat.Abelety), option)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (h *handlers) UpdateCategoryNameUzHandle(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err)
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Data)

	if err != nil {
		return c.Send(err.Error() + "1")
	}
	menu := &telebot.ReplyMarkup{}
	buttons := menu.Row(menu.Data(AdminMessages[lang]["back"], "get_category_info", c.Callback().Data))
	menu.Inline(buttons)
	menu.ResizeKeyboard = true
	return c.Edit(AdminMessages[lang]["name_uz_msg"], menu)
}

func (h *handlers) UpdateCategoryNameRuHandle(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err.Error())
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Data)

	if err != nil {
		return c.Send(err.Error() + "2")
	}
	menu := &telebot.ReplyMarkup{}
	buttons := menu.Row(menu.Data(AdminMessages[lang]["back"], "get_category_info", c.Callback().Data))
	menu.Inline(buttons)
	menu.ResizeKeyboard = true
	return c.Edit(AdminMessages[lang]["name_ru_msg"], menu)
}

func (h *handlers) UpdateCategoryNameEnHandle(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err.Error())
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err.Error() + "3")
	}
	menu := &telebot.ReplyMarkup{}
	buttons := menu.Row(menu.Data(AdminMessages[lang]["back"], "get_category_info", c.Callback().Data))
	menu.Inline(buttons)
	menu.ResizeKeyboard = true
	return c.Send(AdminMessages[lang]["name_en_msg"], menu)
}

func (h *handlers) UpdateCategoryNameTrHandle(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err.Error())
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Data)

	if err != nil {
		return c.Send(err.Error())
	}
	menu := &telebot.ReplyMarkup{}
	buttons := menu.Row(menu.Data(AdminMessages[lang]["back"], "get_category_info", c.Callback().Data))
	menu.Inline(buttons)
	menu.ResizeKeyboard = true
	return c.Edit(AdminMessages[lang]["name_tr_msg"], menu)
}

func (h *handlers) UpdateCategoryNameUz(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	catID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error() + "+")
	}
	name := c.Message().Text
	err = h.storage.UpdateNameUzCategoryById(catID, name)
	if err != nil {
		return c.Send(err.Error() + "4")
	}
	return c.Respond(&telebot.CallbackResponse{
		Text: AdminMessages[lang]["name_uz_updated"]})
}

func (h *handlers) UpdateCategoryNameRu(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	catID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	name := c.Message().Text
	err = h.storage.UpdateNameRuCategoryById(catID, name)
	if err != nil {
		return c.Send(err.Error())
	}
	return c.Respond(&telebot.CallbackResponse{
		Text: AdminMessages[lang]["name_ru_updated"]})
}

func (h *handlers) UpdateCategoryNameEn(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	catID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	name := c.Message().Text
	err = h.storage.UpdateNameEnCategoryById(catID, name)
	if err != nil {
		return c.Send(err.Error())
	}
	return c.Respond(&telebot.CallbackResponse{
		Text: AdminMessages[lang]["name_en_updated"]})
}

func (h *handlers) UpdateCategoryNameTr(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	catID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	name := c.Message().Text
	err = h.storage.UpdateNameTrCategoryById(catID, name)
	if err != nil {
		return c.Send(err.Error())
	}
	return c.Respond(&telebot.CallbackResponse{
		Text: AdminMessages[lang]["name_tr_updated"]})
}

func (h *handlers) UpdateCategoryAvailability(c telebot.Context) error {
	userID := c.Sender().ID
	cat_id := c.Callback().Data
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	err = h.storage.UpdateAbeletyCategoryById(cat_id)
	if err != nil {
		return c.Send(err.Error())
	}

	cat, err := h.storage.GetCategoryByID(cat_id)

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}
	// shortCartID := cat.ID[:8]

	fmt.Println(cat.ID)

	btnUpd_uz := markup.Data(AdminMessages[lang]["btn_upd_uz"], "update_cat_name_uz", cat.ID)
	btnUpd_ru := markup.Data(AdminMessages[lang]["btn_upd_ru"], "update_cat_name_ru", cat.ID)
	btnUpd_en := markup.Data(AdminMessages[lang]["btn_upd_en"], "update_cat_name_en", cat.ID)
	btnUpd_tr := markup.Data(AdminMessages[lang]["btn_upd_tr"], "update_cat_name_tr", cat.ID)
	btnUpd_avail := markup.Data(AdminMessages[lang]["btn_upd_avail"], "update_cat_availability", cat.ID)
	btnUpd_delete := markup.Data(AdminMessages[lang]["delete_category"], "delete_cat", cat.ID)
	btnBack := markup.Data(AdminMessages[lang]["back"], "update_category")

	markup.Inline(
		markup.Row(btnUpd_uz, btnUpd_ru),
		markup.Row(btnUpd_en, btnUpd_tr),
		markup.Row(btnUpd_avail, btnUpd_delete),
		markup.Row(btnBack),
	)

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: markup,
	}

	err = c.Edit(fmt.Sprintf(AdminMessages[lang]["show_category_info"], cat.Name_uz, cat.Name_ru, cat.Name_en, cat.Name_tr, cat.Abelety), option)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (h *handlers) DeleteCategoryHandle(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	menu := &telebot.ReplyMarkup{}

	btnYes := menu.Data(AdminMessages[lang]["btn_yes"], "delete_cat_yes", c.Callback().Data)
	btnNo := menu.Data(AdminMessages[lang]["btn_no"], "get_category_info", c.Callback().Data)

	// Arrange buttons in rows
	menu.Inline(
		menu.Row(btnYes, btnNo),
	)
	menu.ResizeKeyboard = true
	return c.Send(AdminMessages[lang]["delete_cat_msg"], menu)
}

func (h *handlers) DeleteCategory(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	h.storage.DeleteCategoryById(c.Callback().Data)

	menu := &telebot.ReplyMarkup{}

	btnBack := menu.Data(AdminMessages[lang]["back"], "update_category")

	menu.Inline(
		menu.Row(btnBack),
	)
	menu.ResizeKeyboard = true
	return c.Send(AdminMessages[lang]["cat_deleted"], menu)
}

var createCatID int

func (h *handlers) CreateCategoryHandle(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID

	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "back_to_cat_menu"))

	markup.Inline(btnBack)

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}

	msg, err := c.Bot().Send(c.Recipient(), AdminMessages[lang]["category"], options)
	createCatID = msg.ID
	if err != nil {
		fmt.Println(err)
	}
	h.storage.UserMessageStatus(userID, "add_category")
	return nil
}

func (h *handlers) CreateCategory(c telebot.Context) error {
	c.Bot().Delete(&telebot.Message{ID: createCatID, Chat: c.Chat()})
	c.Delete()
	text := c.Message().Text
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	arg := strings.Split(text, ":")

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "back_to_cat_menu"))

	markup.Inline(btnBack)

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}
	if len(arg) != 4 {
		return c.Edit(AdminMessages[lang]["incorrect_category"], options)
	}

	category := models.Category{}

	nameUz := strings.TrimSpace(arg[0])
	nameRu := strings.TrimSpace(arg[1])
	nameEn := strings.TrimSpace(arg[2])
	nameTr := strings.TrimSpace(arg[3])
	fmt.Println(nameUz)

	category.Name_uz = nameUz
	category.Name_ru = nameRu
	category.Name_en = nameEn
	category.Name_tr = nameTr
	category.Abelety = true

	err = h.storage.CreateCategory(&category)

	if err != nil {
		log.Println(err.Error())
		msg, err := c.Bot().Send(c.Recipient(), "You have Dublicate name", options)
		createCatID = msg.ID
		return err
	}

	// return c.Respond(&telebot.CallbackResponse{
	// 	Text: "Category created✅"})

	return c.Send("Category created✅", options)
}

func (h *handlers) GetUsers(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	users, err := h.storage.GetAllUsers()
	if err != nil {
		return c.Send(err.Error())
	}

	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "№")
	f.SetCellValue("Sheet1", "B1", "User ID")
	f.SetCellValue("Sheet1", "C1", "Username")
	f.SetCellValue("Sheet1", "D1", "Phone Number")
	f.SetCellValue("Sheet1", "E1", "Name")
	f.SetCellValue("Sheet1", "F1", "Date")

	row := 2

	for _, data := range users.Users {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), row-1)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), data.TelegramID)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), data.Username)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), data.Phone_Number)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), data.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), data.Created_at.Format("2006-01-02 15:04:05"))
		row++
	}

	f.SetColWidth("Sheet1", "A", "A", 5)  // Adjust column A
	f.SetColWidth("Sheet1", "B", "B", 15) // Adjust column A
	f.SetColWidth("Sheet1", "C", "C", 20) // Adjust column B
	f.SetColWidth("Sheet1", "D", "D", 20) // Adjust column C
	f.SetColWidth("Sheet1", "E", "E", 25) // Adjust column D
	f.SetColWidth("Sheet1", "F", "F", 30) // Adjust column E

	filename := fmt.Sprintf("user_data_%s.xlsx", time.Now())
	err = f.SaveAs(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	// Send file to admin
	doc := &telebot.Document{File: telebot.FromReader(file), FileName: filename}
	_, err = c.Bot().Send(c.Sender(), doc)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Optionally, send a confirmation message to the bot user
	err = h.ShowAdminPanel(c)
	fmt.Println(err)
	return nil
}

func (h *handlers) ShowProductMenu(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	menu := &telebot.ReplyMarkup{}

	// Define the buttons
	btnProd := menu.Data(AdminMessages[lang]["add_product"], "add_product")
	btnAdmins := menu.Data(AdminMessages[lang]["update_product"], "update_product")
	btnBack := menu.Data(AdminMessages[lang]["back"], "back_to_admin_menu")

	// Arrange buttons in rows
	menu.Inline(
		menu.Row(btnProd, btnAdmins),
		menu.Row(btnBack),
	)
	menu.ResizeKeyboard = true

	h.storage.UserMessageStatus(userID, "admin")

	c.EditOrSend("Product menu:", menu)
	return nil
}

func (h *handlers) AddProductHandler(c telebot.Context) error {
	userID := c.Sender().ID
	// Check if the user is an admin
	if !h.storage.CheckAdmin(userID) {
		return c.Send("У вас нет прав для выполнения этой команды.")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "product_menu"))

	markup.Inline(btnBack)

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}

	err = c.Edit(AdminMessages[lang]["product"], options)

	if err != nil {
		fmt.Println(err)
	}
	h.storage.UserMessageStatus(userID, "add_product")
	return nil
}

func (h *handlers) AddProduct(c telebot.Context) error {
	userID := c.Sender().ID
	// fmt.Println("add prod")
	pic := c.Message().Photo
	text := c.Message().Caption
	// fmt.Println(text)
	product := &models.Product{}

	if !h.storage.CheckAdmin(userID) {
		return c.Send("У вас нет прав для выполнения этой команды.")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	args := strings.Split(text, ":")
	// fmt.Println(args)
	if len(args) != 6 {
		return c.Send(AdminMessages[lang]["product_err"])
	}

	nameUz := strings.TrimSpace(args[0])
	nameRu := strings.TrimSpace(args[1])
	nameEn := strings.TrimSpace(args[2])
	nameTr := strings.TrimSpace(args[3])
	description := strings.TrimSpace(args[4])
	fmt.Println(nameUz, nameRu, nameEn, description, args[5])
	price, err := strconv.Atoi(strings.TrimSpace(args[5]))
	if err != nil {
		return c.Send("Ошибка в формате цены. Убедитесь, что это число.")
	}
	// availability, err := strconv.ParseBool(strings.TrimSpace(args[6]))
	// if err != nil {
	// 	return c.Send("Ошибка в формате доступности. Используйте true или false.")
	// }
	product = &models.Product{
		Name_uz:     nameUz,
		Name_ru:     nameRu,
		Name_en:     nameEn,
		Name_tr:     nameTr,
		Description: description,
		Price:       price,
		Abelety:     true,
	}

	reader, err := c.Bot().File(&pic.File)

	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка загрузки фото: %v", err))
	}

	defer reader.Close()

	if _, err := os.Stat("./photos"); os.IsNotExist(err) {
		os.Mkdir("./photos", os.ModePerm)
	}

	// photoURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", c.Bot().Token, pic.FileID)
	photoPath := fmt.Sprintf("./photos/file_%d.jpg", time.Now().Unix())

	file, err := os.Create(photoPath)

	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка создания файла: %v", err))
	}

	defer file.Close()

	_, err = io.Copy(file, reader)

	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка сохранения фото: %v", err))
	}

	product.Photo = photoPath

	_, err = h.storage.CreateProduct(product)
	if err != nil {
		log.Printf("Ошибка добавления продукта: %v", err)
		return c.Send(fmt.Sprintf("Не удалось добавить продукт: %v", err))
	}

	// c.Send(fmt.Sprintf("Продукт '%s' успешно добавлен в категорию '%s'.", product.Name_uz, product.Category_id))

	btn := &telebot.ReplyMarkup{}
	btn.ResizeKeyboard = true
	btn.OneTimeKeyboard = true
	var buttons []telebot.Row

	categories, err := h.storage.GetCategoriesForAdmin()
	if err != nil {
		return c.Send(err)
	}

	productID := helpers.CompresedUUID(uuid.MustParse(product.ID))
	for _, cat := range categories.Categories {
		categoryID := helpers.CompresedUUID(uuid.MustParse(cat.ID))
		button := btn.Data(cat.Name_uz, "add_prod_cat", categoryID+","+productID)
		buttons = append(buttons, btn.Row(button))
	}
	buttons = append(buttons, btn.Row(btn.Data(AdminMessages[lang]["cancel_btn"], "delete_product", product.ID)))
	btn.Inline(buttons...)

	if err := c.Send("Выберите категорию:", btn); err != nil {
		return err
	}

	return nil
}

func (h *handlers) DeleteProductHandle(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	prod_id := c.Callback().Data

	if !h.storage.CheckAdmin(userID) {
		return c.Send("У вас нет прав для выполнения этой команды.")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	err = h.storage.DeleteProductById(prod_id)
	if err != nil {
		log.Println(err)
		return c.Send(err)
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "back_to_prod_menu"))

	markup.Inline(btnBack)

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}

	return c.Send(AdminMessages[lang]["product_deleted"], options)
}

func (h *handlers) AddProductToCategory(c telebot.Context) error {
	userID := c.Sender().ID
	text := c.Callback().Data
	fmt.Println(text)

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	arg := strings.Split(text, ",")

	cat_id := strings.TrimSpace(arg[0])
	prod_id := strings.TrimSpace(arg[1])

	categoryID, err := helpers.DecompresedUUID(cat_id)
	if err != nil {
		return c.Send(err.Error())
	}
	prodID, err := helpers.DecompresedUUID(prod_id)
	if err != nil {
		return c.Send(err.Error())
	}

	err = h.storage.AddProductToCategory(prodID.String(), categoryID.String())

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "product_menu"))

	markup.Inline(btnBack)

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}

	return c.Edit(AdminMessages[lang]["product_created"], options)
}

func (h *handlers) ShowCategoriesToUpdateProducts(c telebot.Context) error {
	userID := c.Sender().ID
	// Check if the user is an admin
	if !h.storage.CheckAdmin(userID) {
		return c.Send("У вас нет прав для выполнения этой команды.")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	categories, err := h.storage.GetCategoriesForAdmin()
	if err != nil {
		return c.Send(err.Error())
	}
	var buttons []telebot.Row
	menu := &telebot.ReplyMarkup{}

	if len(categories.Categories) == 0 {
		return c.Send(AdminMessages[lang]["no_categories"])
	}
	message := " *Update Product*\n"
	for i := 0; i < len(categories.Categories); i += 2 {
		if i+1 < len(categories.Categories) {
			buttons = append(buttons, menu.Row(
				menu.Data(fmt.Sprintf("%s/%s/%s/%s",
					categories.Categories[i].Name_uz,
					categories.Categories[i].Name_ru,
					categories.Categories[i].Name_en,
					categories.Categories[i].Name_tr,
				), "get_product_by_category", categories.Categories[i].ID),
				menu.Data(fmt.Sprintf("%s/%s/%s/%s",
					categories.Categories[i+1].Name_uz,
					categories.Categories[i+1].Name_ru,
					categories.Categories[i+1].Name_en,
					categories.Categories[i+1].Name_tr,
				), "get_product_by_category", categories.Categories[i+1].ID),
			))
		} else {
			buttons = append(buttons, menu.Row(
				menu.Data(fmt.Sprintf("%s/%s/%s/%s",
					categories.Categories[i].Name_uz,
					categories.Categories[i].Name_ru,
					categories.Categories[i].Name_en,
					categories.Categories[i].Name_tr,
				), "get_product_by_category", categories.Categories[i].ID),
			))
		}
	}
	buttons = append(buttons, menu.Row(menu.Data(AdminMessages[lang]["back"], "product_menu")))
	// Define the buttons
	menu.Inline(buttons...)
	menu.ResizeKeyboard = true

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: menu,
	}

	return c.Edit(message, option)
}

func (h *handlers) ShowProductsToUpdate(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	// Check if the user is an admin
	if !h.storage.CheckAdmin(userID) {
		return c.Send("У вас нет прав для выполнения этой команды.")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	prod, err := h.storage.GetProductsByCategoryForAdmin(c.Callback().Data)

	if err != nil {
		return c.Send(err.Error())
	}

	var buttons []telebot.Row
	menu := &telebot.ReplyMarkup{}

	if len(prod.Products) == 0 {
		buttons = append(buttons, menu.Row(menu.Data(AdminMessages[lang]["back"], "update_product")))
		// Define the buttons
		menu.Inline(buttons...)
		menu.ResizeKeyboard = true

		option := &telebot.SendOptions{
			ParseMode:   telebot.ModeMarkdownV2,
			ReplyMarkup: menu,
		}
		return c.Edit(AdminMessages[lang]["no_products"], option)
	}

	message := " *Update Product*\n"
	for i := 0; i < len(prod.Products); i += 2 {
		if i+1 < len(prod.Products) {

			buttons = append(buttons, menu.Row(
				menu.Data(prod.Products[i].Name_uz, "get_product_info", prod.Products[i].ID),
				menu.Data(prod.Products[i+1].Name_uz, "get_product_info", prod.Products[i+1].ID)))
		} else {
			buttons = append(buttons, menu.Row(
				menu.Data(prod.Products[i].Name_uz, "get_product_info", prod.Products[i].ID),
			))
		}
	}
	buttons = append(buttons, menu.Row(menu.Data(AdminMessages[lang]["back"], "update_product")))
	// Define the buttons
	menu.Inline(buttons...)
	menu.ResizeKeyboard = true

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: menu,
	}

	h.storage.UserMessageStatus(userID, "admin")

	return c.Send(message, option)
}

func (h *handlers) GetProductInfo(c telebot.Context) error {
	userID := c.Sender().ID
	productID := c.Callback().Data
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error)
	}
	prod, err := h.storage.GetProductByIdForAdmin(productID)

	if err != nil {
		return c.Send(err.Error())
	}

	cat, err := h.storage.GetCategoryByID(prod.Category_id)

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}
	btnUpd_photo := markup.Data(AdminMessages[lang]["btn_upd_photo"], "update_photo_of_prod", prod.ID)
	btnUpd_cat := markup.Data(AdminMessages[lang]["btn_upd_cat"], "update_cat_of_prod", prod.ID)
	btnUpd_uz := markup.Data(AdminMessages[lang]["btn_upd_uz"], "update_prod_name_uz", prod.ID)
	btnUpd_ru := markup.Data(AdminMessages[lang]["btn_upd_ru"], "update_prod_name_ru", prod.ID)
	btnUpd_en := markup.Data(AdminMessages[lang]["btn_upd_en"], "update_prod_name_en", prod.ID)
	btnUpd_tr := markup.Data(AdminMessages[lang]["btn_upd_tr"], "update_prod_name_tr", prod.ID)
	btnUpd_desc := markup.Data(AdminMessages[lang]["btn_upd_desc"], "update_prod_desc", prod.ID)
	btnUpd_price := markup.Data(AdminMessages[lang]["btn_upd_price"], "update_prod_price", prod.ID)
	btnUpd_avail := markup.Data(AdminMessages[lang]["btn_upd_avail"], "update_prod_availability", prod.ID)
	btnDelete := markup.Data(AdminMessages[lang]["delete_prod"], "delete_prod", prod.ID)
	btnBack := markup.Data(AdminMessages[lang]["back"], "get_product_by_category", prod.Category_id)

	markup.Inline(
		markup.Row(btnUpd_cat, btnUpd_photo),
		markup.Row(btnUpd_uz, btnUpd_ru),
		markup.Row(btnUpd_en, btnUpd_tr),
		markup.Row(btnUpd_desc, btnUpd_price),
		markup.Row(btnUpd_avail, btnDelete),
		markup.Row(btnBack),
	)
	message := fmt.Sprintf("* %s/ %s/ %s/ %s* \n\n", cat.Name_uz, cat.Name_ru, cat.Name_en, cat.Name_tr) + helpers.EscapeMarkdownV2(fmt.Sprintf(AdminMessages[lang]["product_info"], prod.Name_uz, prod.Name_ru, prod.Name_en, prod.Name_tr, prod.Description, prod.Price, prod.Abelety))
	photo := &telebot.Photo{File: telebot.FromDisk(prod.Photo), Caption: message}

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: markup,
	}

	err = c.EditOrSend(photo, option)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (h *handlers) UpdateProductCategoryHandle(c telebot.Context) error {
	userID := c.Sender().ID
	prodID := c.Callback().Data
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	_, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	btn := &telebot.ReplyMarkup{}
	btn.ResizeKeyboard = true
	btn.OneTimeKeyboard = true
	var buttons []telebot.Row

	categories, err := h.storage.GetCategoriesForAdmin()
	if err != nil {
		return c.Send(err)
	}

	productID := helpers.CompresedUUID(uuid.MustParse(prodID))
	for _, cat := range categories.Categories {
		categoryID := helpers.CompresedUUID(uuid.MustParse(cat.ID))
		button := btn.Data(cat.Name_uz, "upd_prod_cat", categoryID+","+productID)
		buttons = append(buttons, btn.Row(button))
	}
	btn.Inline(buttons...)
	return c.Edit("Выберите категорию:", btn)
}

func (h *handlers) UpdateProductCategory(c telebot.Context) error {
	userID := c.Sender().ID
	text := c.Callback().Data

	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	arg := strings.Split(text, ",")

	cat_id := strings.TrimSpace(arg[0])
	prod_id := strings.TrimSpace(arg[1])

	categoryID, err := helpers.DecompresedUUID(cat_id)
	if err != nil {
		return c.Send(err.Error())
	}
	prodID, err := helpers.DecompresedUUID(prod_id)
	if err != nil {
		return c.Send(err.Error())
	}

	err = h.storage.UpdateProductCategoryById(prodID.String(), categoryID.String())

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "get_product_info", prodID.String()))

	markup.Inline(btnBack)

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}

	return c.Edit(AdminMessages[lang]["prod_cat_updated"], options)
}

func (h *handlers) UpdateProductNameUzHandle(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err)
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Data)

	if err != nil {
		return c.Send(err.Error() + "1")
	}
	return c.Send(AdminMessages[lang]["prod_name_uz_msg"])
}

func (h *handlers) UpdateProductNameRuHandle(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err.Error())
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Data)

	if err != nil {
		return c.Send(err.Error() + "2")
	}
	return c.Send(AdminMessages[lang]["prod_name_ru_msg"])
}

func (h *handlers) UpdateProductNameEnHandle(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(userID, c.Callback().Unique)

	if err != nil {
		return c.Send(err.Error())
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Data)

	if err != nil {
		return c.Send(err.Error() + "3")
	}
	return c.Send(AdminMessages[lang]["prod_name_en_msg"])
}

func (h *handlers) UpdateProductNameTrHandle(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err.Error())
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Data)

	if err != nil {
		return c.Send(err.Error())
	}
	return c.Send(AdminMessages[lang]["prod_name_tr_msg"])
}

func (h *handlers) UpdateProductNameUz(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	prodID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error() + "+")
	}
	name := c.Message().Text
	err = h.storage.UpdateProductNameUz(prodID, name)
	if err != nil {
		return c.Send(err.Error() + "4")
	}

	markup := &telebot.ReplyMarkup{}
	btn := markup.Data(AdminMessages[lang]["back"], "get_product_info", prodID)
	markup.Inline(markup.Row(btn))
	markup.ResizeKeyboard = true

	option := &telebot.SendOptions{
		ReplyMarkup: markup,
	}
	return c.Send(AdminMessages[lang]["prod_name_uz_updated"], option)
}

func (h *handlers) UpdateProductNameRu(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	prodID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	name := c.Message().Text
	err = h.storage.UpdateProductNameRu(prodID, name)
	if err != nil {
		return c.Send(err.Error())
	}
	markup := &telebot.ReplyMarkup{}
	btn := markup.Data(AdminMessages[lang]["back"], "get_product_info", prodID)
	markup.Inline(markup.Row(btn))
	markup.ResizeKeyboard = true

	option := &telebot.SendOptions{
		ReplyMarkup: markup,
	}
	return c.Send(AdminMessages[lang]["prod_name_ru_updated"], option)
}

func (h *handlers) UpdateProductNameEn(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	prodID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	name := c.Message().Text
	err = h.storage.UpdateProductNameEn(prodID, name)
	if err != nil {
		return c.Send(err.Error())
	}
	markup := &telebot.ReplyMarkup{}
	btn := markup.Data(AdminMessages[lang]["back"], "get_product_info", prodID)
	markup.Inline(markup.Row(btn))
	markup.ResizeKeyboard = true

	option := &telebot.SendOptions{
		ReplyMarkup: markup,
	}
	return c.Send(AdminMessages[lang]["prod_name_en_updated"], option)
}

func (h *handlers) UpdateProductNameTr(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	prodID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	name := c.Message().Text
	err = h.storage.UpdateProductNameTr(prodID, name)
	if err != nil {
		return c.Send(err.Error())
	}
	markup := &telebot.ReplyMarkup{}
	btn := markup.Data(AdminMessages[lang]["back"], "get_product_info", prodID)
	markup.Inline(markup.Row(btn))
	markup.ResizeKeyboard = true

	option := &telebot.SendOptions{
		ReplyMarkup: markup,
	}
	return c.Send(AdminMessages[lang]["prod_name_tr_updated"], option)
}

func (h *handlers) UpdateProductAvailability(c telebot.Context) error {
	userID := c.Sender().ID
	prod_id := c.Callback().Data
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	err = h.storage.UpdateAbeletyProductById(prod_id)
	if err != nil {
		return c.Send(err.Error())
	}

	prod, err := h.storage.GetProductByIdForAdmin(prod_id)

	if err != nil {
		return c.Send(err.Error())
	}

	cat, err := h.storage.GetCategoryByID(prod.Category_id)

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}
	btnUpd_photo := markup.Data(AdminMessages[lang]["btn_upd_photo"], "update_photo_of_prod", prod.ID)
	btnUpd_cat := markup.Data(AdminMessages[lang]["btn_upd_cat"], "update_cat_of_prod", prod.ID)
	btnUpd_uz := markup.Data(AdminMessages[lang]["btn_upd_uz"], "update_prod_name_uz", prod.ID)
	btnUpd_ru := markup.Data(AdminMessages[lang]["btn_upd_ru"], "update_prod_name_ru", prod.ID)
	btnUpd_en := markup.Data(AdminMessages[lang]["btn_upd_en"], "update_prod_name_en", prod.ID)
	btnUpd_tr := markup.Data(AdminMessages[lang]["btn_upd_tr"], "update_prod_name_tr", prod.ID)
	btnUpd_desc := markup.Data(AdminMessages[lang]["btn_upd_desc"], "update_prod_desc", prod.ID)
	btnUpd_price := markup.Data(AdminMessages[lang]["btn_upd_price"], "update_prod_price", prod.ID)
	btnUpd_avail := markup.Data(AdminMessages[lang]["btn_upd_avail"], "update_prod_availability", prod.ID)
	btnDelete := markup.Data(AdminMessages[lang]["delete_prod"], "delete_prod", prod.ID)
	btnBack := markup.Data(AdminMessages[lang]["back"], "get_product_by_category", prod.Category_id)

	markup.Inline(
		markup.Row(btnUpd_cat, btnUpd_photo),
		markup.Row(btnUpd_uz, btnUpd_ru),
		markup.Row(btnUpd_en, btnUpd_tr),
		markup.Row(btnUpd_desc, btnUpd_price),
		markup.Row(btnUpd_avail, btnDelete),
		markup.Row(btnBack),
	)

	option := &telebot.SendOptions{
		ParseMode:   telebot.ModeMarkdownV2,
		ReplyMarkup: markup,
	}

	err = c.EditCaption(fmt.Sprintf("* %s/ %s/ %s/ %s* \n\n", cat.Name_uz, cat.Name_ru, cat.Name_en, cat.Name_tr)+helpers.EscapeMarkdownV2(fmt.Sprintf(AdminMessages[lang]["product_info"], prod.Name_uz, prod.Name_ru, prod.Name_en, prod.Name_tr, prod.Description, prod.Price, prod.Abelety)), option)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (h *handlers) UpdateProductPhotoHandle(c telebot.Context) error {
	c.Delete()
	prod_id := c.Callback().Data
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, "update_photo")

	if err != nil {
		return c.Send(err)
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, prod_id)

	if err != nil {
		return c.Send(err.Error() + "1")
	}
	markup := &telebot.ReplyMarkup{}
	btn := markup.Data(AdminMessages[lang]["back"], "get_product_info", prod_id)
	markup.Inline(markup.Row(btn))
	markup.ResizeKeyboard = true

	option := &telebot.SendOptions{
		ReplyMarkup: markup,
	}
	return c.Send(AdminMessages[lang]["prod_photo_msg"], option)
}

func (h *handlers) UpdateProductPhoto(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	prodID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error() + "+")
	}

	prod, err := h.storage.GetProductByIdForAdmin(prodID)

	os.Remove(prod.Photo)

	if err != nil {
		return c.Send(err.Error())
	}

	pic := c.Message().Photo
	reader, err := c.Bot().File(&pic.File)

	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка загрузки фото: %v", err))
	}

	defer reader.Close()

	if _, err := os.Stat("./photos"); os.IsNotExist(err) {
		os.Mkdir("./photos", os.ModePerm)
	}

	photoPath := fmt.Sprintf("./photos/file_%d.jpg", time.Now().Unix())

	file, err := os.Create(photoPath)

	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка создания файла: %v", err))
	}

	defer file.Close()

	_, err = io.Copy(file, reader)

	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка копирования файла: %v", err))
	}

	err = h.storage.UpdateProductPhotoById(prodID, photoPath)
	if err != nil {
		return c.Send(err.Error())
	}
	// markup := &telebot.ReplyMarkup{}
	// btn := markup.Data(AdminMessages[lang]["back"], "get_product_info", prodID)
	// markup.Inline(markup.Row(btn))
	// markup.ResizeKeyboard = true

	// option := &telebot.SendOptions{
	// 	ReplyMarkup: markup,
	// }
	return c.Respond(&telebot.CallbackResponse{
		Text: AdminMessages[lang]["prod_photo_updated"],
	})
}

func (h *handlers) UpdateProductDescHandle(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err)
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Data)

	if err != nil {
		return c.Send(err.Error() + "1")
	}
	return c.Send(AdminMessages[lang]["desc_msg"])
}

func (h *handlers) UpdateProductPriceHandle(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	_, err = h.storage.UserMessageStatus(c.Sender().ID, c.Callback().Unique)

	if err != nil {
		return c.Send(err.Error())
	}

	_, err = h.storage.SetDataUserMessageStatus(c.Sender().ID, c.Callback().Data)

	if err != nil {
		return c.Send(err.Error() + "2")
	}
	return c.Send(AdminMessages[lang]["price_msg"])
}

func (h *handlers) UpdateProductDesc(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	prodID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	desc := c.Message().Text
	err = h.storage.UpdateProductDescById(prodID, desc)
	if err != nil {
		return c.Send(err.Error())
	}
	return c.Respond(&telebot.CallbackResponse{
		Text: AdminMessages[lang]["desc_updated"],
	})
}

func (h *handlers) UpdateProductPrice(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}
	prodID, err := h.storage.GetDataUserMessageStatus(userID)
	if err != nil {
		return c.Send(err.Error())
	}
	price, err := strconv.Atoi(c.Message().Text)
	if err != nil {
		return c.Send("Error in price format. Make sure it's a number.")
	}
	err = h.storage.UpdateProductPriceById(prodID, price)
	if err != nil {
		return c.Send(err.Error())
	}
	return c.Respond(&telebot.CallbackResponse{
		Text: AdminMessages[lang]["price_updated"],
	})
}

func (h *handlers) DeleteProduct(c telebot.Context) error {
	c.Delete()
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	err = h.storage.DeleteProductById(c.Callback().Data)
	if err != nil {
		log.Println(err)
		return c.Send(err)
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Data(AdminMessages[lang]["back"], "update_product")

	markup.Inline(markup.Row(btnBack))

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}

	return c.Send(AdminMessages[lang]["product_deleted"], options)
}

func (h *handlers) AddAdminHandle(c telebot.Context) error {
	userID := c.Sender().ID
	// Check if the user is an admin
	if !h.storage.CheckAdmin(userID) {
		return c.Send("У вас нет прав для выполнения этой команды.")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	h.storage.UserMessageStatus(userID, "add_admin")

	markup := &telebot.ReplyMarkup{OneTimeKeyboard: true}

	btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "back_to_admin_menu"))

	markup.Inline(btnBack)

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}

	c.Edit(AdminMessages[lang]["add_admin_msg"], options)

	return nil
}

func (h *handlers) AddAdmin(c telebot.Context) error {
	c.Delete()
	text := c.Text()
	userID := c.Sender().ID
	// Check if the user is an admin
	if !h.storage.CheckAdmin(userID) {
		return c.Send("У вас нет прав для выполнения этой команды.")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	if text == "Back" {
		h.ShowAdminPanel(c)
		return nil
	}
	admin := &models.Admin{}
	args := strings.Split(text, ":")
	if len(args) != 3 {
		return c.Send(Messages[lang]["product_err"])
	}

	telegramID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return c.Send("Ошибка в формате Телеграм ID. Убедитесь, что это число.")
	}
	phone := strings.TrimSpace(args[1])
	password := strings.TrimSpace(args[2])

	admin = &models.Admin{
		Admin_id:     telegramID,
		Phone_Number: phone,
		Password:     password,
	}

	_, err = h.storage.CreateAdmin(admin)

	if err != nil {
		return c.Send(err.Error())
	}

	// markup := &telebot.ReplyMarkup{OneTimeKeyboard: true}

	// btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "back_to_admin_menu"))

	// markup.Inline(btnBack)

	// options := &telebot.SendOptions{
	// 	ReplyMarkup: markup,
	// }

	c.Respond(&telebot.CallbackResponse{
		Text: Messages[lang]["admin_added"],
	})

	return nil
}

func (h *handlers) CloseDay(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		c.Send("У вас нет прав для выполнения этой команды.")
		return nil
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	err = h.storage.CloseDay()

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Data(AdminMessages[lang]["back"], "back_to_admin_menu")

	markup.Inline(markup.Row(btnBack))

	c.Edit(AdminMessages[lang]["day_closed"], markup)
	return nil
}

func (h *handlers) OpenDay(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		c.Send("У вас нет прав для выполнения этой команды.")
		return nil
	}
	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err)
	}
	err = h.storage.OpenDay()

	if err != nil {
		return c.Send(err)
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Data(AdminMessages[lang]["back"], "back_to_admin_menu")

	markup.Inline(markup.Row(btnBack))

	c.Edit(AdminMessages[lang]["day_opened"], markup)
	return nil
}

func (h *handlers) AdminPhotostatus(c telebot.Context) error {
	userId := c.Sender().ID

	status, err := h.storage.GetUserMessageStatus(userId)

	if err != nil {
		return c.Send(err.Error())
	}

	switch status {
	case "add_product":
		return h.AddProduct(c)
	case "update_photo":
		return h.UpdateProductPhoto(c)
	case "adds":
		return h.SendAddToUsers(c)
	default:
		// if h.storage.CheckUserExist(userId) {
		// 	noloc := &telebot.ReplyMarkup{}
		// 	btnBack := noloc.Data(Messages["en"]["back"], "back_to_user_menu")
		// 	noloc.Inline(noloc.Row(btnBack))
		// 	return c.Send("Unknown status", noloc)
		// }
		return c.Send("Unknown status")
	}
}

func (h *handlers) SendAddToUsersHandle(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "back_to_admin_menu"))

	markup.Inline(btnBack)

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}

	c.Edit(AdminMessages[lang]["send_adds"], options)

	h.storage.UserMessageStatus(userID, "adds")
	return nil
}

func (h *handlers) SendAddToUsers(c telebot.Context) error {
	c.Delete()
	text := c.Message().Caption
	photo := c.Message().Photo
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	users, err := h.storage.GetAllUsers()
	if err != nil {
		return c.Send(err.Error())
	}

	photo_add := &telebot.Photo{File: photo.File, Caption: text}

	for _, user := range users.Users {
		_, err := c.Bot().Send(
			&telebot.User{ID: user.TelegramID},
			photo_add,
		)
		if err != nil {
			log.Printf("Error sending message to user %d: %v", user.TelegramID, err)
		}
	}

	// markup := &telebot.ReplyMarkup{}

	// btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "back_to_admin_menu"))

	// markup.Inline(btnBack)

	// options := &telebot.SendOptions{
	// 	ReplyMarkup: markup,
	// }

	return c.Respond(&telebot.CallbackResponse{
		Text: AdminMessages[lang]["adds_sent"],
	})

}

func (h *handlers) ChangeAdminLangHandle(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}

	btnUz := markup.Data("🇺🇿 Uzbek", "change_admin_lang", "uz")
	btnRu := markup.Data("🇷🇺 Russian", "change_admin_lang", "ru")
	btnEn := markup.Data("🇬🇧 English", "change_admin_lang", "en")
	btnTr := markup.Data("🇹🇷 Turkish", "change_admin_lang", "tr")
	btnBack := markup.Data(AdminMessages[lang]["back"], "back_to_admin_menu")

	markup.Inline(
		markup.Row(btnUz, btnRu),
		markup.Row(btnEn, btnTr),
		markup.Row(btnBack),
	)
	markup.ResizeKeyboard = true

	c.EditOrSend("Choose language:", markup)

	return nil
}

func (h *handlers) ChangeAdminLang(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}

	lang := c.Callback().Data

	langs, err := h.storage.ChangeAdminLang(userID, lang)

	if err != nil {
		return c.Send(err.Error())
	}

	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Data(AdminMessages[langs]["back"], "back_to_admin_menu")

	markup.Inline(markup.Row(btnBack))

	return c.Edit("Language changed", markup)
}
