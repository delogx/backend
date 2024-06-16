package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsValidHostName(fl validator.FieldLevel) bool {
	hostnameRegex := `^([a-zA-Z0-9][-a-zA-Z0-9]{0,62})(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\.?$`
	match, _ := regexp.MatchString(hostnameRegex, fl.Field().String())
	return match
}
