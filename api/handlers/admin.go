package handlers

import (
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
		"add_category":       "Add Category ➕",
		"add_product":        "Add Product ➕",
		"add_admin":          "Add Admin 👨‍⚖️",
		"close_day":          "Close Day🔐",
		"open_day":           "Open Day🔓",
		"branch":             "Branch 🏠",
		"get_all_users":      "Get All Users 👥",
		"category":           "Category name: /category :<UZ>:<RU>:<EN>:",
		"incorrect_category": "⁉️Incorrect format. Use: <UZ>,<RU>,<EN>",
		"category_created":   "Category created ✅",
		"back":               "🔙Back",
		"product":            "Use: \n/product :<product name uz>:<product name ru>:<product name en>:<description>:<price>:<availability>",
		"product_err":        "Incorrect format. Use: /category ,<product name uz>,<product name ru>,<product name en>,<description>,<price>,<availability>",
		"category_menu":      "Categories",
		"delete_category":    "Delete category",
		"add_admin_msg":      "Add admin: /admin ,<telegram_id>,<phone_number>,<password>",
		"admin_created":      "Admin created ✅",
		"day_closed":         "Day closed",
	},
	"ru": {
		"add_category":       "Добавить категорию ➕",
		"add_product":        "Добавить продукт ➕",
		"add_admin":          "Добавить админ 👨‍⚖️",
		"close_day":          "Закрыть день🔐",
		"open_day":           "Открыть день🔓",
		"branch":             "Филиал 🏠",
		"get_all_users":      "Получить всех пользователей 👥",
		"category":           "Название категории: /category :<UZ>:<RU>:<EN>:",
		"incorrect_category": "⁉️Неверный формат. Используйте: /category :<UZ>:<RU>:<EN>:",
		"category_created":   "Категория успешно создана ✅",
		"back":               "🔙Назад",
		"product":            "Используйте: \n/product :<имя продукта уз>:<имя продукта ру>:<имя продукта ен>:<описание>:<цена>:<доступность>",
		"product_err":        "Неверный формат. Используйте: /category ,<имя продукта уз>,<имя продукта ру>,<имя продукта ен>,<описание>,<цена>,<доступность>",
		"category_menu":      "Категории",
		"delete_category":    "Удалить категорию",
		"add_admin_msg":      "Добавить админа: /admin ,<telegram_id>,<phone_number>,<password>",
		"admin_created":      "Админ успешно создан ✅",
		"day_closed":         "День закрыт",
	},
	"uz": {
		"add_category":       "Kategoriya qo'shish ➕",
		"add_product":        "Mahsulot qo'shish ➕",
		"add_admin":          "Admin qo'shish 👨‍⚖️",
		"close_day":          "Kunni yopish🔐",
		"open_day":           "Kunni ochish🔓",
		"branch":             "Filial 🏠",
		"get_all_users":      "Barcha foydalanuvchilarni olish 👥",
		"category":           "Kategoriya nomi: /category :<UZ>:<RU>:<EN>:",
		"incorrect_category": "⁉️Iltimos, to'g'ri formatdan foydalaning: /category :<UZ>:<RU>:<EN>:",
		"category_created":   "Kategoriya muvaffaqiyatli yaratildi ✅",
		"back":               "🔙Orqaga",
		"product":            "Mahsulot qo'shish uchun: \n/product :<nomi uz>:<nomi ru>:<nomi en>:<description>:<narxi>:<availability>",
		"product_err":        "Iltimos, to'g'ri formatdan foydalaning: /category ,<nomi uz>,<nomi ru>,<nomi en>,<description>,<narxi>,<availability>",
		"category_menu":      "Kategoriyalar",
		"delete_category":    "Kategoriyani o'chirish",
		"add_admin_msg":      "Admin qo'shish: /admin ,<telegram_id>,<phone_number>,<password>",
		"admin_created":      "Admin muvaffaqiyatli yaratildi ✅",
		"day_closed":         "Kun yopildi",
		"day_opened":         "Kun ochildi",
	},
}

