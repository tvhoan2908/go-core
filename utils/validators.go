package utils

import (
	"fmt"
	"go-core/config"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

var UniqueName validator.Func = func(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	fieldId := fl.Parent().FieldByName("ID")
	fieldUintId := uint64(0)
	if !reflect.ValueOf(fieldId).IsZero() {
		fieldUintId = fieldId.Interface().(uint64)
	}

	if fieldValue == "" {
		return true
	}
	params := strings.Split(fl.Param(), " ")
	if len(params) < 2 {
		return false
	}

	db := config.DB
	var id uint64
	db.Raw("SELECT id FROM "+params[0]+" WHERE "+params[1]+" = ?", fieldValue).Scan(&id)
	fmt.Println("ID", id, fieldUintId)
	if id == 0 || (id > 0 && id == fieldUintId) {
		return true
	}

	return false
}

func HandleErrorsMessage(e error) []map[string]string {
	ve := e.(validator.ValidationErrors)
	invalidFields := make([]map[string]string, 0)

	for _, e := range ve {
		errors := map[string]string{}
		errors[strcase.ToLowerCamel(e.Field())] = messageForTag(e.Tag())

		invalidFields = append(invalidFields, errors)
	}

	return invalidFields
}

// Convert message via tag name
func messageForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required."
	case "unique_name":
		return "This field is exist."
	}

	return ""
}
