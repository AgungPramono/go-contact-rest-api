package response

type ContactResponse struct {
	Id       string `json:"id"`
	FirsName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
