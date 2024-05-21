package event

type Listener interface {
	HandleEvent(Event) error
}

type ListenerFunc func(Event) error

func (f ListenerFunc) HandleEvent(e Event) error {
	return f(e)
}
