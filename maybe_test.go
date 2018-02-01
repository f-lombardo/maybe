package maybe

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOrElse(t *testing.T) {
	f := func(_ onterface{}) (interface{}, error) {
		return "OK", nil
	}
	m := NewMayBe(nil, errors.New("First error"))
	m1 := m.OrElse(f)
	assert.Nil(m1.err)
}
