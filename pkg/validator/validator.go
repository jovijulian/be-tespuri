package validator

import (
	"fmt"
	"strings"

	"api-puri/pkg/i18n"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type ValidationErrorItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(c *fiber.Ctx, s interface{}) []ValidationErrorItem {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}
	lang := i18n.ResolveLanguage(c)
	return FormatValidationErrors(err, lang)
}

func FormatValidationErrors(err error, lang ...i18n.Language) []ValidationErrorItem {
	targetLang := i18n.LangEN
	if len(lang) > 0 {
		targetLang = lang[0]
	}

	var errors []ValidationErrorItem

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, v := range validationErrors {
			field := strings.ToLower(v.Field())
			var msg string

			if targetLang == i18n.LangID {
				switch v.Tag() {
				case "required":
					msg = fmt.Sprintf("Field %s wajib diisi", field)
				case "min":
					msg = fmt.Sprintf("Field %s minimal berisi %s karakter/nilai", field, v.Param())
				case "max":
					msg = fmt.Sprintf("Field %s maksimal berisi %s karakter/nilai", field, v.Param())
				case "gte":
					msg = fmt.Sprintf("Field %s harus lebih besar atau sama dengan %s", field, v.Param())
				case "lte":
					msg = fmt.Sprintf("Field %s harus lebih kecil atau sama dengan %s", field, v.Param())
				case "oneof":
					msg = fmt.Sprintf("Field %s harus salah satu dari: %s", field, v.Param())
				default:
					msg = fmt.Sprintf("Field %s tidak valid", field)
				}
			} else {
				switch v.Tag() {
				case "required":
					msg = fmt.Sprintf("Field '%s' is required", field)
				case "min":
					msg = fmt.Sprintf("Field '%s' must be at least %s characters/value", field, v.Param())
				case "max":
					msg = fmt.Sprintf("Field '%s' must not exceed %s characters/value", field, v.Param())
				case "gte":
					msg = fmt.Sprintf("Field '%s' must be greater than or equal to %s", field, v.Param())
				case "lte":
					msg = fmt.Sprintf("Field '%s' must be less than or equal to %s", field, v.Param())
				case "oneof":
					msg = fmt.Sprintf("Field '%s' must be one of: %s", field, v.Param())
				default:
					msg = fmt.Sprintf("Field '%s' is invalid", field)
				}
			}

			errors = append(errors, ValidationErrorItem{
				Field:   field,
				Message: msg,
			})
		}
	} else if err != nil {
		errors = append(errors, ValidationErrorItem{
			Field:   "general",
			Message: err.Error(),
		})
	}

	return errors
}
