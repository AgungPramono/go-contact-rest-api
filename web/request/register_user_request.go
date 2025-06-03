package request

type RegisterUserRequest struct {
	Username string `validate:"required,max=100" json:"username"`
	Password string `validate:"required,max=100" json:"password"`
	Name     string `validate:"required,max=100" json:"name"`
}
