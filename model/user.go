package model

type User struct {
	Username       string    `gorm:"primary_key:column:username"`
	Password       string    `gorm:"column:password"`
	Name           string    `gorm:"column:name"`
	Token          *string   `gorm:"column:token;unique"`
	TokenExpiredAt int64     `gorm:"column:token_expired_at"`
	Contacts       []Contact `gorm:"foreignKey:Username;references:Username"`
}
