package event

type Event interface {
	Type() Type
}

type Base struct {
	_type Type
}

func NewBase(_type Type) *Base {
	return &Base{
		_type: _type,
	}
}

func (e *Base) Type() Type {
	return e._type
}
