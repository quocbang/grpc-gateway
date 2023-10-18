package stringly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSnakeCase(t *testing.T) {
	assertion := assert.New(t)
	target := "HelloFromQuocBangDev"

	result := ToSnakeCase(target)

	assertion.NotEmpty(result)
	assertion.Equal("hello_from_quoc_bang_dev", result)
}
