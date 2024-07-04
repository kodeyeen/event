package event

import "context"

type Listener interface {
	HandleEvent(context.Context, Event) error
}

type ListenerFunc func(Event) error

func (f ListenerFunc) HandleEvent(ctx context.Context, e Event) error {
	return f(e)
}
