package signup

type SignupRequest struct {
	Email    string `form:"email" binding:"required,userEmailUnique"`
	Password string `form:"password" binding:"required"`
}