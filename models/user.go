package models

// Session struct to represent a session
type Session struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `json:"username"`
	CustomerID uint   `json:"customer_id"`
	SessionID  string `json:"session_id"`
}

// User struct to represent a user (for example purpose)
type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	CustomerID uint   `json:"customer_id"`
}
