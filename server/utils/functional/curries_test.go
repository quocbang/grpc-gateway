package functional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func more(x, y int) bool {
	return x > y
}

func TestTwoCurrying(t *testing.T) {
	assertion := assert.New(t)

	{
		curries := TwoCurrying(more)
		isMoreThan := curries(3)(2)
		assertion.True(isMoreThan)
	}
}

func sum(x, y, z int) int {
	return x + y + z
}

func TestThreeCurrying(t *testing.T) {
	assertion := assert.New(t)

	{
		curries := ThreeCurrying(sum)
		result := curries(1)(2)(3)
		assertion.Equal(6, result)
	}
}
