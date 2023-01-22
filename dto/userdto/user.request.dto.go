package userdto

type UpdateUserRequest struct {
	ID        int32  `json:"id" form:"id" binding:"required"`
	FirstName string `json:"first_name" form:"first_name" binding:"required,min=1"`
	LastName  string `json:"last_name" form:"last_name" binding:"min=1"`
	UserName  string `json:"user_name" form:"user_name" binding:"min=1"`
	Email     string `json:"email" form:"email" binding:"email"`
}
