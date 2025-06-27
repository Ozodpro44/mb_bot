package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	// "github.com/gorilla/mux"
)

type Product struct {
	// ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Photo      string `json:"photo"`
	CategoryID int    `json:"category_id"`
}

func (h *handlers) GetProducts(w http.ResponseWriter, r *http.Request) {
	prod, err := h.storage.GetAllProducts()

	if err != nil {
		http.Error(w, err.Error(), 500)
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	var products []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Price      int    `json:"price"`
		Photo      string `json:"photo"`
		CategoryID string `json:"category_id"`
	}
	for _, p := range prod.Products {
		products = append(products, struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Price      int    `json:"price"`
			Photo      string `json:"photo"`
			CategoryID string `json:"category_id"`
		}{
			ID:         p.ID,
			Name:       p.Name_uz,
			Price:      p.Price,	
			Photo:      p.Photo,
			CategoryID: p.Category_id,
		})
	}
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *handlers) AddProductSite2(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Can't parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	priceStr := r.FormValue("price")
	categoryIDStr := r.FormValue("category_id")
	price, _ := strconv.Atoi(priceStr)
	categoryID, _ := strconv.Atoi(categoryIDStr)

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
	io.Copy(dst, file)

	// Save product to DB with photo name
	// _, err = db.Exec(`INSERT INTO products (name, price, photo, category_id) VALUES ($1, $2, $3, $4)`,
	// 	name, price, "/photos/"+handler.Filename, categoryID)

	// if err != nil {
	// 	http.Error(w, "DB error", 500)
	// 	return
	// }
	fmt.Println(name, price, "/photos/"+handler.Filename, categoryID)

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
