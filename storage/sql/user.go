package postgres

import (
	"bot/models"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (s *Storage) RegisterUser(user *models.User) error {
	_, err := s.db.Exec(`
        INSERT INTO users (id, telegram_id, username, first_name, phone_number, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)`,
		uuid.New(), user.TelegramID, user.Username, user.Name, user.Phone_Number, time.Now())
	if err != nil {
		return fmt.Errorf("failed to register user: %v", err)
	}
	log.Println("user registered ." + user.Phone_Number + user.Name)
	return nil
}

func (s *Storage) GetAllUsers() (*models.Users, error) {
	rows, err := s.db.Query("SELECT telegram_id, username, first_name, phone_number, created_at FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	var users models.Users
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.TelegramID, &user.Username, &user.Name, &user.Phone_Number, &user.Created_at); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users.Users = append(users.Users, &user)
	}
	log.Println("all users got.")
	return &users, nil
}

func (s *Storage) SetLangUser(telegramID int64, lang string) (string, error) {
	_, err := s.db.Exec(`
        INSERT INTO langs (telegram_id, lang)
        VALUES ($1, $2)`, telegramID, lang)
	if err != nil {
		return lang, fmt.Errorf("failed to set lang: %v", err)
	}
	log.Println("User lang set" + strconv.FormatInt(telegramID, 10))
	return lang, nil

}

func (s *Storage) GetLangUser(telegramID int64) (string, error) {
	lang := ""
	err := s.db.QueryRow(`
        SELECT lang
        FROM langs
        WHERE telegram_id = $1`, telegramID).Scan(&lang)
	if err != nil {
		return lang, fmt.Errorf("failed to fetch lang: %v", err)
	}
	log.Println("get user lang" + strconv.Itoa(int(telegramID)))
	return lang, nil
}

func (s *Storage) ChangeLangUser(telegramID int64, lang string) (string, error) {
	_, err := s.db.Exec("UPDATE langs SET lang = $1 WHERE telegram_id = $2", lang, telegramID)
	if err != nil {
		return lang, fmt.Errorf("failed to change lang: %v", err)
	}
	return lang, nil

}

func (s *Storage) CheckUserExist(telegramID int64) bool {
	row := s.db.QueryRow("SELECT 1 FROM users WHERE telegram_id = $1", telegramID)
	var exists bool
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		return false
	}
	log.Println("user exists checked...")
	return exists
}

func (s *Storage) UserMessageStatus(telegramID int64, status string) (string, error) {
	row := s.db.QueryRow("SELECT 1 FROM user_msg_status WHERE telegram_id = $1", telegramID)
	var exists bool
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		_, err = s.db.Exec("INSERT INTO user_msg_status (telegram_id, status) VALUES ($1, $2)", telegramID, status)
		if err != nil {
			return status, err
		}
		return status, nil
	} else {
		_, err = s.db.Exec("UPDATE user_msg_status SET status = $1 WHERE telegram_id = $2", status, telegramID)
		return status, err
	}
	// return status, nil
}

func (s *Storage) SetDataUserMessageStatus(telegramID int64, data string) (string, error) {
	_, err := s.db.Exec("UPDATE user_msg_status SET data = $1 WHERE telegram_id = $2", data, telegramID)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (s *Storage) GetUserMessageStatus(telegramID int64) (string, error) {
	var status string
	err := s.db.QueryRow("SELECT status FROM user_msg_status WHERE telegram_id = $1", telegramID).Scan(&status)
	if err != nil {
		return "", err
	}
	_, err = s.db.Exec("UPDATE user_msg_status SET status = $1  WHERE telegram_id = $2", "not", telegramID)
	if err != nil {
		return "", err
	}
	return status, nil
}

func (s *Storage) GetDataUserMessageStatus(telegramID int64) (string, error) {
	var data string
	err := s.db.QueryRow("SELECT data FROM user_msg_status WHERE telegram_id = $1", telegramID).Scan(&data)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (s *Storage) CreateLocation(location *models.Location) (*models.Location, error) {
	var user_id uuid.UUID
	err := s.db.QueryRow("SELECT id FROM users WHERE telegram_id = $1", location.UserID).Scan(&user_id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	row := s.db.QueryRow("SELECT 1 FROM locations WHERE user_id = $1", user_id)
	var exists bool
	err = row.Scan(&exists)
	if err == sql.ErrNoRows {
		_, err = s.db.Exec(`
			INSERT INTO locations (id, user_id, name_uz, name_ru, name_en, name_tr, lat, lon)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			uuid.New(), user_id, location.Name_uz, location.Name_ru, location.Name_en, location.Name_tr, location.Latitude, location.Longitude)
		if err != nil {
			return nil, fmt.Errorf("failed to create location: %v", err)
		}
		return location, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to check location: %v", err)
	} else {
		_, err = s.db.Exec(`
			UPDATE locations SET name_uz = $1, name_ru = $2, name_en = $3, name_tr = $4, lat = $5, lon = $6
			WHERE user_id = $6`, location.Name_uz, location.Name_ru, location.Name_en, location.Name_tr, location.Latitude, location.Longitude, user_id)
		if err != nil {
			return nil, fmt.Errorf("failed to update location: %v", err)
		}
		return location, nil
	}

}

func (s *Storage) DeleteLocationByUserID(telegramID int64) error {
	_, err := s.db.Exec("DELETE FROM locations WHERE user_id = $1", telegramID)
	if err != nil {
		return fmt.Errorf("failed to delete location: %v", err)
	}
	return nil
}

func (s *Storage) GetLocationByID(telegramId int64) (*models.Location, error) {
	location := &models.Location{}
	err := s.db.QueryRow("SELECT id, user_id, name_uz, name_ru, name_en, name_tr, lat, lon FROM locations WHERE user_id = $1", telegramId).Scan(&location.ID, &location.UserID, &location.Name_uz, &location.Name_ru, &location.Name_en, &location.Name_tr, &location.Latitude, &location.Longitude)
	if err != nil {
		return nil, err
	}
	return location, nil

}
