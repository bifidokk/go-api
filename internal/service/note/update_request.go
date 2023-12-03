package note

type UpdateRequest struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
}
