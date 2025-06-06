package model

type Address struct {
	ID         string  `gorm:"primaryKey;column:id"`               // Primary Key
	ContactID  string  `gorm:"column:contact_id"`                  // Foreign Key from contact
	Street     string  `gorm:"column:street"`                      // Street
	City       string  `gorm:"column:city"`                        // City
	Province   string  `gorm:"column:province"`                    // Province
	Country    string  `gorm:"column:country"`                     // Country (NOT NULL in schema)
	PostalCode string  `gorm:"column:postal_code"`                 // Postal Code
	Contact    Contact `gorm:"foreignKey:ContactID;references:ID"` // Belongs to Contact
}
