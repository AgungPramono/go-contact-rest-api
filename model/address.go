package model

type Address struct {
	ID         string  `gorm:"primaryKey;column:id"`
	Street     string  `json:"street"`
	City       string  `json:"city"`
	Province   string  `json:"province"`
	Country    string  `json:"country"`
	PostalCode string  `gorm:"column:postal_code"`
	ContactID  string  `gorm:"column:contact_id" ` // foreign key
	Contact    Contact `gorm:"foreignKey:ContactID;references:ID"`
}
