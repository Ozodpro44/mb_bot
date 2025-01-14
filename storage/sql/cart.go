package postgres

import (
	"bot/models"
	"fmt"
	"log"
)

func (s *Storage) AddToCart(userID int64, productID string, quantity int) error {
	var existingQuantity int
	err := s.db.QueryRow("SELECT quantity FROM cart WHERE user_id = $1 AND product_id = $2", userID, productID).Scan(&existingQuantity)

	if err == nil {
		// Product already exists in the cart, update the quantity
		newQuantity := existingQuantity + quantity

		_, err = s.db.Exec("UPDATE cart SET quantity = $1 WHERE user_id = $2 AND product_id = $3",
			newQuantity, userID, productID)
		if err != nil {
			return fmt.Errorf("failed to update cart quantity")
		}
	} else {

		// Add to cart
		_, err = s.db.Exec("INSERT INTO cart (user_id, product_id, quantity) VALUES ($1, $2, $3)",
			userID, productID, quantity)
		if err != nil {
			return fmt.Errorf("failed to insert product into cart")
		}
	}
	log.Println("product added to cart" + productID)
	return nil
}

func (s *Storage) GetCart(userID int64) (*models.Cart, error) {
	rows, err := s.db.Query(`
        SELECT p.name_uz, p.name_ru, p.name_en, c.quantity, p.price, p.id
        FROM cart c
        JOIN products p ON c.product_id = p.id
        WHERE c.user_id = $1
		ORDER BY p.id DESC`, userID)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to fetch cart: %v", err)
	}
	defer rows.Close()

	var cartItems models.Cart
	for rows.Next() {
		var item models.Cart_item
		if err := rows.Scan(&item.Name_uz, &item.Name_ru, &item.Name_en, &item.Quantity, &item.Price, &item.ProductID); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("failed to scan cart item: %v", err)
		}
		cartItems.Items = append(cartItems.Items, &item)
	}

	return &cartItems, nil
}

func (s *Storage) RemoveFromCart(userID int64, productID string) error {
	_, err := s.db.Exec("DELETE FROM cart WHERE user_id = $1 AND product_id = $2", userID, productID)
	if err != nil {
		return fmt.Errorf("failed to remove from cart: %v", err)
	}
	return nil
}

func (s *Storage) ClearCart(userID int64) error {
	_, err := s.db.Exec("DELETE FROM cart WHERE user_id = $1", userID)
	if err != nil {
		return fmt.Errorf("failed to clear cart: %v", err)
	}
	return nil
}

func (s *Storage) IncrementCartProductByID(userID int64, productID string) error {
	_, err := s.db.Exec("UPDATE cart SET quantity = quantity + 1 WHERE user_id = $1 AND product_id = $2", userID, productID)
	if err != nil {
		return fmt.Errorf("failed to increment cart product: %v", err)
	}
	return nil
}

func (s *Storage) DecrementCartProductByID(userID int64, productID string) error {
	_, err := s.db.Exec("UPDATE cart SET quantity = quantity - 1 WHERE user_id = $1 AND product_id = $2", userID, productID)
	if err != nil {
		return fmt.Errorf("failed to decrement cart product: %v", err)
	}
	return nil
}

func (s *Storage) UpdateCart(userID int64, cart *models.Cart) error {
	_, err := s.db.Exec("UPDATE cart SET quantity = $1 WHERE user_id = $2 AND product_id = $3", cart.Items[0].Quantity, userID, cart.Items[0].ProductID)
	if err != nil {
		return fmt.Errorf("failed to update cart: %v", err)
	}
	return nil
}
