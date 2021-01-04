package users

import "time"

type Request struct {
	UserID    string    `json:"userID" `
	MobileNo  string    `json:"mobileNo" `
	FirstName string    `json:"firstName" `
	LastName  string    `json:"lastName" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
type Response struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	UserID    string    `json:"userID" gorm:"column:user_id"`
	MobileNo  string    `json:"mobileNo" gorm:"column:mobile_no"`
	FirstName string    `json:"firstName" gorm:"column:first_name"`
	LastName  string    `json:"lastName" gorm:"column:last_name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
