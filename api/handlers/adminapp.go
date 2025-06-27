package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
)

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Photo      string `json:"photo"`
	CategoryID int    `json:"category_id"`
}

// func (h *handlers) GetProducts(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("SELECT id, name, price, photo, category_id FROM products")
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	defer rows.Close()

// 	var products []Product
// 	for rows.Next() {
// 		var p Product
// 		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Photo, &p.CategoryID)
// 		products = append(products, p)
// 	}

// 	json.NewEncoder(w).Encode(products)
// }

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
