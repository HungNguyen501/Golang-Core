package request

type UserInsertRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	ID       string  `json:"id" validate:"required,uuid"`
	Email    *string `json:"email" validate:"email"`
	Name     *string `json:"name"`
	Password *string `json:"password"`
}

type UserListAllRequest struct {
	Limit  *int `query:"limit" validate:"required,min=0,max=100"`
	Offset *int `query:"offset" validate:"required,min=0"`
}
