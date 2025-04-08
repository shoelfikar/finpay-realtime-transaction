package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/shoelfikar/finpay-realtime-transaction/model"
)


func ValidationErrors(err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		return true
	}

	return false
}

func GetErrorMessagevalidator(err interface{}) model.ValidationError {
	errors := err.(validator.ValidationErrors)

	var messages model.ValidationError

	for _, e := range errors {
		message := model.ValidationMessage{
			Field: e.Field(),
			Tag: e.Tag(),
			Value: e.Value(),
			Error: fmt.Sprintf("Field %s %s %s", e.Field(), e.Tag(), e.Param()),
		}

		messages.Errors = append(messages.Errors, &message)
	}

	return messages
}
