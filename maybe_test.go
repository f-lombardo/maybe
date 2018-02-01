package maybe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrElseCallFunctionWithErrorMaybe(t *testing.T) {
	fResult := "OK"
	f := func(result interface{}) (interface{}, error) {
		return result, nil
	}
	m := NewMayBe(nil, errors.New("First error"))
	m1 := m.OrElse(fResult, f)
	assert.Nil(t, m1.err)
	assert.Equal(t, fResult, m1.value)
}

func TestOrElseDoesntCallFunctionWithSomeMaybe(t *testing.T) {
	fResult := "OK"
	f := func(result interface{}) (interface{}, error) {
		return result, nil
	}
	m := NewMayBe("Original", nil)
	m1 := m.OrElse(fResult, f)
	assert.Nil(t, m1.err)
	assert.Equal(t, "Original", m1.value)
}
