package web

type TokenResponse struct {
	Token string `json:"token"`

	ExpiredAt *int64 `json:"expiredAt,omitempty"`
}
