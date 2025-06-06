package model

type Contact struct {
	ID        string    `gorm:"primaryKey;column:id"`                    // Primary Key
	Username  string    `gorm:"column:username"`                         // Foreign Key from users
	FirstName string    `gorm:"column:first_name"`                       // First Name
	LastName  string    `gorm:"column:last_name"`                        // Last Name
	Phone     string    `gorm:"column:phone"`                            // Phone number
	Email     string    `gorm:"column:email"`                            // Email address
	User      User      `gorm:"foreignKey:Username;references:Username"` // Belongs to User
	Addresses []Address `gorm:"foreignKey:ContactID;references:ID"`
}

func (Contact) TableName() string {
	return "contact"
}