// var categoryStep = make(map[int64]string)
// var categoryData = make(map[int64]map[string]string)

func (h *handlers) ShowAdminPanel(c telebot.Context) error {
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
	btnCat := menu.Data(AdminMessages[lang]["category_menu"], "category_menu")
	btnProd := menu.Data(AdminMessages[lang]["add_product"], "add_product")
	btnAdmins := menu.Data(AdminMessages[lang]["add_admin"], "add_admin")
	btnFilial := menu.Data(AdminMessages[lang]["branch"], "branch")

	// Arrange buttons in rows
	menu.Inline(
		menu.Row(btnCat),
		menu.Row(btnProd, btnAdmins),
		menu.Row(btnFilial),
	)
	menu.ResizeKeyboard = true

	c.EditOrSend("Admin panel:", menu)

	return nil
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
	btnProd := menu.Data(AdminMessages[lang]["add_product"], "add_product")
	btnAdmins := menu.Data(AdminMessages[lang]["add_admin"], "add_admin")
	// btnFilial := menu.Data(AdminMessages[lang]["branch"], "branch")
	btnCloseDay := menu.Data(AdminMessages[lang]["close_day"], "close_day")
	btnOpenDay := menu.Data(AdminMessages[lang]["open_day"], "open_day")
	btnGetUsers := menu.Data(AdminMessages[lang]["get_all_users"], "get_all_users")

	// Arrange buttons in rows
	menu.Inline(
		menu.Row(btnCat),
		menu.Row(btnProd, btnAdmins),
		menu.Row(btnCloseDay, btnOpenDay),
		menu.Row(btnGetUsers),
	)
	menu.ResizeKeyboard = true

	c.EditOrSend("Admin panel:", menu)
	return nil
}

func (h *handlers) CreateCategory(c telebot.Context) error {
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

	err = c.Edit(AdminMessages[lang]["category"], options)

	if err != nil {
		fmt.Println(err)
	}
	category := models.Category{}

	c.Bot().Handle("/category", func(ct telebot.Context) error {
		if ct.Sender().ID == userID {

			text := ct.Text()

			arg := strings.Split(text, ",")

			if len(arg) != 4 {
				return c.Edit(AdminMessages[lang]["incorrect_category"], options)
			}

			nameUz := strings.TrimSpace(arg[1])
			nameRu := strings.TrimSpace(arg[2])
			nameEn := strings.TrimSpace(arg[3])
			fmt.Println(nameUz)

			category.Name_uz = nameUz
			category.Name_ru = nameRu
			category.Name_en = nameEn
			category.Abelety = true

			h.storage.CreateCategory(&category)

			ct.Send("Category created", options)
		} else {
			return c.Send("You are not admin")
		}

		return nil
	})
	return nil
}

func (h *handlers) GetUsers(c telebot.Context) error {
	userID := c.Sender().ID
	if !h.storage.CheckAdmin(userID) {
		return c.Send("You are not admin")
	}
	users, err := h.storage.GetAllUsers()
	if err != nil {
		return c.Send(err.Error())
	}

	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "User ID")
	f.SetCellValue("Sheet1", "B1", "Username")
	f.SetCellValue("Sheet1", "C1", "Phone Number")
	f.SetCellValue("Sheet1", "D1", "Name")

	row := 2

	for _, data := range users.Users {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), data.TelegramID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), data.Username)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), data.Phone_Number)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), data.Name)
		row++
	}

	filename := fmt.Sprintf("user_data_%d.xlsx", time.Now().Unix())
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
	return h.ShowCategoryMenu(c)
}

// func (h handlers) AddCategoryHandler(bot *telebot.Bot) func(c telebot.Context) error {
// 	return func(c telebot.Context) error {
// 		// Ожидаем название категории от администратора
// 		if c.Message().Payload == "" {
// 			return c.Send("Введите название новой категории:")
// 		}

