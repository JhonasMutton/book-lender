package validate

import (
	"github.com/JhonasMutton/book-lender/pkg/errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/translations/en"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type Validator struct {
	validate  *validator.Validate
	translate ut.Translator
}

func NewValidator() *Validator {
	enTrans := en.New()
	uni := ut.New(enTrans, enTrans)

	trans, found := uni.GetTranslator("en")
	if !found {
		panic("translation not found")
	}

	validate := validator.New()
	if err := enTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		panic("error to register default translation:" + err.Error())
	}

	return &Validator{
		validate:  validate,
		translate: trans,
	}
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validate.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)

		var messages []string
		for _, m := range errs.Translate(v.translate){
			messages = append(messages, m)
		}

		return errors.New(strings.Join(messages,"; "))
	}

	return nil
}
