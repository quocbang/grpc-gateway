package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyAccountBuildUpdateFields(t *testing.T) {
	assertion := assert.New(t)

	verifyAccount := VerifyAccount{
		Username: "quocbang",
		IsUsed:   true,
	}

	result := verifyAccount.BuildUpdateFields()

	assertion.NotNil(result)
	expected := map[string]interface{}{
		"Username": "quocbang",
		"IsUsed":   true,
	}
	assertion.Equal(expected, result)
}

func TestAccountBuildUpdateFields(t *testing.T) {
	assertion := assert.New(t)

	account := Account{
		Username:       "quocbang",
		IsUserVerified: true,
	}

	result := account.BuildUpdateFields()

	assertion.NotNil(result)
	expected := map[string]interface{}{
		"Username":       "quocbang",
		"IsUserVerified": true,
	}
	assertion.Equal(expected, result)
}
