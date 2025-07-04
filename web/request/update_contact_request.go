package request

type UpdateContactRequest struct {
	Id        string `json:"id" validate:"required"`
	FirstName string `json:"firstName" validate:"required,max=100"`
	LastName  string `json:"lastName" validate:"omitempty,max=100"`
	Phone     string `json:"phone" validate:"omitempty,max=100"`
	Email     string `json:"email" validate:"omitempty,max=100"`
}