// 		categoryName := c.Message().Payload

// 		// Сохранение категории в базе данных
// 		err := h.storage.CreateCategory(categoryName)
// 		if err != nil {
// 			log.Println("Ошибка при добавлении категории:", err)
// 			return c.Send("Ошибка при добавлении категории.")
// 		}

// 		return c.Send(fmt.Sprintf("Категория '%s' успешно добавлена!", categoryName))
// 	}
// }

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

	product := &models.Product{}
	markup := &telebot.ReplyMarkup{}

	btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "back_to_admin_menu"))

	markup.Inline(btnBack)

	options := &telebot.SendOptions{
		ReplyMarkup: markup,
	}

	err = c.Edit(AdminMessages[lang]["product"], options)

	if err != nil {
		fmt.Println(err)
	}

	c.Bot().Handle("/product", func(c telebot.Context) error {
		text := c.Text()

		// if text == "Back" {
		// 	h.ShowAdminPanel(c)
		// 	return nil
		// }
		args := strings.Split(text, ",")
		if len(args) < 6 {
			return c.Send(Messages[lang]["product_err"])
		}

		nameUz := strings.TrimSpace(args[1])
		nameRu := strings.TrimSpace(args[2])
		nameEn := strings.TrimSpace(args[3])
		description := strings.TrimSpace(args[4])
		fmt.Println(nameUz, nameRu, nameEn, description, args[5])
		price, err := strconv.Atoi(strings.TrimSpace(args[5]))
		if err != nil {
			return c.Send("Ошибка в формате цены. Убедитесь, что это число.")
		}
		availability, err := strconv.ParseBool(strings.TrimSpace(args[6]))
		if err != nil {
			return c.Send("Ошибка в формате доступности. Используйте true или false.")
		}

		product = &models.Product{
			Name_uz:     nameUz,
			Name_ru:     nameRu,
			Name_en:     nameEn,
			Description: description,
			Price:       price,
			Abelety:     availability,
		}
		// c.Send("Get category")

		return h.AddProductToCategory(c, product)
	})
	return nil
}

func (h *handlers) AddProductToCategory(c telebot.Context, product *models.Product) error {
	btn := &telebot.ReplyMarkup{}
	btn.ResizeKeyboard = true
	btn.OneTimeKeyboard = true
	var buttons []telebot.Row

	categories, err := h.storage.GetAllCategories()
	if err != nil {
		return c.Send(err.Error())
	}

	// var buttons []telebot.InlineButton
	for _, cat := range categories.Categories {
		button := btn.Data(fmt.Sprintf("%s / %s / %s", cat.Name_uz, cat.Name_ru, cat.Name_en), cat.ID)
		buttons = append(buttons, btn.Row(button))
	}
	btn.Inline(buttons...)
	// Создаем строки кнопок и добавляем в Inline-клавиатуру

	if err := c.Send("Выберите категорию:", btn); err != nil {
		return err
	}

	// Обработчик кнопок
	c.Bot().Handle(telebot.OnCallback, func(c telebot.Context) error {
		uid := c.Callback().Data
		categoryID, err := uuid.Parse(uid[len(uid)-36:])
		if err != nil {
			return err
		}

		product.Category_id = categoryID.String()
		return c.Send("Пожалуйста, 	прикрепите фото продукта.")
	})

	return h.AddPhotoToProduct(c, product)
}

func (h *handlers) AddPhotoToProduct(c telebot.Context, product *models.Product) error {
	c.Bot().Handle(telebot.OnPhoto, func(c telebot.Context) error {
		pic := c.Message().Photo
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

		// err = helpers.DownloadPhoto(photoURL, photoPath)

		if err != nil {
			return c.Send(fmt.Sprintf("Ошибка сохранения фото: %v", err))
		}

		c.Send("Фото получено")

		product.Photo = photoPath

		_, err = h.storage.CreateProduct(product)
		if err != nil {
			log.Printf("Ошибка добавления продукта: %v", err)
			return c.Send(fmt.Sprintf("Не удалось добавить продукт: %v", err))
		}

		// Удаляем продукт из временного хранилища после успешного добавлени

		return c.Send(fmt.Sprintf("Продукт '%s' успешно добавлен в категорию '%s'.", product.Name_uz, product.Category_id))
	})

	return nil
}

