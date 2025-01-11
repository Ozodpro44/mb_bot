package storage

import (
	// "bot/lib/e"
	"bot/models"
	// "crypto/sha1"
	// "io"
)

// import "go.starlark.net/lib/time"

type Storage interface {
	RegisterUser(user *models.User) error
	GetAllUsers() (*models.Users, error)
	CheckUserExist(telegramID int64) bool
	UserMessageStatus(telegramID int64, status string) (string, error)
	GetUserMessageStatus(telegramID int64) (string, error)
	SetDataUserMessageStatus(telegramID int64, data string) (string, error)
	GetDataUserMessageStatus(telegramID int64) (string, error)
	CreateLocation(location *models.Location) (*models.Location, error)
	GetLocationByID(telegramId int64) (*models.Location, error)
	DeleteLocationByUserID(telegramID int64) error
	// // Lang
	SetLangUser(telegramID int64, lang string) (string, error)
	GetLangUser(telegramID int64) (string, error)
	ChangeLangUser(telegramID int64, lang string) (string, error)
	// // Orders
	CreateOrder(userID int64, payment_type string) (string, error)
	GetOrderDetails(orderID string) (*[]models.OrderDetails, error)
	GetOrderByUserID(userID int64) (*[]models.OrderDetails, error)
	GetOrderDetailsByOrderID(orderID string) (*models.OrderDetails, error)
	ChangeOrderStatus(orderID string, status string) (string, error)
	// // Cart
	AddToCart(userID int64, productID string, quantity int) error
	GetCart(userID int64) (*models.Cart, error)
	RemoveFromCart(userID int64, productID string) error
	ClearCart(userID int64) error
	IncrementCartProductByID(userID int64, productID string) error
	DecrementCartProductByID(userID int64, productID string) error
	UpdateCart(userID int64, cart *models.Cart) error
	// // Categories
	CreateCategory(category *models.Category) error
	GetAllCategories() (*models.Categories, error)
	GetCategoryByID(category_id string) (*models.Category, error)
	UpdateCategoryById(category *models.Category) (*models.Category, error)
	UpdateAbeletyCategoryById(category_id string, abelety bool) error
	DeleteCategoryById(category_id string) error
	// // Products
	GetAllProducts() (*models.Products, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	UpdateProductById(product_id string) error
	DeleteProductById(product_id string) error
	GetProductsByCategory(categoryID string) (*models.Products, error)
	GetProductById(product_id string) (*models.Product, error)
	GetProductByName(product_name string) (*models.Product, error)
	// // Admin
	CheckAdmin(telegramID int64) bool
	CreateAdmin(admin *models.Admin) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin) (*models.Admin, error)
	DeleteAdmin(admin_id int64) error
	GetAdminLang(admin_id int64) (string, error)
	SetAdminLang(admin_id int64, lang string) (string, error)
	CloseDay() error
	CheckOpened() (bool, error)
	OpenDay() error
	// // Adds
	CreateAdds(adds *models.Add) (*models.Add, error)
	GetAllAdds() (*models.Adds, error)
	GetAddsById(id string) (*models.Add, error)
}
