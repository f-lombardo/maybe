package maybe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrElseCallsFunctionWithErrorMaybe(t *testing.T) {
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
	originalValue := "Original"
	m := NewMayBe(originalValue, nil)
	m1 := m.OrElse(fResult, f)
	assert.Nil(t, m1.err)
	assert.Equal(t, originalValue, m1.value)
}

func TestAndThenDoesntCallFunctionWithErrorMaybe(t *testing.T) {
	f := func(result interface{}) (interface{}, error) {
		return result, nil
	}
	m := NewMayBe(nil, errors.New("First error"))
	m1 := m.AndThen(f)
	assert.NotNil(t, m1.err)
	assert.Nil(t, m1.value)
}

func TestAndThenCallsFunctionWithSomeMaybe(t *testing.T) {
	f := func(x interface{}) (interface{}, error) {
		v, _ := x.(string)
		return v + "*", nil
	}
	originalValue := "Original"
	m := NewMayBe(originalValue, nil)
	m1 := m.AndThen(f)
	assert.Nil(t, m1.err)
	assert.Equal(t, originalValue+"*", m1.value)
}
