package request

type CreateContactRequest struct {
	FirstName string `validate:"required,max=100" json:"firstName"`
	LastName  string `validate:"omitempty,max=100" json:"lastName"`
	Email     string `validate:"omitempty,max=100" json:"email"`
	Phone     string `validate:"omitempty,max=100" json:"phone"`
}
