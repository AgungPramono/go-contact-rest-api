package model

type Contact struct {
	Id        string `gorm:"primary_key:column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Phone     string `gorm:"column:phone"`
	Email     string `gorm:"column:email"`
	Username  string `gorm:"column:username"` // foreign key
	User      *User  `gorm:"foreignKey:Username;references:Username"`
	//Addresses []Address `gorm:"foreignKey:ContactID;references:ID"`
}

func (c *Contact) TableName() string {
	return "contact"
}
