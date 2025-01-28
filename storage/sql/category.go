package postgres

import (
	"bot/models"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *Storage) CreateCategory(category *models.Category) error {
	_, err := s.db.Exec("INSERT INTO categories (id, name_uz, name_ru, name_en, name_tr, abelety, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", uuid.New(), category.Name_uz, category.Name_ru, category.Name_en, category.Name_tr, category.Abelety, time.Now())
	if err != nil {
		return fmt.Errorf("failed to add category: %v", err)
	}
	return nil
}

func (s *Storage) GetAllCategories() (*models.Categories, error) {
	rows, err := s.db.Query("SELECT id, name_uz, name_ru, name_en, name_tr, abelety, created_at FROM categories WHERE abelety = true ORDER BY name_uz ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %v", err)
	}
	defer rows.Close()
	
	var categories models.Categories
	for rows.Next() {
		category := models.Category{}
		if err := rows.Scan(&category.ID, &category.Name_uz, &category.Name_ru, &category.Name_en, &category.Name_tr, &category.Abelety, &category.Created_at); err != nil {
			return nil, fmt.Errorf("failed to scan category: %v", err)	
		}
		categories.Categories = append(categories.Categories, &category)
	}
	// log.Println(err)
	
	log.Println("Categotegories got")
	return &categories, nil
}

func (s *Storage) GetCategoriesForAdmin() (*models.Categories, error) {
	rows, err := s.db.Query("SELECT id, name_uz, name_ru, name_en, name_tr, abelety, created_at FROM categories ORDER BY name_uz ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %v", err)
	}
	defer rows.Close()
	
	var categories models.Categories
	for rows.Next() {
		category := models.Category{}
		if err := rows.Scan(&category.ID, &category.Name_uz, &category.Name_ru, &category.Name_en, &category.Name_tr, &category.Abelety, &category.Created_at); err != nil {
			return nil, fmt.Errorf("failed to scan category: %v", err)
		}
		categories.Categories = append(categories.Categories, &category)
	}
	// log.Println(err)
	
	log.Println("Categotegories got for admin")
	return &categories, nil
}

func (s *Storage) GetCategoryByID(category_id string) (*models.Category, error) {
	category := &models.Category{}
	err := s.db.QueryRow("SELECT id, name_uz, name_ru, name_en, name_tr, abelety FROM categories WHERE id = $1", category_id).Scan(&category.ID, &category.Name_uz, &category.Name_ru, &category.Name_en, &category.Name_tr, &category.Abelety)
	if err != nil {
		return nil, err
	}
	return category, nil

}

func (s *Storage) UpdateCategoryById(category *models.Category) (*models.Category, error) {
	err := s.db.QueryRow("UPDATE categories SET name_uz = $1, name_ru = $2, name_en = $3 abelety = $4 WHERE id = $5, RETURNING id, name_uz, name_ru, name_en, name_tr, abelety", category.Name_uz, category.Name_ru, category.Name_en, category.Name_tr, category.ID).Scan(&category.ID, &category.Name_uz, &category.Name_ru, &category.Name_en, &category.Name_tr, &category.Abelety)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *Storage) UpdateAbeletyCategoryById(category_id string) error {
	_, err := s.db.Exec("UPDATE categories SET abelety = NOT abelety WHERE id = $1", category_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateNameUzCategoryById(category_id string, name_uz string) error {
	_, err := s.db.Exec("UPDATE categories SET name_uz = $1 WHERE id = $2", name_uz, category_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateNameRuCategoryById(category_id string, name_ru string) error {
	_, err := s.db.Exec("UPDATE categories SET name_ru = $1 WHERE id = $2", name_ru, category_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateNameEnCategoryById(category_id string, name_en string) error {
	_, err := s.db.Exec("UPDATE categories SET name_en = $1 WHERE id = $2", name_en, category_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateNameTrCategoryById(category_id string, name_tr string) error {
	_, err := s.db.Exec("UPDATE categories SET name_tr = $1 WHERE id = $2", name_tr, category_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteCategoryById(category_id string) error {
	_, err := s.db.Exec("DELETE FROM categories WHERE id = $1", category_id)
	if err != nil {
		return err
	}
	return nil
}
