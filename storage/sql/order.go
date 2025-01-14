package postgres

import (
	"bot/lib/helpers"
	"bot/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (s *Storage) CreateOrder(userID int64, payment_type string) (string, error) {
	var totalPrice float64
	var orderID uuid.UUID
	var daily_order_number int
	var order_number int
	var user_id uuid.UUID
	var lat float32
	var lon float32
	var address string
	var phone string

	err := s.db.QueryRow(`
		SELECT id, phone_number
		FROM users
		WHERE telegram_id = $1`, userID).Scan(&user_id, &phone)

	if err != nil {
		return "", err
	}

	log.Println("telegram_id succsefully fetched in CreateOrder")

	err = s.db.QueryRow(`
		SELECT order_number, daily_order_number
		FROM order_numbers
	`).Scan(&order_number, &daily_order_number)

	if err != nil {
		return "", err
	}

	log.Println("order_number succsefully fetched in CreateOrder")

	err = s.db.QueryRow(`
		SELECT lat,lon,name_uz
		FROM locations
		WHERE user_id = $1`, user_id).Scan(&lat, &lon, &address)

	if err != nil {
		return "", err
	}

	log.Println("location succsefully fetched in CreateOrder")

	// Calculate total price
	rows, err := s.db.Query(`
        SELECT c.product_id, c.quantity, p.price
        FROM cart c
        JOIN products p ON c.product_id = p.id
        WHERE c.user_id = $1`, userID)
	if err != nil {
		return "", fmt.Errorf("failed to fetch cart items")
	}

	log.Println("cart items succsefully fetched in CreateOrder")

	defer rows.Close()

	tx, err := s.db.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction")
	}

	log.Println("transaction succsefully started in CreateOrder")

	for rows.Next() {
		var productID uuid.UUID
		var quantity int
		var price float64
		if err := rows.Scan(&productID, &quantity, &price); err != nil {
			tx.Rollback()
			return "", err
		}

		log.Println("cart items succsefully scaned in CreateOrder")

		totalPrice += price * float64(quantity)
	}
	order_id := uuid.New()

	delivery := "0"
	if !helpers.Haversine(41.275030, 69.264482, float64(lat), float64(lon)) {
		delivery = "Yandex Dostavka"
	}
	var status string
	if payment_type == "card" {
		status = "pending"
	} else {
		status = "preparing"
	}

	// Create order
	err = tx.QueryRow("INSERT INTO orders (id,user_id, total_price, order_number, daily_order_number, adress, lat, lon, delivery_price, phone_number,status,payment_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id",
		order_id, user_id, totalPrice, order_number+1, daily_order_number+1, address, lat, lon, delivery, phone, status, payment_type).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	log.Println("order succsefully created in CreateOrder")

	// Move cart items to order_items
	_, err = tx.Exec(`
        INSERT INTO order_items (id, order_id, product_id, quantity)
        SELECT gen_random_uuid(), $1, c.product_id, c.quantity
        FROM cart c
        JOIN products p ON c.product_id = p.id
        WHERE c.user_id = $2`, orderID, userID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	log.Println("cart items succsefully moved to order_items in CreateOrder")

	// Clear cart
	_, err = tx.Exec("DELETE FROM cart WHERE user_id = $1", userID)
	if err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to clear cart")
	}

	log.Println("cart succsefully deleted in CreateOrder")

	order_number++
	daily_order_number++

	// Update the order numbers
	_, err = tx.Exec(`
    UPDATE order_numbers 
    SET order_number = $1, daily_order_number = $2`,
		order_number, daily_order_number)
	if err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to update order numbers: %w", err)
	}

	log.Println("order numbers succsefully updated in CreateOrder")

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Println("transaction succsefully commited in CreateOrder")

	return order_id.String(), nil
}

