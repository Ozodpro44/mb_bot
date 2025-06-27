package handlers

import (
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
	ID         string `json:"id"`
	Name_uz    string `json:"name_uz"`
	Name_ru    string `json:"name_ru"`
	Name_en    string `json:"name_en"`
	Name_tr    string `json:"name_tr"`
	Price      int    `json:"price"`
	Photo      string `json:"photo"`
	CategoryID string `json:"category_id"`
}

func (h *handlers) GetProducts(w http.ResponseWriter, r *http.Request) {
	prod, err := h.storage.GetAllProducts()

	if err != nil {
		http.Error(w, err.Error(), 500)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	var products []struct {
		ID         string `json:"id"`
		Name_uz    string `json:"name_uz"`
		Name_ru    string `json:"name_ru"`
		Name_en    string `json:"name_en"`
		Name_tr    string `json:"name_tr"`
		Price      int    `json:"price"`
		Photo      string `json:"photo"`
		CategoryID string `json:"category_id"`
	}
	for _, p := range prod.Products {
		products = append(products, struct {
			ID         string `json:"id"`
			Name_uz    string `json:"name_uz"`
			Name_ru    string `json:"name_ru"`
			Name_en    string `json:"name_en"`
			Name_tr    string `json:"name_tr"`
			Price      int    `json:"price"`
			Photo      string `json:"photo"`
			CategoryID string `json:"category_id"`
		}{
			ID:         p.ID,
			Name_uz:    p.Name_uz,
			Name_ru:    p.Name_ru,
			Name_en:    p.Name_en,
			Name_tr:    p.Name_tr,
			Price:      p.Price,
			Photo:      "https://mbbot-production.up.railway.app/photos/" + p.Photo,
			CategoryID: p.Category_id,
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

	h.storage.CreateProduct(&models.Product{
		Name_uz:     name_uz,
		Name_ru:     name_ru,
		Name_en:     name_en,
		Name_tr:     name_tr,
		Price:       price,
		Photo:       handler.Filename,
		Category_id: categoryID,
	})

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

	var categories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	for _, c := range cat.Categories {
		categories = append(categories, struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{
			ID:   c.ID,
			Name: c.Name_uz,
		})
	}

	json.NewEncoder(w).Encode(categories)
}
