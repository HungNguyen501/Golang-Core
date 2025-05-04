package echoutils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type IValidator interface {
	Register() (validator.Func, string)
}

type Validators struct {
	val        *validator.Validate
	validators []IValidator
}

func BindAndValidate(ec echo.Context, payload any) error {
	if err := ec.Bind(payload); err != nil {
		return err
	}
	if err := ec.Validate(payload); err != nil {
		return err
	}
	return nil
}

func NewValidators() *Validators {
	return &Validators{
		val:        validator.New(),
		validators: []IValidator{},
	}
}

func (vl *Validators) Setup() error {
	for _, validator := range vl.validators {
		fnc, tag := validator.Register()
		if err := vl.val.RegisterValidation(tag, fnc); err != nil {
			return err
		}
	}
	return nil
}

func (vl *Validators) Validate(requestData any) error {
	if err := vl.val.Struct(requestData); err != nil {
		return err
	}
	return nil
}
