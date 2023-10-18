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
		expected := fmt.Errorf("missing field [Name]")
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
		expected := fmt.Errorf("missing field [Age]")
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
		expected := fmt.Errorf("field [Age] should be greater than or equal to 18 but actual got 12")
		assertion.Equal(expected.Error(), err.Error())
	}
}
