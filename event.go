package event

type Event interface {
	Type() Type
	Payload() any
}

type event struct {
	Event
	_type   Type
	payload any
}

func New(_type Type, payload any) *event {
	return &event{
		_type:   _type,
		payload: payload,
	}
}

func (e *event) Type() Type {
	return e._type
}

func (e *event) Payload() any {
	return e.payload
}
