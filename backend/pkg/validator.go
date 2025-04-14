package pkg

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
)

type ValidatorService interface {
	Validate(s interface{}) []dto.FieldError
}

type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func NewValidator() ValidatorService {
	v := validator.New()

	// Setup English translator
	locale := en.New()
	uni := ut.New(locale, locale)
	trans, _ := uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(v, trans)

	return &Validator{
		validate:   v,
		translator: trans,
	}
}

// Validate performs validation on the given struct and returns field errors if any
func (v *Validator) Validate(s interface{}) []dto.FieldError {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	var fieldErrors []dto.FieldError
	for _, e := range err.(validator.ValidationErrors) {
		// Get the struct field
		fieldName := e.Field()
		field, _ := reflect.TypeOf(s).FieldByName(fieldName)

		// Get the JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			jsonName := strings.Split(jsonTag, ",")[0]
			fieldName = jsonName
		} else {
			fieldName = strings.ToLower(fieldName)
		}

		fieldErrors = append(fieldErrors, dto.FieldError{
			Field:   fieldName,
			Message: e.Translate(v.translator),
		})
	}

	return fieldErrors
}