func (h *handlers) AddAdmin(c telebot.Context) error {
	userID := c.Sender().ID
	// Check if the user is an admin
	if !h.storage.CheckAdmin(userID) {
		return c.Send("У вас нет прав для выполнения этой команды.")
	}

	lang, err := h.storage.GetAdminLang(userID)

	if err != nil {
		return c.Send(err.Error())
	}

	admin := &models.Admin{}

	c.Send(AdminMessages[lang]["add_admin_msg"])
	c.Bot().Handle("/admin", func(c telebot.Context) error {
		text := c.Text()

		if text == "Back" {
			h.ShowAdminPanel(c)
			return nil
		}
		args := strings.Split(text, ",")
		if len(args) < 6 {
			return c.Send(Messages[lang]["product_err"])
		}

		telegramID, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return c.Send("Ошибка в формате Телеграм ID. Убедитесь, что это число.")
		}
		phone := strings.TrimSpace(args[2])
		password := strings.TrimSpace(args[3])

		admin = &models.Admin{
			Admin_id:     telegramID,
			Phone_Number: phone,
			Password:     password,
		}
		// c.Send("Get category")

		_, err = h.storage.CreateAdmin(admin)

		if err != nil {
			return c.Send(err.Error())
		}

		markup := &telebot.ReplyMarkup{OneTimeKeyboard: true}

		btnBack := markup.Row(markup.Data(AdminMessages[lang]["back"], "back_to_admin_menu"))

		markup.Inline(btnBack)

		options := &telebot.SendOptions{
			ReplyMarkup: markup,
		}

		c.Send(AdminMessages[lang]["admin_created"], options)
		return nil
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

	c.Send(AdminMessages[lang]["day_closed"])
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

	c.Send(AdminMessages[lang]["day_opened"])
	return nil
}

// var photoPath string
// pic := c.Message().Photo
// url := pic.FileURL
// photoPath = fmt.Sprintf("./photos/%s.jpg", product.Name_uz)
// err := helpers.DownloadPhoto(url, photoPath)
// if err != nil {
// 	return c.Send(fmt.Sprintf("Ошибка сохранения фото: %v", err))
// }
// c.Send("Photo received")

// _, err = h.storage.CreateProduct(product)
// if err != nil {
// 	log.Printf("Ошибка добавления продукта: %v", err)
// 	return c.Send(fmt.Sprintf("Не удалось добавить продукт: %v", err))
// }

// return c.Send(fmt.Sprintf("Продукт '%s' успешно добавлен в категорию '%s'.", product.Name_uz, product.Category_id))

// func (h handlers) AddProductToCategory(c telebot.Context, product *models.Product) error {
// 	btn := telebot.ReplyMarkup{}

// 	categories, err := h.storage.GetAllCategories()

// 	if err != nil {
// 		return c.Send(err.Error())
// 	}
// 	btn.ResizeKeyboard = true
// 	for _, c := range categories.Categories {
// 		cat := btn.Data(fmt.Sprintf("%s/ %s/ %s", c.Name_uz, c.Name_ru, c.Name_en), c.ID, "create_product")
// 		btn.Row(cat)
// 	}
// 	c.Send(btn,"M")

// 	c.Bot().Handle(telebot.InlineButton{Data: "create_product"}, func(c telebot.Context) error {
// 		category := c.Data()
// 		product.Category_id = category
// 		return nil
// 	})
// 	// Add the product to storage

// 	return c.Send("Пожалуйста, прикрепите фото продукта.")
// }

// for update := range updates {
// 	if update.Message.Photo != nil && update.Message.Chat.ID == c.Chat().ID {
// 		photo := update.Message.Photo // Take the highest resolution photo
// 		// if err != nil {
// 		// 	return c.Send(fmt.Sprintf("Ошибка загрузки фото: %v", err))
// 		// }
// 		photoURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, photo.FilePath)
// 		photoPath = fmt.Sprintf("./photos/%s.jpg", nameUz)

// 		// Download and save the photo locally
// 		err = helpers.DownloadPhoto(photoURL, photoPath)
// 		if err != nil {
// 			return c.Send(fmt.Sprintf("Ошибка сохранения фото: %v", err))
// 		}
// 		break
// 	}
// }

// c.Bot().Handle(telebot.OnPhoto, func(c telebot.Context) error {
// 	if c.Message().Photo != nil {
// 		photo := c.Message().Photo // Take the highest resolution photo
// 		if err != nil {
// 			return c.Send(fmt.Sprintf("Ошибка загрузки фото: %v", err))
// 		}
// 		photoURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, photo.FilePath)
// 		photoPath = fmt.Sprintf("./photos/%s.jpg", nameUz)

// 		// Download and save the photo locally
// 		err = helpers.DownloadPhoto(photoURL, photoPath)
// 		if err != nil {
// 			return c.Send(fmt.Sprintf("Ошибка сохранения фото: %v", err))
// 		}
// 	} else {
// 		return c.Send("Пожалуйста, прикрепите фото продукта.")
// 	}
// 	return nil
// })

// Handle photo upload

// func (h handlers) AddProductHandler(bot *telebot.Bot) func(c telebot.Context) error {
// 	return func(c telebot.Context) error {
// 		// Проверяем состояние администратора
// 		adminState := h.Storage.GetAdminState(c.Sender().ID)

// 		switch adminState {
// 		case "WAITING_CATEGORY":
// 			// Ожидаем выбор категории
// 			category := c.Text()
// 			h.storage.SetAdminState(c.Sender().ID, "WAITING_PRODUCT_NAME", category)
// 			return c.Send("Введите название продукта:")

// 		case "WAITING_PRODUCT_NAME":
// 			// Ожидаем название продукта
// 			productName := c.Text()
// 			h.Storage.SetAdminState(c.Sender().ID, "WAITING_PRODUCT_DESC", productName)
// 			return c.Send("Введите описание продукта:")

// 		case "WAITING_PRODUCT_DESC":
// 			// Ожидаем описание продукта
// 			productDesc := c.Text()
// 			productName := h.Storage.GetTemporaryData(c.Sender().ID, "PRODUCT_NAME")
// 			category := h.Storage.GetTemporaryData(c.Sender().ID, "CATEGORY")

// 			// Сохранение продукта в базе данных
// 			err := h.Storage.AddProduct(category, productName, productDesc)
// 			if err != nil {
// 				log.Println("Ошибка при добавлении продукта:", err)
// 				return c.Send("Ошибка при добавлении продукта.")
// 			}

// 			h.Storage.SetAdminState(c.Sender().ID, "")
// 			return c.Send(fmt.Sprintf("Продукт '%s' успешно добавлен в категорию '%s'!", productName, category))

// 		default:
// 			// Начало добавления продукта
// 			h.Storage.SetAdminState(c.Sender().ID, "WAITING_CATEGORY")
// 			categories := h.Storage.GetCategories()
// 			categoryButtons := &telebot.ReplyMarkup{}
// 			for _, category := range categories {
// 				categoryButtons.Reply(categoryButtons.Row(categoryButtons.Text(category)))
// 			}
// 			return c.Send("Выберите категорию для нового продукта:", categoryButtons)
// 		}
// 	}
// }

// func (h handlers) GetUsers(telegram_id int64) {

// 	admin:= h.storage.CheckAdmin(telegram_id)

// 	if !admin {
// 		return
// 	}

// 	users, err := h.storage.GetAllUsers()

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}

