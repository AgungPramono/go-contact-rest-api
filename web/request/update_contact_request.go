package request

type UpdateContactRequest struct {
	Id        string `json:"_;validate:required"`
	FirstName string `json:"first_name;validate:required;max=100"`
	LastName  string `json:"last_name;validate:max=100"`
	Phone     string `json:"phone;validate:max=100"`
	Email     string `json:"email;validate:max=100"`
}
