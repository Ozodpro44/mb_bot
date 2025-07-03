package postgres

import (
	"bot/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Storage) GetAllProducts() (*models.Products, error) {
	rows, err := s.db.Query(`
        SELECT id, name_uz, name_ru, name_en, name_tr, price, photo, description, categories_id
        FROM products
        WHERE is_active = true
		ORDER BY name_uz ASC`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}
	defer rows.Close()
	var products models.Products
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.Name_uz, &product.Name_ru, &product.Name_en, &product.Name_tr, &product.Price, &product.Photo, &product.Description, &product.Category_id); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		products.Products = append(products.Products, &product)
	}

	return &products, nil
}

func (s *Storage) CreateProduct(product *models.Product) (*models.Product, error) {
	err := s.db.QueryRow(`
        INSERT INTO products (id, name_uz, name_ru, name_en, name_tr, price, photo, description, is_active)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, name_uz, name_ru, name_en, name_tr, price, photo, description, is_active`,
		uuid.New(),
		product.Name_uz,
		product.Name_ru,
		product.Name_en,
		product.Name_tr,
		product.Price,
		product.Photo,
		product.Description,
		product.Abelety).Scan(
			&product.ID, 
			&product.Name_uz, 
			&product.Name_ru, 
			&product.Name_en, 
			&product.Name_tr, 
			&product.Price, 
			&product.Photo, 
			&product.Description, 
			&product.Abelety)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %v", err)
	}
	return product, nil
}

func (s *Storage) AddProductToCategory(productID, categoryID string) error {
	_, err := s.db.Exec("UPDATE products SET categories_id = $1 WHERE id = $2", categoryID, productID)
	if err != nil {
		return fmt.Errorf("failed to add product to category: %v", err)
	}
	return nil
}

func (s *Storage) UpdateProductById(product_id string) error {
	// err := s.db.QueryRow("UPDATE products SET name_uz = $1, name_ru = $2, name_en = $3, price = $4, photo = $5, description = $6 WHERE id = $7", product.Name_uz, product.Name_ru, product.Name_en, product.Price, product.Photo, product.Description, product.ID)
	// if err != nil {
	// 	return fmt.Errorf("failed to update product: %v", err)
	// }
	return nil

}

func (s *Storage) DeleteProductById(product_id string) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id = $1", product_id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}
	return nil
}

func (s *Storage) GetProductsByCategory(categoryID string) (*models.Products, error) {
	rows, err := s.db.Query(`
        SELECT id, name_uz, name_ru, name_en, name_tr, price, photo, description, categories_id
        FROM products
        WHERE categories_id = $1 AND is_active = true
		ORDER BY name_uz ASC`, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}
	defer rows.Close()
	var products models.Products
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.Name_uz, &product.Name_ru, &product.Name_en, &product.Name_tr, &product.Price, &product.Photo, &product.Description, &product.Category_id); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		products.Products = append(products.Products, &product)
	}

	return &products, nil
}

func (s *Storage) GetProductById(product_id string) (*models.Product, error) {
	product := &models.Product{}
	err := s.db.QueryRow("SELECT id, name_uz, name_ru, name_en, name_tr, price, photo, description, categories_id FROM products WHERE id = $1", product_id).Scan(&product.ID, &product.Name_uz, &product.Name_ru, &product.Name_en, &product.Name_tr, &product.Price, &product.Photo, &product.Description, &product.Category_id)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (s *Storage) GetProductByName(product_name string) (*models.Product, error) {
	product := &models.Product{}
	err := s.db.QueryRow("SELECT id, name_uz, name_ru, name_en, name_tr, price, photo, description FROM products WHERE name_uz = $1 OR name_ru = $1 OR name_en = $1", product_name).Scan(&product.ID, &product.Name_uz, &product.Name_ru, &product.Name_en, &product.Name_tr, &product.Price, &product.Photo, &product.Description)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (s *Storage) GetProductsForAdmin() (*models.Products, error) {
	rows, err := s.db.Query("SELECT id, name_uz, name_ru, name_en, name_tr, price, photo, description, is_active, categories_id FROM products ORDER BY name_uz ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}
	defer rows.Close()
	var products models.Products
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.Name_uz, &product.Name_ru, &product.Name_en, &product.Name_tr, &product.Price, &product.Photo, &product.Description, &product.Abelety, &product.Category_id); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		products.Products = append(products.Products, &product)
	}

	return &products, nil
}

func (s *Storage) GetProductsByCategoryForAdmin(category_id string) (*models.Products, error) {
	rows, err := s.db.Query("SELECT id, name_uz, name_ru, name_en, name_tr, price, photo, description, categories_id FROM products WHERE categories_id = $1 ORDER BY name_uz ASC", category_id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}
	defer rows.Close()
	var products models.Products
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.Name_uz, &product.Name_ru, &product.Name_en, &product.Name_tr, &product.Price, &product.Photo, &product.Description, &product.Category_id); err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		products.Products = append(products.Products, &product)
	}

	return &products, nil

}

func (s *Storage) GetProductByIdForAdmin(product_id string) (*models.Product, error) {
	product := &models.Product{}
	err := s.db.QueryRow("SELECT id, name_uz, name_ru, name_en, name_tr, price, photo, description, categories_id, is_active FROM products WHERE id = $1", product_id).Scan(
		&product.ID, 
		&product.Name_uz, 
		&product.Name_ru, 
		&product.Name_en, 
		&product.Name_tr, 
		&product.Price, 
		&product.Photo, 
		&product.Description, 
		&product.Category_id, 
		&product.Abelety)
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (s *Storage) UpdateAbeletyProductById(product_id string) error {
	_, err := s.db.Exec("UPDATE products SET is_active = NOT is_active WHERE id = $1", product_id)
	if err != nil {
		return err
	}
	return nil

}

func (s *Storage) UpdateProductPriceById(product_id string, price int) error {
	_, err := s.db.Exec("UPDATE products SET price = $1 WHERE id = $2", price, product_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateProductDescById(product_id string, description string) error {
	_, err := s.db.Exec("UPDATE products SET description = $1 WHERE id = $2", description, product_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateProductNameUz(product_id string, name_uz string) error {
	_, err := s.db.Exec("UPDATE products SET name_uz = $1 WHERE id = $2", name_uz, product_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateProductNameRu(product_id string, name_ru string) error {
	_, err := s.db.Exec("UPDATE products SET name_ru = $1 WHERE id = $2", name_ru, product_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateProductNameEn(product_id string, name_en string) error {
	_, err := s.db.Exec("UPDATE products SET name_en = $1 WHERE id = $2", name_en, product_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateProductNameTr(product_id string, name_tr string) error {
	_, err := s.db.Exec("UPDATE products SET name_tr = $1 WHERE id = $2", name_tr, product_id)
	if err != nil {
		return err
	}
	return nil

}

func (s *Storage) UpdateProductCategoryById(product_id string, category_id string) error {
	_, err := s.db.Exec("UPDATE products SET categories_id = $1 WHERE id = $2", category_id, product_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateProductPhotoById(product_id string, photo string) error {
	_, err := s.db.Exec("UPDATE products SET photo = $1 WHERE id = $2", photo, product_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateProduct(product_id string, product *models.Product) error {
	_, err := s.db.Exec(`
        UPDATE products 
        SET name_uz = $1, name_ru = $2, name_en = $3, name_tr = $4, price = $5, photo = $6, categories_id = $7
        WHERE id = $8`,
		product.Name_uz,
		product.Name_ru,
		product.Name_en,
		product.Name_tr,
		product.Price,
		product.Photo,
		product.Category_id,
		product_id)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}
	return nil
}