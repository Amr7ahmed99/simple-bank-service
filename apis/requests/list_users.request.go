package requests

type ListUsersRequest struct {
	PageSize int32 `json:"page_size" query:"page_size" binding:"required,min=5,max=10"`
	PageID   int32 `json:"page_id" query:"page_id" binding:"required,min=1"`
}
