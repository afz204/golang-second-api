package dto

type BookUpdateDto struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID string `json:"user_id,omitempty" form:"user_id,omitempty"`
}
type BookCreateDto struct {
	Title string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID string `json:"user_id,omitempty" form:"user_id,omitempty"`
}
