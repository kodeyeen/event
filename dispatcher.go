package event

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Dispatcher struct {
	listeners map[Type][]Listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[Type][]Listener),
	}
}

func (d *Dispatcher) HandleEvent(ctx context.Context, ev Event) error {
	listeners := d.listeners[ev.Type()]

	g, ctx := errgroup.WithContext(ctx)

	for _, listener := range listeners {
		g.Go(func() error {
			return listener.HandleEvent(ctx, ev)
		})
	}

	return g.Wait()
}

func (d *Dispatcher) Listen(_type Type, listener Listener) {
	d.listeners[_type] = append(d.listeners[_type], listener)
}

func (d *Dispatcher) ListenFunc(_type Type, listener func(context.Context, Event) error) {
	d.Listen(_type, ListenerFunc(listener))
}

func (d *Dispatcher) HasListener(_type Type) bool {
	_, ok := d.listeners[_type]
	return ok
}

func (d *Dispatcher) Subscribe(subscriber Subscriber) {
	events := subscriber.SubscribedEvents()

	for _type, listeners := range events {
		d.listeners[_type] = append(d.listeners[_type], listeners...)
	}
}

func (d *Dispatcher) SubscribeFunc(subscriber func() map[Type][]Listener) {
	d.Subscribe(SubscriberFunc(subscriber))
}
