package response

type GeneralResponse[T any] struct {
	Code        int    `json:"code"`
	Message     string `json:"message,omitempty"`
	Data        T      `json:"data,omitempty"`
	ErrorDetail string `json:"error_detail,omitempty"`
}

func ToSuccessResponse[T any](data T) GeneralResponse[T] {
	return GeneralResponse[T]{
		Message: "Success",
		Data:    data,
	}
}

func ToErrorResponse(code int, message string) GeneralResponse[any] {
	return GeneralResponse[any]{
		Code:    code,
		Message: message,
	}
}
