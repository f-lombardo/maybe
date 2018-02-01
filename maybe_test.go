package maybe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrElseCallFunctionWithErrorMaybe(t *testing.T) {
	f := func(_ interface{}) (interface{}, error) {
		return "OK", nil
	}
	m := NewMayBe(nil, errors.New("First error"))
	m1 := m.OrElse(nil, f)
	assert.Nil(t, m1.err)
	assert.Equal(t, "OK", m1.value)
}

func TestOrElseDoesntCallFunctionWithSomeMaybe(t *testing.T) {
	f := func(_ interface{}) (interface{}, error) {
		return "OK", nil
	}
	m := NewMayBe("Original", nil)
	m1 := m.OrElse(nil, f)
	assert.Nil(t, m1.err)
	assert.Equal(t, "Original", m1.value)
}