// 	msg := "User \n"

// 	for i, user := range users.Users {
// 		msg += fmt.Sprintf("%v. id: %v - %s \n %s \n", i+1, user.ID, user.Username, user.Phone_Number)
// 	}

// 	h.tg.SendMessages(msg, telegram_id)
// }

// func (h handlers) SetAdmin(updates *tgbotapi.UpdatesChannel, telegram_id int64) {

// 	adminInfo := models.Admin{}

// 	admin:= h.storage.CheckAdmin(telegram_id)

// 	if !admin {
// 		return
// 	}

// 	keyboard := tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Cancel"),
// 		),
// 	)

// 	h.tg.SendReplyKeyboard("Phone Number:", telegram_id, keyboard)

// 	for update := range *updates {
// 		adminInfo.Phone_Number = update.Message.Text
// 		break
// 	}

// 	h.tg.SendMessages("Phone Number:", telegram_id)

// 	for update := range *updates {
// 		if update.Message.Text == "Cancel" {
// 			return
// 		}
// 		phone, err := helpers.FormatPhoneNumber(update.Message.Text)

// 		if err != nil {
// 			h.tg.SendMessages(err.Error(), telegram_id)
// 			return
// 		} else {
// 			adminInfo.Phone_Number = phone
// 			break
// 		}

