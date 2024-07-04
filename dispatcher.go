package event

import (
	"context"
	"sync"
)

type Dispatcher struct {
	listeners map[Type][]Listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[Type][]Listener),
	}
}

func (d *Dispatcher) HandleEvent(e Event) error {
	listeners := d.listeners[e.Type()]

	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	wg := new(sync.WaitGroup)
	wg.Add(len(listeners))

	for _, listener := range listeners {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				err := listener.HandleEvent(e)
				if err != nil {
					cancel(err)
				}
			}
		}()
	}

	wg.Wait()

	return context.Cause(ctx)
}

func (d *Dispatcher) Listen(_type Type, listener Listener) {
	d.listeners[_type] = append(d.listeners[_type], listener)
}

func (d *Dispatcher) ListenFunc(_type Type, listener func(Event) error) {
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
