package note

type CreateRequest struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
}
