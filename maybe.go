package maybe

type MayBeFunction func(x interface{}) (interface{}, error)

type MayBe struct {
	value interface{}
	err   error
}

func NewMayBe(value interface{}, err error) *MayBe {
	return &MayBe{value, err}
}

func (m *MayBe) OrElse(x interface{}, f MayBeFunction) *MayBe {
	if m.err != nil {
		return NewMayBe(f(x))
	}
	return m
}

func (m *MayBe) AndThen(f MayBeFunction) *MayBe {
	if m.err == nil {
		return NewMayBe(f(m.value))
	}
	return m
}
