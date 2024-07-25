package event

type Event interface {
	Type() Type
	Payload() any
}

type emptyEvt struct {
	Event

	_type Type
}

func Empty(_type Type) Event {
	return &emptyEvt{_type: _type}
}

func (ev *emptyEvt) Type() Type {
	return ev._type
}

func (ev *emptyEvt) Payload() any {
	return nil
}

type payloadEvt struct {
	Event

	_type   Type
	payload any
}

func WithPayload(_type Type, payload any) Event {
	return &payloadEvt{
		_type:   _type,
		payload: payload,
	}
}

func (ev *payloadEvt) Type() Type {
	return ev._type
}

func (ev *payloadEvt) Payload() any {
	return ev.payload
}
