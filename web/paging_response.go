package web

type PagingResponse struct {
	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
	Size        int `json:"size"`
}
