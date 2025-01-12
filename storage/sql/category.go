package postgres

import (
	"bot/models"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *Storage) CreateCategory(category *models.Category) error {
	_, err := s.db.Exec("INSERT INTO categories (id, name_uz, name_ru, name_en, abelety, created_at) VALUES ($1, $2, $3, $4, $5, $6)", uuid.New(), category.Name_uz, category.Name_ru, category.Name_en, category.Abelety, time.Now())
	if err != nil {
		return fmt.Errorf("failed to add category: %v", err)
	}
	return nil
}

func (s *Storage) GetAllCategories() (*models.Categories, error) {
	log.Println("Categoty?")
	rows, err := s.db.Query("SELECT id, name_uz, name_ru, name_en, abelety, created_at FROM categories ORDER BY name_uz ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %v", err)
	}
	defer rows.Close()

	var categories models.Categories
	for rows.Next() {
		category := models.Category{}
		if err := rows.Scan(&category.ID, &category.Name_uz, &category.Name_ru, &category.Name_en, &category.Abelety, &category.Created_at); err != nil {
			return nil, fmt.Errorf("failed to scan category: %v", err)
		}
		categories.Categories = append(categories.Categories, &category)
	}
	log.Println(err)

	return &categories, nil
}

func (s *Storage) GetCategoryByID(category_id string) (*models.Category, error) {
	category := &models.Category{}
	err := s.db.QueryRow("SELECT id, name_uz, name_ru, name_en, abelety FROM categories WHERE id = $1", category_id).Scan(&category.ID, &category.Name_uz, &category.Name_ru, &category.Name_en, &category.Abelety)
	if err != nil {
		return nil, err
	}
	return category, nil

}

func (s *Storage) UpdateCategoryById(category *models.Category) (*models.Category, error) {
	err := s.db.QueryRow("UPDATE categories SET name_uz = $1, name_ru = $2, name_en = $3 WHERE id = $4, RETURNING id, name_uz, name_ru, name_en, abelety", category.Name_uz, category.Name_ru, category.Name_en, category.ID).Scan(&category.ID, &category.Name_uz, &category.Name_ru, &category.Name_en, &category.Abelety)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *Storage) UpdateAbeletyCategoryById(category_id string, abelety bool) error {
	_, err := s.db.Exec("UPDATE categories SET abelety = &1 WHERE id = $2", abelety, category_id)
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
