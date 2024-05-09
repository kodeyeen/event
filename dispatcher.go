package event

import (
	"reflect"
	"slices"
)

type Dispatcher struct {
	listeners map[Type][]listener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[Type][]listener),
	}
}

func Dispatch[T any](d *Dispatcher, evtType Type, evt T) bool {
	listeners := d.listeners[evtType]

	for _, l := range listeners {
		handler, ok := l.handler.(func(T) bool)
		if !ok {
			continue
		}

		callNext := handler(evt)

		if l.once {
			d.Off(evtType, l.handler)
		}

		if !callNext {
			return false
		}
	}

	return true
}

func (d *Dispatcher) On(evtType Type, handler any) {
	listeners := d.listeners[evtType]

	listeners = append(listeners, listener{
		handler: handler,
		once:    false,
	})

	d.listeners[evtType] = listeners
}

func (d *Dispatcher) Once(evtType Type, handler any) {
	listeners := d.listeners[evtType]

	listeners = append(listeners, listener{
		handler: handler,
		once:    true,
	})

	d.listeners[evtType] = listeners
}

func (d *Dispatcher) Off(evtType Type, handler any) {
	listeners := d.listeners[evtType]

	idx := slices.IndexFunc(listeners, func(l listener) bool {
		return reflect.ValueOf(l.handler).Pointer() == reflect.ValueOf(handler).Pointer()
	})

	if idx == -1 {
		return
	}

	d.listeners[evtType] = slices.Delete(listeners, idx, idx+1)
}

func (d *Dispatcher) Has(evtType Type) bool {
	_, ok := d.listeners[evtType]
	return ok
}
