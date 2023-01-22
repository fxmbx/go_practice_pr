package userdto

import "github.com/fxmbx/go_practice_pr/models"

type UserResponse struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Token     string `json:"token,omitempty"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FitstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
	}
}
