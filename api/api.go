package api

import (
	"bot/api/handlers"
	"bot/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/telebot.v3"
)

type Options struct {
	Tg      telebot.Settings
	Storage storage.Storage
	R       *mux.Router
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

	bot.Handle(telebot.OnLocation, h.HandleLocation)

	bot.Handle(telebot.OnPhoto, h.AdminPhotostatus)

	bot.Handle(handlers.Messages["uz"]["lang_btn"], h.ChangeLanguage)

	bot.Handle(handlers.Messages["ru"]["lang_btn"], h.ChangeLanguage)

	bot.Handle(handlers.Messages["en"]["lang_btn"], h.ChangeLanguage)

	bot.Handle(handlers.Messages["tr"]["lang_btn"], h.ChangeLanguage)

	bot.Handle(handlers.Messages["uz"]["order_btn"], h.ShowMenu)

	bot.Handle(handlers.Messages["ru"]["order_btn"], h.ShowMenu)

	bot.Handle(handlers.Messages["en"]["order_btn"], h.ShowMenu)

	bot.Handle(handlers.Messages["tr"]["order_btn"], h.ShowMenu)

	bot.Handle(handlers.Messages["uz"]["my_orders"], h.ShowUserOrders)

	bot.Handle(handlers.Messages["ru"]["my_orders"], h.ShowUserOrders)

	bot.Handle(handlers.Messages["en"]["my_orders"], h.ShowUserOrders)

	bot.Handle(handlers.Messages["tr"]["my_orders"], h.ShowUserOrders)

	bot.Handle(handlers.Messages["uz"]["about_us"], h.SendAboutUs)

	bot.Handle(handlers.Messages["ru"]["about_us"], h.SendAboutUs)

	bot.Handle(handlers.Messages["en"]["about_us"], h.SendAboutUs)

	bot.Handle(handlers.Messages["tr"]["about_us"], h.SendAboutUs)

	// bot.Handle(&telebot.InlineButton{Unique: "payment_type"}, h.SetPaymentType)

	bot.Handle(&telebot.InlineButton{Unique: "continue_order"}, h.ShowMenu)

	bot.Handle(&telebot.InlineButton{Unique: "payment_type"}, h.CompleteOrder)

	bot.Handle(&telebot.InlineButton{Unique: "true-location"}, h.ChoosePaymentType)

	bot.Handle(&telebot.InlineButton{Unique: "false-location"}, h.HandleFalseLocation)

	bot.Handle(&telebot.InlineButton{Unique: "language_add"}, h.GetUserName)

	bot.Handle(&telebot.InlineButton{Unique: "confirm_order"}, h.RequestLocation)

	bot.Handle(&telebot.InlineButton{Unique: "lang_btn"}, h.ChangeLanguage)

	bot.Handle(&telebot.InlineButton{Unique: "my_orders"}, h.ShowUserOrders)

	bot.Handle(&telebot.InlineButton{Unique: "order_btn"}, h.ShowMenu)

	bot.Handle(&telebot.InlineButton{Unique: "about_us"}, h.SendAboutUs)

	bot.Handle(&telebot.InlineButton{Unique: "show_cart"}, h.SendCart)

	bot.Handle(&telebot.InlineButton{Unique: "get_category_by_id"}, h.ShowProducts)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_products_menu"}, h.ShowProducts)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_categories"}, h.ShowMenu)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_user_menu"}, h.ShowUserMenu)

	bot.Handle(&telebot.InlineButton{Unique: "add_category"}, h.CreateCategoryHandle)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_admin_menu"}, h.ShowCategoryMenu)

	bot.Handle(&telebot.InlineButton{Unique: "add_product"}, h.AddProductHandler)

	bot.Handle(&telebot.InlineButton{Unique: "clear_cart"}, h.ClearCart)

	bot.Handle(&telebot.InlineButton{Unique: "decrement_cart_product"}, h.HandleDecrement)

	bot.Handle(&telebot.InlineButton{Unique: "increment_cart_product"}, h.HandleIncrement)

	// bot.Handle("/admin", h.ShowAdminPanel)

	bot.Handle(&telebot.InlineButton{Unique: "get_product_by_id"}, h.ShowProductByID)

	bot.Handle(&telebot.InlineButton{Unique: "category_menu"}, h.ShowCategoryMenu)

	bot.Handle(&telebot.InlineButton{Unique: "change_status_preparing"}, h.ChangeOrderStatus)

	bot.Handle(&telebot.InlineButton{Unique: "change_status_deliver"}, h.ChangeOrderStatus)

	bot.Handle(&telebot.InlineButton{Unique: "change_status_completed"}, h.ChangeOrderStatus)

	bot.Handle(&telebot.InlineButton{Unique: "change_status_canceled"}, h.ChangeOrderStatus)

	bot.Handle(&telebot.InlineButton{Unique: "get_all_users"}, h.GetUsers)

	bot.Handle(&telebot.InlineButton{Unique: "close_day"}, h.CloseDay)

	bot.Handle(&telebot.InlineButton{Unique: "open_day"}, h.OpenDay)

	bot.Handle(&telebot.InlineButton{Unique: "add_admin"}, h.AddAdminHandle)

	bot.Handle(&telebot.InlineButton{Unique: "language_change"}, h.SetChangeLang)

	bot.Handle(&telebot.InlineButton{Unique: "update_category"}, h.ShowCategoryToUpdate)

	bot.Handle(&telebot.InlineButton{Unique: "get_category_info"}, h.GetCategoryInfo)

	bot.Handle(&telebot.InlineButton{Unique: "update_cat_name_uz"}, h.UpdateCategoryNameUzHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_cat_name_ru"}, h.UpdateCategoryNameRuHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_cat_name_en"}, h.UpdateCategoryNameUzHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_cat_name_tr"}, h.UpdateCategoryNameTrHandle)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_cat_menu"}, h.ShowCategoryMenu)

	bot.Handle(&telebot.InlineButton{Unique: "update_cat_availability"}, h.UpdateCategoryAvailability)

	bot.Handle(&telebot.InlineButton{Unique: "delete_cat"}, h.DeleteCategoryHandle)

	bot.Handle(&telebot.InlineButton{Unique: "delete_cat_yes"}, h.DeleteCategory)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_admin_menu"}, h.ShowAdminPanel)

	bot.Handle(&telebot.InlineButton{Unique: "product_menu"}, h.ShowProductMenu)

	bot.Handle(&telebot.InlineButton{Unique: "update_product"}, h.ShowCategoriesToUpdateProducts)

	bot.Handle(&telebot.InlineButton{Unique: "get_product_by_category"}, h.ShowProductsToUpdate)

	bot.Handle(&telebot.InlineButton{Unique: "add_prod_cat"}, h.AddProductToCategory)

	bot.Handle(&telebot.InlineButton{Unique: "get_product_info"}, h.GetProductInfo)

	bot.Handle(&telebot.InlineButton{Unique: "update_prod_name_uz"}, h.UpdateProductNameUzHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_prod_name_ru"}, h.UpdateProductNameRuHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_prod_name_en"}, h.UpdateProductNameEnHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_prod_name_tr"}, h.UpdateProductNameTrHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_prod_desc"}, h.UpdateProductDescHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_prod_price"}, h.UpdateProductPriceHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_prod_availability"}, h.UpdateProductAvailability)

	bot.Handle(&telebot.InlineButton{Unique: "delete_prod"}, h.DeleteProduct)

	bot.Handle(&telebot.InlineButton{Unique: "update_cat_of_prod"}, h.UpdateProductCategoryHandle)

	bot.Handle(&telebot.InlineButton{Unique: "upd_prod_cat"}, h.UpdateProductCategory)

	bot.Handle(&telebot.InlineButton{Unique: "cancel_order"}, h.CancelOrder)

	bot.Handle(&telebot.InlineButton{Unique: "send_adds"}, h.SendAddToUsersHandle)

	bot.Handle(&telebot.InlineButton{Unique: "change_admin_lang"}, h.ChangeAdminLang)

	bot.Handle(&telebot.InlineButton{Unique: "admin_lang"}, h.ChangeAdminLangHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update_photo_of_prod"}, h.UpdateProductPhotoHandle)

	bot.Handle(&telebot.InlineButton{Unique: "update"}, h.UpdateAdminPanel)

	bot.Handle(&telebot.InlineButton{Unique: "back_to_user_menu_from_orders"}, h.BackToMainMenuFromOrders)

	log.Println("Bot started...")

	go bot.Start()

	o.R.HandleFunc("/api/get-products", h.GetProducts).Methods("GET")
	o.R.HandleFunc("/api/products", h.AddProductSite2).Methods("POST")
	// o.R.HandleFunc("/api/products/{id}", UpdateProduct).Methods("PUT")
	// o.R.HandleFunc("/api/products/{id}", DeleteProduct).Methods("DELETE")

	// Categories
	o.R.HandleFunc("/api/categories", h.GetCategories).Methods("GET")
	// o.R.HandleFunc("/api/categories", AddCategory).Methods("POST")
	// o.R.HandleFunc("/api/categories/{id}", UpdateCategory).Methods("PUT")
	// o.R.HandleFunc("/api/categories/{id}", DeleteCategory).Methods("DELETE")

	// Admins
	o.R.HandleFunc("/api/check-admin", h.CheckAdmin).Methods("POST")
	// o.R.HandleFunc("/api/admins", GetAdmins).Methods("GET")
	// o.R.HandleFunc("/api/admins", GetAdmins).Methods("GET")
	// o.R.HandleFunc("/api/admins", h.AddAdmin).Methods("POST")
	// o.R.HandleFunc("/api/admins/{id}", UpdateAdmin).Methods("PUT")
	// o.R.HandleFunc("/api/admins/{id}", DeleteAdmin).Methods("DELETE")

	o.R.PathPrefix("/photos/").Handler(http.StripPrefix("/photos/", http.FileServer(http.Dir("./photos"))))


	http.ListenAndServe(":8080", enableCORS(o.R))
}

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
