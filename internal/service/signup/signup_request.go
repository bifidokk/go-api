package signup

type SignupRequest struct {
	Email    string `form:"email" binding:"required,email,unique"`
	Password string `form:"password" binding:"required"`
}
