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
		return err
	}
	return nil
}
