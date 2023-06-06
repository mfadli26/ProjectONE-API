package validator

import (
	"html"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
)

func isNilValue(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return field.IsNil()
	default:
		return false
	}
}

func excludedIf(fl validator.FieldLevel) bool {

	arrParams := strings.Split(fl.Param(), " ")

	if len(arrParams)%2 == 1 {
		return false
	}

	field := isNilValue(fl.Field())
	if !field {

		for index := 0; index < len(arrParams); index += 2 {

			otherFieldName := strings.Split(fl.Param(), " ")[index]
			otherFieldValCheck := strings.Split(fl.Param(), " ")[index+1]
			var otherFieldVal reflect.Value

			if fl.Parent().Kind() == reflect.Ptr {
				otherFieldVal = fl.Parent().Elem().FieldByName(otherFieldName)
			} else {
				otherFieldVal = fl.Parent().FieldByName(otherFieldName)
			}

			if otherFieldValCheck == otherFieldVal.String() || otherFieldValCheck == otherFieldVal.Elem().String() {
				return false
			}
		}
	}

	return true
}

func AlphaWithSpaceAndDot(fl validator.FieldLevel) bool {

	params := fl.Field().String()

	regexParams := "^[a-zA-Z. ]*$"

	re := regexp.MustCompile(regexParams)

	response := re.MatchString(params)

	return response
}

func PasswordAtleastOneUppercaseLowercaseSymbolAndNumber(fl validator.FieldLevel) bool {

	params := fl.Field().String()
	number, upper, lower, special := false, false, false, false

	for _, c := range params {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		default:
		}
	}
	return number && upper && lower && special
}

func NotXssScript(fl validator.FieldLevel) bool {

	params := fl.Field().String()

	p := bluemonday.StrictPolicy()

	modifiedParams := p.Sanitize(params)
	modifiedParams = html.UnescapeString(modifiedParams)

	if len(params) != len(modifiedParams) {
		return false
	}

	return true
}
