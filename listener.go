package event

import "context"

type Listener interface {
	HandleEvent(context.Context, Event) error
}

type ListenerFunc func(context.Context, Event) error

func (f ListenerFunc) HandleEvent(ctx context.Context, ev Event) error {
	return f(ctx, ev)
}