// 	}

// 	h.tg.SendMessages("User id", telegram_id)

// 	for update := range *updates {
// 		if update.Message.Text == "Cancel" {
// 			return
// 		} else {
// 			adminInfo.Admin_id = update.Message.Text
// 			break
// 		}
// 	}

// 	adminSet, err := h.storage.CreateAdmin(&adminInfo)

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}

// 	msg := tgbotapi.NewMessage(telegram_id, fmt.Sprintf("Admin sucessfully added: \n %s \n %s \n %s", adminSet.Admin_id, adminSet.Phone_Number, adminSet.Phone_Number))
// 	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // Pass `true` to remove for the current chat
// 	h.tg.SendMsg(msg)
// }

// func (h handlers) SendToAllUser(updates *tgbotapi.UpdatesChannel, telegram_id int64) {

// 	admin:= h.storage.CheckAdmin(telegram_id)

// 	if !admin {
// 		log.Println("not admin")
// 		return
// 	}

// 	keyboard := tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Adds"),
// 			tgbotapi.NewKeyboardButton("Create Add"),
// 		),
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Cancel"),
// 		),
// 	)

// 	h.tg.SendReplyKeyboard("Choose:", telegram_id, keyboard)

// 	for update := range *updates {
// 		if update.CallbackData() != "" {
// 			h.tg.SendMessages("Unknown command", telegram_id)
// 		}
// 		if update.Message == nil {
// 			continue
// 		}
// 		switch update.Message.Text {
// 		case "Adds":
// 			h.Adds(updates, telegram_id)
// 			return
// 		case "Create Add":
// 			h.CreateAdd(updates, telegram_id)
// 			return
// 		case "Cancel":
// 			return
// 		default:
// 			h.tg.SendMessages("Unknown command", telegram_id)
// 		}
// 	}

// }

// func (h handlers) Adds(updates *tgbotapi.UpdatesChannel, telegram_id int64) {
// 	IDs := []int{}
// 	inline := tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Cancel"),
// 		),
// 	)

// 	IDs = append(IDs, h.tg.SendReplyKeyboard("Choose Add:", telegram_id, inline))

// 	adds, err := h.storage.GetAllAdds()

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}

// 	for _, add := range adds.Adds {
// 		photoFile, err := os.Open(add.Photo)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		defer photoFile.Close()

// 		button := tgbotapi.NewInlineKeyboardButtonData("Send This", add.ID)
// 		button2 := tgbotapi.NewInlineKeyboardButtonData("Delte", add.ID)

