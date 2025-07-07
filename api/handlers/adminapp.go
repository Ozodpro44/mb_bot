package handlers

import (
	// "bot/lib/helpers"
	"bot/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	// "github.com/gorilla/mux"
)

type Product struct {
	ID          string `json:"id"`
	Name        Name   `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Photo       string `json:"photo"`
	Abelety     bool   `json:"stock"`
	CategoryID  string `json:"category_id"`
}

type Name struct {
	Uz string `json:"uz"`
	Ru string `json:"ru"`
	En string `json:"en"`
	Tr string `json:"tr"`
}

type Category struct {
	ID   string `json:"id"`
	Name Name   `json:"name"`
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
	OrderID  string  `json:"id"`
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
	Name      string  `json:"name"`
	Revenue   float64 `json:"revenue"`
	Percetage float64 `json:"percetage"`
	Status    string  `json:"status"`
}

type SalesOverview struct {
	SevenDays  []int `json:"7d"`
	ThirtyDays []int `json:"30d"`
	NinetyDays []int `json:"90d"`
}

func (h *handlers) Dashboard(w http.ResponseWriter, r *http.Request) {
	dashboard, err := h.storage.GetDashboard()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(dashboard)

	// var topProducts []TopProducts
	// for _, p := range dashboard.TopProducts {
	// 	topProducts = append(topProducts, TopProducts{
	// 		Name:    p.Name,
	// 		Sold:    p.Sold,
	// 		Revenue: p.Revenue,
	// 		Change:  p.Change,
	// 	})
	// }

	// var recentOrders []RecentOrders
	// for _, o := range dashboard.RecentOrders {
	// 	recentOrders = append(recentOrders, RecentOrders{
	// 		OrderID:  o.OrderID,
	// 		Customer: o.Customer,
	// 		Total:    o.Total,
	// 		Status:   o.Status,
	// 		Time:     o.Time,
	// 	})
	// }

	// var categoryPerformance []CategoryPerformance
	// for _, c := range dashboard.CategoryPerformance {
	// 	categoryPerformance = append(categoryPerformance, CategoryPerformance{
	// 		Name:       c.Name,
	// 		Revenue:    c.Revenue,
	// 		Percentage: c.Percentage,
	// 		Change:     c.Change,
	// 	})
	// }

	// var peakHours []PeakHours
	// for _, p := range dashboard.PeakHours {
	// 	peakHours = append(peakHours, PeakHours{
	// 		Hour:       p.Hour,
	// 		Orders:     p.Orders,
	// 		Percentage: p.Percentage,
	// 	})
	// }

	// var branchPerformance []BranchPerformance
	// for _, b := range dashboard.BranchPerformance {
	// 	branchPerformance = append(branchPerformance, BranchPerformance{
	// 		Name:      b.Name,
	// 		Revenue:   b.Revenue,
	// 		Percetage: b.Percetage,
	// 		Status:    b.Status,
	// 	})
	// }

	// response := Dashboard{
	// 	TotalOrders:    dashboard.TotalOrders,
	// 	TotalRevenue:   dashboard.TotalRevenue,
	// 	TotalUsers:     dashboard.TotalUsers,
	// 	AvgOrderValue:  dashboard.AvgOrderValue,
	// 	OrdersToday:    dashboard.OrdersToday,
	// 	RevenueToday:   dashboard.RevenueToday,
	// 	SatisfactionRate: dashboard.SatisfactionRate,
	// 	ActiveBranches: dashboard.ActiveBranches,		
	// 	TotalBranches:  dashboard.TotalBranches,
	// 	Trends: Trends{
	// 		Orders:  dashboard.Trends.Orders,
	// 		Revenue: dashboard.Trends.Revenue,
	// 		Users:   dashboard.Trends.Users,
	// 		Aov:     dashboard.Trends.Aov,
	// 	},
	// 	TopProducts:         topProducts,
	// 	RecentOrders:        recentOrders,
	// 	CategoryPerformance: categoryPerformance,
	// 	PeakHours:           peakHours,
	// 	BranchPerformance:   branchPerformance,
	// 	SalesOverview: SalesOverview{
	// 		SevenDays:  dashboard.SalesOverview.SevenDays,
	// 		ThirtyDays: dashboard.SalesOverview.ThirtyDays,
	// 		NinetyDays: dashboard.SalesOverview.NinetyDays,
	// 	},
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboard)

}

func (h *handlers) GetProducts(w http.ResponseWriter, r *http.Request) {
	prod, err := h.storage.GetProductsForAdmin()

	if err != nil {
		http.Error(w, err.Error(), 500)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	var products []Product
	for _, p := range prod.Products {
		if _, err := os.Stat("./photos/" + p.Photo); os.IsNotExist(err) {
			p.Photo = "no_photo.jpg"
		}
		products = append(products, Product{
			ID: p.ID,
			Name: Name{
				Uz: p.Name_uz,
				Ru: p.Name_ru,
				En: p.Name_en,
				Tr: p.Name_tr,
			},
			Description: p.Description,
			Price:       p.Price,
			Photo:       "https://mbbot-production.up.railway.app/photos/" + p.Photo,
			Abelety:     p.Abelety,
			CategoryID:  p.Category_id,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *handlers) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Path[len("/api/products-by-category/"):]
	prod, err := h.storage.GetProductsByCategoryForAdmin(categoryID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var products []Product
	for _, p := range prod.Products {
		if _, err := os.Stat("./photos/" + p.Photo); os.IsNotExist(err) {
			p.Photo = "no_photo.jpg"
		}
		products = append(products, Product{
			ID: p.ID,
			Name: Name{
				Uz: p.Name_uz,
				Ru: p.Name_ru,
				En: p.Name_en,
				Tr: p.Name_tr,
			},
			Description: p.Description,
			Price:       p.Price,
			Photo:       "https://mbbot-production.up.railway.app/photos/" + p.Photo,
			Abelety:     p.Abelety,
			CategoryID:  p.Category_id,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *handlers) AddProductSite2(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Can't parse form", http.StatusBadRequest)
		return
	}

	name_uz := r.FormValue("name_uz")
	name_ru := r.FormValue("name_ru")
	name_en := r.FormValue("name_en")
	name_tr := r.FormValue("name_tr")
	priceStr := r.FormValue("price")
	categoryID := r.FormValue("category_id")
	price, _ := strconv.Atoi(priceStr)

	file, handler, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Photo required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file
	dst, err := os.Create("./photos/" + handler.Filename)
	if err != nil {
		http.Error(w, "Can't save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Can't save file", http.StatusInternalServerError)
		return
	}

	prod, _ := h.storage.CreateProduct(&models.Product{
		Name_uz: name_uz,
		Name_ru: name_ru,
		Name_en: name_en,
		Name_tr: name_tr,
		Price:   price,
		Photo:   handler.Filename,
	})
	h.storage.AddProductToCategory(prod.ID, categoryID)

	w.WriteHeader(http.StatusCreated)
}

func (h *handlers) AddProductSite(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Bad input", 400)
		return
	}

	// _, err := db.Exec("INSERT INTO products (name, price, photo, category_id) VALUES ($1, $2, $3, $4)",
	// 	p.Name, p.Price, p.Photo, p.CategoryID)
	// if err != nil {
	// 	http.Error(w, "DB error", 500)
	// 	return
	// }
	fmt.Println(p)
	w.WriteHeader(http.StatusCreated)
}

// func (h *handlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]

// 	var p Product
// 	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
// 		http.Error(w, "Invalid", 400)
// 		return
// 	}

// 	_, err := db.Exec("UPDATE products SET name=$1, price=$2, photo=$3, category_id=$4 WHERE id=$5",
// 		p.Name, p.Price, p.Photo, p.CategoryID, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// }

// func (h *handlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]
// 	_, err := db.Exec("DELETE FROM products WHERE id=$1", id)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 	}
// }

func (h *handlers) GetCategories(w http.ResponseWriter, r *http.Request) {
	// rows, err := h.db.Query("SELECT id, name FROM categories")
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// defer rows.Close()
	cat, err := h.storage.GetAllCategories()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var categories []Category
	for _, c := range cat.Categories {
		categories = append(categories, Category{
			ID: c.ID,
			Name: Name{
				Uz: c.Name_uz,
				Ru: c.Name_ru,
				En: c.Name_en,
				Tr: c.Name_tr,
			},
		})
	}

	json.NewEncoder(w).Encode(categories)
}

func (h *handlers) CheckAdmin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TelegramID int64 `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println(req.TelegramID)

	isAdmin := h.storage.CheckAdmin(req.TelegramID)

	fmt.Println(isAdmin)

	if isAdmin {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"role": "admin"})
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"role": "user"})
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(map[string]bool{"is_admin": })
}

