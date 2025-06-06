package request

type SearchContactRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
}
