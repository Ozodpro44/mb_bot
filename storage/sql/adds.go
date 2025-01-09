package postgres

import (
	"bot/models"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *Storage) CreateAdds(adds *models.Add) (*models.Add, error) {
	var id uuid.UUID = uuid.New()
	var created_at time.Time = time.Now()
	err := s.db.QueryRow(
		`INSERT INTO adds (id, photo, text, created_at) 
		 VALUES ($1, $2, $3, $4) 
		 RETURNING id, photo, text`,
		id, adds.Photo, adds.Text, created_at,
	).Scan(&id, &adds.Photo, &adds.Text)
	
	if err != nil {
		// Handle error
		log.Println("Error inserting data:", err)
		return nil, err
	}
	log.Println("Data inserted successfully")
	return adds, err
}

func (s *Storage) GetAllAdds() (*models.Adds, error) {
	rows, err := s.db.Query("SELECT id, photo, text FROM adds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	adds := &models.Adds{}
	for rows.Next() {
		add := &models.Add{}
		if err := rows.Scan(&add.ID, &add.Photo, &add.Text); err != nil {
			return nil, err
		}
		adds.Adds = append(adds.Adds, add)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println("Adds sucsefully fetched")
	return adds, nil

}
func (s *Storage) GetAddsById(id string) (*models.Add, error) {
	add := &models.Add{}
	err := s.db.QueryRow("SELECT id, photo, text FROM adds WHERE id = $1", id).Scan(&add.ID, &add.Photo, &add.Text)
	if err != nil {
		return nil, err
	}
	log.Println("Add sucsefully fetched")
	return add, nil
}
