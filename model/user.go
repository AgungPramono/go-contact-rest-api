package model

type User struct {
	Username       string  `gorm:"primaryKey;column:username"` // Primary Key
	Password       string  `gorm:"column:password"`            // Password
	Name           string  `gorm:"column:name"`                // Name
	Token          *string `gorm:"column:token;unique"`        // Token (unique constraint)
	TokenExpiredAt int64   `gorm:"column:token_expired_at"`    // Token expiration timestamp
	//Contacts       []Contact `gorm:"foreignKey:Username;references:Username"` // One-to-Many with Contact
}
