package maybe

type MayBeFunction func() (interface{}, error)

type MayBe struct {
	value interface{}
	err   error
}

func NewMayBe(i interface{}, e error) *MayBe {
	return &MayBe{i, e}
}

func (m *MayBe) OrElse(f MayBeFunction) *MayBe {
	if m.err == nil {
		return NewMayBe(f())
	}
	return m
}
