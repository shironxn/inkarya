package pkg

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/shironxn/inkarya/internal/delivery/http/dto"
)

type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func NewValidator() *Validator {
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

func (v *Validator) Validate(s interface{}) []dto.FieldError {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	var fieldErrors []dto.FieldError
	for _, e := range err.(validator.ValidationErrors) {
		fieldErrors = append(fieldErrors, dto.FieldError{
			Field:   e.Field(),
			Message: e.Translate(v.translator),
		})
	}

	return fieldErrors
}
