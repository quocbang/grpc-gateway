package validator

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// ValidateStructWithoutTag is validate each fields in the struct with custom
// value and map. Before call this function ensure your struct is an expose struct.
//
//	 example:
//			type Test {
//			   Name string
//			   Age int
//			}
//			data := Test{
//				Name: "test",
//				Age: 18
//			}
//
//			map[string]string{
//					"Name": "required",
//					"Age": "required,min=18",
//			}
//			rules := map[string]string {
//				"Name": "required"
//				"Age": "required,min=18"
//			}
//
// called example: ValidateStructWithoutTag[Test](rules, data)
func ValidateStructWithoutTag[T any](data any, rules map[string]string) error {
	var (
		t        T
		validate = validator.New()
	)
	if kind := reflect.TypeOf(t).Kind(); kind != reflect.Struct && kind != reflect.Pointer {
		return fmt.Errorf("unexpected type, expected is struct or pointer of struct")
	} else if kind == reflect.Pointer {
		if reflect.TypeOf(t).Elem().Kind() != reflect.Struct {
			return fmt.Errorf("unexpected type, expected is struct or pointer of struct")
		}
	}
	validate.RegisterStructValidationMapRules(rules, t)

	// validate struct
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorsMessage error
		for _, validationError := range validationErrors {
			if errorsMessage != nil {
				errorsMessage = fmt.Errorf("%v, %v", errorsMessage, customError(validationError)) // append multiple error.
			} else {
				errorsMessage = fmt.Errorf("%v", customError(validationError))
			}
		}
		return errorsMessage

	}
	return nil
}

func customError(validationError validator.FieldError) error {
	fieldName := validationError.Field()
	errTag := validationError.Tag()
	params := validationError.Param()
	value := validationError.Value()

	switch errTag {
	case "required":
		return fmt.Errorf("missing field [%v]", fieldName)
	case "min", "gte":
		if validationError.Kind() == reflect.String {
			return fmt.Errorf("length of field [%s] should be greater than or equal to %s but actual got %d", fieldName, params, len(value.(string)))
		} else {
			return fmt.Errorf("field [%s] should be greater than or equal to %s but actual got %v", fieldName, params, value)
		}
	case "max", "lte":
		if validationError.Kind() == reflect.String {
			return fmt.Errorf("length of field [%s] should be less than or equal to %s but actual got %v", fieldName, params, len(value.(string)))
		} else {
			return fmt.Errorf("field [%s] should be less than or equal to %s but actual got %v", fieldName, params, value)
		}
	case "eq":
		return fmt.Errorf("field [%s] should be equal %s but actual got %v", fieldName, params, value)
	case "ne":
		return fmt.Errorf("field [%s] should not be equal %s", fieldName, params)
	case "gt":
		return fmt.Errorf("field [%s] should be greater than %s but actual got %v", fieldName, params, value)
	case "lt":
		return fmt.Errorf("field [%s] should be less than %s but actual got %v", fieldName, params, value)
	case "email":
		return fmt.Errorf("field [%s] should be a valid email address example: abc@gmail.com but actual got %v", fieldName, value)
	default:
		return validationError.(error) // return default error.
	}
}
