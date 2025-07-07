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
	Name_tr     string
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
	ID         string `json:"id"`
	Name_uz    string `json:"name"`
	Name_ru    string
	Name_en    string
	Name_tr    string
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
	OrderID            string
	Daily_order_number int
	Order_number       int
	Address            *Location
	Branch             *Branch
	UserID             string
	TotalPrice         int
	Status             string
	CreatedAt          string
	Delivery_type      string
	Items              []*Item
	PhoneNumber        string
	Payment_type       string
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
	Name_tr  string
	Quantity int
	Price    int
}

type Cart_item struct {
	ProductID string
	Name_uz   string
	Name_ru   string
	Name_en   string
	Name_tr   string
	Quantity  int
	Price     int
}

type Cart struct {
	Count int32
	Items []*Cart_item
}

type Admin struct {
	Id           string
	Admin_id     int64
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
	UserID    int64
	Name_uz   string
	Name_ru   string
	Name_en   string
	Name_tr   string
	Latitude  float32
	Longitude float32
}

type GeocodingResponse struct {
	Address struct {
		Country string `json:"country"`
		State   string `json:"state"`
		Suburb  string `json:"suburb"`
		Road    string `json:"road"`
	} `json:"address"`
}

type GetOrderMsg struct {
	MsgID  int
	UserID int64
	Lang   string
}

type Dashboard struct {
	TotalOrders         int                   `json:"totalOrders"`  // This Month
	TotalRevenue        float64               `json:"totalRevenue"` // This Month
	TotalUsers          int                   `json:"totalUsers"`
	AvgOrderValue       float64               `json:"avgOrderValue"` // This Month
	OrdersToday         int                   `json:"ordersToday"`
	RevenueToday        float64               `json:"revenueToday"`
	SatisfactionRate    float64               `json:"satisfactionRate"`
	ActiveBranches      int                   `json:"activeBranches"`
	TotalBranches       int                   `json:"totalBranches"`
	Trends              Trends                `json:"trends"`
	TopProducts         []TopProducts         `json:"topProducts"`
	RecentOrders        []RecentOrders        `json:"recentOrders"`
	CategoryPerformance []CategoryPerformance `json:"categoryPerformance"`
	PeakHours           []PeakHours           `json:"peakHours"`
	BranchPerformance   []BranchPerformance   `json:"branchPerformance"`
	SalesOverview       SalesOverview         `json:"salesOverview"`
}

type Trends struct {
	Orders  float64 `json:"orders"`
	Revenue float64 `json:"revenue"`
	Users   float64 `json:"users"`
	Aov     float64 `json:"aov"`
}

type TopProducts struct {
	Name    string  `json:"name"`
	Sold    int     `json:"sold"`
	Revenue float64 `json:"revenue"`
	Change  float64 `json:"change"`
}

type RecentOrders struct {
	OrderID  int     `json:"id"`
	Customer string  `json:"customer"`
	Total    float64 `json:"total"`
	Status   string  `json:"status"`
	Time     string  `json:"time"`
}

type CategoryPerformance struct {
	Name       string  `json:"name"`
	Revenue    float64 `json:"revenue"`
	Percentage float64 `json:"percentage"`
	Change     float64 `json:"change"`
}

type PeakHours struct {
	Hour       string  `json:"hour"`
	Orders     int     `json:"orders"`
	Percentage float64 `json:"percentage"`
}

type BranchPerformance struct {
	Name       string  `json:"name"`
	Revenue    float64 `json:"revenue"`
	Percentage float64 `json:"percentage"`
	Status     string  `json:"status"`
}

type SalesOverview struct {
	SevenDays  []int `json:"7d"`
	ThirtyDays []int `json:"30d"`
	NinetyDays []int `json:"90d"`
}
