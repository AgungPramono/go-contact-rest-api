package web

type ApiResponse struct {
	Data    interface{}     `json:"data,omitempty"`
	Status  bool            `json:"status,omitempty"`
	Message string          `json:"message,omitempty"`
	Errors  string          `json:"errors,omitempty"`
	Paging  *PagingResponse `json:"paging,omitempty"`
}
