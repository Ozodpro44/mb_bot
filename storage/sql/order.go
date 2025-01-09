package postgres

import (
	"bot/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (s *Storage) CreateOrder(userID int64) (string, error) {
	var totalPrice float64
	var orderID uuid.UUID
	var daily_order_number int
	var order_number int
	var user_id uuid.UUID

	err := s.db.QueryRow(`
		SELECT id
		FROM users
		WHERE telegram_id = $1`, userID).Scan(&user_id)

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

	// Create order
	err = tx.QueryRow("INSERT INTO orders (id,user_id, total_price, order_number, daily_order_number) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		order_id, user_id, totalPrice, order_number+1, daily_order_number+1).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	log.Println("order succsefully created in CreateOrder")

	// Move cart items to order_items
	_, err = tx.Exec(`
        INSERT INTO order_items (id, order_id, product_id, quantity)
        SELECT $1, $2, c.product_id, c.quantity
        FROM cart c
        JOIN products p ON c.product_id = p.id
        WHERE c.user_id = $3`, uuid.New(), orderID, userID)
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
        SELECT p.name, oi.quantity, oi.price
        FROM order_items oi
        JOIN products p ON oi.product_id = p.id 
        WHERE oi.order_id = $1`, orderID)
	if err != nil {
		return &orders, fmt.Errorf("failed to fetch order items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Name_uz, &item.Name_ru, &item.Name_en, &item.Quantity, &item.Price); err != nil {
			return &orders, fmt.Errorf("failed to scan order item: %v", err)
		}
		orders = append(orders, order)
		orders[0].Items = append(orders[0].Items, &item)

	}

	return &orders, nil
}

func (s *Storage) GetOrderByUserID(userID int64) (*[]models.OrderDetails, error) {
	var orders []models.OrderDetails
	var order models.OrderDetails
	user_id := uuid.UUID{}

	err := s.db.QueryRow(`
		SELECT id
		FROM users
		WHERE telegram_id = $1`, userID).Scan(&user_id)

	if err != nil {
		return &orders, fmt.Errorf("failed to fetch user: %v", err)
	}

	// Fetch order details
	rows, err := s.db.Query(`
        SELECT id, order_number, user_id, total_price, status, created_at
        FROM orders
        WHERE user_id = $1
		ORDER BY id DESC`, user_id)
	if err != nil {
		return &orders, fmt.Errorf("failed to fetch order: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		order = models.OrderDetails{}
		err = rows.Scan(&order.OrderID, &order.Order_number, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt)
		if err != nil {
			return &orders, fmt.Errorf("failed to scan order: %v", err)
		}
		// Fetch order items
		rows_items, err := s.db.Query(
			`SELECT p.name_uz, p.name_ru, p.name_en, oi.quantity
			FROM order_items oi
			JOIN products p ON oi.product_id = p.id
			WHERE oi.order_id = $1`, order.OrderID)
		if err != nil {
			return &orders, fmt.Errorf("failed to fetch order items: %v", err)
		}
		defer rows_items.Close()
		for rows_items.Next() {
			items := models.Item{}
			err = rows_items.Scan(&items.Name_uz, &items.Name_ru, &items.Name_en, &items.Quantity)
			if err != nil {
				return &orders, fmt.Errorf("failed to scan order items: %v", err)
			}
			order.Items = append(order.Items, &items)
		}
		orders = append(orders, order)
	}

	return &orders, nil

}

func (s *Storage) GetOrderDetailsByOrderID(orderID string) (*models.OrderDetails, error) {
	order := &models.OrderDetails{}
	err := s.db.QueryRow(`
        SELECT o.id, o.order_number, o.user_id, o.total_price, o.status, o.created_at
        FROM orders o
        WHERE o.id = $1`, orderID).Scan(&order.OrderID, &order.Order_number, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order: %v", err)
	}

	rows, err := s.db.Query(`
        SELECT p.name_uz, p.name_ru, p.name_en, oi.quantity, p.price
        FROM order_items oi
        JOIN products p ON oi.product_id = p.id
        WHERE oi.order_id = $1`, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Item{}
		if err := rows.Scan(&item.Name_uz, &item.Name_ru, &item.Name_en, &item.Quantity, &item.Price); err != nil {
			return nil, fmt.Errorf("failed to scan order item: %v", err)
		}
		order.Items = append(order.Items, &item)
	}

	return order, nil

}
	
