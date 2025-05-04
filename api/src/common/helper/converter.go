package helper

import (
	"golang-core/api/src/client/request"
	"golang-core/api/src/client/response"
	"golang-core/api/src/infrastructure/repository/entity"
)

func FromUserUpdateRequest(request request.UserUpdateRequest, user *entity.User) {
	if email := request.Email; email != nil {
		user.Email = email
	}
	if name := request.Name; name != nil {
		user.Name = name
	}
	if password := request.Password; password != nil {
		user.Password = password
	}
}

func ToListUserInforResponse(users []entity.User) []response.UserInfoResponse {
	var result []response.UserInfoResponse
	for _, i := range users {
		result = append(result, *response.NewUserInfoResponse(&i))
	}
	return result
}
