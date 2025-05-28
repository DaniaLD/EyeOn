package errorhandler

import "github.com/go-playground/validator/v10"

func tagMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "The field is required"
	}
	return fe.Error() // default error
}
