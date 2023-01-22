package dto

type RequestQuery struct {
	PageID   int32  `json:"page_id" binding:"required,min=1"`
	PageSize int32  `json:"page_size" binding:"required,min=1"`
	SortBy   string `json:"sortBy"`
}