func (s *Storage) GetOrderDetails(orderID string) (*[]models.OrderDetails, error) {
	var orders []models.OrderDetails
	order := models.OrderDetails{}

	// Fetch order details
	err := s.db.QueryRow(`
        SELECT id, user_id, total_price, status, created_at
        FROM orders
        WHERE id = $1`, orderID).Scan(&order.OrderID, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt)
	if err != nil {
		return &orders, fmt.Errorf("failed to fetch order: %v", err)
	}

	// Fetch order items
	rows, err := s.db.Query(`
        SELECT p.name_uz, p.name_ru, p.name_en, p.name_tr, oi.quantity, oi.price
        FROM order_items oi
        JOIN products p ON oi.product_id = p.id 
        WHERE oi.order_id = $1`, orderID)
	if err != nil {
		return &orders, fmt.Errorf("failed to fetch order items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Name_uz, &item.Name_ru, &item.Name_en, &item.Name_tr, &item.Quantity, &item.Price); err != nil {
			return &orders, fmt.Errorf("failed to scan order item: %v", err)
		}
		orders = append(orders, order)
		//orders[0].Items = append(orders[0].Items, &item)

	}

	return &orders, nil
}

func (s *Storage) GetOrderByUserID(userID int64) (*[]models.OrderDetails, error) {
	var orders []models.OrderDetails
	var order models.OrderDetails
	user_id := uuid.UUID{}
	var adress models.Location

	err := s.db.QueryRow(`
		SELECT id
		FROM users
		WHERE telegram_id = $1`, userID).Scan(&user_id)

	if err != nil {
		return &orders, fmt.Errorf("failed to fetch user: %v", err)
	}

	// Fetch order details
	rows, err := s.db.Query(`
        SELECT id, order_number,daily_order_number, user_id, total_price, status, created_at, delivery_price, phone_number,lat,lon,adress, payment_type
        FROM orders
        WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT 1`, user_id)
	if err != nil {
		return &orders, fmt.Errorf("failed to fetch order: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		order = models.OrderDetails{}
		err = rows.Scan(&order.OrderID, &order.Order_number, &order.Daily_order_number, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.Delivery_type, &order.PhoneNumber, &adress.Latitude, &adress.Longitude, &adress.Name_uz, &order.Payment_type)
		if err != nil {
			return &orders, fmt.Errorf("failed to scan order: %v", err)
		}
		// Fetch order items
		rows_items, err := s.db.Query(
			`SELECT p.name_uz, p.name_ru, p.name_en, p.name_tr, oi.quantity
			FROM order_items oi
			JOIN products p ON oi.product_id = p.id
			WHERE oi.order_id = $1`, order.OrderID)
		if err != nil {
			return &orders, fmt.Errorf("failed to fetch order items: %v", err)
		}

		for rows_items.Next() {
			items := models.Item{}
			err = rows_items.Scan(&items.Name_uz, &items.Name_ru, &items.Name_en, &items.Name_tr, &items.Quantity)
			if err != nil {
				return &orders, fmt.Errorf("failed to scan order items: %v", err)
			}
			order.Items = append(order.Items, &items)
		}
		order.Address = &adress
		orders = append(orders, order)
		defer rows_items.Close()
	}

	return &orders, nil

}

func (s *Storage) GetOrderDetailsByOrderID(orderID string) (*models.OrderDetails, error) {
	order := &models.OrderDetails{}
	adress := &models.Location{}

	err := s.db.QueryRow(`
        SELECT o.id, o.order_number,o.daily_order_number, o.user_id, o.total_price, o.status, o.created_at, o.adress, o.lat, o.lon, o.delivery_price, o.payment_type, o.phone_number
        FROM orders o
        WHERE o.id = $1`, orderID).Scan(&order.OrderID, &order.Order_number, &order.Daily_order_number, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt, &adress.Name_uz, &adress.Latitude, &adress.Longitude, &order.Delivery_type, &order.Payment_type, &order.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order: %v", err)
	}

	rows, err := s.db.Query(`
        SELECT p.name_uz, p.name_ru, p.name_en, p.name_tr, oi.quantity, p.price
        FROM order_items oi
        JOIN products p ON oi.product_id = p.id
        WHERE oi.order_id = $1`, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Item{}
		if err := rows.Scan(&item.Name_uz, &item.Name_ru, &item.Name_en, &item.Name_tr, &item.Quantity, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to scan order item: %v", err)
		}
		order.Items = append(order.Items, &item)
	}
	order.Address = adress

	return order, nil

}

func (s *Storage) ChangeOrderStatus(orderID string, status string) (string, error) {
	_, err := s.db.Exec("UPDATE orders SET status = $1 WHERE id = $2", status, orderID)
	if err != nil {
		return "", err
	}
	return status, nil
}