func (h *handlers) DeleteProductSite(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/api/delete-products/"):]
	err := h.storage.DeleteProductById(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}

func (h *handlers) UpdateProductSite(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Can't parse form", http.StatusBadRequest)
		return
	}

	productID := r.URL.Path[len("/api/update-products/"):]

	name_uz := r.FormValue("name_uz")
	name_ru := r.FormValue("name_ru")
	name_en := r.FormValue("name_en")
	name_tr := r.FormValue("name_tr")
	priceStr := r.FormValue("price")
	categoryID := r.FormValue("category_id")
	price, _ := strconv.Atoi(priceStr)

	// Check if a new photo is uploaded
	file, handler, err := r.FormFile("photo")
	var photoPath string
	if err == nil { // Photo is uploaded
		defer file.Close()
		dst, err := os.Create("./photos/" + handler.Filename)
		if err != nil {
			http.Error(w, "Can't save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Can't save file", http.StatusInternalServerError)
			return
		}
		photoPath = handler.Filename
	} else if err == http.ErrMissingFile { // No new photo, keep existing
		prod, err := h.storage.GetProductByIdForAdmin(productID)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		photoPath = prod.Photo
	} else { // Other error
		http.Error(w, "Error processing photo", http.StatusInternalServerError)
		return
	}

	err = h.storage.UpdateProduct(productID, &models.Product{
		Name_uz:     name_uz,
		Name_ru:     name_ru,
		Name_en:     name_en,
		Name_tr:     name_tr,
		Price:       price,
		Photo:       photoPath,
		Category_id: categoryID,
	})
	if err != nil {
		http.Error(w, "DB error", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
