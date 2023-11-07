package signup

type SignupRequest struct {
	Email    string `form:"email" binding:"required,email,userEmailUnique"`
	Password string `form:"password" binding:"required"`
}
