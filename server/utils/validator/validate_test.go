package validator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStructWithoutTag(t *testing.T) {
	assertion := assert.New(t)
	type Test struct {
		Name string
		Age  int
	}
	rules := map[string]string{
		"Name": "required",
		"Age":  "required,min=18",
	}

	// good case.
	{
		// Arrange
		data := Test{
			Name: "test_name",
			Age:  18,
		}

		// Act
		err := ValidateStructWithoutTag[Test](data, rules)

		// Assert
		assertion.NoError(err)
	}

	// bad cases
	{ // missing Name field
		// Arrange
		data := Test{
			Age: 18,
		}

		// Act
		err := ValidateStructWithoutTag[Test](data, rules)

		// Assert
		assertion.Error(err)
		expected := fmt.Errorf("Key: 'Test.Name' Error:Field validation for 'Name' failed on the 'required' tag")
		assertion.Equal(expected.Error(), err.Error())
	}
	{ // missing Age field
		// Arrange
		data := Test{
			Name: "test_name",
		}

		// Act
		err := ValidateStructWithoutTag[Test](data, rules)

		// Assert
		assertion.Error(err)
		expected := fmt.Errorf("Key: 'Test.Age' Error:Field validation for 'Age' failed on the 'required' tag")
		assertion.Equal(expected.Error(), err.Error())
	}
	{
		// Arrange
		data := Test{
			Name: "test_name",
			Age:  12,
		}

		// Act
		err := ValidateStructWithoutTag[Test](data, rules)

		// Assert
		assertion.Error(err)
		expected := fmt.Errorf("Key: 'Test.Age' Error:Field validation for 'Age' failed on the 'min' tag")
		assertion.Equal(expected.Error(), err.Error())
	}
}
