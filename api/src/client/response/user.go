package response

import (
	"time"

	"golang-core/api/src/infrastructure/repository/entity"
)

type UserInfoResponse struct {
	Id        *string    `json:"id"`
	Email     *string    `json:"email"`
	Name      *string    `json:"name"`
	Password  *string    `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewUserInfoResponse(user *entity.User) *UserInfoResponse {
	if user == nil {
		return nil
	}
	userId := user.ID.String()
	return &UserInfoResponse{
		Id:        &userId,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
