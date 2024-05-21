package event

import (
	"slices"
)

type Dispatcher struct {
	listeners     map[Type]map[int][]Listener
	minPriorities map[Type]int
	maxPriorities map[Type]int
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners:     make(map[Type]map[int][]Listener),
		minPriorities: make(map[Type]int),
		maxPriorities: make(map[Type]int),
	}
}

func (d *Dispatcher) Dispatch(e Event) error {
	minPriority := d.minPriorities[e.Type()]
	maxPriority := d.maxPriorities[e.Type()]
	listenerPriorities := d.listeners[e.Type()]

	for i := maxPriority; i >= minPriority; i-- {
		listeners := listenerPriorities[i]

		for _, listener := range listeners {
			err := listener.HandleEvent(e)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *Dispatcher) Listen(_type Type, listener Listener, priority int) {
	if priority < d.minPriorities[_type] {
		d.minPriorities[_type] = priority
	} else if priority > d.maxPriorities[_type] {
		d.maxPriorities[_type] = priority
	}

	listeners, ok := d.listeners[_type]
	if !ok {
		listeners = make(map[int][]Listener)
	}

	listeners[priority] = append(listeners[priority], listener)

	d.listeners[_type] = listeners
}

func (d *Dispatcher) ListenFunc(_type Type, listener func(Event) error, priority int) {
	d.Listen(_type, ListenerFunc(listener), priority)
}

func (d *Dispatcher) Unlisten(_type Type, listener Listener) {
	listenerPriorities := d.listeners[_type]

	for priority, listeners := range listenerPriorities {
		idx := slices.IndexFunc(listeners, func(lis Listener) bool {
			return lis == listener
		})

		if idx != -1 {
			d.listeners[_type][priority] = slices.Delete(listeners, idx, idx+1)
		}
	}
}

func (d *Dispatcher) UnlistenFunc(_type Type, listener func(Event) error) {
	d.Unlisten(_type, ListenerFunc(listener))
}

func (d *Dispatcher) HasListener(_type Type) bool {
	_, ok := d.listeners[_type]
	return ok
}

func (d *Dispatcher) Subscribe(subscriber Subscriber) {
	events := subscriber.SubscribedEvents()

	for _type, listenerPriorities := range events {
		for priority, listeners := range listenerPriorities {
			for _, listener := range listeners {
				d.Listen(_type, listener, priority)
			}
		}
	}
}

func (d *Dispatcher) SubscribeFunc(subscriber func() map[Type]map[int][]Listener) {
	d.Subscribe(SubscriberFunc(subscriber))
}

func (d *Dispatcher) Unsubscribe(subscriber Subscriber) {
	events := subscriber.SubscribedEvents()

	for _type, listenerPriorities := range events {
		for _, listeners := range listenerPriorities {
			for _, listener := range listeners {
				d.Unlisten(_type, listener)
			}
		}
	}
}

func (d *Dispatcher) UnsubscribeFunc(subscriber func() map[Type]map[int][]Listener) {
	d.Unsubscribe(SubscriberFunc(subscriber))
}
