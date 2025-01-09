package models

import (
	"time"

	"github.com/google/uuid"
)

type ClientsOrder struct {
	OrderID       uuid.UUID
	ClientOrderID uuid.UUID
	UserName      string
	Data          time.Time
}

type AdminOrder struct {
	OrderID uuid.UUID
	Order   string
	Price   string
}

type Order struct {
	OrderID      uuid.UUID
	Order_number int
	UserName     string
	Data         time.Time
	Order        string
	Price        string
}

type GetOrders struct {
	Count  int32
	Orders []*Order
}

type GetAdminOrders struct {
	AdminOrders []*AdminOrder
}

type GetAdminClient struct {
	UserName string
	Order    string
	Data     string
	Price    string
}

// const orders = `1. Type`

type Products struct {
	Count    int32
	Products []*Product
}

type Product struct {
	ID          string
	Name_uz     string
	Name_ru     string
	Name_en     string
	Price       int
	Photo       string
	Description string
	Created_at  time.Time
	Category_id string
	Abelety     bool
}

type Categories struct {
	Count      int32
	Categories []*Category
}

type Category struct {
	ID         string
	Name_uz    string
	Name_ru    string
	Name_en    string
	Abelety    bool
	Created_at time.Time
}

type Users struct {
	Count int32
	Users []*User
}

type User struct {
	ID           string
	TelegramID   int64
	Username     string
	Name         string
	Phone_Number string
	Lang         string
	Created_at   time.Time
}

type OrderDetails struct {
	OrderID            uuid.UUID
	Daily_order_number int
	Order_number       int
	Address            *Location
	Branch             *Branch
	UserID             uuid.UUID
	TotalPrice         int
	Status             string
	CreatedAt          string
	Delivery_type      string
	Delivery_price     int
	Items              []*Item
}

type GetBranch struct {
	Count   int32
	Branchs []*Branch
}

type Branch struct {
	ID   string
	Name string
	Lat  float64
	Lon  float64
}

type Item struct {
	Name_uz  string
	Name_ru  string
	Name_en  string
	Quantity int
	Price    int
}

type Cart_item struct {
	ProductID string
	Name_uz   string
	Name_ru   string
	Name_en   string
	Quantity  int
	Price     int
}

type Cart struct {
	Count int32
	Items []*Cart_item
}

type Admin struct {
	Id           string
	Admin_id     string
	Phone_Number string
	Password     string
	Created_at   time.Time
}

type Add struct {
	ID         string
	Text       string
	Photo      string
	Created_at time.Time
}

type Adds struct {
	Count int32
	Adds  []*Add
}

type GetLocation struct {
	Count     int32
	Locations []*Location
}

type Location struct {
	ID        string
	UserID    string
	Name_uz   string
	Name_ru   string
	Name_en   string
	Latitude  float64
	Longitude float64
}
