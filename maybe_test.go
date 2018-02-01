package maybe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrElse(t *testing.T) {
	f := func(_ interface{}) (interface{}, error) {
		return "OK", nil
	}
	m := NewMayBe(nil, errors.New("First error"))
	m1 := m.OrElse(nil, f)
	assert.Nil(t, m1.err)
}
