package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateFieldOK(t *testing.T) {
	c := require.New(t)
	type Test struct {
		Name string `json:"name" validate:"required"`
		Age  int    `json:"age" validate:"required"`
	}
	test := Test{
		Name: "test",
	}

	actual := ValidateFields(test)

	expected := map[string]string{
		"age": "required",
	}

	c.Equal(expected, actual)
}
