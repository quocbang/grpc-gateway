package activate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHTML(t *testing.T) {
	assertion := assert.New(t)

	// good case
	{
		// Arrange
		id := "test_id"
		secretCode := "test_secret_code"
		activateService := NewHTMLActivateService(id, secretCode)

		// Act
		htmlBody, err := activateService.GenerateHTML()

		// Assert
		assertion.NoError(err)
		assertion.NotNil(htmlBody)
	}

	// bad cases
	{ // missing id
		// Arrange
		activateService := NewHTMLActivateService("", "test_secret_code")

		// Act
		_, err := activateService.GenerateHTML()

		// Assert
		assertion.Error(err)
		expected := fmt.Errorf("missing id")
		assertion.Equal(expected, err)
	}
	{ // missing secret code
		// Arrange
		activateService := NewHTMLActivateService("test_id", "")

		// Act
		_, err := activateService.GenerateHTML()

		// Assert
		assertion.Error(err)
		expected := fmt.Errorf("missing secret code")
		assertion.Equal(expected, err)
	}

}
