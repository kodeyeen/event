package event

type Type string

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

func (e *emptyEvt) Type() Type {
	return e._type
}

func (e *emptyEvt) Payload() any {
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

func (e *payloadEvt) Type() Type {
	return e._type
}

func (e *payloadEvt) Payload() any {
	return e.payload
}
