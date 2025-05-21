package custom_validator

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) (messages []string, ok bool, err error) {
	if !isStruct(data) {
		return nil, false, errors.New("data is not a struct")
	}

	validate := validator.New()
	e := validate.Struct(data)
	if e != nil {
		for _, err := range e.(validator.ValidationErrors) {
			messages = append(messages, err.Error())
		}
		return messages, false, nil
	}
	return nil, true, nil
}

func isStruct(value interface{}) bool {
	t := reflect.TypeOf(value)
	return t.Kind() == reflect.Struct
}
