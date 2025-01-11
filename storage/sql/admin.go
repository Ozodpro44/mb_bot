package postgres

import (
	"bot/models"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func (s *Storage) CheckAdmin(telegramID int64) bool {
	son := 0 
	err := s.db.QueryRow("SELECT 1 FROM admins WHERE telegram_id = $1", telegramID).Scan(&son)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		return false
	}
	return true

}

func (s *Storage) CreateAdmin(admin *models.Admin) (*models.Admin, error) {
	var id uuid.UUID = uuid.New()
	var created_at time.Time = time.Now()
	err := s.db.QueryRow("INSERT INTO admins (id, telegram_id, name, password) VALUES ($1, $2, $3, $4) RETURNING id, telegram_id, phone_number, password", id, admin.Admin_id, admin.Phone_Number, admin.Password, created_at).Scan(&admin.Id, &admin.Admin_id, &admin.Password)
	if err != nil {
		return nil, err
	}
	return admin, nil

}

func (s *Storage) UpdateAdmin(admin *models.Admin) (*models.Admin, error) {
	err := s.db.QueryRow("UPDATE admins SET phone_number = $1, password = $2 WHERE id = $3", admin.Phone_Number, admin.Password, admin.Id).Scan(&admin.Phone_Number, &admin.Password)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *Storage) DeleteAdmin(admin_id int64) error {
	err := s.db.QueryRow("DELETE FROM admins WHERE id = $1", admin_id).Scan()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetAdminLang(admin_id int64) (string, error) {
	var lang string
	err := s.db.QueryRow("SELECT lang FROM admins WHERE telegram_id = $1", admin_id).Scan(&lang)
	if err != nil {
		return "", err
	}
	return lang, nil
}

func (s *Storage) SetAdminLang(admin_id int64, lang string) (string, error) {
	err := s.db.QueryRow("UPDATE admins SET lang = $1 WHERE telegram_id = $2", lang, admin_id).Scan()
	if err != nil {
		return "", err
	}
	return lang, nil
}

func (s *Storage) CloseDay() error {
	_, err := s.db.Exec("UPDATE order_numbers SET daily_order_number = 0")
	if err != nil {
		return err
	}

	_, err = s.db.Exec("UPDATE branch SET opened = false WHERE id = 'a7c96256-961a-4694-8991-622851e75a96'")
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) CheckOpened() (bool, error) {
	var opened bool
	err := s.db.QueryRow("SELECT opened FROM branch WHERE id = 'a7c96256-961a-4694-8991-622851e75a96'").Scan(&opened)
	if err != nil {
		return false, err
	}
	return opened, nil
}

func (s *Storage) OpenDay() error {
	_, err := s.db.Exec("UPDATE branch SET opened = true WHERE id = 'a7c96256-961a-4694-8991-622851e75a96'")
	if err != nil {
		return err
	}
	return nil

}
