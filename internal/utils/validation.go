package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validation(request interface{}) map[string]string {
	errors := make(map[string]string)
	if err := validate.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(request).Elem().FieldByName(err.StructField())
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = strings.ToLower(err.StructField())
			}
			errors[jsonTag] = err.Tag()
		}
	}

	return errors
}
