package errorhandler

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func GlobalErrorWrapper(err error) []GlobalErrorsDto {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		formattedErrors := make([]GlobalErrorsDto, len(ve))
		for i, fe := range ve {
			formattedErrors[i] = GlobalErrorsDto{fe.Field(), tagMessage(fe)}
		}
		return formattedErrors
	} else {
		formattedErrors := make([]GlobalErrorsDto, 1)
		formattedErrors[0] = GlobalErrorsDto{Message: err.Error()}
		return formattedErrors
	}
}
