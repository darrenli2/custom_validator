package custom_validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStructWithValidData(t *testing.T) {
	type User struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}

	user := User{Name: "John", Email: "john@example.com"}
	_, ok, err := ValidateStruct(user)

	assert.Equal(t, err, nil)
	assert.Equal(t, ok, true)
}

func TestValidateStructWithInvalidData(t *testing.T) {
	type User struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}

	user := User{Name: "John"}
	errs, ok, err := ValidateStruct(user)

	assert.Equal(t, err, nil)
	assert.Equal(t, ok, false)
	assert.Equal(t, errs, []string{"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag"})
}

func TestValidateStructWithInvalidStruct(t *testing.T) {
	// Create a slice of ints
	slice := []int{1, 2, 3}

	_, ok, err := ValidateStruct(slice)

	assert.Equal(t, err, errors.New("data is not a struct"))
	assert.Equal(t, ok, false)
}
