package _internal

type PaginationQuery struct {
	// from 0
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type PaginationResponse struct {
	Total int `json:"total"`
	Page  int `json:"page"`
}

type ResponseBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListResponseBody struct {
	ResponseBody
	Pagination *PaginationResponse `json:"pagination"`
}
