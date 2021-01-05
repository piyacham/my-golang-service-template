package users

import "time"

type Request struct {
	UserID    string `json:"userID" `
	MobileNo  string `json:"mobileNo" `
	FirstName string `json:"firstName" `
	LastName  string `json:"lastName" `
}
type Response struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	UserID    string    `json:"userID" `
	MobileNo  string    `json:"mobileNo" `
	FirstName string    `json:"firstName" `
	LastName  string    `json:"lastName" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
type User struct {
	UserID    string
	MobileNo  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
