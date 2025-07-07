package postgres

import (
	"bot/models"
	"database/sql"
	"fmt"

	// "time"

	"github.com/google/uuid"
)

func (s *Storage) GetDashboard() (*models.Dashboard, error) {
	var dashboard models.Dashboard

	// TotalOrders (This Month)
	err := s.db.QueryRow(`
		SELECT COUNT(id) FROM orders 
		WHERE created_at >= date_trunc('month', current_date);
	`).Scan(&dashboard.TotalOrders)
	if err != nil {
		return nil, fmt.Errorf("failed to get total orders: %v", err)
	}

	// TotalRevenue (This Month)
	err = s.db.QueryRow(`
		SELECT COALESCE(SUM(total_price), 0) FROM orders 
		WHERE created_at >= date_trunc('month', current_date);
	`).Scan(&dashboard.TotalRevenue)
	if err != nil {
		return nil, fmt.Errorf("failed to get total revenue: %v", err)
	}

	// TotalUsers
	err = s.db.QueryRow(`
		SELECT COUNT(id) FROM users;
	`).Scan(&dashboard.TotalUsers)
	if err != nil {
		return nil, fmt.Errorf("failed to get total users: %v", err)
	}

	// AvgOrderValue (This Month)
	err = s.db.QueryRow(`
		SELECT COALESCE(AVG(total_price), 0) FROM orders 
		WHERE created_at >= date_trunc('month', current_date);
	`).Scan(&dashboard.AvgOrderValue)
	if err != nil {
		return nil, fmt.Errorf("failed to get average order value: %v", err)
	}

	// OrdersToday
	err = s.db.QueryRow(`
		SELECT COUNT(id) FROM orders 
		WHERE created_at >= current_date;
	`).Scan(&dashboard.OrdersToday)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders today: %v", err)
	}

	// RevenueToday
	err = s.db.QueryRow(`
		SELECT COALESCE(SUM(total_price), 0) FROM orders 
		WHERE created_at >= current_date;
	`).Scan(&dashboard.RevenueToday)
	if err != nil {
		return nil,
			fmt.Errorf("failed to get revenue today: %v", err)
	}

	// SatisfactionRate (dummy for now)
	dashboard.SatisfactionRate = 4.5 // Placeholder

	// ActiveBranches
	err = s.db.QueryRow(`
		SELECT COUNT(id) FROM branch WHERE opened = true;
	`).Scan(&dashboard.ActiveBranches)
	if err != nil {
		return nil, fmt.Errorf("failed to get active branches: %v", err)
	}

	// TotalBranches
	err = s.db.QueryRow(`
		SELECT COUNT(id) FROM branch;
	`).Scan(&dashboard.TotalBranches)
	if err != nil {
		return nil, fmt.Errorf("failed to get total branches: %v", err)
	}

	// // Trends (dummy for now)
	// dashboard.Trends = models.Trends{
	// 	Orders:  10.5, // Placeholder
	// 	Revenue: 12.3, // Placeholder
	// 	Users:   8.2,  // Placeholder
	// 	Aov:     5.1,  // Placeholder
	// }	// Trends
	err = s.db.QueryRow(`
		SELECT 
			COALESCE(((SELECT COUNT(id) FROM orders WHERE created_at >= date_trunc('month', current_date)) - 
			(SELECT COUNT(id) FROM orders WHERE created_at >= date_trunc('month', current_date) - interval '1 month' AND created_at < date_trunc('month', current_date))) * 100.0 / 
			NULLIF((SELECT COUNT(id) FROM orders WHERE created_at >= date_trunc('month', current_date) - interval '1 month' AND created_at < date_trunc('month', current_date)), 0), 0) AS orders_change,

			COALESCE(((SELECT SUM(total_price) FROM orders WHERE created_at >= date_trunc('month', current_date)) - 
			(SELECT SUM(total_price) FROM orders WHERE created_at >= date_trunc('month', current_date) - interval '1 month' AND created_at < date_trunc('month', current_date))) * 100.0 / 
			NULLIF((SELECT SUM(total_price) FROM orders WHERE created_at >= date_trunc('month', current_date) - interval '1 month' AND created_at < date_trunc('month', current_date)), 0), 0) AS revenue_change,

			COALESCE(((SELECT COUNT(id) FROM users WHERE created_at >= date_trunc('month', current_date)) - 
			(SELECT COUNT(id) FROM users WHERE created_at >= date_trunc('month', current_date) - interval '1 month' AND created_at < date_trunc('month', current_date))) * 100.0 / 
			NULLIF((SELECT COUNT(id) FROM users WHERE created_at >= date_trunc('month', current_date) - interval '1 month' AND created_at < date_trunc('month', current_date)), 0), 0) AS users_change,
			
			COALESCE(((SELECT AVG(total_price) FROM orders WHERE created_at >= date_trunc('month', current_date)) - 
			(SELECT AVG(total_price) FROM orders WHERE created_at >= date_trunc('month', current_date) - interval '1 month' AND created_at < date_trunc('month', current_date))) * 100.0 / 
			NULLIF((SELECT AVG(total_price) FROM orders WHERE created_at >= date_trunc('month', current_date) - interval '1 month' AND created_at < date_trunc('month', current_date)), 0), 0) AS aov_change;
	`).Scan(&dashboard.Trends.Orders, &dashboard.Trends.Revenue, &dashboard.Trends.Users, &dashboard.Trends.Aov)
	if err != nil {
		return nil, fmt.Errorf("failed to get trends: %v", err)
	}

	// TopProducts
	rows, err := s.db.Query(`
		SELECT 
			p.name_en, 
			SUM(ci.quantity) AS sold_count, 
			SUM(ci.quantity * p.price) AS total_revenue
		FROM orders od
		JOIN order_items ci ON od.id = ci.order_id
		JOIN products p ON ci.product_id = p.id
		WHERE od.created_at >= date_trunc('month', current_date)
		GROUP BY p.name_en
		ORDER BY sold_count DESC
		LIMIT 5;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get top products: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var tp models.TopProducts
		tp.Change = 0.0 // Placeholder
		if err := rows.Scan(&tp.Name, &tp.Sold, &tp.Revenue); err != nil {
			return nil, fmt.Errorf("failed to scan top product: %v", err)
		}
		dashboard.TopProducts = append(dashboard.TopProducts, tp)
	}

	//
	// RecentOrders
	rows, err = s.db.Query(`
		SELECT 
			od.id, 
			u.username, 
			od.total_price, 
			od.status, 
			od.created_at
		FROM orders od
		JOIN users u ON od.user_id = u.id
		ORDER BY od.created_at DESC
		LIMIT 5;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get recent orders: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var ro models.RecentOrders
		if err := rows.Scan(&ro.OrderID, &ro.Customer, &ro.Total, &ro.Status, &ro.Time); err != nil {
			return nil, fmt.Errorf("failed to scan recent order: %v", err)
		}
		dashboard.RecentOrders = append(dashboard.RecentOrders, ro)
	}

	// CategoryPerformance
	rows, err = s.db.Query(`
		SELECT 
			c.name_en, 
			SUM(ci.quantity * p.price) AS category_revenue,
			(SUM(ci.quantity * p.price) * 100.0 / (SELECT COALESCE(SUM(total_price), 1) FROM orders WHERE created_at >= date_trunc('month', current_date))) AS percentage
		FROM orders od
		JOIN order_items ci ON od.id = ci.order_id
		JOIN products p ON ci.product_id = p.id
		JOIN categories c ON p.categories_id = c.id
		WHERE od.created_at >= date_trunc('month', current_date)
		GROUP BY c.name_en
		ORDER BY category_revenue DESC
		LIMIT 5;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get category performance: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var cp models.CategoryPerformance
		cp.Change = 0.0 // Placeholder
		if err := rows.Scan(&cp.Name, &cp.Revenue, &cp.Percentage); err != nil {
			return nil, fmt.Errorf("failed to scan category performance: %v", err)
		}
		dashboard.CategoryPerformance = append(dashboard.CategoryPerformance, cp)
	}

	// PeakHours
	rows, err = s.db.Query(`
		SELECT 
			TO_CHAR(created_at, 'HH24') AS hour, 
			COUNT(id) AS orders_count,
			(COUNT(id) * 100.0 / (SELECT COUNT(id) FROM orders WHERE created_at >= date_trunc('month', current_date))) AS percentage
		FROM orders
		WHERE created_at >= date_trunc('month', current_date)
		GROUP BY hour
		ORDER BY orders_count DESC
		LIMIT 5;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get peak hours: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var ph models.PeakHours
		if err := rows.Scan(&ph.Hour, &ph.Orders, &ph.Percentage); err != nil {
			return nil, fmt.Errorf("failed to scan peak hour: %v", err)
		}
		dashboard.PeakHours = append(dashboard.PeakHours, ph)
	}

	// BranchPerformance
	rows, err = s.db.Query(`
		SELECT 
			b.name, 
			SUM(od.total_price) AS branch_revenue,
			(SUM(od.total_price) * 100.0 / (SELECT COALESCE(SUM(total_price), 1) FROM orders WHERE created_at >= date_trunc('month', current_date))) AS percentage
		FROM orders od
		JOIN branch b ON od.branch_id = b.id
		WHERE od.created_at >= date_trunc('month', current_date)
		GROUP BY b.name
		ORDER BY branch_revenue DESC
		LIMIT 5;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get branch performance: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var bp models.BranchPerformance
		bp.Status = "Good" // Placeholder
		if err := rows.Scan(&bp.Name, &bp.Revenue, &bp.Percentage); err != nil {
			return nil, fmt.Errorf("failed to scan branch performance: %v", err)
		}
		dashboard.BranchPerformance = append(dashboard.BranchPerformance, bp)
	}

	// SalesOverview (dummy for now)
	// dashboard.SalesOverview = models.SalesOverview{
	// 	SevenDays:  []int{10, 20, 15, 25, 30, 20, 10},                                                                                                       // Placeholder
	// 	ThirtyDays: []int{100, 120, 110, 130, 150, 140, 130, 120, 110, 100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 5, 15, 25, 35, 45, 55, 65, 75, 85, 95, 105}, // Placeholder
	// 	NinetyDays: make([]int, 90),                                                                                                                         // Placeholder
	// }
	// for i := 0; i < 90; i++ {
	// 	dashboard.SalesOverview.NinetyDays[i] = 0 // Initialize with 0
	// }

	dashboard.SalesOverview.SevenDays = make([]int, 7)
	dashboard.SalesOverview.ThirtyDays = make([]int, 30)
	dashboard.SalesOverview.NinetyDays = make([]int, 90)

	rows, err = s.db.Query(`
		SELECT 
			EXTRACT(DAY FROM created_at) AS day, 
			COUNT(id) AS orders_count
		FROM orders
		WHERE created_at >= current_date - interval '7 days'
		GROUP BY EXTRACT(DAY FROM created_at)
		ORDER BY day;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get 7-day sales overview: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var day int
		var count int
		if err := rows.Scan(&day, &count); err != nil {
			return nil, fmt.Errorf("failed to scan 7-day sales overview: %v", err)
		}
		// Assuming the last 7 days are represented, map day to index 0-6
		// This needs more robust date handling for real-world scenarios
		dashboard.SalesOverview.SevenDays[day%7] = count
	}

	rows, err = s.db.Query(`
		SELECT 
			EXTRACT(DAY FROM created_at) AS day, 
			COUNT(id) AS orders_count
		FROM orders
		WHERE created_at >= current_date - interval '30 days'
		GROUP BY EXTRACT(DAY FROM created_at)
		ORDER BY day;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get 30-day sales overview: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var day int
		var count int
		if err := rows.Scan(&day, &count); err != nil {
			return nil, fmt.Errorf("failed to scan 30-day sales overview: %v", err)
		}
		dashboard.SalesOverview.ThirtyDays[day%30] = count
	}

	rows, err = s.db.Query(`
		SELECT 
			EXTRACT(DAY FROM created_at) AS day, 
			COUNT(id) AS orders_count
		FROM orders
		WHERE created_at >= current_date - interval '90 days'
		GROUP BY EXTRACT(DAY FROM created_at)
		ORDER BY day;
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get 90-day sales overview: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var day int
		var count int
		if err := rows.Scan(&day, &count); err != nil {
			return nil, fmt.Errorf("failed to scan 90-day sales overview: %v", err)
		}
		dashboard.SalesOverview.NinetyDays[day%90] = count
	}
	fmt.Println("dashboard:")
	fmt.Println(dashboard)
	return &dashboard, nil
}

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
	// var created_at time.Time = time.Now()
	err := s.db.QueryRow(`INSERT INTO admins (id, telegram_id, phone_number, password) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, telegram_id, phone_number, password`, id, admin.Admin_id, admin.Phone_Number, admin.Password).Scan(
		&admin.Id,
		&admin.Admin_id,
		&admin.Phone_Number,
		&admin.Password)
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

func (s *Storage) ChangeAdminLang(admin_id int64, lang string) (string, error) {
	_, err := s.db.Exec("UPDATE admins SET lang = $1 WHERE telegram_id = $2", lang, admin_id)
	if err != nil {
		return "", err
	}
	return lang, nil

}