// 		// Arrange buttons in a keyboard layout
// 		keyboard := tgbotapi.NewInlineKeyboardMarkup(
// 			tgbotapi.NewInlineKeyboardRow(button, button2),
// 		)

// 		// Create a new photo message
// 		photo := tgbotapi.NewPhoto(telegram_id, tgbotapi.FileReader{
// 			Name:   "photo.jpg",
// 			Reader: photoFile,
// 		})
// 		photo.Caption = add.Text
// 		photo.ParseMode = "Markdown"
// 		photo.ReplyMarkup = keyboard
// 		IDs = append(IDs, h.tg.SendMsg(photo))
// 	}

// 	var Add_ID string

// 	for update := range *updates {
// 		if update.CallbackQuery != nil {
// 			Add_ID = update.CallbackQuery.Data
// 			break
// 		}else if update.Message.Text == "Cancel" {
// 			h.SendToAllUser(updates, telegram_id)
// 			return
// 		}
// 	}

// 	users, err := h.storage.GetAllUsers()

// 	if err != nil {
// 		log.Printf("Error on get users: %v",telegram_id)
// 		return
// 	}

// 	add, err := h.storage.GetAddsById(Add_ID)

// 	if err != nil {
// 		log.Printf("Error on get add: %v",telegram_id)
// 		return
// 	}

// 	for _, user := range users.Users {
// 		photoFile, err := os.Open(add.Photo)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		defer photoFile.Close()

// 		// Create a new photo message
// 		photo := tgbotapi.NewPhoto(user.TelegramID, tgbotapi.FileReader{
// 			Name:   "photo.jpg",
// 			Reader: photoFile,
// 		})
// 		photo.Caption = add.Text
// 		photo.ParseMode = "Markdown"
// 		h.tg.SendMsg(photo)
// 	}
// 	for _, id := range IDs {
// 		h.tg.DeleteMessage(telegram_id, id)
// 	}
// 	h.tg.SendMessages("Successfully sent", telegram_id)
// 	h.SendToAllUser(updates, telegram_id)
// }

// func (h handlers) CreateAdd(updates *tgbotapi.UpdatesChannel, telegram_id int64) {
// 	add := models.Add{}
// 	h.tg.SendMessages("Photo:", telegram_id)

// 	for update := range *updates {
// 		if update.Message != nil { // Check if the update contains a message
// 			if update.Message.Photo != nil { // Check if the message contains a photo
// 				photos := update.Message.Photo
// 				photo := photos[len(photos)-1] // Get the highest resolution photo

// 				// Get file details from Telegram
// 				fileConfig := tgbotapi.FileConfig{FileID: photo.FileID}
// 				file, err := h.tg.GetFile(fileConfig)
// 				if err != nil {
// 					fmt.Println("Failed to get file:", err)
// 					continue
// 				}

// 				// Construct the file URL
// 				fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", "7635834906:AAF-inAvfxCydE5o1mCtDHoDcI3_0j5bIo8", file.FilePath)

// 				// Download the photo
// 				err = helpers.DownloadPhoto(fileURL, file.FilePath)
// 				if err != nil {
// 					fmt.Println("Failed to download photo:", err)
// 					continue
// 				}

// 				add.Photo = file.FilePath

// 				fmt.Println("Photo downloaded successfully!")

// 				// Extract caption (description)
// 				caption := update.Message.Caption
// 				if caption != "" {
// 					fmt.Println("Photo description:", caption)
// 					add.Text = caption
// 					break
// 				} else {
// 					fmt.Println("No description provided.")
// 				}
// 			}
// 		}
// 	}
// 	_, err := h.storage.CreateAdds(&add)

// 	if err != nil {
// 		h.tg.SendMessages(err.Error(), telegram_id)
// 		return
// 	}

// 	h.tg.SendMessages("Add successfully created", telegram_id)
// 	h.SendToAllUser(updates, telegram_id)
// }
