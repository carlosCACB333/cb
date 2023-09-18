package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateFields(data interface{}) map[string]string {

	validate := validator.New()
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	messages := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		st := reflect.TypeOf(data)
		var field string
		json_name, e := st.FieldByName(err.Field())
		if e {
			field = json_name.Tag.Get("json")
		} else {
			field = strings.ToLower(err.Field())
		}
		messages[field] = err.Tag()
	}
	return messages

}
