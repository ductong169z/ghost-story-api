package response

// Pagination holds metadata about the pagination state
type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	TotalPages  int `json:"total_pages"`
	HasMore     bool `json:"has_more"`
}

// PaginatedResponse is a generic response structure for paginated data
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}
